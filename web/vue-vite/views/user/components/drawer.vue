<template>
  <el-drawer v-model="drawerProps.visible" :direction="direction" :title="drawerProps.title+' '+funcs.lang('User')" :before-close="handleClose">
    <template #default>
      <el-form
          ref="ruleFormRef"
          label-width="100px"
          label-suffix=" :"
          :rules="rules"
          :model="drawerProps.row"
      >
        <el-form-item :label="funcs.lang('Username')" prop="username">
          <el-input v-model="drawerProps.row.username" :placeholder="funcs.lang('Please enter the username')" clearable></el-input>
        </el-form-item>
        <el-form-item :label="funcs.lang('FullName')" prop="full_name">
          <el-input v-model="drawerProps.row.full_name" :placeholder="funcs.lang('Please enter the full name')" clearable></el-input>
        </el-form-item>
        <el-form-item :label="funcs.lang('Gender')" prop="gender">
          <el-select  v-model="drawerProps.row.gender" :placeholder="funcs.lang('Please select the gender')" clearable>
            <el-option
                v-for="item in genderEnum"
                :key="item.value"
                :label="item.label"
                :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item :label="funcs.lang('User Status')" prop="status">
          <el-switch
              v-model="drawerProps.row.status"
              @click="swithUserStatus"
              class="mb-2"
              :active-text="funcs.lang('Enabled')"
              :inactive-text="funcs.lang('Disabled')"
              :active-value="1"
              :inactive-value="2"
          />
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
import userModule from '@/app/modules/admin/user';
import createService from '@/utils/service';
import {genderEnum} from '@/enums/user';

const userService = createService(userModule);
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

const emit = defineEmits(['updateUserStatus', 'dataChange']);
const swithUserStatus = () => {
  emit('updateUserStatus', props.drawerProps.row.is_publish);
}
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