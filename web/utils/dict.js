export function getDict(dict, value) {
    const item = dict.find(i => i.value === value);
    return item ? item.label : '';
}