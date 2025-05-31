<script setup>
import {computed, defineAsyncComponent, ref, defineEmits} from 'vue';
import {cloneDeep, get, set} from 'lodash-es';

const IconSelector = defineAsyncComponent(() => import('/@/components/iconSelector/index.vue'));
const props = defineProps({
  model: {
    type: Object,
    required: true
  },
  formConfig: {
    type: Array,
    required: true
  },
  formProps: {
    type: Object,
    default: () => ({})
  },
  rules: {
    type: Object,
    default: () => ({})
  },
  gutter: {
    type: Number,
    default: 20
  },
  dev: {
    type: Boolean,
    default: false
  }
});

const formRef = ref(null);
defineExpose({
  formRef,
  resetFields: () => formRef.value?.resetFields?.(),
  validate: (...args) => formRef.value?.validate?.(...args),
});

const resolveComponent = (item) => {
  const map = {
    input: 'el-input',
    textarea: 'el-input',
    select: 'el-select',
    radio: 'el-radio-group',
    checkbox: 'el-checkbox-group',
    switch: 'el-switch',
    date: 'el-date-picker',
    datetime: 'el-date-picker',
    cascader: 'el-cascader',
    inputNumber: 'el-input-number',
    slider: 'el-slider',
    time: 'el-time-picker',
    upload: 'el-upload',
    icon: IconSelector,
  };

  if (item.type === 'textarea') {
    item.attrs = {type: 'textarea', ...(item.attrs || {})};
  }

  return map[item.type] || 'el-input';
};

// 预生成所有字段的 computed 映射
const computedMap = ref({});
props.formConfig.forEach((item) => {
  if (item.prop) {
    computedMap.value[parseFiled(item.prop)] = computed({
      get() {
        return get(props.model, item.prop);
      },
      set(val) {
        const cloneData = cloneDeep(props.model);
        set(cloneData, decodeField(item.prop), val)
        emit('update:model',cloneData);
      },
    });
  }
});

/**
 * @description 转换filed的值，将.或[]转换成__,因为el-form-item的prop不能有.或[]
 */
function parseFiled(field) {
  return (field || '').replace(/\./g, '__').replace(/\[([0-9]+)\]/, '__$1');
}
/**
 * @description decodeField 解码字段,将__转换成.
 */
function decodeField(field) {
  return (field || '').replace(/__/g, '.');
}
const getOptions = (item) => {
  return typeof item.options === 'function' ? item.options() : item.options || []
};
const emit = defineEmits(['update:model'])
</script>

<template>
  <el-form
      :model="computedMap"
      ref="formRef"
      :rules="rules"
      v-bind="formProps"
  >
    <el-row
        :gutter="gutter"
    >
      <template
          v-for="(item, index) in formConfig"
          :key="index"
      >
        <el-col
            v-if="!item.hidden"
            :span="item.span || 12"
            class="mb20"
        >
          <el-form-item
              :label="item.label"
              :prop="item.prop"
              :rules="item.rules"
              :label-width="item.labelWidth || undefined"
          >
            <component
                v-if="!item.slot"
                :is="resolveComponent(item)"
                v-model="computedMap[parseFiled(item.prop)]"
                :props="item.props"
                :options="item.type === 'cascader' ? getOptions(item) : undefined"
                v-bind="item.attrs"
                v-on="item.events || {}"
            >
              <!-- 渲染 options -->
              <template v-if="item.options">
                <template
                    v-for="opt in getOptions(item)"
                    :key="opt.value"
                >
                  <el-option
                      v-if="item.type === 'select'"
                      :label="opt.label"
                      :value="opt.value"
                  />
                  <el-radio
                      v-else-if="item.type === 'radio'"
                      :value="opt.value"
                      :label="opt.label"
                  />
                  <el-checkbox
                      v-else-if="item.type === 'checkbox'"
                      :value="opt.value"
                      :label="opt.label"
                  />
                </template>
              </template>
              <!-- cascader 默认插槽 -->
              <template
                  v-if="item.type === 'cascader' && item.slotDefault"
                  #default="{ node, data }"
              >
                <component
                    :is="item.slotDefault"
                    :node="node"
                    :data="data"
                />
              </template>
            </component>
            <slot
                v-else
                :name="item.slot"
            />
          </el-form-item>
        </el-col>
      </template>
    </el-row>
  </el-form>
  <template v-if="dev">
    <el-divider>调试信息</el-divider>
    <pre style="background: #f5f5f5; padding: 12px; font-size: 12px; overflow: auto;">
    <strong>model:</strong>
    {{ JSON.stringify(model, null, 2) }}
  </pre>
    <pre style="background: #f5f5f5; padding: 12px; font-size: 12px; overflow: auto;">
    <strong>formConfig:</strong>
    {{ JSON.stringify(formConfig, null, 2) }}
  </pre>
  </template>
</template>
