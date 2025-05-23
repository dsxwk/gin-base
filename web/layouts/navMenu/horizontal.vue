<template>
	<div class="el-menu-horizontal-warp">
		<el-menu router :default-active="state.defaultActive" background-color="transparent" mode="horizontal">
			<template v-for="val in menuLists">
				<el-sub-menu :index="val.path" v-if="val.children && val.children.length > 0" :key="val.path">
					<template #title>
						<SvgIcon :name="val.meta.icon" />
						<span>{{ $t(val.meta.title) }}</span>
					</template>
					<SubItem :chil="val.children" />
				</el-sub-menu>
				<template v-else>
					<el-menu-item :index="val.path" :key="val.path">
						<template #title v-if="!val.meta.isLink || (val.meta.isLink && val.meta.isIframe)">
							<SvgIcon :name="val.meta.icon" />
							{{ $t(val.meta.title) }}
						</template>
						<template #title v-else>
							<a class="w100" @click.prevent="onALinkClick(val)">
								<SvgIcon :name="val.meta.icon" />
								{{ $t(val.meta.title) }}
							</a>
						</template>
					</el-menu-item>
				</template>
			</template>
		</el-menu>
	</div>
</template>

<script setup name="navMenuHorizontal">
import { defineAsyncComponent, reactive, computed, onBeforeMount } from 'vue';
import { useRoute, onBeforeRouteUpdate, RouteRecordRaw } from 'vue-router';
import { storeToRefs } from 'pinia';
import { useRoutesList } from '/@/stores/routesList';
import { useThemeConfig } from '/@/stores/themeConfig';
import other from '/@/utils/other';
import mittBus from '/@/utils/mitt';

// 引入组件
const SubItem = defineAsyncComponent(() => import('/@/layouts/navMenu/subItem.vue'));

// 定义父组件传过来的值
const props = defineProps({
	// 菜单列表
	menuList: {
		type,
		default: () => [],
	},
});

// 定义变量内容
const stores = useRoutesList();
const storesThemeConfig = useThemeConfig();
const { routesList } = storeToRefs(stores);
const { themeConfig } = storeToRefs(storesThemeConfig);
const route = useRoute();
const state = reactive({
	defaultActive: '',
});

// 获取父级菜单数据
const menuLists = computed(() => {
	return props.menuList;
});
// 路由过滤递归函数
const filterRoutesFun = (arr) => {
	return arr
		.filter((item) => !item.meta?.isHide)
		.map((item) => {
			item = Object.assign({}, item);
			if (item.children) item.children = filterRoutesFun(item.children);
			return item;
		});
};
// 传送当前子级数据到菜单中
const setSendClassicChildren = (path) => {
	const currentPathSplit = path.split('/');
	let currentData = { children: [] };
	filterRoutesFun(routesList.value).map((v, k) => {
		if (v.path === `/${currentPathSplit[1]}`) {
			v['k'] = k;
			currentData['item'] = { ...v };
			currentData['children'] = [{ ...v }];
			if (v.children) currentData['children'] = v.children;
		}
	});
	return currentData;
};
// 设置页面当前路由高亮
const setCurrentRouterHighlight = (currentRoute) => {
  const { path, meta } = currentRoute;

  if (!path) {
    console.warn('当前路由没有 path 属性');
    return;
  }

  if (themeConfig.value.layout === 'classic') {
    // 如果是 classic 布局，高亮第一级路径
    state.defaultActive = `/${path.split('/')[1]}`;
  } else {
    // 如果不是 classic 布局
    const pathSplit = meta?.isDynamic ? meta.isDynamicPath?.split('/') : path.split('/');
    if (pathSplit && pathSplit.length >= 4 && meta?.isHide) {
      // 如果路径分割后长度大于等于 4 且 meta.isHide 为 true
      state.defaultActive = pathSplit.slice(0, 3).join('/');
    } else {
      // 否则高亮完整路径
      state.defaultActive = path;
    }
  }
};
// 打开外部链接
const onALinkClick = (val) => {
	other.handleOpenLink(val);
};
// 页面加载前
onBeforeMount(() => {
	setCurrentRouterHighlight(route);
});
// 路由更新时
onBeforeRouteUpdate((to) => {
	// 修复：https://gitee.com/lyt-top/vue-next-admin/issues/I3YX6G
	setCurrentRouterHighlight(to);
	// 修复经典布局开启切割菜单时，点击tagsView后左侧导航菜单数据不变的问题
	let { layout, isClassicSplitMenu } = themeConfig.value;
	if (layout === 'classic' && isClassicSplitMenu) {
		mittBus.emit('setSendClassicChildren', setSendClassicChildren(to.path));
	}
});
</script>

<style scoped lang="scss">
.el-menu-horizontal-warp {
	flex: 1;
	overflow: hidden;
	margin-right: 30px;
	:deep(.el-scrollbar__bar.is-vertical) {
		display: none;
	}
	:deep(a) {
		width: 100%;
	}
	.el-menu.el-menu--horizontal {
		display: flex;
		height: 100%;
		width: 100%;
		box-sizing: border-box;
	}
}
</style>
