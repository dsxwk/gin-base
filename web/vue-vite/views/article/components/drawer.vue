<template>
  <el-drawer v-model="drawerProps.visible" :direction="direction" :title="drawerProps.title+' '+funcs.lang('Article')" :before-close="handleClose">
    <template #default>
      <el-form
          ref="ruleFormRef"
          label-width="100px"
          label-suffix=" :"
          :rules="rules"
          :model="drawerProps.row"
      >
        <el-form-item :label="funcs.lang('Title')" prop="title">
          <el-input v-model="drawerProps.row.title" :placeholder="funcs.lang('Please enter the title')" clearable></el-input>
        </el-form-item>
        <el-form-item :label="funcs.lang('Content')" prop="content">
          <el-input type="textarea" v-model="drawerProps.row.content" :placeholder="funcs.lang('Please enter the content')" clearable></el-input>
        </el-form-item>
        <el-form-item :label="funcs.lang('Data Source')" prop="data_source">
          <el-select  v-model="drawerProps.row.data_source" :placeholder="funcs.lang('Please select data source')" clearable>
            <el-option
                v-for="item in dataSourceDict"
                :key="item.value"
                :label="item.label"
                :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item :label="funcs.lang('Is Publish')" prop="is_publish">
          <el-switch
              v-model="drawerProps.row.is_publish"
              @click="swithIsPublish"
              class="mb-2"
              :active-text="funcs.lang('Published')"
              :inactive-text="funcs.lang('Unpublished')"
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
import {dataSourceDict} from '@/app/modules/admin/article/dict';
import Functions from '@/utils/functions';
import articleModule from '@/app/modules/admin/article';
import createService from '@/utils/service';

const articleService = createService(articleModule);
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

const emit = defineEmits(['updateIsPublish', 'dataChange']);
const swithIsPublish = () => {
  emit('updateIsPublish', props.drawerProps.row.is_publish);
}
const direction = ref('rtl');
const rules = reactive({
  title: [{ required: true, message: funcs.lang('Please enter the title') }],
  content: [{ required: true, message: funcs.lang('Please enter the content') }],
  data_source: [{ required: true, message: funcs.lang('Please select data source') }],
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
    await articleService.update(props.drawerProps.row);
    console.log('update', props.drawerProps.row.id);
  } else {
    await articleService.create(props.drawerProps.row);
    console.log('create');
  }
  emit('dataChange');
  props.drawerProps.visible = false;
}
</script>