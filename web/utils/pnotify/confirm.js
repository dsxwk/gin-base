import {alert, defaultModules} from '@pnotify/core';
import * as PNotifyConfirm from '@pnotify/confirm';
import * as PNotify from '@pnotify/core';
import '@pnotify/countdown/dist/PNotifyCountdown.css';

defaultModules.set(PNotifyConfirm, {});

const defaultDelay = 2000;

function setOpts(text, type, title, delay, resolve) {
    return {
        text: text === '' ? 'Are you sure?' : text,
        // icon: 'fas fa-question-circle',
        type: type === '' ? 'notice' : type,
        title: title === '' ? '提示信息' : title,
        // 时间
        delay: delay ?? 2000,
        hide: false,
        closer: false,
        sticker: false,
        destroy: true,
        stack: new PNotify.Stack({
            dir1: 'down',
            modal: true,
            firstpos1: 25,
            overlayClose: false
        }),
        modules: new Map([
            ...PNotify.defaultModules,
            [PNotifyConfirm, {
                confirm: true,
                buttons: [
                    {
                        'text': '确认',
                        click: notice => {
                            notice.close();
                            resolve(true);
                        }
                    },
                    {
                        'text': '取消',
                        click: notice => {
                            notice.close();
                            resolve(false);
                        }
                    }
                ]
            }]
        ])
    };
}

class pnotifyConfirm {
    static success(text = '', title = '', delay = defaultDelay) {
        return new Promise((resolve, reject) => {
            alert(setOpts(text, 'success', title, delay, resolve));
        });
    };
    static info(text = '', title = '', delay = defaultDelay) {
        return new Promise((resolve, reject) => {
            alert(setOpts(text, 'info', title, delay, resolve));
        });
    };
    static notice(text = '', title = '', delay = defaultDelay) {
        return new Promise((resolve, reject) => {
            alert(setOpts(text, 'notice', title, delay, resolve));
        });
    };
    static error(text = '', title = '', delay = defaultDelay) {
        return new Promise((resolve, reject) => {
            alert(setOpts(text, 'error', title, delay, resolve));
        });
    };
}

export default pnotifyConfirm;