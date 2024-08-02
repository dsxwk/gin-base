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
<style lang="scss">
.el-sub-menu .el-sub-menu__title:hover {
  color: var(--el-menu-hover-text-color) !important;
  background-color: transparent !important;
}
.el-menu--collapse {
  .is-active {
    .el-sub-menu__title {
      color: #ffffff !important;
      background-color: var(--el-color-primary) !important;
    }
  }
}
.el-menu-item {
  &:hover {
    color: var(--el-menu-hover-text-color);
  }
  &.is-active {
    /*color: var(--el-menu-active-color) !important;*/
    color: #009688 !important;
    /*background-color: var(--el-menu-active-bg-color) !important;*/
    background-color: #e6f5f3 !important;
    &::before {
      position: absolute;
      top: 0;
      bottom: 0;
      width: 4px;
      content: "";
      /*background-color: var(--el-color-primary);*/
      background-color: #009688;
      /* 新增 */
      left: 0;
    }
  }
}
.vertical,
.classic,
.transverse {
  .el-menu-item {
    &.is-active {
      &::before {
        left: 0;
      }
    }
  }
}
.columns {
  .el-menu-item {
    &.is-active {
      &::before {
        right: 0;
      }
    }
  }
}
</style>