import {alert, defaultModules} from '@pnotify/core';
import * as PNotifyConfirm from '@pnotify/confirm';
import * as PNotify from '@pnotify/core';
import '@pnotify/countdown/dist/PNotifyCountdown.css';

defaultModules.set(PNotifyConfirm, {});

const pnotifyConfirm = function pnotifyConfirm(text = 'Are you sure?') {
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
};

export default pnotifyConfirm;