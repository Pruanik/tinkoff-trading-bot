import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";

import "./../node_modules/uikit/dist/css/uikit.min.css";
import "./../node_modules/uikit/dist/js/uikit.min.js";
import "./../node_modules/uikit/dist/js/uikit-icons.min.js";

const app = createApp(App);

app.use(router);

app.mount("#app");
