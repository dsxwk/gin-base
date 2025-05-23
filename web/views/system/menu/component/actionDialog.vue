<template></template>
<script setup name="systemMenuActionDialog">
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