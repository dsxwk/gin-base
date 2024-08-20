<template>
  <div class="search-menu">
    <i :title="funcs.lang('Menu Search')" :class="'iconfont icon-sousuo'" class="toolBar-icon" @click="handleOpen"></i>
    <el-dialog class="search-dialog" v-model="isShowSearch" :width="600" :show-close="false" top="10vh">
      <el-input
        v-model="searchMenu"
        ref="menuInputRef"
        :placeholder="funcs.lang('Menu Search')+': '+funcs.lang('Support menu names and paths')"
        size="large"
        clearable
        :prefix-icon="Search"
      ></el-input>
      <div v-if="searchList.length" class="menu-list" ref="menuListRef">
        <div
          v-for="item in searchList"
          :key="item.path"
          :class="['menu-item', { 'menu-active': item.path === activePath }]"
          @mouseenter="mouseoverMenuItem(item)"
          @click="handleClickMenu()"
        >
          <div class="menu-lf">
            <el-icon class="menu-icon">
              <!-- 动态解析图标组件 -->
              <component :is="resolveIconComponent(item.meta.icon)"></component>
            </el-icon>
            <span class="menu-title">{{ item.meta.title }}</span>
          </div>
          <i :class="'iconfont icon-huiche'" class="menu-enter" @click="handleOpen"></i>
        </div>
      </div>
      <el-empty v-else class="mt20 mb20" :image-size="100" :description="funcs.lang('No Menu')" />
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, computed, nextTick, watch, onMounted, onUnmounted, resolveComponent } from 'vue';
import { Search } from '@element-plus/icons-vue';
import { useRouter } from 'vue-router';
import { useDebounceFn } from '@vueuse/core';
import { menuJson } from '@/utils/data/menu';
import Functions from '@/utils/functions';

const funcs = new Functions();
const router = useRouter();
const menuList = computed(() => (menuJson.filter(item => !item.meta.isHide)));

// 动态解析图标组件
const resolveIconComponent = (iconName) => {
  return iconName ? resolveComponent(iconName) : null;
};

onMounted(() => {
  document.addEventListener('keydown', keyboardOperation);
});

onUnmounted(() => {
  document.removeEventListener('keydown', keyboardOperation);
});

const activePath = ref('');
const mouseoverMenuItem = (menu) => {
  activePath.value = menu.path;
};

const menuInputRef = ref(null);
const isShowSearch = ref(false);
const searchMenu = ref('');
const handleOpen = () => {
  isShowSearch.value = true;
  nextTick(() => {
    setTimeout(() => {
      menuInputRef.value?.focus();
    });
  });
};

const searchList = ref([]);
const filterMenu = (menu, searchValue) => {
  const filteredMenu = [];

  menu.forEach(item => {
    // 检查当前菜单项是否匹配搜索条件
    const isMatch = item.path.toLowerCase().includes(searchValue.toLowerCase()) || item.meta.title.toLowerCase().includes(searchValue.toLowerCase());

    // 如果当前菜单项匹配，或者其子菜单中有匹配项，则将其加入结果列表
    if (isMatch) {
      filteredMenu.push(item);
    }

    // 递归检查子菜单
    if (item.children && item.children.length > 0) {
      const filteredChildren = filterMenu(item.children, searchValue);
      if (filteredChildren.length > 0) {
        filteredMenu.push(...filteredChildren);
      }
    }
  });

  return filteredMenu;
};

const updateSearchList = () => {
  searchList.value = searchMenu.value ? filterMenu(menuList.value, searchMenu.value) : [];
  activePath.value = searchList.value.length ? searchList.value[0].path : '';
};

const debouncedUpdateSearchList = useDebounceFn(updateSearchList, 300);

watch(searchMenu, debouncedUpdateSearchList);

const menuListRef = ref(null);
const keyPressUpOrDown = (direction) => {
  const length = searchList.value.length;
  if (length === 0) return;
  const index = searchList.value.findIndex(item => item.path === activePath.value);
  const newIndex = (index + direction + length) % length;
  activePath.value = searchList.value[newIndex].path;
  nextTick(() => {
    if (!menuListRef.value?.firstElementChild) {
      return;
    }
    const menuItemHeight = menuListRef.value.firstElementChild.clientHeight + 12 || 0;
    menuListRef.value.scrollTop = newIndex * menuItemHeight;
  });
};

const keyboardOperation = (event) => {
  if (event.key === 'ArrowUp') {
    event.preventDefault();
    keyPressUpOrDown(-1);
  } else if (event.key === 'ArrowDown') {
    event.preventDefault();
    keyPressUpOrDown(1);
  } else if (event.key === 'Enter') {
    event.preventDefault();
    handleClickMenu();
  }
};

const handleClickMenu = () => {
  const menu = searchList.value.find(item => item.path === activePath.value);
  if (!menu) {
    return;
  }
  if (menu.meta?.isLink) {
    window.open(menu.meta.isLink, '_blank');
  } else {
    router.push(menu.path);
    searchMenu.value = '';
    isShowSearch.value = false;
  }
};
</script>

<style scoped lang="scss">
@import './index.scss';
</style>
