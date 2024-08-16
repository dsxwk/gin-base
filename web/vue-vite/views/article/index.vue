<template>
  <TablePlus
      ref="tablePlus"
      title="列表"
      :columns="columns"
      :request-api="getList"
      :init-param="initParam"
      :pagination="true"
      :data-callback="dataCallback"
      :reset-callback="resetCallback"
      row-key="id"
  >
    <!-- 表格 header 按钮 -->
    <template #tableHeader="scope">
      <el-button type="primary" :icon="CirclePlus" class="mb10">新建</el-button>
      <el-button
          type="primary"
          plain
          :disabled="!scope.isSelected"
          class="mb10"
          @click="batchPublish(scope.selectedListIds, 2)"
      >
        批量发布
      </el-button>
      <el-button
          type="danger"
          :icon="Delete"
          plain
          :disabled="!scope.isSelected"
          class="mb10"
          @click="batchDelete(scope.selectedListIds)"
      >
        批量删除
      </el-button>
    </template>
    <!-- 表格操作 -->
    <template #operation="scope">
      <el-button
          type="primary"
          link
          :icon="EditPen"
          @click="openDrawerEdit(scope.row)"
      >
        编辑
      </el-button>
      <el-button type="primary" link :icon="Delete">删除</el-button>
    </template>
  </TablePlus>
</template>
<script setup lang="tsx">
import { ref, reactive } from 'vue';
import { CirclePlus, Delete, EditPen } from '@element-plus/icons-vue';
import articleModule from '@/app/modules/admin/article';
import createService from '@/utils/service';

const data = ref([]);
const articleService = createService(articleModule);

const tablePlus = ref()
const initParam = reactive({
  pageNum: 1,
  pageSize: 10
});
const getList = async () => {
  const result = await articleService.list({page: 1, pageSize: 10});
  return {
    total: result.data.total,
    pageNo: result.data.page,
    pageSize: result.data.page_size,
    list: result.data.list
  };
}
const dataCallback = (data: any) => {
  console.log("data", data);
  return {
    list: data?.list,
    total: data?.total,
    pageNum: data?.page,
    pageSize: data?.pageSize
  }
}
// 批量删除用户信息
const batchDelete = async (articleIds: string[]) => {
  console.log("articleIds", articleIds)
  // await fetchAPI()
  tablePlus.value?.clearSelection()
  tablePlus.value?.getTableList()
}
const batchPublish = async (articleIds: string[], status: number) => {
  console.log("articleIds: string[], status: number", articleIds, status)
  // await fetchAPI()
  tablePlus.value?.clearSelection()
  tablePlus.value?.getTableList()
}
// 表格配置项
const columns = [
  { type: "selection", fixed: "left", width: 80 },
  { type: "index", label: "序号", width: 80 },
  {
    prop: "title",
    label: "标题"
  },
  {
    prop: "content",
    label: "内容",
    search: {
      el: "input",
      props: { maxlength: 30, placeholder: "请输入文章内容" }
    }
  },
  /*{
    prop: "origin",
    label: "数据来源",
    width: 160,
    enum: [
      {
        label: "文章库",
        value: 1
      },
      {
        label: "自建新增 ",
        value: 2
      }
    ],
    search: { el: "tree-select", props: { filterable: true } }
  },*/
  {
    prop: "publishTime",
    label: "发布时间",
    width: 180,
    render: (scope) => {
      return <div>{scope.row.publishTime || "- -"}</div>
    }
  },

  {
    prop: "publishStatus",
    label: "是否发布",
    width: 160,
    enum: [
      {
        label: "已发布",
        value: 2
      },
      {
        label: "未发布",
        value: 1
      }
    ],
    search: { el: "tree-select", props: { filterable: true } },
    render: (scope) => {
      return (
          <>
            <el-switch
                model-value={scope.row.publishStatus}
                active-text={scope.row.publishStatus === 2 ? "已发布" : "未发布"}
                active-value={2}
                inactive-value={1}
                onClick={doPublish}
            />
          </>
      )
    }
  },
  { prop: "operation", label: "操作", fixed: "right", width: 200 }
]
const resetCallback = () => {
  console.log("resetCallBack")
  tablePlus.value?.getTableList()
}
// 上架
const doPublish = async (params: any) => {
  console.log("doPublish", params)
  // await fetchAPI()
  tablePlus.value?.getTableList()
}
const openDrawerEdit = async (row: Partial<any>) => {
  console.log(row)
}
</script>