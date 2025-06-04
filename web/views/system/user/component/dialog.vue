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
<script setup name="systemUserDialog">
import {nextTick, onMounted, reactive, ref} from 'vue';
import {storeToRefs} from 'pinia';
import {useRoutesList} from '/@/stores/routesList';
import {i18n} from '/@/static/i18n';
import {userApi} from '/@/api/user';
import {ElMessage} from 'element-plus';
import ConfigForm from '/@/components/form/index.vue';
import {genderDict, statusDict} from '/@/dict/user';

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
const api = userApi();

const state = reactive({
  roles: [],
  selectedRoleIds: [],
  ruleForm: {
    fullName: '',
    avatar: '',
    username: '',
    email: '',
    password: '',
    nickname: '',
    gender: 1,
    age: 0,
    status: 1,
    userRoles: []
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
      label: '用户名',
      prop: 'username',
      type: 'input',
      attrs: {
        placeholder: '请输入用户名',
        clearable: true
      },
      rules: [
        {
          required: true,
          message: '请输入用户名',
          trigger: 'blur'
        }
      ]
    },
    {
      label: '姓名',
      prop: 'fullName',
      type: 'input',
      attrs: {
        placeholder: '请输入姓名',
        clearable: true
      },
      rules: [
        {
          required: true,
          message: '请输入用姓名',
          trigger: 'blur'
        }
      ]
    },
    {
      label: '性别',
      prop: 'gender',
      type: 'radio',
      options: genderDict
    },
    {
      label: '头像',
      prop: 'avatar',
      type: 'input',
      attrs: {
        placeholder: '请输入头像地址',
        clearable: true
      }
    },
    {
      label: '用户角色',
      prop: 'userRoles',
      type: 'select',
      options: state.roles.map(role => ({label: role.name, value: role.id})),
      attrs: {
        multiple: true,
        clearable: true,
        placeholder: '请选择角色',
        class: 'w100'
      }
    },
    {
      label: '邮箱',
      prop: 'email',
      type: 'input',
      attrs: {
        placeholder: '请输入邮箱',
        clearable: true
      }
    },
    {
      label: '密码',
      prop: 'password',
      type: 'input',
      hidden: state.dialog.type !== 'add',
      attrs: {
        placeholder: '请输入密码',
        clearable: true
      },
      rules: [
        {
          required: true,
          message: '请输入密码',
          trigger: 'blur'
        }
      ]
    },
    {
      label: '昵称',
      prop: 'nickname',
      type: 'input',
      attrs: {
        placeholder: '请输入昵称',
        clearable: true
      },
      rules: [
        {
          required: true,
          message: '请输入昵称',
          trigger: 'blur'
        }
      ]
    },
    {
      label: '年龄',
      prop: 'age',
      type: 'input',
      attrs: {
        type: 'number',
        placeholder: '请输入年龄',
        clearable: true
      }
    },
    {
      label: '状态',
      prop: 'status',
      type: 'radio',
      options: statusDict
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

const openDialog = (type, row) => {
  state.ruleForm = {
    fullName: '',
    avatar: '',
    username: '',
    email: '',
    password: '',
    nickname: '',
    gender: 1,
    age: 0,
    status: 1,
    userRoles: []
  };
  if (type === 'edit') {
    Object.keys(state.ruleForm).forEach((key) => {
      if (row.hasOwnProperty(key)) {
        state.ruleForm[key] = row[key];
      }
    });
    // 设置角色 ID 数组用于 select 默认选中
    state.selectedRoleIds = row.userRoles?.map(item => item.roleId) || [];
    delete state.ruleForm['password'];
    state.dialog.title = '修改用户';
    state.dialog.submitTxt = '修 改';
  } else {
    state.selectedRoleIds = [];
    state.dialog.title = '新增用户';
    state.dialog.submitTxt = '新 增';
  }
  state.dialog.type = type;
  state.dialog.isShowDialog = true;
  // 清空表单，此项需加表单验证才能使用
  nextTick(() => {
    dialogFormRef.value && dialogFormRef.value.resetFields();
  });
};

const closeDialog = () => {
  state.dialog.isShowDialog = false;
  state.selectedRoleIds = [];
};

const onCancel = () => {
  closeDialog();
};

const onSubmit = async () => {
  state.ruleForm.userRoles = state.ruleForm.userRoles.map(roleId => {
    const role = state.roles.find(r => r.id === roleId);
    return {
      userId: props.row.id ?? 0,
      roleId: roleId,
      name: role ? role.name : ''
    };
  });

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

const getRoles = async () => {
  const data = await api.roleList({page: 1, pageSize: 10, isPage: false});
  state.roles = data.data;
};

onMounted(() => {
  getRoles();
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
