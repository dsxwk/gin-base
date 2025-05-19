<template>
  <div class="table-demo-container layout-padding">
    <div class="table-demo-padding layout-padding-view layout-padding-auto">
      <TableSearch :search="state.tableData.search" @search="onSearch" />
      <Table
          ref="tableRef"
          v-bind="state.tableData"
          @delRow="onTableDelRow"
          @sortHeader="onSortHeader"
          row-key="id"
          :tree-props="{ children: 'children', hasChildren: 'hasChildren' }"
      >
        <template #tools>
          <div class="table-tool">
            <el-button size="default" type="primary" @click="onOpenAddMenu('add')">
              <el-icon>
                <ele-FolderAdd />
              </el-icon>
              新增菜单
            </el-button>
          </div>
        </template>
        <template #operation="{row}">
          <div class="flex items-center">
            <el-button type="primary" size="small" @click="onOpenEditMenu('edit', row)">编辑</el-button>
            <el-popconfirm title="确定删除吗？" @confirm="onTableDelRow(row)">
              <template #reference>
                <el-button size="small" type="danger">删除</el-button>
              </template>
            </el-popconfirm>
          </div>
        </template>
        <template #dialog>
          <MenuDialog ref="menuDialogRef" @refresh="getTableData(state.tableData.param)" />
        </template>
      </Table>
    </div>
  </div>
</template>

<script setup name="systemMenu">
import {defineAsyncComponent, reactive, ref, onMounted, h} from 'vue';
import { ElMessage } from 'element-plus';
import {menuJson} from '/@/static/menu';
import SvgIcon from '/@/components/svgIcon/index.vue';

// 引入组件
const Table = defineAsyncComponent(() => import('/@/components/table/index.vue'));
const TableSearch = defineAsyncComponent(() => import('/@/components/table/component/search.vue'));
const MenuDialog = defineAsyncComponent(() => import('/@/views/system/menu/component/dialog.vue'));

// 定义变量内容
const tableRef = ref();
const menuDialogRef = ref();
const state = reactive({
  tableData: {
    // 列表数据（必传）
    data: [],
    // 表头内容（必传，注意格式）
    header: [
      { key: 'menuType', colWidth: '', title: '菜单类型', type: 'text', isCheck: true,
        render: (scope) => {
          return scope.row?.menuType;
        }
      },
      { key: 'title', colWidth: '', title: '菜单名称', type: 'text', isCheck: true },
      { key: 'name', colWidth: '', title: '路由名称', type: 'text', isCheck: true },
      { key: 'path', colWidth: '', title: '路由路径', type: 'text', isCheck: true },
      { key: 'redirect', colWidth: '', title: '重定向', type: 'text', isCheck: true },
      { key: 'icon', colWidth: '', title: '菜单图标', isCheck: true,
        render: (scope) => {
          /*return h('el-icon', {
            class: scope.row?.meta?.icon
          });*/
          return h(SvgIcon, {
            name: scope.row?.meta?.icon,
          });
        },
      },
      { key: 'componentAlias', colWidth: '', title: '组件路径', type: 'text', isCheck: true },
      { key: 'isLink', colWidth: '', width: '70', height: '40', title: '链接地址', type: 'text', isCheck: true },
      { key: 'isHide', colWidth: '', width: '70', height: '40', title: '是否隐藏', type: 'select', isCheck: true },
    ],
    // 配置项（必传）
    config: {
      total: 0, // 列表总数
      loading: true, // loading 加载
      isBorder: true, // 是否显示表格边框
      isSerialNo: true, // 是否显示表格序号
      isSelection: true, // 是否显示表格多选
      isOperate: true, // 是否显示表格操作栏
      fixed: 'right', // 固定操作列
      operationWith: 200, // 固定操作列宽度
      isPage: true, // 是否不分页,默认不传或者为false时展示分页,为true时不展示分页
    },
    // 搜索表单，动态生成（传空数组时，将不显示搜索，注意格式）
    search: [
      { label: '菜单类型', prop: 'menuType', placeholder: '请输入菜单类型', required: true, type: 'input' },
      { label: '菜单名称', prop: 'title', placeholder: '请输入菜单名称', required: false, type: 'input' },
      { label: '路由名称', prop: 'name', placeholder: '请输入路由名称', required: false, type: 'input' },
      {
        label: '是否隐藏',
        prop: 'isHide',
        placeholder: '请选择',
        required: false,
        type: 'select',
        options: [
          { label: '隐藏', value: true },
          { label: '不隐藏', value: false },
        ],
      },
    ],
    // 搜索参数、搜索时传给后台的值,`getTableData` 中使用）
    param: {},
    // 打印标题
    printName: 'ginBaseAdmin 表格打印演示',
  },
});
// 打开新增菜单弹窗
const onOpenAddMenu = (type) => {
  menuDialogRef.value.openDialog(type);
};
// 打开编辑菜单弹窗
const onOpenEditMenu = (type, row) => {
  menuDialogRef.value.openDialog(type, row);
};
// 初始化列表数据
const getTableData = async (param) => {
  state.tableData.config.loading = true;
  state.tableData.data = menuJson.data;
  setTimeout(() => {
    state.tableData.config.loading = false;
  }, 500);
};
// 搜索点击时表单回调
const onSearch = (data) => {
  const filterData = Object.fromEntries(
      Object.entries(data).filter(([key, value]) => value !== null && value !== undefined && value !== '')
  );
  state.tableData.param = Object.assign({}, state.tableData.param, filterData);
  getTableData(state.tableData.param);
};
// 删除当前项回调
const onTableDelRow = (row) => {
  ElMessage.success(`删除${row.name}成功！`);
  state.tableData.data = state.tableData.data.filter((item) => item.id !== row.id);
};
// 拖动显示列排序回调
const onSortHeader = (data) => {
  state.tableData.header = data;
};
// 页面加载时
onMounted(() => {
  getTableData(state.tableData.param);
});
</script>

<style scoped lang="scss">
.table-demo-container {
  .table-demo-padding {
    padding: 15px;
    .table-demo {
      flex: 1;
      overflow: hidden;
    }
  }
}
</style>