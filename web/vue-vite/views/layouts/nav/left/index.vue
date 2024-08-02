<template>
  <div class="tool-bar-lf">
    <el-icon class="collapse-icon" @click="changeCollapse">
      <component :is="isCollapse ? 'expand' : 'fold'"></component>
    </el-icon>
  </div>
</template>
<script setup>
defineProps({
  isCollapse: Boolean
});

const emit = defineEmits(['toggle-collapse', 'show-span']);

function changeCollapse() {
  emit('toggle-collapse');
  emit('show-span');
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