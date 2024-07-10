import "lbl-toast/dist/index.css";
import Toast from "lbl-toast";

const toast = {
    loading: function (style = 'style2') {
        return Toast.loading({
            style
        });
    },
    clear: function () {
        return Toast.clear();
    }
};
export default toast