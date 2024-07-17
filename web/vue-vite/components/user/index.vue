<script setup>
import { Functions } from '@utils/functions/functions';
import { onMounted, reactive, ref } from 'vue';
import 'element-plus/theme-chalk/index.css';
import { ElConfigProvider, ElTable, ElTableColumn, ElPagination, ElButton } from 'element-plus';
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

const data = reactive({
  funcs: new Functions(),
  apiUrl: confs.apiUrl,
  cdnUrl: confs.cdnUrl,
  userData: [],
  user: {},
});

document.title = data.funcs.lang('User List');
const userService = createService(userModule, request);

onMounted(async () => {
  data.userData = await userService.list({page: 1, pageSize: 10});
});

function update(id) {

  console.log(id);
}

function del(id) {
  console.log(id);
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
            <h1>{{funcs.lang('User List')}}</h1>
          </div>
          <div class="col-sm-6">
            <ol class="breadcrumb float-sm-right">
              <li class="breadcrumb-item active">{{funcs.lang('User Manage')}}</li>
              <li class="breadcrumb-item active">{{funcs.lang('User List')}}</li>
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
          <form id="search" class="form-inline">
            <div class="form-group mx-sm-3 mb-2">
              <label class="font-weight-normal" for="username">{{funcs.lang('Username')}}&nbsp;</label>
              <input type="text" class="form-control" name="username" id="username" :placeholder="funcs.lang('Please Entry Username')">
            </div>
            <div class="form-group mx-sm-3 mb-2">
              <label class="font-weight-normal" for="nickname">{{funcs.lang('Nickname')}}&nbsp;</label>
              <input type="text" class="form-control" name="nickname" id="nickname"
                     :placeholder="funcs.lang('Please Entry Nickname')">
            </div>
            <div class="form-group mx-sm-3 mb-2">
              <button @click="search" class="btn btn-primary" type="button">{{funcs.lang('Search')}}</button>
            </div>
          </form>
        </div>
        <!-- /.card-header -->
        <div class="card-body">
          <button @click="showModal('addModal')" class="btn btn-primary mb-2" type="button">{{funcs.lang('Add User')}}</button>
          <!-- 表格 -->
          <el-table :data="data.userData?.data?.list" border style="width: 100%">
            <el-table-column type="selection" width="55" />
            <el-table-column prop="id" label="ID" sortable width="100"></el-table-column>
            <el-table-column prop="username" :label="funcs.lang('Username')" width="180"></el-table-column>
            <el-table-column prop="nickname" :label="funcs.lang('Nickname')"></el-table-column>
            <el-table-column prop="age" :label="funcs.lang('Age')"></el-table-column>
            <el-table-column prop="image" :label="funcs.lang('Avatar')"></el-table-column>
            <el-table-column prop="email" :label="funcs.lang('Email')"></el-table-column>
            <el-table-column prop="action" :label="funcs.lang('Action')">
              <el-button type="primary">修改</el-button>
              <el-button type="danger">删除</el-button>
            </el-table-column>
          </el-table>

          <!-- 分页 -->
          <el-config-provider :locale="funcs.getLang() !== 'zh-cn' ? en : zhCn">
            <el-pagination
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
                style="margin-top: 20px;"
            />
          </el-config-provider>
        </div>
        <!-- /.card-body -->
      </div>
      <!-- /.card -->

      <div id="addModal" class="modal fade" tabindex="-1">
        <div class="modal-dialog modal-lg" style="width: 85%;">
          <div class="modal-content">
            <div class="modal-header">
              <h5 v-if="JSON.stringify(data.user) !== '{}'" class="modal-title">{{funcs.lang('Modify User')}}</h5>
              <h5 v-else class="modal-title">{{funcs.lang('Add User')}}</h5>
              <button @click="clear" type="button" class="close" data-dismiss="modal" aria-label="Close">
                <span aria-hidden="true">&times;</span>
              </button>
            </div>
            <div class="modal-body">
              <form id="my-form" class="form-horizontal">
                <input v-if="JSON.stringify(data.user) !== '{}'" type="hidden" name="id" :value="data.user?.id">
                <div class="form-group">
                  <label for="username" class="col-sm-2 control-label font-weight-normal">{{funcs.lang('Username')}}</label>
                  <div class="col-sm-10">
                    <input name="username" type="text" class="form-control" id="username" :placeholder="funcs.lang('Please Entry Username')" disabled :value="data.user?.username">
                  </div>
                </div>
                <div class="form-group">
                  <label for="nickname" class="col-sm-2 control-label font-weight-normal">{{funcs.lang('Nickname')}}</label>
                  <div class="col-sm-10">
                    <input name="nickname" type="text" class="form-control" id="nickname" :placeholder="funcs.lang('Please Entry Nickname')" :value="data.user?.nickname">
                  </div>
                </div>
                <div class="form-group">
                  <label for="avatar" class="col-sm-2 control-label font-weight-normal">{{funcs.lang('Avatar')}}</label>
                  <div class="col-sm-10">
                    <input type="file" class="input-file form-control" id="avatar">
                    <input name="image" type="hidden" :value="data.user?.image" id="image">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 control-label font-weight-normal">{{funcs.lang('Gender')}}</label>
                  <div class="icheck-primary col-sm-10">
                    <div class="icheck-primary icheck-inline">
                      <input v-if="data.user?.sex === 0" name="sex" type="radio" id="sex1" checked value="0"/>
                      <input v-else name="sex" type="radio" id="sex1" value="0"/>
                      <label class="font-weight-normal" for="sex1">{{funcs.lang('Male')}}</label>
                    </div>
                    <div class="icheck-primary icheck-inline">
                      <input v-if="data.user?.sex === 1" name="sex" type="radio" id="sex2" checked value="1"/>
                      <input v-else name="sex" type="radio" id="sex2" value="1"/>
                      <label class="font-weight-normal" for="sex2">{{funcs.lang('Female')}}</label>
                    </div>
                  </div>
                </div>
              </form>
            </div>
            <div class="modal-footer">
              <button type="button" @click="clear" class="btn btn-secondary" data-dismiss="modal">{{funcs.lang('Close')}}</button>
              <button v-if="JSON.stringify(data.user) !== '{}'" type="button" class="btn btn-primary">{{funcs.lang('Modify')}}</button>
              <button v-else type="button" class="btn btn-primary">{{funcs.lang('Create')}}</button>
            </div>
          </div>
        </div>
      </div>

    </section>
    <!-- /.content -->
  </div>
  <!-- /.content-wrapper -->
</template>