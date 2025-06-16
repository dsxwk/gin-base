<template>
  <div class="system-menu-dialog-container">
    <el-dialog :title="state.dialog.title" v-if="state.dialog.isShowDialog" v-model="state.dialog.isShowDialog" width="769px">
      <ConfigForm
          ref="dialogFormRef"
          v-model:model="state.ruleForm"
          :form-config="getFormData()"
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
<script setup name="articleDialog">
import {nextTick, onMounted, reactive, ref} from 'vue';
import {storeToRefs} from 'pinia';
import {useRoutesList} from '/@/stores/routesList';
import {i18n} from '/@/static/i18n';
import {articleApi} from '/@/api/article';
import {ElMessage} from 'element-plus';
import ConfigForm from '/@/components/form/index.vue';

const props = defineProps({
  row: {
    type: Object,
    required: true,
    default: () => ({})
  }
});
const emit = defineEmits(['refresh']);
const dialogFormRef = ref();
const stores = useRoutesList();
const {routesList} = storeToRefs(stores);
const api = articleApi();

const state = reactive({
  roles: [],
  selectedRoleIds: [],
  ruleForm: {
    title: '',
    user: {
      fullName: '',
    },
    category: {
      name: '',
    },
    content: '',
    tag: '',
  },
  dialog: {
    isShowDialog: false,
    type: 'add',
    title: '',
    submitTxt: ''
  }
});

const getFormData = () => {
  return [
    {
      label: '标题',
      prop: 'title',
      type: 'input',
      span: 24,
      attrs: {
        placeholder: '请输入标题',
        clearable: true
      },
      rules: [
        {
          required: true,
          message: '请输入用标题',
          trigger: 'blur'
        }
      ]
    },
    {
      label: '标签',
      prop: 'tag',
      type: 'input',
      span: 24,
      attrs: {
        placeholder: '请输入标签',
        clearable: true
      }
    },
    {
      label: '内容',
      prop: 'content',
      type: 'textarea',
      span: 24,
      attrs: {
        placeholder: '请输入内容',
        clearable: true
      },
      rules: [
        {
          required: true,
          message: '请输入内容',
          trigger: 'blur'
        }
      ]
    },
  ];
};
const rules = {};
const getData = (routes) => {
  const arr = [];
  routes.forEach((val) => {
    val['title'] = i18n.global.t(val.meta?.title);
    arr.push({...val});
    if (val.children) getData(val.children);
  });
  return arr;
};

const openDialog = async (type, row) => {
  state.ruleForm = {
    title: '',
    user: {
      fullName: '',
    },
    category: {
      name: '',
    },
    content: '',
    tag: '',
  };
  if (type === 'edit') {
    const data = await detail(row.id);
    Object.keys(state.ruleForm).forEach((key) => {
      if (data.hasOwnProperty(key)) {
        state.ruleForm[key] = data[key];
      }
    });
    state.dialog.title = '修改文章';
    state.dialog.submitTxt = '修 改';
  } else {
    state.selectedRoleIds = [];
    state.dialog.title = '新增文章';
    state.dialog.submitTxt = '新 增';
  }
  state.dialog.type = type;
  state.dialog.isShowDialog = true;
  // 清空表单，此项需加表单验证才能使用
  await nextTick(() => {
    dialogFormRef.value && dialogFormRef.value.resetFields();
  });
};

const closeDialog = () => {
  state.dialog.isShowDialog = false;
};

const onCancel = () => {
  closeDialog();
};

const onSubmit = async () => {
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
onMounted(() => {
  state.menuData = getData(routesList.value);
});

defineExpose({openDialog});
</script>

<style scoped>
.dialog-footer {
  display: flex;
  justify-content: flex-end;
}
</style>
