import {alert, defaultModules} from '@pnotify/core';
import * as PNotifyConfirm from '@pnotify/confirm';
import * as PNotify from '@pnotify/core';
// import '@pnotify/core/dist/PNotify.css';
// import '@pnotify/core/dist/BrightTheme.css';
import '@pnotify/countdown/dist/PNotifyCountdown.css';

defaultModules.set(PNotifyConfirm, {});

// pnotify 弹窗插件
// export default function pnotify(text: string, type: "success" | "notice" | "info" | "error", title?: string, delay?: number) {
//     let opts = {
//         text: text ? text : '',
//         type: type ? type : 'success',
//         title: title ? title : '提示信息',
//         // 时间
//         delay: delay ? delay : 2000,
//     };

//     return alert(opts);
// }

/*export default function pnotifyConfirm(text = 'Are you sure?') {
    return new Promise((resolve, reject) => {
        let notice = alert({
            text: text,
            // icon: 'fas fa-question-circle',
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
                }]
            ])
        });

        notice.on('pnotify:confirm', () => {
            resolve(true); // 解决Promise并传递true
        });

        notice.on('pnotify:cancel', () => {
            resolve(false); // 解决Promise并传递false
        });
    });
}*/

export default function pnotifyConfirm(text = 'Are you sure?') {
    return new Promise((resolve, reject) => {
        alert({
            text: text,
            // icon: 'fas fa-question-circle',
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
        });
    });
}