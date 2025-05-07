import "normalize.css";
import "./assets/css/index.less";

import { createApp } from "vue";
import App from "./App.vue";

import router from "./router";

import registerIcons from "./utils/register-icon";

const app = createApp(App);
app.use(registerIcons);
app.use(router);
app.mount("#app");
