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
            <el-dropdown-item @click="closeTabsOnSide('left')">
              <el-icon><DArrowLeft /></el-icon>关闭左侧
            </el-dropdown-item>
            <el-dropdown-item @click="closeTabsOnSide('right')">
              <el-icon><DArrowRight /></el-icon>关闭右侧
            </el-dropdown-item>
            <el-dropdown-item divided @click="closeOthersTab(route.fullPath)">
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
import { ref, watch, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import {HOME_URL} from '@/config';
import Functions from '@/utils/functions';

const funcs = new Functions();
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

// 初始设置首页tab
const initializeHomeTab = () => {
  const homeRoute = {
    title: funcs.lang('Home'),
    path: '/dashboard?lang=' + funcs.getLang(),
    // 首页不可关闭
    close: false,
    icon: 'HomeFilled'
  };
  const homeTabExists = tabsMenuList.value.some((tab) => {
    return tab.path === homeRoute.path;
  });
  if (!homeTabExists) {
    tabsMenuList.value.unshift(homeRoute);
  }
};

onMounted(() => {
  initializeHomeTab();
});

watch(route, (newRoute) => {
  const exists = tabsMenuList.value.some((tab) => {
    return tab.path === newRoute.fullPath;
  });
  if (!exists) {
    tabsMenuList.value.push({
      title: newRoute.meta.title || newRoute.name,
      path: newRoute.fullPath,
      // 控制标签页是否可以关闭
      close: newRoute.meta.close !== false,
      icon: newRoute.meta.icon || ''
    });
  }
  tabsMenuValue.value = newRoute.fullPath;

  // 确保首页tab始终在最左边
  const homeTabIndex = tabsMenuList.value.findIndex((tab) => {
    return tab.path === HOME_URL;
  });
  if (homeTabIndex > 0) {
    const homeTab = tabsMenuList.value.splice(homeTabIndex, 1)[0];
    tabsMenuList.value.unshift(homeTab);
  }
}, { immediate: true });
const tabClick = (tabItem) => {
  const fullPath = tabItem.props.name;
  router.push(fullPath);
};
const tabRemove = (removedPath) => {
  if (removedPath === HOME_URL) {
    return;
  }
  // 查找被关闭的标签页在 tabsMenuList 中的位置
  const index = tabsMenuList.value.findIndex((tab) => {
    return tab.path === removedPath;
  });
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
  location.reload();
};
const maximize = () => {
  const elem = document.querySelector('main');
  if (!document.fullscreenElement) {
    elem.requestFullscreen().catch(err => console.error(`Error attempting to enable full-screen mode: ${err.message}`));
  } else {
    document.exitFullscreen();
  }
}
const closeCurrentTab = () => {
  tabRemove(tabsMenuValue.value);
};
const closeTabsOnSide = (side) => {
  const index = tabsMenuList.value.findIndex((tab) => {
    return tab.path === tabsMenuValue.value;
  });
  if (index !== -1) {
    if (side === 'left') {
      // 保留当前标签页及其右侧标签页，确保首页不会被删除
      tabsMenuList.value = tabsMenuList.value.filter((_, i) => i >= index || tabsMenuList.value[i].path === HOME_URL);
    } else if (side === 'right') {
      // 保留当前标签页及其左侧标签页，确保首页不会被删除
      tabsMenuList.value = tabsMenuList.value.filter((_, i) => i <= index || tabsMenuList.value[i].path === HOME_URL);
    }
  }
};
const closeOthersTab = () => {
  const currentTab = tabsMenuList.value.find((tab) => {
    return tab.path === tabsMenuValue.value;
  });
  const homeTab = tabsMenuList.value.find((tab) => {
    return tab.path === HOME_URL;
  });

  if (currentTab) {
    // 保证首页在第一个位置，然后保留当前选中的标签页
    if (homeTab) {
      tabsMenuList.value = [homeTab, currentTab].filter((v, i, a) => a.findIndex((t) => {
        return t.path === v.path;
      }) === i);
    } else {
      tabsMenuList.value = [currentTab];
    }
  }
}
const closeAllTab = () => {
  tabsMenuList.value = tabsMenuList.value.filter((tab) => {
    return tab.path === HOME_URL;
  });
  router.push(HOME_URL);
};
</script>
<style lang="scss" scoped>
@import "./index.scss";
</style>