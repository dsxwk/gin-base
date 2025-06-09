<template>
  <div class="system-menu-dialog-container">
    <el-dialog :title="state.dialog.title" v-if="state.dialog.isShowDialog" v-model="state.dialog.isShowDialog"
               width="769px">
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
<script setup name="systemMenuActionOperationDialog">
import {computed, nextTick, onMounted, reactive, ref} from 'vue';
import {ElMessage} from 'element-plus';
import ConfigForm from "/@/components/form/index.vue";
import {actionTypeDict, btnSizeDict, btnStyleDict, btnTypeDict, isConfirmDict} from '/@/dict/menu';
import {menuApi} from '/@/api/menu';

const api = menuApi();
const props = defineProps({
  menuId: {
    type: Number,
    required: true,
    default: 0
  },
  row: {
    type: Object,
    default: () => ({})
  }
});
const state = reactive({
  roles: [],
  selectedRoleIds: [],
  ruleForm: {
    menuId: "", // 菜单id
    type: "", // 类型 1=header 2=operation
    btnType: "", // 按钮类型 text|btn
    btnStyle: "", // 按钮样式
    btnSize: "", // 按钮尺寸
    isConfirm: 2, // 是否确认
    name: "", // 功能名称
    isLink: false, // 是否为链接 1=是 2=否
    sort: 0, // 排序
    actionRoles: [], // 功能角色
  },
  dialog: {
    isShowDialog: false,
    type: '',
    title: '',
    submitTxt: '',
  },
});

const formData = computed(() => [
  {
    label: "菜单id",
    prop: "menuId",
    type: "input",
    attrs: {
      disabled: true,
      placeholder: "请输入菜单id",
      clearable: true
    },
    rules: [
      {
        required: true,
        message: "请输入菜单id",
        trigger: "blur"
      },
    ],
  },
  {
    label: "类型",
    prop: "type",
    type: "select",
    options: actionTypeDict,
    attrs: {
      placeholder: "请选择类型",
      clearable: true
    },
    rules: [
      {
        required: true,
        message: "请选择类型",
        trigger: "blur"
      }
    ]
  },
  {
    label: "按钮类型",
    prop: "btnType",
    type: "select",
    col: 12,
    options: btnTypeDict,
    attrs: {
      placeholder: '请选择按钮类型',
      clearable: true,
      class: 'w100'
    },
    rules: [
      {
        required: true,
        message: "请输入按钮类型",
        trigger: "blur"
      },
    ],
  },
  {
    label: "按钮样式",
    prop: "btnStyle",
    type: "select",
    col: 12,
    options: btnStyleDict,
    attrs: {
      placeholder: '请选择按钮样式',
      clearable: true,
      class: 'w100'
    },
    rules: [
      {
        required: true,
        message: "请输入按钮样式",
        trigger: "blur"
      },
    ],
  },
  {
    label: "按钮尺寸",
    prop: "btnSize",
    type: "select",
    col: 12,
    options: btnSizeDict,
    attrs: {
      placeholder: '请选择按钮尺寸',
      clearable: true,
      class: 'w100'
    },
    rules: [
      {
        required: true,
        message: "请输入按钮尺寸",
        trigger: "blur"
      },
    ],
  },
  {
    label: "是否确认",
    prop: "isConfirm",
    type: "select",
    col: 12,
    options: isConfirmDict,
    attrs: {
      placeholder: '请选择是否确认',
      clearable: true,
      class: 'w100'
    },
    rules: [
      {
        required: true,
        message: "请输入是否确认",
        trigger: "blur"
      },
    ],
  },
  {
    label: "功能名称",
    prop: "name",
    type: "input",
    attrs: {
      placeholder: "请输入功能名称",
      clearable: true
    },
    rules: [
      {
        required: true,
        message: "请输入功能名称",
        trigger: "blur"
      },
    ],
  },
  {
    label: "是否为链接",
    prop: "isLink",
    type: "switch",
    labelWidth: "120px",
    attrs: {
      placeholder: "请选择是否为链接",
      clearable: true
    },
    rules: [
      {
        required: true,
        message: "请选择是否为链接",
        trigger: "blur"
      },
    ],
  },
  {
    label: '功能角色',
    prop: 'actionRoles',
    type: 'select',
    col: 12,
    options: () => {
      return state.roles.map(role => ({label: role.name, value: role.id}));
    },
    attrs: {
      placeholder: '请选择角色',
      multiple: true,
      clearable: true,
      class: 'w100'
    },
  },
  {
    label: "排序",
    prop: "sort",
    type: "inputNumber",
    attrs: {
      type: "number",
      placeholder: "请输入排序",
      clearable: true
    },
    rules: [],
  }
]);

const rules = {};
// 定义子组件向父组件传值/事件
const emit = defineEmits(['refresh']);

// 定义变量内容
const dialogFormRef = ref();
// 打开弹窗
const openDialog = async (type, row) => {
  state.ruleForm = {
    menuId: props.menuId, // 菜单id
    type: "", // 类型 1=header 2=operation
    btnType: "", // 按钮类型 text|btn
    btnStyle: "", // 按钮样式
    btnSize: "", // 按钮尺寸
    isConfirm: 2, // 是否确认
    name: "", // 功能名称
    isLink: false, // 是否为链接 1=是 2=否
    sort: 0, // 排序
    actionRoles: [],
  };
  if (type === 'edit') {
    const data = await detail(row.id);
    Object.keys(state.ruleForm).forEach((key) => {
      if (data.hasOwnProperty(key)) {
        state.ruleForm[key] = data[key];
      }
    });
    // 设置角色 ID 数组用于 select 默认选中
    state.selectedRoleIds = state.ruleForm.actionRoles?.map(item => item.roleId) || [];
    state.dialog.title = '修改功能';
    state.dialog.submitTxt = '修 改';
  } else {
    state.selectedRoleIds = [];
    state.dialog.title = '新增功能';
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
  state.selectedRoleIds = [];
  state.dialog.isShowDialog = false;
};
// 取消
const onCancel = () => {
  closeDialog();
};
// 提交
const onSubmit = async () => {
  state.ruleForm.actionRoles = state.ruleForm.actionRoles?.map(roleId => {
    const role = state.roles.find(r => r.id === roleId);
    return {
      menuId: props.menuId ?? 0,
      roleId: roleId,
      name: role ? role.name : ''
    };
  }) ?? [];

  dialogFormRef.value.validate(async (valid) => {
    if (!valid) return;

    let msg = '';
    if (state.dialog.type === 'add') {
      await api.createAction(state.ruleForm);
      msg = '创建成功';
    } else {
      state.ruleForm.menuId = props.menuId;
      state.ruleForm.id = props.row.id;
      state.ruleForm.actionId = props.row.id;
      await api.updateAction(state.ruleForm);
      msg = '更新成功';
    }
    ElMessage.success(msg);
    closeDialog();
    emit('refresh');
  });
};
// 功能详情
const detail = async (id) => {
  const res = await api.actionDetail({id: id, menuId: props.menuId});
  const data = res.data;
  data.actionRoles = data.actionRoles.map(role => role.roleId);
  return data;
};
// 获取角色
const getRoles = async () => {
  const data = await api.roleList({page: 1, pageSize: 10, isPage: false});
  state.roles = data.data;
};
// 页面加载时
onMounted(() => {
  getRoles();
});
// 暴露变量
defineExpose({
  openDialog,
});
</script>