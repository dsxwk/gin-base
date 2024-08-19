import {alert, defaultModules} from '@pnotify/core';
import * as PNotifyCountdown from '@pnotify/countdown';
import '@pnotify/core/dist/PNotify.css';
import '@pnotify/core/dist/BrightTheme.css';
import Functions from '@/utils/functions';

defaultModules.set(PNotifyCountdown, {});

const funcs = new Functions();
// pnotify 弹窗插件
const pnotify = function pnotify(text, type, title, delay) {
    let opts = {
        text: text ? text : '',
        type: type ? type : 'success',
        title: title ? title : funcs.lang('Reminder Information'),
        // 时间
        delay: delay ? delay : 2000,
        modules: new Map([
            ...defaultModules,
            [
                PNotifyCountdown,
                {
                    // Countdown Module Options
                },
            ],
        ])
    };

    return alert(opts);
};

export default pnotify;