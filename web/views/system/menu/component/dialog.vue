<template>
	<div class="system-menu-dialog-container">
		<el-dialog :title="state.dialog.title" v-model="state.dialog.isShowDialog" width="769px">
			<el-form ref="dialogFormRef" :model="state.ruleForm" size="default" label-width="80px">
				<el-row :gutter="35">
					<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
						<el-form-item label="上级菜单">
							<el-cascader
								:options="state.menuData"
								:props="{ checkStrictly: true, value: 'path', label: 'title' }"
								placeholder="请选择上级菜单"
								clearable
								class="w100"
								v-model="state.ruleForm.menuSuperior"
							>
								<template #default="{ node, data }">
									<span>{{ data.title }}</span>
									<span v-if="!node.isLeaf"> ({{ data.children.length }}) </span>
								</template>
							</el-cascader>
						</el-form-item>
					</el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="菜单名称">
							<el-input v-model="state.ruleForm.meta.title" placeholder="格式：message.router.xxx" clearable></el-input>
						</el-form-item>
					</el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="路由名称">
              <el-input v-model="state.ruleForm.name" placeholder="路由中的 name 值" clearable></el-input>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="路由路径">
              <el-input v-model="state.ruleForm.path" placeholder="路由中的 path 值" clearable></el-input>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="重定向">
              <el-input v-model="state.ruleForm.redirect" placeholder="请输入路由重定向" clearable></el-input>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="菜单图标">
              <IconSelector placeholder="请输入菜单图标" v-model="state.ruleForm.meta.icon" />
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="组件路径">
              <el-input v-model="state.ruleForm.component" placeholder="组件路径" clearable></el-input>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="链接地址">
              <el-input
                  v-model="state.ruleForm.meta.isLink"
                  placeholder="外链/内嵌时链接地址（http:xxx.com）"
                  clearable
                  :disabled="!state.ruleForm.isLink"
              >
              </el-input>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="权限标识">
              <el-select v-model="state.ruleForm.meta.roles" multiple placeholder="取角色管理" clearable class="w100">
                <el-option label="admin" value="admin"></el-option>
                <el-option label="common" value="common"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
					<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
						<el-form-item label="菜单排序">
							<el-input-number v-model="state.ruleForm.sort" controls-position="right" placeholder="请输入排序" class="w100" />
						</el-form-item>
					</el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="是否隐藏">
              <el-radio-group v-model="state.ruleForm.meta.isHide">
                <el-radio :value="true">隐藏</el-radio>
                <el-radio :value="false">不隐藏</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="页面缓存">
              <el-radio-group v-model="state.ruleForm.meta.isKeepAlive">
                <el-radio :value="true">缓存</el-radio>
                <el-radio :value="false">不缓存</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="是否固定">
              <el-radio-group v-model="state.ruleForm.meta.isAffix">
                <el-radio :value="true">固定</el-radio>
                <el-radio :value="false">不固定</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="是否外链">
              <el-radio-group v-model="state.ruleForm.isLink" :disabled="state.ruleForm.meta.isIframe">
                <el-radio :value="true">是</el-radio>
                <el-radio :value="false">否</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="是否内嵌">
              <el-radio-group v-model="state.ruleForm.meta.isIframe" @change="onSelectIframeChange">
                <el-radio :value="true">是</el-radio>
                <el-radio :value="false">否</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
				</el-row>
			</el-form>
			<template #footer>
				<span class="dialog-footer">
					<el-button @click="onCancel" size="default">取 消</el-button>
					<el-button type="primary" @click="onSubmit" size="default">{{ state.dialog.submitTxt }}</el-button>
				</span>
			</template>
		</el-dialog>
	</div>
</template>

<script setup name="systemMenuDialog">
import { defineAsyncComponent, reactive, onMounted, ref } from 'vue';
import { storeToRefs } from 'pinia';
import { useRoutesList } from '/@/stores/routesList';
import { i18n } from '/@/static/i18n';

// 定义子组件向父组件传值/事件
const emit = defineEmits(['refresh']);

// 引入组件
const IconSelector = defineAsyncComponent(() => import('/@/components/iconSelector/index.vue'));

// 定义变量内容
const dialogFormRef = ref();
const stores = useRoutesList();
const { routesList } = storeToRefs(stores);
const state = reactive({
	// 参数请参考 `/router/route.js` 中的 `dynamicRoutes` 路由菜单格式
	ruleForm: {
		menuSuperior: [], // 上级菜单
		name: '', // 路由名称
		component: '', // 组件路径
		componentAlias: '', // 组件路径别名
		isLink: false, // 是否外链
		sort: 0, // 菜单排序
		path: '', // 路由路径
		redirect: '', // 路由重定向，有子集 children 时
		meta: {
			title: '', // 菜单名称
			icon: '', // 菜单图标
			isHide: false, // 是否隐藏
			isKeepAlive: true, // 是否缓存
			isAffix: false, // 是否固定
			isLink: '', // 外链/内嵌时链接地址（http:xxx.com），开启外链条件，`1、isLink: 链接地址不为空`
			isIframe: false, // 是否内嵌，开启条件，`1、isIframe:true 2、isLink：链接地址不为空`
			roles: '', // 权限标识，取角色管理
		},
		btnPower: '', // 菜单类型为按钮时，权限标识
	},
	menuData: [], // 上级菜单数据
	dialog: {
		isShowDialog: false,
		type: '',
		title: '',
		submitTxt: '',
	},
});

// 获取 pinia 中的路由
const getData = (routes) => {
	const arr = [];
	routes.map((val) => {
		val['title'] = i18n.global.t(val.meta?.title);
		arr.push({ ...val });
		if (val.children) getData(val.children);
	});
	return arr;
};
// 打开弹窗
const openDialog = (type, row) => {
	if (type === 'edit') {
		// 模拟数据，实际请走接口
		row.menuType = 'menu';
		row.sort = Math.floor(Math.random() * 100);
		state.ruleForm = JSON.parse(JSON.stringify(row));
		state.dialog.title = '修改菜单';
		state.dialog.submitTxt = '修 改';
	} else {
		state.dialog.title = '新增菜单';
		state.dialog.submitTxt = '新 增';
		// 清空表单，此项需加表单验证才能使用
		// nextTick(() => {
		// 	dialogFormRef.value.resetFields();
		// });
	}
	state.dialog.type = type;
	state.dialog.isShowDialog = true;
};
// 关闭弹窗
const closeDialog = () => {
	state.dialog.isShowDialog = false;
};
// 是否内嵌下拉改变
const onSelectIframeChange = () => {
	if (state.ruleForm.meta.isIframe) state.ruleForm.isLink = true;
	else state.ruleForm.isLink = false;
};
// 取消
const onCancel = () => {
	closeDialog();
};
// 提交
const onSubmit = () => {
	closeDialog(); // 关闭弹窗
	emit('refresh');
	// if (state.dialog.type === 'add') { }
	// setBackEndControlRefreshRoutes() // 刷新菜单，未进行后端接口测试
};
// 页面加载时
onMounted(() => {
	state.menuData = getData(routesList.value);
});

// 暴露变量
defineExpose({
	openDialog,
});
</script>
