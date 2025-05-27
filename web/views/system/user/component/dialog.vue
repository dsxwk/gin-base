<template>
  <div class="system-menu-dialog-container">
    <el-dialog :title="state.dialog.title" v-model="state.dialog.isShowDialog" width="769px">
      <el-form ref="dialogFormRef" :model="state.ruleForm" size="default" label-width="80px">
        <el-row :gutter="35">
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="姓名">
              <el-input v-model="state.ruleForm.full_name" placeholder="请输入姓名" clearable></el-input>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="性别">
              <el-radio-group v-model.number="state.ruleForm.gender">
                <el-radio :value="1">男</el-radio>
                <el-radio :value="2">女</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="头像">
              <el-input v-model="state.ruleForm.avatar" placeholder="" clearable></el-input>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="用户名">
              <el-input v-model="state.ruleForm.username" placeholder="请输入用户名" clearable></el-input>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="邮箱">
              <el-input v-model="state.ruleForm.email" placeholder="" clearable></el-input>
            </el-form-item>
          </el-col>
          <template v-if="state.dialog.type === 'add'">
            <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
              <el-form-item label="密码">
                <el-input v-model="state.ruleForm.password" placeholder="" clearable></el-input>
              </el-form-item>
            </el-col>
          </template>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="昵称">
              <el-input v-model="state.ruleForm.nickname" placeholder="" clearable></el-input>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="年龄">
              <el-input v-model.number="state.ruleForm.age" placeholder="" clearable></el-input>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="状态">
              <el-radio-group v-model.number="state.ruleForm.status">
                <el-radio :value="1">启用</el-radio>
                <el-radio :value="2">停用</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
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
import {reactive, onMounted, ref, nextTick} from 'vue';
import {storeToRefs} from 'pinia';
import {useRoutesList} from '/@/stores/routesList';
import {i18n} from '/@/static/i18n';
import {userApi} from '/@/api/user';
import {ElMessage} from 'element-plus';
import {deepClone} from '/@/utils/other.js';

const props = defineProps({
  row: {
    type: Object,
    required: true,
    default: () => ({})
  }
})
const api = userApi();
// 定义子组件向父组件传值/事件
const emit = defineEmits(['refresh']);

// 定义变量内容
const dialogFormRef = ref();
const stores = useRoutesList();
const {routesList} = storeToRefs(stores);
const state = reactive({
  // 参数请参考 `/src/router/route.ts` 中的 `dynamicRoutes` 路由菜单格式
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
  },
  dialog: {
    isShowDialog: false,
    type: '',
    title: '',
    submitTxt: '',
  },
});
const defaultForm = {
  fullName: '',
  avatar: '',
  username: '',
  email: '',
  password: '',
  nickname: '',
  gender: 1,
  age: 0,
  status: 1,
};
// 获取 pinia 中的路由
const getData = (routes) => {
  const arr = [];
  routes.map((val) => {
    val['title'] = i18n.global.t(val.meta?.title);
    arr.push({...val});
    if (val.children) getData(val.children);
  });
  return arr;
};

// 打开弹窗
const openDialog = (type, row) => {
  if (type === 'edit') {
    Object.keys(state.ruleForm).forEach(key => {
      if (row.hasOwnProperty(key)) {
        if (typeof state.ruleForm[key] === 'object' && state.ruleForm[key] !== null) {
          Object.assign(state.ruleForm[key], row[key]);
        } else {
          state.ruleForm[key] = row[key];
        }
      }
    });
    delete state.ruleForm['password'];
    state.dialog.title = '修改用户';
    state.dialog.submitTxt = '修 改';
  } else {
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
// 关闭弹窗
const closeDialog = () => {
  state.dialog.isShowDialog = false;
  state.ruleForm = deepClone(defaultForm);
  dialogFormRef.value && dialogFormRef.value.resetFields();
};
// 取消
const onCancel = () => {
  closeDialog();
};
// 提交
const onSubmit = async () => {
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
};
// 页面加载时
onMounted(() => {
  state.menuData = getData(routesList.value);
});

// 暴露变量
defineExpose({
  openDialog,
});
</script>
