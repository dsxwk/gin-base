<template>
  <div class="tabs-box">
    <div class="tabs-menu">
      <el-tabs v-model="tabsMenuValue" type="card" @tab-click="tabClick" @tab-remove="tabRemove">
        <el-tab-pane v-for="item in tabsMenuList" :key="item.path" :label="item.title" :name="item.path" :closable="item.close && item.path !== HOME_URL">
          <template #label>
            <el-icon v-if="item.icon && tabsIcon" class="tabs-icon">
              <component :is="item.icon"></component>
            </el-icon>
            {{ item.title }}
          </template>
        </el-tab-pane>
      </el-tabs>
      <el-dropdown trigger="click" :teleported="false">
        <div class="more-button">
          <i :class="'iconfont icon-xiala'"></i>
        </div>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item @click="refresh">
              <el-icon><Refresh /></el-icon>刷新
            </el-dropdown-item>
            <el-dropdown-item @click="maximize">
              <el-icon><FullScreen /></el-icon>最大化
            </el-dropdown-item>
            <el-dropdown-item divided @click="closeCurrentTab">
              <el-icon><Remove /></el-icon>关闭当前
            </el-dropdown-item>
            <el-dropdown-item @click="closeTabsOnSide(route.fullPath, 'left')">
              <el-icon><DArrowLeft /></el-icon>关闭左侧
            </el-dropdown-item>
            <el-dropdown-item @click="closeTabsOnSide(route.fullPath, 'right')">
              <el-icon><DArrowRight /></el-icon>关闭右侧
            </el-dropdown-item>
            <el-dropdown-item divided @click="closeMultipleTab(route.fullPath)">
              <el-icon><CircleClose /></el-icon>关闭其他
            </el-dropdown-item>
            <el-dropdown-item @click="closeAllTab">
              <el-icon><FolderDelete /></el-icon>关闭全部
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </div>
</template>
<script setup>
import { ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import {HOME_URL} from '@/config/configs.js';

const route = useRoute();
const router = useRouter();
const tabsMenuValue = ref(route.fullPath);
const tabsIcon = ref(true);
/*const tabsMenuList = ref([
  {
    "icon": "HomeFilled",
    "title": "首页",
    "path": "/dashboard",
    "name": "home",
    "close": false,
    "isKeepAlive": true
  },
  {
    "icon": "User",
    "title": "用户列表",
    "path": "/user",
    "name": "userList",
    "close": true,
    "isKeepAlive": true
  }
]);*/
const tabsMenuList = ref([]);

watch(route, (newRoute) => {
  const exists = tabsMenuList.value.some(tab => tab.path === newRoute.fullPath);
  if (!exists) {
    tabsMenuList.value.push({
      title: newRoute.meta.title || newRoute.name,
      path: newRoute.fullPath,
      close: newRoute.meta.close || true,
      icon: newRoute.meta.icon || ''
    });
  }
  tabsMenuValue.value = newRoute.fullPath;
}, { immediate: true });
const tabClick = (tabItem) => {
  const fullPath = tabItem.props.name;
  router.push(fullPath);
};
const tabRemove = (removedPath) => {
  // 查找被关闭的标签页在 tabsMenuList 中的位置
  const index = tabsMenuList.value.findIndex(tab => tab.path === removedPath);
  if (index !== -1) {
    // 从 tabsMenuList 中移除被关闭的标签页
    tabsMenuList.value.splice(index, 1);
  }
  // 如果被关闭的标签页是当前激活的标签页
  if (tabsMenuValue.value === removedPath && tabsMenuList.value.length > 0) {
    // 切换到下一个标签页，如果有的话
    const nextTab = tabsMenuList.value[Math.max(0, index - 1)];
    router.push(nextTab.path);
  }
}
const refresh = () => {

};
const maximize = () => {
  
}
const closeCurrentTab = () => {

};
const closeTabsOnSide = () => {
  
}
const closeMultipleTab = () => {
  
}
const closeAllTab = () => {

};
</script>
<style lang="scss" scoped>
@import "./index.scss";
</style>