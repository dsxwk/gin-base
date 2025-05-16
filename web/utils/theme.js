/**
 * 颜色转换函数
 * @method hexToRgb hex 颜色转 rgb 颜色
 * @method rgbToHex rgb 颜色转 Hex 颜色
 * @method getDarkColor 加深颜色值
 * @method getLightColor 变浅颜色值
 */
export function useChangeColor() {
	// 公共正则表达式
	const hexReg = /^\#?[0-9A-Fa-f]{6}$/;
	const rgbReg = /^\d{1,3}$/;

	// hex 颜色转 rgb 颜色
	const hexToRgb = (str) => {
		if (!hexReg.test(str)) {
			console.warn('输入错误的 hex 颜色值');
			return null;
		}
		str = str.replace('#', '');
		const hexs = str.match(/../g).map((hex) => parseInt(hex, 16));
		return hexs;
	};

	// rgb 颜色转 Hex 颜色
	const rgbToHex = (r, g, b) => {
		if (!rgbReg.test(r.toString()) || !rgbReg.test(g.toString()) || !rgbReg.test(b.toString())) {
			console.warn('输入错误的 rgb 颜色值');
			return '';
		}
		const hexs = [r, g, b].map((num) => num.toString(16).padStart(2, '0'));
		return `#${hexs.join('')}`;
	};

	// 加深颜色值
	const getDarkColor = (color, level) => {
		if (!hexReg.test(color)) {
			console.warn('输入错误的 hex 颜色值');
			return '';
		}
		const rgb = hexToRgb(color);
		if (!rgb) return '';
		const darkenedRgb = rgb.map((value) => Math.floor(value * (1 - level)));
		return rgbToHex(darkenedRgb[0], darkenedRgb[1], darkenedRgb[2]);
	};

	// 变浅颜色值
	const getLightColor = (color, level) => {
		if (!hexReg.test(color)) {
			console.warn('输入错误的 hex 颜色值');
			return '';
		}
		const rgb = hexToRgb(color);
		if (!rgb) return '';
		const lightenedRgb = rgb.map((value) => Math.floor((255 - value) * level + value));
		return rgbToHex(lightenedRgb[0], lightenedRgb[1], lightenedRgb[2]);
	};

	return {
		hexToRgb,
		rgbToHex,
		getDarkColor,
		getLightColor,
	};
}
