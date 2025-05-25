<template>
  <div class="system-menu-dialog-container">
    <el-dialog :title="state.dialog.title" v-model="state.dialog.isShowDialog" width="769px">
      actionDialog
    </el-dialog>
  </div>
</template>
<script setup name="systemMenuActionDialog">
import { reactive } from 'vue';

const state = reactive({
  dialog: {
    isShowDialog: false,
    type: '',
    title: '',
    submitTxt: '',
  },
});
// 打开弹窗
const openDialog = (row) => {
  state.dialog.isShowDialog = true;
};
// 关闭弹窗
const closeDialog = () => {
  state.dialog.isShowDialog = false;
};
// 暴露变量
defineExpose({
  openDialog,
});
</script>