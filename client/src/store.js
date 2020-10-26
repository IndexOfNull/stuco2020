import { createStore } from "vuex";
import globalSettings from "./globalSettings";

export default createStore({
  state() {
    return {
      redirectError: "",
      codeInput: "",
      code: Object,
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
    codeInput(state, code) {
      state.codeInput = code;
    },
    pushCode(state, code) {
      state.code = code;
    },
    codeUsed(state) {
      state.code.times_used += 1;
    }
  }
});
