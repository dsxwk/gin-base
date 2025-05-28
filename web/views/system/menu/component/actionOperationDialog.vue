<template>
  <div class="system-menu-dialog-container">
    <el-dialog :title="state.dialog.title" v-model="state.dialog.isShowDialog" width="769px">
      <el-form ref="dialogFormRef" :model="state.ruleForm" size="default" label-width="80px">
        systemMenuActionOperationDialog
      </el-form>
    </el-dialog>
  </div>
</template>
<script setup  name="systemMenuActionOperationDialog">
import {onMounted, ref, reactive, nextTick} from 'vue';
import {ElMessage} from 'element-plus';
import {deepClone} from '/@/utils/other.js';

const state = reactive({
  ruleForm: {},
  dialog: {
    isShowDialog: false,
    type: '',
    title: '',
    submitTxt: '',
  },
});
const defaultForm = {};
// 定义子组件向父组件传值/事件
const emit = defineEmits(['refresh']);

// 定义变量内容
const dialogFormRef = ref();
// 打开弹窗
const openDialog = (type, row) => {
  if (type === 'edit') {
    Object.keys(state.ruleForm).forEach(key => {
      if (row.hasOwnProperty(key)) {
        if (typeof state.ruleForm[key] === 'object' && state.ruleForm[key] !== null) {
          Object.assign(state.ruleForm[key], row[key]);
        } else {
          state.ruleForm[key] = row[key];
        }
      }
    });
    state.dialog.title = '修改功能';
    state.dialog.submitTxt = '修 改';
  } else {
    state.dialog.title = '新增功能';
    state.dialog.submitTxt = '新 增';
  }
  state.dialog.type = type;
  state.dialog.isShowDialog = true;
  // 清空表单，此项需加表单验证才能使用
  nextTick(() => {
    dialogFormRef.value && dialogFormRef.value.resetFields();
  });
};
// 关闭弹窗
const closeDialog = () => {
  state.dialog.isShowDialog = false;
  state.ruleForm = deepClone(defaultForm);
  dialogFormRef.value && dialogFormRef.value.resetFields();
};
// 取消
const onCancel = () => {
  closeDialog();
};
// 提交
const onSubmit = async () => {
  let msg = '';
  if (state.dialog.type === 'add') {
    msg = '创建成功';
  } else {
    msg = '更新成功';
  }
  ElMessage.success(msg);
  closeDialog();
  emit('refresh');
};
// 页面加载时
onMounted(() => {
  
});
// 暴露变量
defineExpose({
  openDialog,
});
</script>