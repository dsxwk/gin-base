<template>
	<div class="system-menu-dialog-container">
		<el-dialog :title="state.dialog.title" v-if="state.dialog.isShowDialog" v-model="state.dialog.isShowDialog" width="769px">
      <ConfigForm
          ref="dialogFormRef"
          v-model:model="state.ruleForm"
          :form-config="formData"
          :rules="rules"
          :form-props="{
            labelWidth: '80px',
            size: 'default'
          }"
      >
      </ConfigForm>
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
import {onMounted, reactive, ref, markRaw, nextTick} from 'vue';
import {storeToRefs} from 'pinia';
import {useRoutesList} from '/@/stores/routesList';
import {i18n} from '/@/static/i18n';
import {ElMessage} from 'element-plus';
import {menuApi} from '/@/api/menu';
import {initBackEndControlRoutes} from '/@/router/backEnd.js';
import ConfigForm from '/@/components/form/index.vue';
import {isAffixDict, isHideDict, isIframeDict, isKeepAliveDict, isLinkDict} from '/@/dict/menu';
import CascaderLabel from '/@/components/form/CascaderLabel.vue';

const props = defineProps({
  row: {
    type: Object,
    required: true,
    default: () => ({})
  }
});

// 定义子组件向父组件传值/事件
const emit = defineEmits(['refresh']);

const api = menuApi();

// 定义变量内容
const dialogFormRef = ref();
const stores = useRoutesList();
const { routesList } = storeToRefs(stores);
const state = reactive({
  roles: [],
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
      roles: [], // 权限标识，取角色管理
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
// 是否内嵌下拉改变
const onSelectIframeChange = () => {
  if (state.ruleForm.meta.isIframe) state.ruleForm.isLink = true;
  else state.ruleForm.isLink = false;
};
const formData = ref([
  {
    label: '上级菜单',
    prop: 'menuSuperior',
    type: 'cascader',
    options: () => {
      return state.menuData;
    },
    props: {
      checkStrictly: true,
      value: 'path',
      label: 'title',
    },
    attrs: {
      placeholder: '请选择上级菜单',
      clearable: true,
      class: 'w100',
    },
    slotDefault: markRaw(CascaderLabel),
  },
  {
    label: '菜单名称',
    prop: 'meta.title',
    type: 'input',
    col: 12,
    attrs: {
      placeholder: '格式：message.router.xxx',
      clearable: true
    },
    rules: [
      {
        required: true,
        message: "请输入菜单名称",
        trigger: "blur"
      },
    ]
  },
  {
    label: '路由名称',
    prop: 'name',
    type: 'input',
    col: 12,
    attrs: {
      placeholder: '路由中的 name 值',
      clearable: true
    },
    rules: [
      {
        required: true,
        message: "请输入路由名称",
        trigger: "blur"
      },
    ]
  },
  {
    label: '路由路径',
    prop: 'path',
    type: 'input',
    col: 12,
    attrs: {
      placeholder: '路由中的 path 值',
      clearable: true
    },
    rules: [
      {
        required: true,
        message: "请输入路由路径",
        trigger: "blur"
      },
    ]
  },
  {
    label: '重定向',
    prop: 'redirect',
    type: 'input',
    col: 12,
    attrs: {
      placeholder: '请输入路由重定向',
      clearable: true
    }
  },
  {
    label: '菜单图标',
    prop: 'meta.icon',
    type: 'icon', // 自定义类型
    col: 12,
    attrs: {
      placeholder: '请选择菜单图标'
    }
  },
  {
    label: '组件路径',
    prop: 'component',
    type: 'input',
    col: 12,
    attrs: {
      placeholder: '请输入组件路径',
      clearable: true
    },
    rules: [
      {
        required: true,
        message: "请输入组件路径",
        trigger: "blur"
      },
    ]
  },
  {
    label: '链接地址',
    prop: 'meta.isLink',
    type: 'input',
    col: 12,
    attrs: () => {
      return {
        placeholder: '外链/内嵌时链接地址（http:xxx.com）',
        clearable: true,
        disabled: !state.ruleForm.isLink
      };
    }
  },
  {
    label: '菜单角色',
    prop: 'meta.roles',
    type: 'select',
    col: 12,
    options: () => {
      return state.roles.map(role => ({label: role.name, value: role.id}));
    },
    attrs: {
      placeholder: '请选择角色',
      multiple: true,
      clearable: true,
      class: 'w100'
    },
  },
  {
    label: '菜单排序',
    prop: 'sort',
    type: 'input-number',
    col: 12,
    attrs: {
      placeholder: '请输入排序',
      controlsPosition: 'right',
      class: 'w100'
    }
  },
  {
    label: '是否隐藏',
    prop: 'meta.isHide',
    type: 'radio',
    col: 12,
    options: isHideDict
  },
  {
    label: '页面缓存',
    prop: 'meta.isKeepAlive',
    type: 'radio',
    col: 12,
    options: isKeepAliveDict
  },
  {
    label: '是否固定',
    prop: 'meta.isAffix',
    type: 'radio',
    col: 12,
    options: isAffixDict
  },
  {
    label: '是否外链',
    prop: 'isLink',
    type: 'radio',
    col: 12,
    attrs: {
      disabled: state.ruleForm.meta.isIframe
    },
    options: isLinkDict
  },
  {
    label: '是否内嵌',
    prop: 'meta.isIframe',
    type: 'radio',
    col: 12,
    attrs: {
      onChange: onSelectIframeChange
    },
    options: isIframeDict
  }
]);
const rules = {};
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
// 递归查找父级 path 数组
function findMenuPathById(data, targetId, pathArr = []) {
  for (const item of data) {
    const newPathArr = [...pathArr, item.path];
    if (item.id === targetId) {
      return newPathArr;
    }
    if (item.children && item.children.length) {
      const found = findMenuPathById(item.children, targetId, newPathArr);
      if (Array.isArray(found) && found.length) return found;
    }
  }
  return [];
}
// 递归查找菜单项
function findMenuByPath(data, path) {
  for (const item of data) {
    if (item.path === path) return item;
    if (item.children && item.children.length) {
      const found = findMenuByPath(item.children, path);
      if (found) return found;
    }
  }
  return null;
}
// 打开弹窗
const openDialog = async (type, row) => {
  state.ruleForm = {
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
      roles: [], // 权限标识，取角色管理
    },
    btnPower: '', // 菜单类型为按钮时，权限标识
  };
  if (type === 'edit') {
    const data = await detail(row.id);
    Object.keys(state.ruleForm).forEach((key) => {
      if (data.hasOwnProperty(key)) {
        state.ruleForm[key] = data[key];
      }
    });
    // 设置上级菜单默认选中
    if (row.pid) {
      state.ruleForm.menuSuperior = findMenuPathById(state.menuData, row.pid);
    } else {
      state.ruleForm.menuSuperior = [];
    }
		state.dialog.title = '修改菜单';
		state.dialog.submitTxt = '修 改';
	} else {
		state.dialog.title = '新增菜单';
		state.dialog.submitTxt = '新 增';
	}
	state.dialog.type = type;
	state.dialog.isShowDialog = true;
  // 清空表单，此项需加表单验证才能使用
  await nextTick(() => {
    dialogFormRef.value && dialogFormRef.value.resetFields();
  });
};
// 关闭弹窗
const closeDialog = () => {
	state.dialog.isShowDialog = false;
};
// 取消
const onCancel = () => {
	closeDialog();
};
// 提交
const onSubmit = async () => {
  // 获取上级菜单 id
  if (state.ruleForm.menuSuperior && state.ruleForm.menuSuperior.length > 0) {
    const lastPath = state.ruleForm.menuSuperior[state.ruleForm.menuSuperior.length - 1];
    const menu = findMenuByPath(state.menuData, lastPath);
    state.ruleForm.pid = menu ? menu.id : 0;
  } else {
    state.ruleForm.pid = 0; // 顶级菜单
  }

  state.ruleForm.menuRoles = state.ruleForm.meta.roles?.map(roleId => {
    const role = state.roles.find(r => r.id === roleId);
    return {
      menuId: props.row.id ?? 0,
      roleId: roleId,
      name: role ? role.name : ''
    };
  }) ?? [];

  dialogFormRef.value.validate(async (valid) => {
    if (!valid) return;

    let msg = '';
    if (state.dialog.type === 'add') {
      await api.create(state.ruleForm);
      msg = '创建成功';
    } else {
      state.ruleForm.id = props.row.id;
      await api.update(state.ruleForm);
      msg = '更新成功';
    }
    ElMessage.success(msg);
    closeDialog();
    emit('refresh');
    await initBackEndControlRoutes();
  });
};
// 获取角色
const getRoles = async () => {
  const data = await api.roleList({page:1, pageSize: 10, isPage: false});
  state.roles = data.data;
};
// 详情
const detail = async (id) => {
  const res = await api.detail({id: id});

  return res.data;
};
// 页面加载时
onMounted(() => {
  getRoles();
	state.menuData = getData(routesList.value);
});
// 暴露变量
defineExpose({
	openDialog,
});
</script>
