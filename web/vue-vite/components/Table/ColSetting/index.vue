<template>
  <!-- 列设置 -->
  <el-drawer v-model="drawerVisible" title="列设置" size="450px">
    <div class="table-main">
      <el-table :data="colSetting" :border="true" row-key="prop" default-expand-all :tree-props="{ children: '_children' }">
        <el-table-column prop="label" align="center" label="列名" />
        <el-table-column v-slot="scope" prop="isShow" align="center" label="显示">
          <el-switch v-model="scope.row.isShow"></el-switch>
        </el-table-column>
        <el-table-column v-slot="scope" prop="sortable" align="center" label="排序">
          <el-switch v-model="scope.row.sortable"></el-switch>
        </el-table-column>
        <template #empty>
          <div class="table-empty">
            <img src="@/styles/assets/images/notData.png" alt="notData" />
            <div>暂无可配置列</div>
          </div>
        </template>
      </el-table>
    </div>
  </el-drawer>
</template>

<script setup>
import { ref, watch } from 'vue';
const props = defineProps({
  colSetting: Object,
  isOpen: Boolean
});
const emit = defineEmits(['update:isOpen']);
const drawerVisible = ref(false);
// 监听 isOpen 的变化并更新 drawerVisible
watch(() => props.isOpen, (newVal) => {
  drawerVisible.value = newVal;
});
// 当 drawerVisible 改变时，同步更新 isOpen
watch(drawerVisible, (newVal) => {
  emit('update:isOpen', newVal);
});
</script>

<style scoped lang="scss">
.cursor-move {
  cursor: move;
}
</style>