<template>
  <div class="table-box">
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
        row-key="name"
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
      <!-- 菜单图标 -->
      <template #icon="scope">
        <el-icon :size="18" v-if="scope.row.meta.icon">
          <component :is="scope.row.meta.icon"></component>
        </el-icon>
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
    <MeunDrawer :drawerProps="drawerProps" @dataChange="dataChange"/>
    <ColSetting :colSetting="colSetting" v-model:isOpen="isOpen"/>
  </div>
</template>
<script setup>
import {CirclePlus, Delete, EditPen, Operation, Refresh, Search} from '@element-plus/icons-vue';
import ColSetting from '@/components/Table/ColSetting/index.vue';
import MeunDrawer from './components/drawer.vue';
import Functions from '@/utils/functions';
import pnotifyConfirm from '@/utils/pnotify/confirm';
import {reactive, ref, watch} from "vue";
import createService from '@/utils/service.js';
import menuModule from '@/app/modules/admin/menu';
import {menuJson} from '@/utils/data/menu';

const funcs = new Functions();
const isShowSearch = ref(true);
const menuService = createService(menuModule);
const drawerProps = reactive({
  row: {},
  title: '',
  visible: false,
});
const tablePlus = ref();
const initParam = reactive({
  page: 1,
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
  // 手动触发 TablePlus 更新数据
  tablePlus.value?.getTableList();
}
const getList = async (params) => {
  // return await menuService.list(params);
  return await {
    data: {
      list: menuJson,
      total: menuJson.length,
      page: 1,
      pageSize: menuJson.length
    }
  };
};
const dataCallback = (data) => {
  return {
    list: data?.data?.list,
    total: data?.data?.total,
    pageNum: data?.data?.page,
    pageSize: data?.data?.pageSize
  }
};
const create = () => {
  drawerProps.title = funcs.lang('Create');
  drawerProps.row = {};
  drawerProps.visible = true;
}
const batchDelete = async (ids) => {
  console.log("ids", ids)
  tablePlus.value?.clearSelection()
  tablePlus.value?.getTableList()
};
const columns = [
  { type: "index", label: funcs.lang('Index'), width: 120 },
  { prop: "meta.title", label: funcs.lang('Menu Name'), width:200, align: "left", search: { el: "input", props: {placeholder:funcs.lang('Please enter the Menu Name')} } },
  { prop: "meta.icon", label: funcs.lang('Icon'),  width:100},
  { prop: "name", label: funcs.lang('Menu Alias'), width:200, search: { el: "input", props: {placeholder:funcs.lang('Please enter the Menu Alias')} } },
  { prop: "path", label: funcs.lang('Route'), width: 300, search: { el: "input", props: {placeholder:funcs.lang('Please enter the Route')} } },
  { prop: "component", label: funcs.lang('Component Path'), width: 300 },
  { prop: "operation", label: funcs.lang('Operation'), width: 250, fixed: "right" }
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
const edit = async (row) => {
  drawerProps.title = funcs.lang('Update');
  console.log(row);
  if (typeof row.component === "function") {
    row.component = row.component.toString();
  }
  drawerProps.row = row;
  drawerProps.visible = true;
};
const del = async (row) => {
  pnotifyConfirm(funcs.lang('Are you sure you want to delete?'), 'error').then(res => {
    if (res) {
      userService.delete({id: row.id});
      tablePlus.value?.getTableList();
    } else {
      console.log('取消', row.id);
    }
  });
  console.log('delete', row.id);
}
</script>