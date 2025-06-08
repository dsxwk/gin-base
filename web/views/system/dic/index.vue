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
            <el-button size="default" type="primary" @click="onOpenAddDict('add')">
              <el-icon>
                <ele-FolderAdd/>
              </el-icon>
              新增字典
            </el-button>
          </div>
        </template>
        <template #operation="{row}">
          <div class="flex items-center">
            <el-button type="primary" size="small" @click="onOpenEditDict('edit', row)">编辑</el-button>
            <el-popconfirm title="确定删除吗？" @confirm="onTableDelRow(row)">
              <template #reference>
                <el-button size="small" type="danger">删除</el-button>
              </template>
            </el-popconfirm>
          </div>
        </template>
        <template #dialog>
          <DictDialog ref="dictDialogRef" @refresh="getTableData(state.tableData.param)" :row="listRow" :dictId="dictId"/>
        </template>
      </Table>
    </div>
  </div>
</template>

<script setup name="systemDict">
import {defineAsyncComponent, h, onMounted, reactive, ref} from 'vue';
import {ElMessage} from 'element-plus';
import {dictApi} from '/@/api/dict';

// 引入组件
const Table = defineAsyncComponent(() => import('/@/components/table/index.vue'));
const TableSearch = defineAsyncComponent(() => import('/@/components/table/component/search.vue'));
const DictDialog = defineAsyncComponent(() => import('/@/views/system/dic/component/dialog.vue'));

const api = dictApi();
// 定义变量内容
const tableRef = ref();
const dictDialogRef = ref();
const dictId = ref();
const listRow = ref();
const state = reactive({
  tableData: {
    // 列表数据（必传）
    data: [],
    // 表头内容（必传，注意格式）
    header: [
      {key: 'id', colWidth: '', title: 'ID', type: 'text', isCheck: true},
      {key: 'pid', colWidth: '', title: '父级id', type: 'text', isCheck: true},
      {key: 'name', colWidth: '140', title: '字典名称(英文)', type: 'text', isCheck: true},
      {key: 'title', colWidth: '140', title: '字典名称(中文)', type: 'text', isCheck: true},
      {key: 'label', colWidth: '', title: '映射值', type: 'text', isCheck: true},
      {key: 'extend', colWidth: '100', title: '扩展字段', type: 'text', isCheck: true},
      {key: 'sort', colWidth: '', title: '排序', type: 'text', isCheck: true},
      {key: 'desc', colWidth: '', title: '描述', type: 'text', isCheck: true},
      {key: 'status', colWidth: '', title: '状态', type: 'text', isCheck: true},
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
      isPrintTool: true, // 是否显示打印工具
      isExcelTool: true, // 是否显示导出Excel工具
      isRefresh: true, // 是否显示刷新
      fixed: 'right', // 固定操作列
      operationWith: 200, // 固定操作列宽度
      isPage: false, // 是否显示分页
    },
    // 搜索表单，动态生成（传空数组时，将不显示搜索，注意格式）
    search: [
      {label: '名称(英文)', prop: 'name', placeholder: '请输入字典名称(英文)', required: false, type: 'input'},
      {
        label: '状态',
        prop: 'status',
        placeholder: '请选择',
        required: false,
        type: 'select',
        options: [
          {label: '启用', value: 1},
          {label: '禁用', value: 2},
        ],
      },
    ],
    // 搜索参数、搜索时传给后台的值,`getTableData` 中使用）
    param: {},
    // 打印标题
    printName: 'ginBaseAdmin 表格打印演示',
  },
});
// 打开新增弹窗
const onOpenAddDict = (type) => {
  dictDialogRef.value.openDialog(type);
};
// 打开编辑弹窗
const onOpenEditDict = (type, row) => {
  listRow.value = row;
  dictDialogRef.value.openDialog(type, row);
};
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
const onTableDelRow = async (row) => {
  await api.delete({id: row.id});
  ElMessage.success(`删除成功！`);
  state.tableData.data = state.tableData.data.filter((item) => item.id !== row.id);
  await getTableData(state.tableData.param);
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