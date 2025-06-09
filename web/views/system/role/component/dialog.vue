<template>
  <div class="system-menu-dialog-container">
    <el-dialog :title="state.dialog.title" v-if="state.dialog.isShowDialog" v-model="state.dialog.isShowDialog" width="769px">
      <ConfigForm
          ref="dialogFormRef"
          v-model:model="state.ruleForm"
          :form-config="getFormData()"
          :rules="rules"
          :form-props="{
            labelWidth: '80px',
            size: 'default'
          }"
      >
      </ConfigForm>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="onCancel" size="default">取 消</el-button>
          <el-button type="primary" @click="onSubmit" size="default">{{ state.dialog.submitTxt }}</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>
<script setup name="systemRoleDialog">
import {computed, nextTick, onMounted, reactive, ref} from 'vue';
import {roleApi} from '/@/api/role';
import {ElMessage} from 'element-plus';
import ConfigForm from '/@/components/form/index.vue';
import {statusDict} from '/@/dict/role';

const props = defineProps({
  row: {
    type: Object,
    required: true,
    default: () => ({})
  }
});
const emit = defineEmits(['refresh']);
const dialogFormRef = ref();
const api = roleApi();

const state = reactive({
  ruleForm: {
    name: '',
    desc: '',
    status: 1,
  },
  dialog: {
    isShowDialog: false,
    type: 'add',
    title: '',
    submitTxt: ''
  }
});

const getFormData = computed(() => {
  return [
    {
      label: '角色名称',
      prop: 'name',
      type: 'input',
      attrs: {
        placeholder: '请输入角色名称',
        clearable: true
      },
      rules: [
        {
          required: true,
          message: '请输入角色名称',
          trigger: 'blur'
        }
      ]
    },
    {
      label: '描述',
      prop: 'desc',
      type: 'input',
      attrs: {
        placeholder: '请输入描述',
        clearable: true
      },
      rules: []
    },
    {
      label: '状态',
      prop: 'status',
      type: 'radio',
      options: statusDict
    },
  ];
});
const rules = {};

const openDialog = async (type, row) => {
  state.ruleForm = {
    name: '',
    desc: '',
    status: 1,
  };
  if (type === 'edit') {
    const data = await detail(row.id);
    Object.keys(state.ruleForm).forEach((key) => {
      if (data.hasOwnProperty(key)) {
        state.ruleForm[key] = data[key];
      }
    });
    state.dialog.title = '修改角色';
    state.dialog.submitTxt = '修 改';
  } else {
    state.dialog.title = '新增角色';
    state.dialog.submitTxt = '新 增';
  }
  state.dialog.type = type;
  state.dialog.isShowDialog = true;
  // 清空表单，此项需加表单验证才能使用
  await nextTick(() => {
    dialogFormRef.value && dialogFormRef.value.resetFields();
  });
};

const closeDialog = () => {
  state.dialog.isShowDialog = false;
  state.selectedRoleIds = [];
};

const onCancel = () => {
  closeDialog();
};

const onSubmit = async () => {
  dialogFormRef.value.validate(async (valid) => {
    if (!valid) return;

    let msg = '';
    if (state.dialog.type === 'add') {
      await api.create(state.ruleForm);
      msg = '创建成功';
    } else {
      state.ruleForm.id = props.row.id;
      await api.update(state.ruleForm);
      msg = '更新成功';
    }
    ElMessage.success(msg);
    closeDialog();
    emit('refresh');
  });
};
// 详情
const detail = async (id) => {
  const res = await api.detail({id: id});

  return res.data;
};
onMounted(() => {});

defineExpose({openDialog});
</script>

<style scoped>
.dialog-footer {
  display: flex;
  justify-content: flex-end;
}
</style>
