<template>
  <el-drawer v-model="drawerProps.visible" :direction="direction" title="添加文章" :before-close="handleClose">
    <template #default>
      <el-form
          ref="ruleFormRef"
          label-width="100px"
          label-suffix=" :"
          :rules="rules"
          :model="drawerProps.row"
      >
        <el-form-item label="标题" prop="title">
          <el-input v-model="drawerProps.row.title" placeholder="请输入标题" clearable></el-input>
        </el-form-item>
        <el-form-item label="内容" prop="content">
          <el-input type="textarea" v-model="drawerProps.row.content" placeholder="请输入内容" clearable></el-input>
        </el-form-item>
      </el-form>
    </template>
    <template #footer>
      <div style="flex: auto">
        <el-button @click="cancelClick">cancel</el-button>
        <el-button type="primary" @click="confirmClick">confirm</el-button>
      </div>
    </template>
  </el-drawer>
</template>

<script setup>
import { ref, reactive } from 'vue';
import { ElMessageBox } from 'element-plus';

defineProps({
  drawerProps: {
    type: Object,
    default: () => {
      return {
        row: {
          id: '',
          title: '',
          content: '',
        },
        title: '',
      };
    },
  },
});

const direction = ref('rtl');
const rules = reactive({
  title: [{ required: true, message: "请输入标题" }],
  content: [{ required: true, message: "请输入内容" }],
});
const handleClose = (done) => {
  ElMessageBox.confirm('Are you sure you want to close this?')
      .then(() => {
        done();
      })
      .catch(() => {
        // catch error
      });
};
function cancelClick() {
  drawer.value = false;
}
function confirmClick() {
  ElMessageBox.confirm(`Are you confirm to chose this?`)
      .then(() => {
        drawer.value = false;
      })
      .catch(() => {
        // catch error
      });
}
</script>