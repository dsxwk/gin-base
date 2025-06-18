<template>
  <div class="table-search-container" v-if="props.search.length > 0">
    <el-form ref="tableSearchRef" :model="state.form" size="default" label-width="100px" class="table-form">
      <el-row>
        <el-col :xs="24" :sm="12" :md="8" :lg="6" :xl="4" class="mb20" v-for="(val, key) in search" :key="key" v-show="key === 0 || state.isToggle">
          <template v-if="val.type !== ''">
            <el-form-item
                :label="val.label"
                :prop="val.prop"
                :rules="[{ required: val.required, message: `${val.label}不能为空`, trigger: val.type === 'input' ? 'blur' : 'change' }]"
            >
              <el-input v-model="state.form[val.prop]" :placeholder="val.placeholder" clearable v-if="val.type === 'input'" style="width: 100%" />
              <el-date-picker
                  v-model="state.form[val.prop]"
                  type="date"
                  :placeholder="val.placeholder"
                  v-else-if="val.type === 'date'"
                  style="width: 100%"
              />
              <el-date-picker
                  v-model="state.form[val.rangeProp ? val.rangeProp.join('_') : val.prop]"
                  type="daterange"
                  :start-placeholder="'开始日期'"
                  :end-placeholder="'结束日期'"
                  v-if="val.type === 'daterange'"
                  style="width: 100%"
                  value-format="YYYY-MM-DD"
              />
              <el-select v-model="state.form[val.prop]" :placeholder="val.placeholder" v-else-if="val.type === 'select'" style="width: 100%">
                <el-option v-for="item in val.options" :key="item.value" :label="item.label" :value="item.value"> </el-option>
              </el-select>
            </el-form-item>
          </template>
        </el-col>
        <el-col :xs="24" :sm="12" :md="8" :lg="6" :xl="4" class="mb20">
          <el-form-item class="table-form-btn" :label-width="search.length <= 1 ? '10px' : '100px'">
            <template #label v-if="search.length > 1">
              <div class="table-form-btn-toggle ml10" @click="state.isToggle = !state.isToggle">
                <span>{{ state.isToggle ? '收筛选' : '展筛选' }}</span>
                <SvgIcon :name="state.isToggle ? 'ele-ArrowUp' : 'ele-ArrowDown'" />
              </div>
            </template>
            <div>
              <el-button size="default" type="primary" @click="onSearch(tableSearchRef)">查询 </el-button>
              <el-button size="default" type="info" class="ml10" @click="onReset(tableSearchRef)"> 重置 </el-button>
            </div>
          </el-form-item>
        </el-col>
      </el-row>
    </el-form>
  </div>
</template>

<script setup name="tableSearch">
import { reactive, ref, onMounted } from 'vue';

// 定义父组件传过来的值
const props = defineProps({
  // 搜索表单
  search: {
    type: Array,
    default: () => [],
  },
});

// 定义子组件向父组件传值/事件
const emit = defineEmits(['search']);

// 定义变量内容
const tableSearchRef = ref();
const state = reactive({
  form: {},
  isToggle: false,
});

// 查询
const onSearch = (formEl) => {
  if (!formEl) return;
  formEl.validate((valid) => {
    if (valid) {
      // 过滤掉无效的搜索条件
      const filterForm = Object.fromEntries(
          Object.entries(state.form).filter(([key, value]) => {
            return value !== null && value !== undefined && value !== '' && !(Array.isArray(value) && value.length === 0);
          })
      );
      emit('search', filterForm);
    } else {
      return false;
    }
  });
};
// 重置
const onReset = (formEl) => {
  if (!formEl) return;
  formEl.resetFields();
  emit('search', state.form);
};
// 初始化 form 字段，取自父组件 search.prop
const initFormField = () => {
  if (props.search.length <= 0) return false;
  // props.search.forEach((v) => (state.form[v.prop] = ''));
  props.search.forEach((v) => {
    if (v.type === 'daterange' && v.rangeProp) {
      state.form[v.rangeProp.join('_')] = [];
    } else {
      state.form[v.prop] = '';
    }
  });
};
// 页面加载时
onMounted(() => {
  initFormField();
});
</script>

<style scoped lang="scss">
.table-search-container {
  display: flex;
  .table-form {
    flex: 1;
    .table-form-btn-toggle {
      white-space: nowrap;
      user-select: none;
      display: flex;
      align-items: center;
      color: var(--el-color-primary);
    }
  }
}
</style>