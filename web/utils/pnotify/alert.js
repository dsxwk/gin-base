import {alert, defaultModules} from '@pnotify/core';
import * as PNotifyCountdown from '@pnotify/countdown';
import '@pnotify/core/dist/PNotify.css';
import '@pnotify/core/dist/BrightTheme.css';

defaultModules.set(PNotifyCountdown, {});

const defaultDelay = 2000;

function setOpts(text, type, title, delay) {
    return {
        text: text === '' ? 'Success' : text,
        type: type === '' ? 'success' : type,
        title: title === '' ? '提示信息' : title,
        // 时间
        delay: delay ?? 2000,
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
}

class pnotify {
    static success(text = '', title = '', delay = defaultDelay) {
        return alert(setOpts(text, 'success', title, delay));
    };
    static info(text = '', title = '', delay = defaultDelay) {
        return alert(setOpts(text, 'info', title, delay));
    };
    static notice(text = '', title = '', delay = defaultDelay) {
        return alert(setOpts(text, 'notice', title, delay));
    };
    static error(text = '', title = '', delay = defaultDelay) {
        return alert(setOpts(text, 'error', title, delay));
    };
}

export default pnotify;