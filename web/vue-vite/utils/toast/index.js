import "lbl-toast/dist/index.css";
import Toast from "lbl-toast";

let body = document.querySelector('body');

const toast = {
    loading: function (style = 'style2') {
        body.style.opacity = '0.5';
        return Toast.loading({
            style
        });
    },
    clear: function () {
        body.style.opacity = '1';
        return Toast.clear();
    }
};

export default toast;