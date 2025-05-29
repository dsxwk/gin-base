<script setup>
import {ref} from 'vue'

const props = defineProps({
  model: Object,
  formConfig: Array,
  formProps: Object,
  rules: Object,
  gutter: {
    type: Number,
    default: 20
  }
})

const formRef = ref(null)
defineExpose({ formRef })

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
    upload: 'el-upload',
  }

  if (item.type === 'textarea') {
    item.attrs = {
      type: 'textarea',
      ...(item.attrs || {}),
    }
  }

  return map[item.type] || 'el-input'
}
</script>

<template>
  <el-form :model="model" ref="formRef" :rules="rules" v-bind="formProps">
    <el-row :gutter="gutter">
      <template v-for="(item, index) in formConfig" :key="index">
        <el-col :span="item.span || 12" v-if="!item.hidden" class="mb20">
          <el-form-item
              :label="item.label"
              :prop="item.prop"
              :rules="item.rules"
          >
            <!-- 默认组件渲染 -->
            <component
                v-if="!item.slot"
                :is="resolveComponent(item)"
                v-model="model[item.prop]"
                v-bind="item.attrs"
            >
              <!-- 渲染 options -->
              <template v-if="item.options" v-for="opt in item.options" :key="opt.value">
                <el-option
                    v-if="item.type === 'select'"
                    :label="opt.label"
                    :value="opt.value"
                />
                <el-radio
                    v-else-if="item.type === 'radio'"
                    :value="opt.value"
                >
                  {{ opt.label }}
                </el-radio>
                <el-checkbox
                    v-else-if="item.type === 'checkbox'"
                    :value="opt.value"
                >
                  {{ opt.label }}
                </el-checkbox>
              </template>
            </component>

            <!-- 自定义插槽 -->
            <slot v-else :name="item.slot" />
          </el-form-item>
        </el-col>
      </template>
    </el-row>
  </el-form>
</template>
