import {alert, defaultModules} from '@pnotify/core';
import * as PNotifyConfirm from '@pnotify/confirm';
import * as PNotify from '@pnotify/core';
import '@pnotify/countdown/dist/PNotifyCountdown.css';
import Functions from '@/utils/functions';

defaultModules.set(PNotifyConfirm, {});

const funcs = new Functions();
const pnotifyConfirm = function pnotifyConfirm(text = 'Are you sure?', type, title, delay) {
    return new Promise((resolve, reject) => {
        alert({
            text: text,
            // icon: 'fas fa-question-circle',
            type: type ? type : 'notice',
            title: title ? title : funcs.lang('Reminder Information'),
            // 时间
            delay: delay ? delay : 2000,
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
                            'text': funcs.lang('Confirm'),
                            click: notice => {
                                notice.close();
                                resolve(true);
                            }
                        },
                        {
                            'text': funcs.lang('Cancel'),
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