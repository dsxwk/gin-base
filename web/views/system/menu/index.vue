<template>
  <div class="table-demo-container layout-padding">
    <div class="table-demo-padding layout-padding-view layout-padding-auto">
      <TableSearch :search="state.tableData.search" @search="onSearch"/>
      <Table
          dev
          ref="tableRef"
          v-bind="state.tableData"
          @delRow="onTableDelRow"
          @sortHeader="onSortHeader"
          @pageChange="onTablePageChange"
          row-key="id"
          :tree-props="{ children: 'children', hasChildren: 'hasChildren' }"
      >
        <template #tools>
          <div class="table-tool">
            <el-button size="default" type="primary" @click="onOpenAddMenu('add')">
              <el-icon>
                <ele-FolderAdd/>
              </el-icon>
              新增菜单
            </el-button>
          </div>
        </template>
        <template #operation="{row}">
          <div class="flex items-center">
            <el-button type="primary" size="small" @click="onOpenEditMenu('edit', row)">编辑</el-button>
            <el-button type="primary" size="small" @click="onOpenMenuAction(row)">功能</el-button>
            <el-popconfirm title="确定删除吗？" @confirm="onTableDelRow(row)">
              <template #reference>
                <el-button size="small" type="danger">删除</el-button>
              </template>
            </el-popconfirm>
          </div>
        </template>
        <template #dialog>
          <MenuDialog ref="menuDialogRef" @refresh="getTableData(state.tableData.param)" :row="listRow" :menuId="menuId"/>
          <ActionDialog ref="menuActionDialogRef" :menuId="menuId" />
        </template>
      </Table>
    </div>
  </div>
</template>

<script setup name="systemMenu">
import {defineAsyncComponent, h, onMounted, reactive, ref} from 'vue';
import {ElMessage} from 'element-plus';
import {menuApi} from '/@/api/menu';
import SvgIcon from '/@/components/svgIcon/index.vue';
import {i18n} from '/@/static/i18n';
import {getDict} from '/@/utils/dict.js';
import {isHideDict, isLinkDict} from '/@/dict/menu';

// 引入组件
const Table = defineAsyncComponent(() => import('/@/components/table/index.vue'));
const TableSearch = defineAsyncComponent(() => import('/@/components/table/component/search.vue'));
const MenuDialog = defineAsyncComponent(() => import('/@/views/system/menu/component/dialog.vue'));
const ActionDialog = defineAsyncComponent(() => import('/@/views/system/menu/component/actionDialog.vue'));

const api = menuApi();
// 定义变量内容
const tableRef = ref();
const menuDialogRef = ref();
const menuActionDialogRef = ref();
const menuId = ref();
const listRow = ref();
const state = reactive({
  tableData: {
    // 列表数据（必传）
    data: [],
    // 表头内容（必传，注意格式）
    header: [
      {key: 'id', colWidth: '', title: 'ID', type: 'text', isCheck: true},
      {key: 'pid', colWidth: '', title: '父级id', type: 'text', isCheck: true},
      {
        key: 'meta.title', colWidth: '', title: '菜单名称', type: 'text', isCheck: true,
        render: (scope) => {
          return i18n.global.t(scope.row?.meta?.title || '');
        }
      },
      {
        key: 'meta.icon', colWidth: '', title: '菜单图标', isCheck: true,
        render: (scope) => {
          /*return h('el-icon', {
            class: scope.row?.meta?.icon
          });*/
          return h(SvgIcon, {
            name: scope.row?.meta?.icon,
          });
        },
      },
      {key: 'path', colWidth: '', title: '路由路径', type: 'text', isCheck: true},
      {key: 'name', colWidth: '', title: '路由名称', type: 'text', isCheck: true},
      {key: 'redirect', colWidth: '', title: '重定向', type: 'text', isCheck: true},
      {
        key: 'isLink', colWidth: '', title: '是否外链', type: 'text', isCheck: true,
        render: (scope) => {
          return getDict(isLinkDict, scope.row?.isLink);
        }
      },
      {key: 'component', colWidth: '', title: '组件路径', type: 'text', isCheck: true},
      {key: 'meta.isLink', colWidth: '', width: '70', height: '40', title: '链接地址', type: 'text', isCheck: true},
      {
        key: 'meta.isHide', colWidth: '', width: '70', height: '40', title: '是否隐藏', type: 'select', isCheck: true,
        render: (scope) => {
          return getDict(isHideDict, scope.row?.meta?.isHide);
        }
      },
      {key: 'menuRoles', colWidth: '', width: '70', height: '40', title: '菜单角色', isCheck: true,
        render: (scope) => {
          return scope.row?.menuRoles?.length > 0 ? scope.row?.menuRoles.map(item => item.name).join(',') : '';
        }
      },
      {key: 'createdAt', colWidth: '', title: '创建时间', type: 'text', isCheck: true},
      {key: 'updatedAt', colWidth: '', title: '更新时间', type: 'text', isCheck: true},
    ],
    // 配置项（必传）
    config: {
      total: 0, // 列表总数
      loading: true, // loading 加载
      isBorder: true, // 是否显示表格边框
      isSerialNo: false, // 是否显示表格序号
      isSelection: true, // 是否显示表格多选
      isOperate: true, // 是否显示表格操作栏
      isPrintTool: true, // 是否显示打印工具
      isExcelTool: true, // 是否显示导出Excel工具
      isRefresh: true, // 是否显示刷新
      fixed: 'right', // 固定操作列
      operationWith: 200, // 固定操作列宽度
      isPage: false, // 是否显示分页
    },
    // 搜索表单，动态生成（传空数组时，将不显示搜索，注意格式）
    search: [
      {label: '路由名称', prop: 'name', placeholder: '请输入路由名称', required: false, type: 'input'},
      {
        label: '是否隐藏',
        prop: 'meta.isHide',
        placeholder: '请选择',
        required: false,
        type: 'select',
        options: [
          {label: '隐藏', value: true},
          {label: '不隐藏', value: false},
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
  listRow.value = row;
  menuDialogRef.value.openDialog(type, row);
};
// 菜单功能
const onOpenMenuAction = (row) => {
  menuId.value = row.id;
  menuActionDialogRef.value.openDialog(row);
}
// 初始化列表数据
const getTableData = async (param) => {
  state.tableData.config.loading = true;
  let response = await api.list(param);
  state.tableData.data = response?.data;
  state.tableData.config.loading = false;
};
// 搜索点击时表单回调
const onSearch = (data) => {
  Object.entries(data).forEach(([key, value]) => {
    if (value !== undefined) {
      state.tableData.param[key] = value;
    } else {
      delete state.tableData.param[key];
    }
  });
  getTableData(state.tableData.param);
};
// 分页改变时回调
const onTablePageChange = (page) => {
  state.tableData.param.page = page.page;
  state.tableData.param.pageSize = page.pageSize;
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