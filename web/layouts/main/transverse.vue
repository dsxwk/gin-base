<template>
	<el-container class="layout-container flex-center layout-backtop">
		<LayoutHeader />
		<LayoutMain ref="layoutMainRef" />
	</el-container>
</template>

<script setup name="layoutTransverse">
import { defineAsyncComponent, ref, watch, nextTick, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { storeToRefs } from 'pinia';
import { useThemeConfig } from '/@/stores/themeConfig';

// 引入组件
const LayoutHeader = defineAsyncComponent(() => import('/@/layouts/component/header.vue'));
const LayoutMain = defineAsyncComponent(() => import('/@/layouts/component/main.vue'));

// 定义变量内容
const layoutMainRef = ref();
const storesThemeConfig = useThemeConfig();
const { themeConfig } = storeToRefs(storesThemeConfig);
const route = useRoute();

// 更新子级 scrollbar
const updateScrollbar = () => {
  if (
      layoutMainRef &&
      layoutMainRef.value &&
      layoutMainRef.value.layoutMainScrollbarRef
  ) {
    layoutMainRef.value.layoutMainScrollbarRef.update();
  }
};

// 重置滚动条高度，由于组件是异步引入的
const initScrollBarHeight = () => {
  setTimeout(() => {
    updateScrollbar();
    if (
        layoutMainRef &&
        layoutMainRef.value &&
        layoutMainRef.value.layoutMainScrollbarRef &&
        layoutMainRef.value.layoutMainScrollbarRef.wrapRef
    ) {
      layoutMainRef.value.layoutMainScrollbarRef.wrapRef.scrollTop = 0;
    }
  }, 500);
};
// 页面加载时
onMounted(() => {
	initScrollBarHeight();
});
// 监听路由的变化，切换界面时，滚动条置顶
watch(
	() => route.path,
	() => {
		initScrollBarHeight();
	}
);
// 监听 themeConfig 配置文件的变化，更新菜单 el-scrollbar 的高度
watch(
	() => themeConfig.value.isTagsview,
	() => {
		nextTick(() => {
			updateScrollbar();
		});
	},
	{
		deep: true,
	}
);
</script>
