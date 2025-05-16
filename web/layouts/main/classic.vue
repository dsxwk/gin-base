<template>
	<el-container class="layout-container flex-center">
		<LayoutHeader />
		<el-container class="layout-mian-height-50">
			<LayoutAside />
			<div class="flex-center layout-backtop">
				<LayoutTagsView v-if="isTagsview" />
				<LayoutMain ref="layoutMainRef" />
			</div>
		</el-container>
	</el-container>
</template>

<script setup name="layoutClassic">
import { defineAsyncComponent, computed, ref, watch, nextTick, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { storeToRefs } from 'pinia';
import { useThemeConfig } from '/@/stores/themeConfig';

// 引入组件
const LayoutAside = defineAsyncComponent(() => import('/@/layouts/component/aside.vue'));
const LayoutHeader = defineAsyncComponent(() => import('/@/layouts/component/header.vue'));
const LayoutMain = defineAsyncComponent(() => import('/@/layouts/component/main.vue'));
const LayoutTagsView = defineAsyncComponent(() => import('/@/layouts/navBars/tagsView/tagsView.vue'));

// 定义变量内容
const layoutMainRef = ref();
const route = useRoute();
const storesThemeConfig = useThemeConfig();
const { themeConfig } = storeToRefs(storesThemeConfig);

// 判断是否显示 tasgview
const isTagsview = computed(() => {
	return themeConfig.value.isTagsview;
});
// 重置滚动条高度，更新子级 scrollbar
const updateScrollbar = () => {
	layoutMainRef.value?.layoutMainScrollbarRef.update();
};
// 重置滚动条高度，由于组件是异步引入的
const initScrollBarHeight = () => {
  setTimeout(() => {
    updateScrollbar();
    // 检查 layoutMainRef 是否存在并安全访问其属性
    if (layoutMainRef && layoutMainRef.value && layoutMainRef.value.layoutMainScrollbarRef) {
      const wrapRef = layoutMainRef.value.layoutMainScrollbarRef.wrapRef;
      if (wrapRef) {
        wrapRef.scrollTop = 0;
      }
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
// 监听 themeConfig  isTagsview 配置文件的变化，更新菜单 el-scrollbar 的高度
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
