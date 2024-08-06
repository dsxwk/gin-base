<template>
  <template v-for="items in menuList" :key="items.path">
    <el-sub-menu v-if="items.children?.length" :index="items.path">
      <template #title>
        <el-icon v-if="items.meta.icon">
          <component :is="items.meta.icon"></component>
        </el-icon>
        <span class="sle">{{ items.meta.title }}</span>
      </template>
      <!-- 递归组件名需与文件名一致 -->
      <Menus :menuList="items.children" />
    </el-sub-menu>
    <el-menu-item v-else :index="items.path" @click="handleClickMenu(items)">
      <el-icon v-if="items.meta?.icon">
        <component :is="items.meta.icon"></component>
      </el-icon>
      <template #title>
        <span class="sle">{{ items.meta?.title }}</span>
      </template>
    </el-menu-item>
  </template>
</template>
<script setup>
import { useRouter } from "vue-router";

defineProps({
  menuList: {
    type: Array,
    default: () => [],
  },
});

const router = useRouter();
const handleClickMenu = (items) => {
  if (items.meta.isLink) {
    return window.open(items.meta.isLink, "_blank");
  } else {
    router.push(items.path);
  }
};
</script>
<style lang="scss" scoped>
@import "./index.scss";
</style>