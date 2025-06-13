import { useUserInfo } from '/@/stores/userInfo';
import { judementSameArr } from '/@/utils/arrayOperation';

/**
 * 用户权限指令
 * @directive 单个权限验证（v-auth="xxx"）
 * @directive 多个权限验证，满足一个则显示（v-auths="[xxx,xxx]"）
 * @directive 多个权限验证，全部满足则显示（v-auth-all="[xxx,xxx]"）
 */
export function authDirective(app) {
	// // 单个权限验证（v-auth="xxx"）
	// app.directive('auth', {
	// 	mounted(el, binding) {
	//
	// 		const route = el.__vueParentComponent?.appContext.config.globalProperties.$route;
	// 		const stores = useUserInfo();
	//
	// 		// 当前路由的按钮权限列表
	// 		const metaAuthBtnList = route.meta?.authBtnList?.map(item => {
	// 			return item?.authValue
	// 		}) || [];
	//
	// 		// 用户权限和路由meta权限两重校验，举例是取交集
	// 		const hasPermission =
	// 			stores.userInfos.authBtnList.includes(binding.value) &&
	// 			metaAuthBtnList.includes(binding.value);
	//
	// 		if (!hasPermission) {
	// 			el?.parentNode.removeChild(el);
	// 		}
	// 	},
	// });
	//
	// // 多个权限验证，满足一个则显示（v-auths="[xxx,xxx]"）
	// app.directive('auths', {
	// 	mounted(el, binding) {
	// 		const route = el.__vueParentComponent?.appContext.config.globalProperties.$route;
	// 		const stores = useUserInfo();
	//
	// 		// 当前路由的按钮权限列表
	// 		const metaAuthBtnList = route.meta?.authBtnList?.map(item => {
	// 			return item?.authValue
	// 		}) || [];
	//
	// 		const hasPermission = binding.value.some(
	// 			v => stores.userInfos.authBtnList.includes(v) && metaAuthBtnList.includes(v)
	// 		);
	//
	// 		if (!hasPermission) {
	// 			el?.parentNode.removeChild(el);
	// 		}
	// 	},
	// });
	//
	// // 多个权限验证，全部满足则显示（v-auth-all="[xxx,xxx]"）
	// app.directive('auth-all', {
	// 	mounted(el, binding) {
	// 		const route = el.__vueParentComponent?.appContext.config.globalProperties.$route;
	// 		const stores = useUserInfo();
	//
	// 		// 当前路由的按钮权限列表
	// 		const metaAuthBtnList = route.meta?.authBtnList?.map(item => {
	// 			return item?.authValue
	// 		}) || [];
	//
	// 		// 需要满足所有传入的权限，同时也在当前路由的meta权限内
	// 		const neededAuths = binding.value.filter(v => metaAuthBtnList.includes(v));
	//
	// 		const hasPermission =
	// 			judementSameArr(neededAuths, stores.userInfos.authBtnList) &&
	// 			neededAuths.length === binding.value.length;
	//
	// 		if (!hasPermission) {
	// 			el.parentNode.removeChild(el);
	// 		}
	// 	},
	// });

	// 单个权限验证（v-auth="xxx"）
	app.directive('auth', {
		mounted(el, binding) {
			const route = el.__vueParentComponent?.appContext.config.globalProperties.$route;
			// 只用 meta.authBtnList 判断
			const metaAuthBtnList = route.meta?.authBtnList?.map(item => item?.authValue).filter(Boolean) || [];

			const hasPermission = metaAuthBtnList.includes(binding.value);

			if (!hasPermission) {
				el?.parentNode?.removeChild(el);
			}
		},
	});

	// 多个权限验证，满足一个则显示（v-auths="[xxx,xxx]"）
	app.directive('auths', {
		mounted(el, binding) {
			const route = el.__vueParentComponent?.appContext.config.globalProperties.$route;
			const metaAuthBtnList = route.meta?.authBtnList?.map(item => item?.authValue).filter(Boolean) || [];

			const hasPermission = binding.value?.some(v => metaAuthBtnList.includes(v));

			if (!hasPermission) {
				el?.parentNode?.removeChild(el);
			}
		},
	});

	// 多个权限验证，全部满足则显示（v-auth-all="[xxx,xxx]"）
	app.directive('auth-all', {
		mounted(el, binding) {
			const route = el.__vueParentComponent?.appContext.config.globalProperties.$route;
			const metaAuthBtnList = route.meta?.authBtnList?.map(item => item?.authValue).filter(Boolean) || [];

			const neededAuths = binding.value?.filter(v => metaAuthBtnList.includes(v));

			const hasPermission = neededAuths.length === binding.value.length;

			if (!hasPermission) {
				el?.parentNode?.removeChild(el);
			}
		},
	});
}
