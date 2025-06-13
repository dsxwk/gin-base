// 通用函数
import useClipboard from 'vue-clipboard3';
import { ElMessage } from 'element-plus';
import { formatDate } from '/@/utils/formatTime';
import { useI18n } from 'vue-i18n';

export default function () {
	const { t } = useI18n();
	const { toClipboard } = useClipboard();

	// 百分比格式化
	const percentFormat = (row, column, cellValue) => {
		return cellValue ? `${cellValue}%` : '-';
	};
	// 列表日期时间格式化
	const dateFormatYMD = (row, column, cellValue) => {
		if (!cellValue) return '-';
		return formatDate(new Date(cellValue), 'YYYY-mm-dd');
	};
	// 列表日期时间格式化
	const dateFormatYMDHMS = (row, column, cellValue) => {
		if (!cellValue) return '-';
		return formatDate(new Date(cellValue), 'YYYY-mm-dd HH:MM:SS');
	};
	// 列表日期时间格式化
	const dateFormatHMS = (row, column, cellValue) => {
		if (!cellValue) return '-';
		let time = 0;
		if (typeof row === 'number') time = row;
		if (typeof cellValue === 'number') time = cellValue;
		return formatDate(new Date(time * 1000), 'HH:MM:SS');
	};
	// 小数格式化
	const scaleFormat = (value = '0', scale = 4) => {
		return Number.parseFloat(value).toFixed(scale);
	};
	// 小数格式化
	const scale2Format = (value = '0') => {
		return Number.parseFloat(value).toFixed(2);
	};
	// 点击复制文本
	const copyText = (text) => {
		return new Promise((resolve, reject) => {
			try {
				//复制
				toClipboard(text);
				//下面可以设置复制成功的提示框等操作
				ElMessage.success(t('message.layout.copyTextSuccess'));
				resolve(text);
			} catch (e) {
				//复制失败
				ElMessage.error(t('message.layout.copyTextError'));
				reject(e);
			}
		});
	};
	return {
		percentFormat,
		dateFormatYMD,
		dateFormatYMDHMS,
		dateFormatHMS,
		scaleFormat,
		scale2Format,
		copyText,
	};
}

export function sprintf(template, ...args) {
	return template.replace(/%s/g, () => args.shift());
}

export function arrayColumnDeep(array, valueKey = null, indexKey = null, unique = false) {
	const result = [];

	for (const item of array) {
		const values = valueKey !== null ? getDeepValues(item, valueKey.split('.')) : [item];
		const key = indexKey !== null ? getNestedValue(item, indexKey) : null;

		if (indexKey !== null && (typeof key === 'string' || typeof key === 'number')) {
			result[key] = values;
		} else {
			result.push(...values);
		}
	}

	return unique ? [...new Set(result.flat(Infinity))] : result;
}

function getNestedValue(obj, path) {
	const keys = typeof path === 'string' ? path.split('.') : path;
	let current = obj;
	for (const key of keys) {
		if (current && typeof current === 'object' && key in current) {
			current = current[key];
		} else {
			return undefined;
		}
	}
	return current;
}

function getDeepValues(obj, pathParts) {
	const key = pathParts[0];
	const rest = pathParts.slice(1);

	if (Array.isArray(obj)) {
		return obj.flatMap(o => getDeepValues(o, pathParts));
	}

	if (typeof obj !== 'object' || obj === null) return [];

	if (!(key in obj)) return [];

	const sub = obj[key];

	if (rest.length === 0) {
		return Array.isArray(sub) ? sub : [sub];
	}

	return getDeepValues(sub, rest);
}
