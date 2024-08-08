<template>
  <TablePlus
      ref="tablePlus"
      title="列表"
      :columns="columns"
      :request-api="getTableList"
      :init-param="initParam"
      :pagination="true"
      :data-callback="dataCallback"
      :reset-callback="resetCallback"
      row-key="articleId"
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
import Layout from '@/views/layouts/index.vue';
import {onMounted, ref, reactive} from 'vue';
import {CirclePlus, Delete, EditPen} from '@element-plus/icons-vue';
import createService from '@/utils/service';
import request from '@/utils/request';
import userModule from '@/app/modules/admin/user';

const userService = createService(userModule, request);
onMounted(async () => {
  await userService.list({page: 1, pageSize: 10});
});

const rowItem = {
  articleId: 0,
  title: '文章标题',
  origin: 1,
  publishTime: '2024-09-01 12:00:00',
  publishStatus: 2
};

type rowItemType = typeof rowItem;

interface paramsReq {
  pageSize: number
  pageNum: number
}

const tablePlus = ref();
const initParam = reactive({
  type: 'yyy',
  pageSize: 10,
  hotWordList: []
});
// 模拟获取列表接口
const articleList = (params: paramsReq) => {
  return new Promise((resolve) => {
    const entities: rowItemType[] = []
    for (let i = 0; i < 100; i++) {
      entities.push({
        articleId: i,
        title: '文章标题' + i,
        origin: 1,
        publishTime: '2024-09-01 12:00:00',
        publishStatus: 2
      })
    }
    const ressult = {
      total: 100,
      pageNo: params.pageNum,
      pageSize: params.pageSize,
      entities
    }
    setTimeout(() => {
      resolve(ressult)
    }, 1500)
  })
};
const dataCallback = (data: any) => {
  data?.entities.forEach((item: any) => {
    item.mechanismValue = [item.category, item.column]
  })
  return {
    list: data?.entities,
    total: data.total,
    pageNum: data.pageNo,
    pageSize: data.pageSize
  }
};
const getTableList = (params: paramsReq) => {
  const newParams = JSON.parse(JSON.stringify(params))
  // 请求前参数可以在这里处理
  return articleList(newParams)
};
// 批量删除用户信息
const batchDelete = async (articleIds: string[]) => {
  console.log("articleIds", articleIds)
  // await fetchAPI()
  tablePlus.value?.clearSelection()
  tablePlus.value?.getTableList()
};
const batchPublish = async (articleIds: string[], status: number) => {
  console.log("articleIds: string[], status: number", articleIds, status)
  // await fetchAPI()
  tablePlus.value?.clearSelection()
  tablePlus.value?.getTableList()
};
// 表格配置项
const columns = [
  {type: "selection", fixed: "left", width: 80},
  {type: "index", label: "序号", width: 80},
  {
    prop: "title",
    label: "文章标题0"
  },
  {
    prop: "title",
    label: "文章标题",
    search: {
      el: "input",
      props: {maxlength: 30, placeholder: "请输入文章标题"}
    }
  },
  {
    prop: "title",
    label: "文章标题2",
    search: {
      el: "input",
      order: 10,
      props: {maxlength: 30, placeholder: "请输入文章标题2"}
    }
  },
  {
    prop: "title",
    label: "文章标题3",
    isShow: false,
    search: {
      el: "input",
      order: 1,
      props: {placeholder: "请输入文章标题3"}
    }
  },
  {
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
    search: {el: "tree-select", props: {filterable: true}}
  },

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
    search: {el: "tree-select", props: {filterable: true}},
    render: (scope) => {
      return (
          <>
            <el-switch
                model-value={scope.row.publishStatus}
                active-text={scope.row.publishStatus === 2 ? '已发布' : '未发布'}
                active-value={2}
                inactive-value={1}
                onClick={doPublish}
            />
          </>
      )
    }
  },
  {prop: "operation", label: '操作', fixed: 'right', width: 200}
];
const resetCallback = () => {
  console.log("resetCallBack")
  initParam.hotWordList = []
  tablePlus.value?.getTableList()
};
// 上架
const doPublish = async (params: any) => {
  console.log("doPublish", params)
  // await fetchAPI()
  tablePlus.value?.getTableList()
};
const openDrawerEdit = async (row: Partial<any>) => {
  console.log(row)
};
</script>