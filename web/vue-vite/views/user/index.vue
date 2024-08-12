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
import {ref, reactive} from 'vue';
import {CirclePlus, Delete, EditPen} from '@element-plus/icons-vue';
import createService from '@/utils/service';
import request from '@/utils/request';
import userModule from '@/app/modules/admin/user';

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

const dataCallback = (data: any) => {
  data?.entities.forEach((item: any) => {
    item.mechanismValue = [item.category, item.column]
  })
  return {
    list: data?.list,
    total: data.total,
    pageNum: data.page,
    pageSize: data.pageSize
  }
};
const userService = createService(userModule, request);
const getTableList = async (params: paramsReq) => {
  const newParams = JSON.parse(JSON.stringify(params))
  // 请求前参数可以在这里处理
  return await userService.list(newParams)
};
// 批量删除用户信息
const batchDelete = async (articleIds: string[]) => {
  console.log("articleIds", articleIds)
  // await fetchAPI()
  tablePlus.value?.clearSelection()
  tablePlus.value?.getTableList()
};
// 表格配置项
const columns = [
  {type: "selection", fixed: "left", width: 80},
  {type: "index", label: "ID", width: 80},
  {
    prop: "username",
    label: "用户名",
    width: 200,
    props: {maxlength: 30, placeholder: "请输入用户名"}
  },
  {
    prop: "full_name",
    label: "姓名",
    width: 200,
    search: {
      el: "input",
      props: {maxlength: 30, placeholder: "请输入姓名"}
    }
  },
  {
    prop: "nickname",
    label: "昵称",
    width: 200,
    search: {
      el: "input",
      order: 10,
      props: {maxlength: 30, placeholder: "请输入昵称"}
    }
  },
  {
    prop: "email",
    label: "邮箱",
    width: 200,
    isShow: false,
    search: {
      el: "input",
      order: 1,
      props: {placeholder: "请输入邮箱"}
    }
  },
  {
    prop: "gender",
    label: "性别",
    width: 160,
    enum: [
      {
        label: "男",
        value: 1
      },
      {
        label: "女",
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
    prop: "operation",
    label: '操作',
    fixed: 'right',
    width: 200
  }
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