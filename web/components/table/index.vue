<template>
	<div class="table-container layout-pd">
    <div class="flex items-center" style="justify-content: space-between;">
      <div class="table-tools-bar" v-if="slots.tools">
        <slot name="tools"></slot>
      </div>
      <div class="table-header-right-tool">
        <SvgIcon name="iconfont icon-dayin" :size="19" title="打印" @click="onPrintTable" />
        <SvgIcon name="iconfont icon-yunxiazai_o" :size="22" title="导出" @click="onImportTable" />
        <SvgIcon name="iconfont icon-shuaxin" :size="22" title="刷新" @click="onRefreshTable" />
        <el-popover
            placement="top-end"
            trigger="click"
            transition="el-zoom-in-top"
            popper-class="table-tool-popper"
            :width="300"
            :persistent="false"
            @show="onSetTable"
        >
          <template #reference>
            <SvgIcon name="iconfont icon-quanjushezhi_o" :size="22" title="设置" />
          </template>
          <template #default>
            <div class="tool-box">
              <el-tooltip content="拖动进行排序" placement="top-start">
                <SvgIcon name="fa fa-question-circle-o" :size="17" class="ml11" color="#909399" />
              </el-tooltip>
              <el-checkbox
                  v-model="state.checkListAll"
                  :indeterminate="state.checkListIndeterminate"
                  class="ml10 mr1"
                  label="列显示"
                  @change="onCheckAllChange"
              />
              <el-checkbox v-model="getConfig.isSerialNo" class="ml12 mr1" label="序号" />
              <el-checkbox v-model="getConfig.isSelection" class="ml12 mr1" label="多选" />
            </div>
            <el-scrollbar>
              <div ref="toolSetRef" class="tool-sortable">
                <div class="tool-sortable-item" v-for="v in header" :key="v.key" :data-key="v.key">
                  <i class="fa fa-arrows-alt handle cursor-pointer"></i>
                  <el-checkbox v-model="v.isCheck" size="default" class="ml12 mr8" :label="v.title" @change="onCheckChange" />
                </div>
              </div>
            </el-scrollbar>
          </template>
        </el-popover>
      </div>
    </div>
		<el-table
			:data="data"
			:border="setBorder"
			v-bind="$attrs"
			row-key="id"
			stripe
			style="width: 100%"
			v-loading="config.loading"
			@selection-change="onSelectionChange"
		>
			<el-table-column type="selection" :reserve-selection="true" width="40" align="center" v-if="config.isSelection" />
			<el-table-column type="index" label="序号" width="60" v-if="config.isSerialNo" />
			<el-table-column
				v-for="(item, index) in setHeader"
				:key="index"
				show-overflow-tooltip
				:prop="item.key"
				:width="item.colWidth"
				:label="item.title"
			>
				<template v-slot="scope">
					<template v-if="item.type === 'image'">
						<el-image
							:style="{ width: `${item.width}px`, height: `${item.height}px` }"
							:src="scope.row[item.key]"
							:zoom-rate="1.2"
							:preview-src-list="[scope.row[item.key]]"
							preview-teleported
							fit="cover"
						/>
					</template>
          <template v-else-if="item.render">
            <component
                :is="item.render"
                v-bind="scope"
                v-if="item.render"
            >
            </component>
          </template>
					<template v-else>
						{{ scope.row[item.key] }}
					</template>
				</template>
			</el-table-column>
			<el-table-column label="操作" :width="config.operationWith" v-if="config.isOperate" :fixed="config.fixed">
				<template v-slot="scope">
          <slot name="operation" :row="scope.row" v-if="slots.operation">
            <el-popconfirm title="确定删除吗？" @confirm="onDelRow(scope.row)">
              <template #reference>
                <el-button text type="danger">删除</el-button>
              </template>
            </el-popconfirm>
          </slot>
				</template>
			</el-table-column>
			<template #empty>
				<el-empty description="暂无数据" />
			</template>
		</el-table>
		<div class="table-footer mt15">
			<el-pagination
				v-model:current-page="state.page.page"
				v-model:page-size="state.page.pageSize"
				:pager-count="5"
				:page-sizes="[10, 50, 100, 200, 500, 1000]"
				:total="config.total"
				layout="total, sizes, prev, pager, next, jumper"
				background
				@size-change="onHandleSizeChange"
				@current-change="onHandleCurrentChange"
			>
			</el-pagination>
		</div>
    <div class="table-dialog" v-if="slots.dialog">
      <slot name="dialog"></slot>
    </div>
	</div>
</template>

<script setup name="netxTable">
import {reactive, onMounted, computed, nextTick, ref, useSlots} from 'vue';
import { ElMessage } from 'element-plus';
import printJs from 'print-js';
import table2excel from 'js-table2excel';
import Sortable from 'sortablejs';
import { storeToRefs } from 'pinia';
import { useThemeConfig } from '/@/stores/themeConfig';
import '/@/theme/tableTool.scss';

const slots = useSlots();
// 定义父组件传过来的值
const props = defineProps({
	// 列表内容
	data: {
		type: Array,
		default: () => [],
	},
	// 表头内容
	header: {
		type: Array,
		default: () => [],
	},
	// 配置项
	config: {
		type: Object,
		default: () => {},
	},
	// 打印标题
	printName: {
		type: String,
		default: () => '',
	},
});

// 定义子组件向父组件传值/事件
const emit = defineEmits(['delRow', 'pageChange', 'sortHeader']);

// 定义变量内容
const toolSetRef = ref();
const storesThemeConfig = useThemeConfig();
const { themeConfig } = storeToRefs(storesThemeConfig);
const state = reactive({
	page: {
		page: 1,
		pageSize: 10,
	},
	selectList: [],
	checkListAll: true,
	checkListIndeterminate: false,
});
// 设置边框显示/隐藏
const setBorder = computed(() => {
	return !!props.config.isBorder;
});
// 获取父组件 配置项（必传）
const getConfig = computed(() => {
	return props.config;
});
// 设置 tool header 数据
const setHeader = computed(() => {
	return props.header.filter((v) => v.isCheck);
});
// tool 列显示全选改变时
const onCheckAllChange = (val) => {
	if (val) props.header.forEach((v) => (v.isCheck = true));
	else props.header.forEach((v) => (v.isCheck = false));
	state.checkListIndeterminate = false;
};
// tool 列显示当前项改变时
const onCheckChange = () => {
	const headers = props.header.filter((v) => v.isCheck).length;
	state.checkListAll = headers === props.header.length;
	state.checkListIndeterminate = headers > 0 && headers < props.header.length;
};
// 表格多选改变时，用于导出
const onSelectionChange = (val) => {
	state.selectList = val;
};
// 删除当前项
const onDelRow = (row) => {
	emit('delRow', row);
};
// 分页改变
const onHandleSizeChange = (val) => {
	state.page.pageSize = val;
	emit('pageChange', state.page);
};
// 分页改变
const onHandleCurrentChange = (val) => {
	state.page.page = val;
	emit('pageChange', state.page);
};
// 搜索时，分页还原成默认
const pageReset = () => {
	state.page.page = 1;
	state.page.pageSize = 10;
	emit('pageChange', state.page);
};
// 打印
const onPrintTable = () => {
	// https://printjs.crabbly.com/#documentation
	// 自定义打印
	let tableTh = '';
	let tableTrTd = '';
	let tableTd = {};
	// 表头
	props.header.forEach((v) => {
		tableTh += `<th class="table-th">${v.title}</th>`;
	});
	// 表格内容
	props.data.forEach((val, key) => {
		if (!tableTd[key]) tableTd[key] = [];
		props.header.forEach((v) => {
			if (v.type === 'text') {
				tableTd[key].push(`<td class="table-th table-center">${val[v.key]}</td>`);
			} else if (v.type === 'image') {
				tableTd[key].push(`<td class="table-th table-center"><img src="${val[v.key]}" style="width:${v.width}px;height:${v.height}px;"/></td>`);
			}
		});
		tableTrTd += `<tr>${tableTd[key].join('')}</tr>`;
	});
	// 打印
	printJs({
		printable: `<div style=display:flex;flex-direction:column;text-align:center><h3>${props.printName}</h3></div><table border=1 cellspacing=0><tr>${tableTh}${tableTrTd}</table>`,
		type: 'raw-html',
		css: ['//at.alicdn.com/t/c/font_2298093_rnp72ifj3ba.css', '//unpkg.com/element-plus/dist/index.css'],
		style: `@media print{.mb15{margin-bottom:15px;}.el-button--small i.iconfont{font-size: 12px !important;margin-right: 5px;}}; .table-th{word-break: break-all;white-space: pre-wrap;}.table-center{text-align: center;}`,
	});
};
// 导出
const onImportTable = () => {
	if (state.selectList.length <= 0) return ElMessage.warning('请先选择要导出的数据');
	table2excel(props.header, state.selectList, `${themeConfig.value.globalTitle} ${new Date().toLocaleString()}`);
};
// 刷新
const onRefreshTable = () => {
	emit('pageChange', state.page);
};
// 设置
const onSetTable = () => {
	nextTick(() => {
		const sortable = Sortable.create(toolSetRef.value, {
			handle: '.handle',
			dataIdAttr: 'data-key',
			animation: 150,
			onEnd: () => {
				const headerList = [];
				sortable.toArray().forEach((val) => {
					props.header.forEach((v) => {
						if (v.key === val) headerList.push({ ...v });
					});
				});
				emit('sortHeader', headerList);
			},
		});
	});
};

// 暴露变量
defineExpose({
	pageReset,
});
</script>

<style scoped lang="scss">
.table-container {
	display: flex;
	flex-direction: column;
  .table-tools-bar {
    padding-bottom: 10px;
  }
  .table-header-right-tool {
    display: flex;
    align-items: center;
    justify-content: flex-end;
    i {
      margin-right: 10px;
      cursor: pointer;
      color: var(--el-text-color-regular);
      &:last-of-type {
        margin-right: 0;
      }
    }
  }
	.el-table {
		flex: 1;
	}
	.table-footer {
		display: flex;
	}
}
</style>
