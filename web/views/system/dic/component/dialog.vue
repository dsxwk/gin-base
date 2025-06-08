<template>
	<div class="container">
		<el-dialog :title="state.dialog.title" v-if="state.dialog.isShowDialog" v-model="state.dialog.isShowDialog" width="769px">
      <ConfigForm
          ref="dialogFormRef"
          v-model:model="state.ruleForm"
          :form-config="formData"
          :rules="rules"
          :form-props="{
            labelWidth: '80px',
            size: 'default'
          }"
      >
      </ConfigForm>
			<template #footer>
				<span class="dialog-footer">
					<el-button @click="onCancel" size="default">取 消</el-button>
					<el-button type="primary" @click="onSubmit" size="default">{{ state.dialog.submitTxt }}</el-button>
				</span>
			</template>
		</el-dialog>
	</div>
</template>

<script setup name="systemDictDialog">
import {onMounted, reactive, ref, markRaw, nextTick} from 'vue';
import {ElMessage} from 'element-plus';
import {dictApi} from '/@/api/dict';
import ConfigForm from '/@/components/form/index.vue';
import CascaderLabel from '/@/components/form/CascaderLabel.vue';

const props = defineProps({
  row: {
    type: Object,
    required: true,
    default: () => ({})
  }
});

// 定义子组件向父组件传值/事件
const emit = defineEmits(['refresh']);

const api = dictApi();

// 定义变量内容
const dialogFormRef = ref();
const state = reactive({
  roles: [],
	// 参数请参考 `/router/route.js` 中的 `dynamicRoutes` 路由字典格式
	ruleForm: {
    dictSuperior: [], // 上级字典
    name: '', // 字典名称(英文)
    title: '', // 字典名称(中文)
    label: '', // 映射值
    sort: 0, // 排序
    extend: {}, // 扩展字段
    desc: "", // 描述
    status: 1, // 状态 1=启用 2=禁用
  },
	dictData: [], // 上级字典数据
	dialog: {
		isShowDialog: false,
		type: '',
		title: '',
		submitTxt: '',
	},
});
const formData = ref([
  {
    label: '英文名称',
    prop: 'name',
    type: 'input',
    col: 12,
    attrs: {
      placeholder: '请输入字典名称(英文)',
      clearable: true
    },
    rules: [
      {
        required: true,
        message: "请输入字典名称(英文)",
        trigger: "blur"
      },
    ]
  },
  {
    label: '中文名称',
    prop: 'title',
    type: 'input',
    col: 12,
    attrs: {
      placeholder: '请输入字典名称(中文)',
      clearable: true
    },
    rules: [
      {
        required: true,
        message: "请输入字典名称(中文)",
        trigger: "blur"
      },
    ]
  },
  {
    label: '映射值',
    prop: 'label',
    type: 'input',
    col: 12,
    attrs: {
      placeholder: '请输入映射值',
      clearable: true
    },
    rules: []
  },
  {
    label: '上级字典',
    prop: 'dictSuperior',
    type: 'cascader',
    options: () => {
      return state.dictData;
    },
    props: {
      checkStrictly: true,
      value: 'path',
      label: 'title',
    },
    attrs: {
      placeholder: '请选择上级字典',
      clearable: true,
      class: 'w100',
    },
    slotDefault: markRaw(CascaderLabel),
  },
  {
    label: '状态',
    prop: 'status',
    type: 'select',
    col: 12,
    options: [
      {
        label: "启用",
        value: 1,
      },
      {
        label: "禁用",
        value: 2,
      },
    ],
    attrs: {
      placeholder: '请选择状态',
      clearable: true,
      class: 'w100'
    },
  },
  {
    label: '扩展字典',
    prop: 'extend',
    type: 'textarea',
    col: 24,
    attrs: {
      placeholder: '请输入扩展字典(json字符串)',
      clearable: true,
      class: 'w100',
      rows: 3
    },
  },
  {
  label: '描述',
    prop: 'desc',
    type: 'textarea',
    col: 24,
    attrs: {
      placeholder: '请输入描述',
      clearable: true,
      class: 'w100',
      rows: 3
    },
  }
]);
const rules = {};
// 递归查找父级 path 数组
function findPathById(data, targetId, pathArr = []) {
  for (const item of data) {
    const newPathArr = [...pathArr, item.path];
    if (item.id === targetId) {
      return newPathArr;
    }
    if (item.children && item.children.length) {
      const found = findPathById(item.children, targetId, newPathArr);
      if (Array.isArray(found) && found.length) return found;
    }
  }
  return [];
}
// 递归查找字典项
function findByPath(data, path) {
  for (const item of data) {
    if (item.path === path) return item;
    if (item.children && item.children.length) {
      const found = findByPath(item.children, path);
      if (found) return found;
    }
  }
  return null;
}
// 打开弹窗
const openDialog = async (type, row) => {
  state.ruleForm = {
    dictSuperior: [], // 上级字典
    name: '', // 字典名称(英文)
    title: '', // 字典名称(中文)
    label: '', // 映射值
    sort: 0, // 排序
    extend: [], // 扩展字段
    desc: "", // 描述
    status: 1, // 状态 1=启用 2=禁用
  };
  if (type === 'edit') {
    const data = await detail(row.id);
    Object.keys(state.ruleForm).forEach((key) => {
      if (data.hasOwnProperty(key)) {
        state.ruleForm[key] = data[key];
      }
    });
    if (typeof data.extend === 'object' && data.extend !== null) {
      state.ruleForm.extend = JSON.stringify(data.extend, null, 2);
    } else {
      state.ruleForm.extend = data.extend || '';
    }
    // 设置上级字典默认选中
    if (row.pid) {
      state.ruleForm.dictSuperior = findPathById(state.dictData, row.pid);
    } else {
      state.ruleForm.dictSuperior = [];
    }
		state.dialog.title = '修改字典';
		state.dialog.submitTxt = '修 改';
	} else {
		state.dialog.title = '新增字典';
		state.dialog.submitTxt = '新 增';
	}
	state.dialog.type = type;
	state.dialog.isShowDialog = true;
  // 清空表单，此项需加表单验证才能使用
  await nextTick(() => {
    dialogFormRef.value && dialogFormRef.value.resetFields();
  });
};
// 关闭弹窗
const closeDialog = () => {
	state.dialog.isShowDialog = false;
};
// 取消
const onCancel = () => {
	closeDialog();
};
// 提交
const onSubmit = async () => {
  // 获取上级字典 id
  if (state.ruleForm.dictSuperior && state.ruleForm.dictSuperior.length > 0) {
    const lastPath = state.ruleForm.dictSuperior[state.ruleForm.dictSuperior.length - 1];
    const dict = findByPath(state.dictData, lastPath);
    state.ruleForm.pid = dict ? dict.id : 0;
  } else {
    state.ruleForm.pid = 0; // 顶级字典
  }

  if (state.ruleForm.extend !== '' && typeof state.ruleForm.extend === 'string') {
    try {
      state.ruleForm.extend = JSON.parse(state.ruleForm.extend);
    } catch (e) {
      ElMessage.error('扩展字典格式错误，请输入合法的 JSON 字符串');
      return;
    }
  }

  dialogFormRef.value.validate(async (valid) => {
    if (!valid) return;

    let msg = '';
    if (state.dialog.type === 'add') {
      await api.create(state.ruleForm);
      msg = '创建成功';
    } else {
      state.ruleForm.id = props.row.id;
      await api.update(state.ruleForm);
      msg = '更新成功';
    }
    ElMessage.success(msg);
    closeDialog();
    emit('refresh');
  });
};
// 详情
const detail = async (id) => {
  const res = await api.detail({id: id});

  return res.data;
};
const getData = async () => {
  const data = await api.list();
  state.dictData = data.data;
};
// 页面加载时
onMounted(() => {
	getData();
});
// 暴露变量
defineExpose({
	openDialog,
});
</script>
