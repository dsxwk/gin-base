<template>
  <TablePlus
      ref="tablePlus"
      :title="funcs.lang('List')"
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
      <el-button type="primary" :icon="CirclePlus" class="mb10">{{ funcs.lang('Create') }}</el-button>
      <el-button
          type="primary"
          plain
          :disabled="!scope.isSelected"
          class="mb10"
          @click="batchPublish(scope.selectedListIds, 2)"
      >
        {{ funcs.lang('Batch Publish') }}
      </el-button>
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
          @click="openDrawerEdit(scope.row)"
      >
        {{ funcs.lang('Edit') }}
      </el-button>
      <el-button type="primary" link :icon="Delete">{{ funcs.lang('Delete') }}</el-button>
    </template>
    <template #toolButton="scope">
      <el-button :icon="Refresh" circle @click="getList" />
      <el-button v-if="columns.length" :icon="Operation" circle />
      <el-button
          :icon="Search"
          circle
      />
    </template>
  </TablePlus>
</template>
<script setup lang="tsx">
import { ref, reactive, h } from 'vue';
import { CirclePlus, Delete, EditPen, Refresh, Operation, Search } from '@element-plus/icons-vue';
import articleModule from '@/app/modules/admin/article';
import createService from '@/utils/service';
import Functions from '@/utils/functions';

const funcs = new Functions();
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
  { type: "index", label: funcs.lang('Index'), width: 80 },
  {
    prop: "title",
    label: funcs.lang('Title')
  },
  {
    prop: "content",
    label: funcs.lang('Content'),
    width: 160,
    search: {
      el: "input",
      props: { maxlength: 30, placeholder: funcs.lang('Please enter the article content') }
    }
  },
  {
    prop: "origin",
    label: funcs.lang('Data Source'),
    width: 160,
    enum: [
      {
        label: funcs.lang('Article Library'),
        value: 1
      },
      {
        label: funcs.lang('Self Built'),
        value: 2
      }
    ],
    search: { el: "tree-select", props: { filterable: true, placeholder: funcs.lang('Please select') } }
  },
  {
    prop: "created_at",
    label: funcs.lang('Create Time'),
    width: 180,
    render: (scope) => {
      return h('div', scope.row.created_at || '-');
    }
  },
  {
    prop: "publish_status",
    label: funcs.lang('Is Publish'),
    width: 160,
    enum: [
      {
        label: funcs.lang('Published'),
        value: 2
      },
      {
        label: funcs.lang('Unpublished'),
        value: 1
      }
    ],
    search: { el: "tree-select", props: { filterable: true, placeholder: funcs.lang('Please select') } },
    render: (scope) => {
      console.log("scope", scope)
      return h(
          'el-switch',
          {
            modelValue: scope.row.publishStatus,
            'active-text': scope.row.publishStatus === 2 ? funcs.lang('Published') : funcs.lang('Unpublished'),
            'active-value': 2,
            'inactive-value': 1,
            onClick: () => doPublish(scope.row)
          }
      );
    }
  },
  { prop: "operation", label: funcs.lang('Operation'), fixed: "right", width: 200 }
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