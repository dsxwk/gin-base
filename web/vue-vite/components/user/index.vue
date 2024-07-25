<script setup>
import {Functions} from '@utils/functions/functions';
import {onMounted, reactive, ref} from 'vue';
import pnotify from '@utils/pnotify/pnotify';
import 'element-plus/theme-chalk/index.css';
import {
  ElConfigProvider,
  ElTable,
  ElTableColumn,
  ElPagination,
  ElButton,
  ElDialog,
  ElForm,
  ElFormItem,
  ElInput,
  ElRadioGroup,
  ElRadio
} from 'element-plus';
import zhCn from 'element-plus/es/locale/lang/zh-cn';
import en from 'element-plus/es/locale/lang/en';
import confs from "@config/configs";
import createService from "@utils/service";
import request from "@utils/request";
import userModule from "@modules/admin/user/user";

let funcs = new Functions();
const currentPage4 = ref(1);
const pageSize4 = ref(10); // 正确定义 pageSize4 变量
const size = 'small'; // 根据需要设置大小
const disabled = false; // 根据需要设置禁用状态
const background = true; // 根据需要设置背景色
// 正确定义 handleSizeChange 方法
const handleSizeChange = (val) => {
  pageSize4.value = val;
};

const handleCurrentChange = (val) => {
  currentPage4.value = val;
};

const addForm = ref(null);
const updateForm = ref(null);
const searchForm = ref(null);
const addFormVisible = ref(false);
const updateFormVisible = ref(false);

const data = reactive({
  funcs: new Functions(),
  apiUrl: confs.apiUrl,
  cdnUrl: confs.cdnUrl,
  userData: [],
  user: {
    id: 0,
    username: '',
    full_name: '',
    nickname: '',
    gender: 1,
    password: '',
    age: 0,
  },
  userSearch: {
    username: '',
    nickname: '',
  }
});

document.title = data.funcs.lang('User List');
const userService = createService(userModule, request);

onMounted(async () => {
  await getUserList();
});

async function getUserList() {
  data.userData = await userService.list({page: 1, pageSize: 10});
}

async function showForm(func) {
  if (func === 'add') {
    addFormVisible.value = true;
    addForm.value?.resetFields();
  } else {
    updateFormVisible.value = true;
  }
}

async function submitAdd() {
  await addForm.value.validate();
  await userService.create(data.user);
  pnotify(funcs.lang('Created Successfully'));
  addFormVisible.value = false;
  await getUserList();
}

async function update(id) {
  await showForm('update');
  await getUserInfo(id);
}

async function getUserInfo(id) {
  const result = await userService.detail({id: id});
  data.user = result.data;
}

async function submitUpdate() {
  await updateForm.value.validate();
  await userService.update(data.user);
  pnotify(funcs.lang('Updated Successfully'));
  updateFormVisible.value = false;
  await getUserList();
}

async function del(id) {
  await userService.delete({id: id});
  pnotify(funcs.lang('Deleted Successfully'));
  await getUserList();
}

async function search() {
  // 参数合并
  data.userData = await userService.list({ ...data.userSearch, page: 1, pageSize: 10 });
}

function getGender(gender) {
  return gender === 1 ? '男' : gender === 2 ? '女' : '未知';
}
</script>
<template>
  <!-- Content Wrapper. Contains page content -->
  <div class="content-wrapper">
    <!-- Content Header (Page header) -->
    <section class="content-header">
      <div class="container-fluid">
        <div class="row mb-2">
          <div class="col-sm-6">
            <h1>{{ funcs.lang('User List') }}</h1>
          </div>
          <div class="col-sm-6">
            <ol class="breadcrumb float-sm-right">
              <li class="breadcrumb-item active">{{ funcs.lang('User Manage') }}</li>
              <li class="breadcrumb-item active">{{ funcs.lang('User List') }}</li>
            </ol>
          </div>
        </div>
      </div><!-- /.container-fluid -->
    </section>

    <!-- Main content -->
    <section class="content">

      <!-- Default box -->
      <div class="card">
        <div class="card-header">
          <ElForm ref="searchForm" :inline="true" class="demo-form-inline" :model="data.userSearch">
            <ElFormItem v-model="data.userSearch.username" :label="funcs.lang('Username')" pop="username">
              <ElInput :placeholder="funcs.lang('Please Entry Username')"/>
            </ElFormItem>
            <ElFormItem v-model="data.userSearch.nickname" :label="funcs.lang('Nickname')" pop="nickname">
              <ElInput :placeholder="funcs.lang('Please Entry Nickname')"/>
            </ElFormItem>
            <ElFormItem>
              <ElButton type="primary" @click="search">{{funcs.lang('Query')}}</ElButton>
            </ElFormItem>
          </ElForm>
        </div>
        <!-- /.card-header -->
        <div class="card-body">
          <ElButton @click="showForm('add')" type="primary" class="mb-2">{{funcs.lang('Add User')}}</ElButton>
          <!-- 表格 -->
          <ElTable :data="data.userData?.data?.list" border>
            <ElTableColumn type="selection" width="55"/>
            <ElTableColumn prop="id" label="ID" sortable width="100"></ElTableColumn>
            <ElTableColumn prop="username" :label="funcs.lang('Username')" width="180"></ElTableColumn>
            <ElTableColumn prop="full_name" :label="funcs.lang('FullName')"></ElTableColumn>
            <ElTableColumn prop="nickname" :label="funcs.lang('Nickname')"></ElTableColumn>
            <ElTableColumn prop="age" :label="funcs.lang('Age')"></ElTableColumn>
            <ElTableColumn prop="gender" :label="funcs.lang('Gender')">
              <template #default="scope">
                {{ getGender(scope.row.gender) }}
              </template>
            </ElTableColumn>
            <ElTableColumn prop="email" :label="funcs.lang('Email')"></ElTableColumn>
            <ElTableColumn prop="action" :label="funcs.lang('Action')">
              <template #default="scope">
                <ElButton @click="update(scope.row.id)" type="primary">{{ funcs.lang('Modify') }}</ElButton>
                <ElButton @click="del(scope.row.id)" type="danger">{{funcs.lang('Delete')}}</ElButton>
              </template>
            </ElTableColumn>
          </ElTable>

          <!-- 分页 -->
          <ElConfigProvider :locale="funcs.getLang() !== 'zh-cn' ? en : zhCn">
            <ElPagination
                :current-page="currentPage4"
                :page-size="pageSize4"
                :page-sizes="[10, 20, 30, 40, 50, 100]"
                :size="size"
                :disabled="disabled"
                :background="background"
                layout="total, sizes, prev, pager, next, jumper"
                v-if="data.userData?.data?.total"
                :total="data.userData?.data?.total"
                @size-change="handleSizeChange"
                @current-change="handleCurrentChange"
                class="mt-3"
            />
          </ElConfigProvider>
        </div>
        <!-- /.card-body -->
      </div>
      <!-- /.card -->

      <div class="card">
        <ElDialog v-model="updateFormVisible" :title="funcs.lang('Modify User')"
                  width="500">
          <ElForm ref="updateForm" label-width="auto" :model="data.user">
            <ElFormItem :label="funcs.lang('Username')" prop="username"
                        :rules="[
                {
                  required: true,
                  message: funcs.lang('Please Entry Username'),
                  trigger: 'blur',
                }
              ]"
            >
              <ElInput v-model="data.user.username" autocomplete="off"/>
            </ElFormItem>
            <ElFormItem :label="funcs.lang('FullName')" prop="full_name"
                        :rules="[
                {
                  required: true,
                  message: funcs.lang('Please Entry FullName'),
                  trigger: 'blur',
                }
              ]"
            >
              <ElInput v-model="data.user.full_name" autocomplete="off"/>
            </ElFormItem>
            <ElFormItem :label="funcs.lang('Nickname')" prop="nickname"
                        :rules="[
                {
                  required: true,
                  message: funcs.lang('Please Entry Nickname'),
                  trigger: 'blur',
                }
              ]"
            >
              <ElInput v-model="data.user.nickname" autocomplete="off"/>
            </ElFormItem>
            <ElFormItem :label="funcs.lang('Gender')" prop="gender"
                        :rules="[
                {
                  required: true,
                  message: funcs.lang('Please Select Gender'),
                  trigger: 'blur',
                }
              ]"
            >
              <ElRadioGroup v-model="data.user.gender">
                <ElRadio :value="1">{{ funcs.lang('Male') }}</ElRadio>
                <ElRadio :value="2">{{ funcs.lang('Female') }}</ElRadio>
              </ElRadioGroup>
            </ElFormItem>
          </ElForm>
          <template #footer>
            <div class="dialog-footer">
              <ElButton type="primary" @click="submitUpdate()">{{ funcs.lang('Modify') }}</ElButton>
              <ElButton @click="updateFormVisible = false">{{ funcs.lang('Cancel') }}</ElButton>
            </div>
          </template>
        </ElDialog>
        <ElDialog v-model="addFormVisible" :title="funcs.lang('Add User')" width="500">
          <ElForm ref="addForm" label-width="auto" :model="data.user">
            <ElFormItem :label="funcs.lang('Username')" prop="username"
                        :rules="[
                {
                  required: true,
                  message: funcs.lang('Please Entry Username'),
                  trigger: 'blur',
                }
              ]"
            >
              <ElInput v-model="data.user.username" autocomplete="off"/>
            </ElFormItem>
            <ElFormItem :label="funcs.lang('Password')" prop="password"
                        :rules="[
                {
                  required: true,
                  message: funcs.lang('Please Entry Password'),
                  trigger: 'blur',
                }
              ]"
            >
              <ElInput v-model="data.user.password" autocomplete="off"/>
            </ElFormItem>
            <ElFormItem :label="funcs.lang('Nickname')" prop="nickname"
                        :rules="[
                {
                  required: true,
                  message: funcs.lang('Please Entry Nickname'),
                  trigger: 'blur',
                }
              ]"
            >
              <ElInput v-model="data.user.nickname" autocomplete="off"/>
            </ElFormItem>
            <ElFormItem :label="funcs.lang('FullName')" prop="full_name"
                        :rules="[
                {
                  required: true,
                  message: funcs.lang('Please Entry FullName'),
                  trigger: 'blur',
                }
              ]"
            >
              <ElInput v-model="data.user.full_name" autocomplete="off"/>
            </ElFormItem>
            <ElFormItem :label="funcs.lang('Gender')" prop="gender"
                        :rules="[
                {
                  required: true,
                  message: funcs.lang('Please Select Gender'),
                  trigger: 'blur',
                }
              ]"
            >
              <ElRadioGroup v-model="data.user.gender">
                <ElRadio :value="1">{{ funcs.lang('Male') }}</ElRadio>
                <ElRadio :value="2">{{ funcs.lang('Female') }}</ElRadio>
              </ElRadioGroup>
            </ElFormItem>
          </ElForm>
          <template #footer>
            <div class="dialog-footer">
              <ElButton type="primary" @click="submitAdd()">{{funcs.lang('Create')}}</ElButton>
              <ElButton @click="addFormVisible = false">{{funcs.lang('Cancel')}}</ElButton>
            </div>
          </template>
        </ElDialog>
      </div>

    </section>
    <!-- /.content -->
  </div>
  <!-- /.content-wrapper -->
</template>