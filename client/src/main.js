import { createApp } from "vue";
import { createStore } from "vuex";
import App from "./App.vue";
import router from "./router";

import "@/assets/styles/index.scss";
import "@/assets/styles/tailwind.css";

const store = createStore({
  state () {
    return {
      redirectError: "",
      code: ""
    }
  },
  mutations: {
    error (state, err) {
      state.redirectError = err
    },
    clearerror (state) {
      state.redirectError = ""
    },
    code (state, code) {
      state.code = code
    }
  }
})


createApp(App)
  .use(router).use(store)
  .mount("#app");
