<template>
  <div class="system-menu-dialog-container">
    <el-dialog :title="state.dialog.title" v-model="state.dialog.isShowDialog" width="769px">
      <div class="table-demo-padding layout-padding-view layout-padding-auto">
        <Table
            ref="tableMenuActionRef"
            v-bind="state.tableData"
            class="table-demo"
            @pageChange="onTablePageChange"
        >
          <template #tools>
            <div class="table-tool">
              <el-button size="default" type="primary">
                <el-icon>
                  <ele-FolderAdd/>
                </el-icon>
                新增功能
              </el-button>
            </div>
          </template>
          <template #operation="{row}">
            <div class="flex items-center">
              <el-button type="primary" size="small">编辑</el-button>
              <el-popconfirm title="确定删除吗？">
                <template #reference>
                  <el-button size="small" type="danger">删除</el-button>
                </template>
              </el-popconfirm>
            </div>
          </template>
        </Table>
      </div>
    </el-dialog>
  </div>
</template>
<script setup name="systemMenuActionDialog">
import {defineAsyncComponent, reactive, ref} from 'vue';
import {getDict} from "/@/utils/dict.js";
import {actionIsLinkEnum, actionTypeDict} from '/@/dict/menu/index.js';
import {menuApi} from '/@/api/menu';

// 引入组件
const Table = defineAsyncComponent(() => import('/@/components/table/index.vue'));

const api = menuApi();
const props = defineProps({
  row: {
    type: Object,
    required: true,
    default: () => ({})
  }
});
const tableMenuActionRef = ref();
const state = reactive({
  tableData: {
    // 列表数据（必传）
    data: [],
    // 表头内容（必传，注意格式）
    header: [
      { key: 'id', colWidth: '', title: 'ID', type: 'text', isCheck: true },
      { key: 'menuId', colWidth: '', title: '菜单id', type: 'text', isCheck: true },
      { key: 'type', colWidth: '', title: '类型', type: 'text', isCheck: true,
        render: (scope) => {
          return getDict(actionTypeDict, scope.row?.type);
        }
      },
      { key: 'name', colWidth: '', title: '功能名称', type: 'text', isCheck: true },
      { key: 'isLink', colWidth: '', title: '是否为链接', type: 'text', isCheck: true,
        render: (scope) => {
          return getDict(actionIsLinkEnum, scope.row?.isLink);
        }
      },
      { key: 'sort', colWidth: '', title: '排序', type: 'text', isCheck: true },
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
      isPrintTool: false, // 是否显示打印工具
      isExcelTool: false, // 是否显示导出Excel工具
      isRefresh: false, // 是否显示刷新
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
// 分页改变时回调
const onTablePageChange = (page) => {
  state.tableData.param.page = page.page;
  state.tableData.param.pageSize = page.pageSize;
  getTableData(state.tableData.param);
};
// 打开弹窗
const openDialog = (row) => {
  state.tableData.param.id = row.id;
  getTableData(state.tableData.param);
  state.dialog.isShowDialog = true;
};
// 关闭弹窗
const closeDialog = () => {
  state.dialog.isShowDialog = false;
};
// 初始化列表数据
const getTableData = async (param) => {
  const data = await api.actionList(param);
  state.tableData.data = data.data;
  state.tableData.config.loading = false;
};
// 暴露变量
defineExpose({
  openDialog,
});
</script>