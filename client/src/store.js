import { createStore } from "vuex";
import globalSettings from "./globalSettings";

export default createStore({
  state() {
    return {
      redirectError: "",
      code: "",
      globalSettings: globalSettings
    };
  },
  mutations: {
    error(state, err) {
      state.redirectError = err;
    },
    clearerror(state) {
      state.redirectError = "";
    },
    code(state, code) {
      state.code = code;
    }
  }
});
