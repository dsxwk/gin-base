<template>
  <el-drawer v-model="drawerProps.visible" :direction="direction" :title="drawerProps.title+' '+funcs.lang('Menu')" :before-close="handleClose">
    <template #default>
      <el-form
          ref="ruleFormRef"
          label-width="100px"
          label-suffix=" :"
          :rules="rules"
          :model="drawerProps.row"
      >
        <el-form-item :label="funcs.lang('Menu Name')" prop="meta.title">
          <el-input v-model="drawerProps.row.meta.title" :placeholder="funcs.lang('Please enter the Menu Name')" clearable></el-input>
        </el-form-item>
        <el-form-item :label="funcs.lang('Menu Alias')" prop="name">
          <el-input v-model="drawerProps.row.name" :placeholder="funcs.lang('Please enter the Menu Alias')" clearable></el-input>
        </el-form-item>
        <el-form-item :label="funcs.lang('Icon')" prop="meta.icon">
          <el-input v-model="drawerProps.row.name" :placeholder="funcs.lang('Please enter the Icon')" clearable></el-input>
        </el-form-item>
        <el-form-item :label="funcs.lang('Route')" prop="path">
          <el-input v-model="drawerProps.row.path" :placeholder="funcs.lang('Please enter the Route')" clearable></el-input>
        </el-form-item>
        <el-form-item :label="funcs.lang('Component Path')" prop="component">
          <el-input v-model="drawerProps.row.component" :placeholder="funcs.lang('Please enter the Component Path')" clearable></el-input>
        </el-form-item>
      </el-form>
    </template>
    <template #footer>
      <div style="flex: auto">
        <el-button @click="cancel">cancel</el-button>
        <el-button type="primary" @click="confirm">confirm</el-button>
      </div>
    </template>
  </el-drawer>
</template>

<script setup>
import { ref, reactive } from 'vue';
import { ElMessageBox } from 'element-plus';
import Functions from '@/utils/functions';
import menuModule from '@/app/modules/admin/menu';
import createService from '@/utils/service';

const menuService = createService(menuModule);
const funcs = new Functions();
const props = defineProps({
  drawerProps: {
    type: Object,
    default: () => {
      return {
        row: {
          id: 0,
          title: '',
          content: '',
        },
        title: '',
        visible: false,
      };
    },
  },
});

const emit = defineEmits(['dataChange']);
const direction = ref('rtl');
const rules = reactive({
  username: [{ required: true, message: funcs.lang('Please enter the username') }],
  full_name: [{ required: true, message: funcs.lang('Please enter the full name') }],
  gender: [{ required: true, message: funcs.lang('Please select') }],
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
function cancel() {
  props.drawerProps.visible = false;
}
async function confirm() {
  // 创建和更新
  if (props.drawerProps.row.id) {
    console.log('update', props.drawerProps.row);
    await userService.update(props.drawerProps.row);
  } else {
    console.log('create', props.drawerProps.row);
    await userService.create(props.drawerProps.row);
  }
  emit('dataChange');
  props.drawerProps.visible = false;
}
</script>