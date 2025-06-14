<template>
  <div class="system-menu-dialog-container">
    <el-dialog :title="state.dialog.title" v-if="state.dialog.isShowDialog" v-model="state.dialog.isShowDialog" width="769px">
      <div class="table-demo-padding layout-padding-view layout-padding-auto">
        <Table
            ref="tableMenuActionRef"
            v-bind="state.tableData"
            class="table-demo"
            @pageChange="onTablePageChange"
            @delRow="onTableDelRow"
            row-key="id"
            :tree-props="{ children: 'children', hasChildren: 'hasChildren' }"
        >
          <template #tools>
            <div class="table-tool">
              <el-button v-auth="'sys.menu.action.add'" size="default" type="primary" @click="onOpenAdd('add')">
                <el-icon>
                  <ele-FolderAdd/>
                </el-icon>
                新增功能
              </el-button>
            </div>
          </template>
          <template #operation="{row}">
            <div class="flex items-center">
              <el-button v-auth="'sys.menu.action.edit'" type="primary" size="small" @click="onOpenEdit('edit', row)">编辑</el-button>
              <el-popconfirm title="确定删除吗？" @confirm="onTableDelRow(row)">
                <template #reference>
                  <el-button v-auth="'sys.menu.action.del'" size="small" type="danger">删除</el-button>
                </template>
              </el-popconfirm>
            </div>
          </template>
          <template #dialog>
            <ActionOperationDialog ref="dialogRef" @refresh="getTableData(state.tableData.param)" :row="listRow" :menuId="props.menuId"/>
          </template>
        </Table>
      </div>
    </el-dialog>
  </div>
</template>
<script setup name="systemMenuActionDialog">
import {defineAsyncComponent, reactive, ref} from 'vue';
import {getDict} from "/@/utils/dict.js";
import {actionIsLinkEnum, actionTypeDict, isConfirmDict} from '/@/dict/menu/index.js';
import {menuApi} from '/@/api/menu';
import {ElMessage} from 'element-plus';

// 引入组件
const Table = defineAsyncComponent(() => import('/@/components/table/index.vue'));
const ActionOperationDialog = defineAsyncComponent(() => import('/@/views/system/menu/component/ActionOperationDialog.vue'));

const api = menuApi();
const props = defineProps({
  menuId: {
    type: Number,
    required: true,
    default: 0
  }
});
const listRow = ref();
const dialogRef = ref();
const tableMenuActionRef = ref();
const state = reactive({
  tableData: {
    // 列表数据（必传）
    data: [],
    // 表头内容（必传，注意格式）
    header: [
      { key: 'id', colWidth: '', title: 'ID', type: 'text', isCheck: true },
      { key: 'pid', colWidth: '100', title: '父级id', type: 'text', isCheck: true },
      { key: 'label', colWidth: '100', title: '功能名称', type: 'text', isCheck: true },
      { key: 'authValue', colWidth: '120', title: '权限标识', type: 'text', isCheck: true },
      { key: 'parent', colWidth: '100', title: '父级功能', type: 'text', isCheck: true,
        render: (scope) => {
          return scope.row?.parent?.label
        }
      },
      { key: 'menuId', colWidth: '100', title: '菜单id', type: 'text', isCheck: true },
      { key: 'type', colWidth: '100', title: '类型', type: 'text', isCheck: true,
        render: (scope) => {
          return getDict(actionTypeDict, scope.row?.type);
        }
      },
      { key: 'btnType', colWidth: '100', title: '按钮类型', type: 'text', isCheck: true },
      { key: 'btnStyle', colWidth: '100', title: '按钮样式', type: 'text', isCheck: true },
      { key: 'btnSize', colWidth: '100', title: '按钮大小', type: 'text', isCheck: true },
      { key: 'isConfirm', colWidth: '100', title: '是否确认', type: 'text', isCheck: true,
        render: (scope) => {
          return getDict(isConfirmDict, scope.row?.isConfirm);
        }
      },
      { key: 'isLink', colWidth: '120', title: '是否为链接', type: 'text', isCheck: true,
        render: (scope) => {
          return getDict(actionIsLinkEnum, scope.row?.isLink);
        }
      },
      { key: 'sort', colWidth: '', title: '排序', type: 'text', isCheck: true },
      {key: 'actionRoles', colWidth: '100', width: '', height: '40', title: '功能角色', isCheck: true,
        render: (scope) => {
          return scope.row?.actionRoles?.length > 0 ? scope.row?.actionRoles.map(item => item.name).join(',') : '';
        }
      },
      {key: 'createdAt', colWidth: '100', title: '创建时间', type: 'text', isCheck: true},
      {key: 'updatedAt', colWidth: '100', title: '更新时间', type: 'text', isCheck: true},
    ],
    // 配置项（必传）
    config: {
      total: 0, // 列表总数
      loading: true, // loading 加载
      isBorder: true, // 是否显示表格边框
      isSerialNo: false, // 是否显示表格序号
      isSelection: true, // 是否显示表格多选
      isOperate: true, // 是否显示表格操作栏
      isPrintTool: false, // 是否显示打印工具
      isExcelTool: false, // 是否显示导出Excel工具
      isRefresh: true, // 是否显示刷新
      fixed: 'right', // 固定操作列
      operationWith: 200, // 固定操作列宽度
      isPage: false, // 是否显示分页
    },
    // 搜索表单，动态生成（传空数组时，将不显示搜索，注意格式）
    search: [],
    // 搜索参数、搜索时传给后台的值,`getTableData` 中使用）
    param: {},
    // 打印标题
    printName: 'ginBaseAdmin 表格打印演示',
  },
  dialog: {
    isShowDialog: false,
    type: '',
    title: '',
    submitTxt: '',
  },
});
// 打开新增弹窗
const onOpenAdd = (type) => {
  dialogRef.value.openDialog(type);
};
// 分页改变时回调
const onTablePageChange = (page) => {
  state.tableData.param.page = page.page;
  state.tableData.param.pageSize = page.pageSize;
  getTableData(state.tableData.param);
};
// 打开弹窗
const openDialog = (row) => {
  state.tableData.param.menuId = row.id;
  getTableData(state.tableData.param);
  state.dialog.isShowDialog = true;
};
const onOpenEdit = (type, row) => {
  listRow.value = row;
  dialogRef.value.openDialog(type, row);
};
// 删除当前项回调
const onTableDelRow = async (row) => {
  await api.deleteAction({id: row.id, menuId: row.menuId, actionId: row.id});
  ElMessage.success(`删除成功！`);
  state.tableData.data = state.tableData.data.filter((item) => item.id !== row.id);
};
// 初始化列表数据
const getTableData = async (param) => {
  param.menuId = param.menuId ?? props.menuId;
  const data = await api.actionList(param);
  state.tableData.data = data.data;
  state.tableData.config.loading = false;
};
// 暴露变量
defineExpose({
  openDialog,
});
</script>