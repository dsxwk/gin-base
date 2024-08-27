<template>
  <TablePlus
      :isShowSearch="isShowSearch"
      ref="tablePlus"
      :title="funcs.lang('List')"
      :columns="columns"
      :request-api="getList"
      :init-param="initParam"
      :pagination="true"
      :data-callback="dataCallback"
      :reset-callback="resetCallback"
      :operationBtnText="operationBtnText"
      :emptyListText="funcs.lang('No Data')"
      row-key="id"
  >
    <!-- 表格 header 按钮 -->
    <template #tableHeader="scope">
      <el-button type="primary" :icon="CirclePlus" class="mb10" @click="create">{{ funcs.lang('Create') }}</el-button>
      <el-button
          type="danger"
          :icon="Delete"
          plain
          :disabled="!scope.isSelected"
          class="mb10"
          @click="batchDelete(scope.selectedListIds)"
      >
        {{ funcs.lang('Batch Delete') }}
      </el-button>
    </template>
    <!-- 表格操作 -->
    <template #operation="scope">
      <el-button
          type="primary"
          link
          :icon="EditPen"
          @click="edit(scope.row)"
      >
        {{ funcs.lang('Edit') }}
      </el-button>
      <el-button type="primary" link :icon="Delete" @click="del(scope.row)">{{ funcs.lang('Delete') }}</el-button>
    </template>
    <template #toolButton="scope">
      <el-button :icon="Refresh" circle @click="reload"/>
      <el-button v-if="columns.length" :icon="Operation" circle @click="openColSetting"/>
      <el-button :icon="Search" circle @click="isShowSearch = !isShowSearch"/>
    </template>
  </TablePlus>
  <UserDrawer :drawerProps="drawerProps" @updateIsPublish="updateUserStatus" @dataChange="dataChange"/>
  <ColSetting :colSetting="colSetting" v-model:isOpen="isOpen"/>
</template>
<script setup>
import { ref, reactive, watch, h } from 'vue';
import { ElSwitch } from 'element-plus';
import { CirclePlus, Delete, EditPen, Refresh, Operation, Search } from '@element-plus/icons-vue';
import Functions from '@/utils/functions';
import userModule from '@/app/modules/admin/user';
import createService from '@/utils/service';
import UserDrawer from './components/drawer.vue';
import ColSetting from '@/components/Table/ColSetting/index.vue';
import pnotifyConfirm from '@/utils/pnotify/confirm';
import {genderDict} from '@/app/modules/admin/user/dict';

const funcs = new Functions();
const isShowSearch = ref(true);
const userService = createService(userModule);
const drawerProps = reactive({
  row: {
    is_publish: 1
  },
  title: '',
  visible: false,
});
const tablePlus = ref();
const initParam = reactive({
  pageNo: 1,
  pageNum: 1,
  pageSize: 10
});
const operationBtnText = reactive({
  search: funcs.lang('Search'),
  reset: funcs.lang('Reset'),
  ArrowDown: funcs.lang('Expand'),
  ArrowUp: funcs.lang('Collapse'),
});
const reload = () => {
  location.reload();
};
watch(initParam, () => {
  dataChange();
}, { deep: true });
const dataChange = () => {
  initParam.page = initParam.page || initParam.pageNum;
  initParam.page_size = initParam.pageSize || initParam.pageSize;
  // 手动触发 TablePlus 更新数据
  tablePlus.value?.getTableList();
}
const getList = async (params) => {
  params.page = params.page || initParam.pageNum;
  params.page_size = params.pageSize || initParam.pageSize;
  return await userService.list(params);
};
const dataCallback = (data) => {
  return {
    list: data?.data?.list,
    total: data?.data?.total,
    pageNum: data?.data?.page,
    pageSize: data?.data?.page_size
  }
};
const create = () => {
  drawerProps.title = funcs.lang('Create');
  drawerProps.row = {};
  drawerProps.visible = true;
}
const updateUserStatus = (value) => {
  drawerProps.row.status = value;
};
const batchDelete = async (ids) => {
  console.log("ids", ids)
  tablePlus.value?.clearSelection()
  tablePlus.value?.getTableList()
};
// 表格配置项
const columns = [
  { type: "selection", fixed: "left", width: 80 },
  { type: "index", label: funcs.lang('Index'), width: 120 },
  { type: "expand", label: funcs.lang('Expand'), width: 120,
    render: (scope) => {
      return h('span', JSON.stringify(scope.row) || '-');
    }
  },
  {
    prop: "id",
    width: 120,
    label: funcs.lang('ID')
  },
  {
    prop: "username",
    width: 200,
    label: funcs.lang('Username'),
    search: {
      el: "input",
      props: { maxlength: 30, placeholder: funcs.lang('Please enter the username') }
    }
  },
  {
    prop: "full_name",
    label: funcs.lang('FullName'),
    width: 200,
    search: {
      el: "input",
      props: { maxlength: 30, placeholder: funcs.lang('Please enter the full name') }
    }
  },
  {
    prop: "email",
    label: funcs.lang('Email'),
    width: 200,
    search: {
      el: "input",
      props: { maxlength: 30, placeholder: funcs.lang('Please enter the email') }
    }
  },
  {
    prop: "age",
    label: funcs.lang('Age'),
    width: 200,
    search: {
      el: "input",
      props: { maxlength: 30, placeholder: funcs.lang('Please enter the age') }
    }
  },
  {
    prop: "gender",
    label: funcs.lang('Gender'),
    width: 200,
    enum: genderDict,
    search: { el: "tree-select", props: { filterable: true, placeholder: funcs.lang('Please select') } },
  },
  {
    prop: "status",
    label: funcs.lang('User Status'),
    width: 200,
    search: { el: "tree-select", props: { filterable: true, placeholder: funcs.lang('Please select') } },
    render: (scope) => {
      return h(
          ElSwitch,
          {
            // 根据状态设置开关的初始值
            modelValue: scope.row.status,
            // 开启时的文字描述
            'active-text': funcs.lang('Enabled'),
            // 关闭时的文字描述
            'inactive-text': funcs.lang('Disabled'),
            // 开启时的值
            'active-value': 1,
            // 关闭时的值
            'inactive-value': 2,
            onChange: (value) => {
              // 切换开关状态时调用 switchStatus 方法
              switchStatus(scope.row, value);
            }
          }
      );
    }
  },
  {
    prop: "created_at",
    label: funcs.lang('Create Time'),
    width: 200,
    render: (scope) => {
      return h('span', scope.row.created_at || '-');
    }
  },
  { prop: 'operation', label: funcs.lang('Operation'), fixed: 'right', width: 200 }
];
// 列设置,需要过滤掉不需要设置的列
const colSetting = columns.filter(item => {
  const { type, prop } = item;
  return !['selection', 'index', 'expand'].includes(type) && prop !== 'operation';
});
const isOpen = ref(false);
const openColSetting = () => {
  isOpen.value = !isOpen.value;
};
const resetCallback = () => {
  console.log('resetCallBack');
  tablePlus.value?.getTableList();
};
const switchStatus = async (row, newValue) => {
  console.log('id=' + row.id + ' User status changed:', newValue === 1 ? funcs.lang('Enabled') : funcs.lang('Disabled'));

  // 更新api...
  // 更新表格中当前行的 status 值
  row.status = newValue;
};
const edit = async (row) => {
  drawerProps.title = funcs.lang('Update');
  drawerProps.row = row;
  drawerProps.visible = true;
};
const del = async (row) => {
  pnotifyConfirm(funcs.lang('Are you sure you want to delete?'), 'error').then(res => {
    if (res) {
      userService.delete({id: row.id});
      getList(initParam);
    } else {
      console.log('取消', row.id);
    }
  });
  console.log('delete', row.id);
}
</script>