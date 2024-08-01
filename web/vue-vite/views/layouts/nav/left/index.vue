<template>
  <div class="tool-bar-lf">
    <el-icon class="collapse-icon" @click="changeCollapse">
      <component :is="isCollapse ? 'expand' : 'fold'"></component>
    </el-icon>
  </div>
</template>
<script setup>
import { ref } from 'vue';

const data = defineProps({
  isCollapse: {
    type: Boolean,
    default: false,
  },
})

const isCollapse = ref(data.isCollapse);

function changeCollapse() {
  if (isCollapse.value) {
    isCollapse.value = false;
  } else {
    isCollapse.value = true;
  }
}
</script>
<style scoped lang="scss">
.tool-bar-lf {
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
  white-space: nowrap;
}
.collapse-icon {
  margin-right: 20px;
  font-size: 22px;
  color: var(--el-header-text-color);
  cursor: pointer;
}
</style>