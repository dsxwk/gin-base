import {defineStore} from 'pinia';
// import Cookies from 'js-cookie';
import {Session} from '/@/utils/storage';

/**
 * 用户信息
 * @methods setUserInfos 设置用户信息
 */
export const useUserInfo = defineStore('userInfo', {
	state: () => ({
		userInfos: {
			username: '',
			avatar: '',
			time: 0,
			roles: [],
			authBtnList: [],
		},
	}),
	actions: {
		async setUserInfos(userInfo) {
			Session.set('userInfo', userInfo);
			this.userInfos = userInfo;
			// 存储用户信息到浏览器缓存
			// if (Session.get('userInfo')) {
			// 	this.userInfos = Session.get('userInfo');
			// } else {
			// 	// this.userInfos = await this.getApiUserInfo();
			// 	this.userInfos = {};
			// }
		},
		// 模拟接口数据
		// https://gitee.com/lyt-top/vue-next-admin/issues/I5F1HP
		// async getApiUserInfo() {
		// 	return new Promise((resolve) => {
		// 		setTimeout(() => {
		// 			// 模拟数据，请求接口时，记得删除多余代码及对应依赖的引入
		// 			const username = Cookies.get('username');
		// 			// 模拟数据
		// 			let defaultRoles = [];
		// 			let defaultAuthBtnList = [];
		// 			// admin 页面权限标识，对应路由 meta.roles，用于控制路由的显示/隐藏
		// 			let adminRoles = ['admin'];
		// 			// admin 按钮权限标识
		// 			let adminAuthBtnList = ['btn.add', 'btn.del', 'btn.edit', 'btn.link'];
		// 			// test 页面权限标识，对应路由 meta.roles，用于控制路由的显示/隐藏
		// 			let testRoles = ['common'];
		// 			// test 按钮权限标识
		// 			let testAuthBtnList = ['btn.add', 'btn.link'];
		// 			// 不同用户模拟不同的用户权限
		// 			if (username === 'admin') {
		// 				defaultRoles = adminRoles;
		// 				defaultAuthBtnList = adminAuthBtnList;
		// 			} else {
		// 				defaultRoles = testRoles;
		// 				defaultAuthBtnList = testAuthBtnList;
		// 			}
		// 			// 用户信息模拟数据
		// 			const userInfos = {
		// 				username: username,
		// 				avatar:
		// 					username === 'admin'
		// 						? 'https://img2.baidu.com/it/u=1978192862,2048448374&fm=253&fmt=auto&app=138&f=JPEG?w=504&h=500'
		// 						: 'https://img2.baidu.com/it/u=2370931438,70387529&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=500',
		// 				time: new Date().getTime(),
		// 				roles: defaultRoles,
		// 				authBtnList: defaultAuthBtnList,
		// 			};
		// 			Session.set('userInfo', userInfos);
		// 			resolve(userInfos);
		// 		}, 0);
		// 	});
		// },
	},
});
