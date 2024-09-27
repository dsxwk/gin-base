import Functions from '@/utils/functions';

const funcs = new Functions();
/**
 * @description: 数据来源
 */
export const dataSourceEnum = [
    {
        label: funcs.lang('Article Library'),
        value: 1
    },
    {
        label: funcs.lang('Self Built'),
        value: 2
    }
];

/**
 * @description: 是否发布
 */
export const isPublishEnum = [
    {
        label: funcs.lang('Published'),
        value: 1
    },
    {
        label: funcs.lang('Unpublished'),
        value: 2
    }
];