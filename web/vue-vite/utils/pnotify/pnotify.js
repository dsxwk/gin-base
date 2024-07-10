import {alert, defaultModules} from '@pnotify/core';
import * as PNotifyCountdown from '@pnotify/countdown';
import '@pnotify/core/dist/PNotify.css';
import '@pnotify/core/dist/BrightTheme.css';

defaultModules.set(PNotifyCountdown, {});

// pnotify 弹窗插件
export default function pnotify(text, type, title, delay) {
    let opts = {
        text: text ? text : '',
        type: type ? type : 'success',
        title: title ? title : '提示信息',
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
}