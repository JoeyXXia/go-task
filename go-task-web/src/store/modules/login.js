import { defineStore } from "pinia";

const useLonginStore = defineStore("login", {
  state: () => ({
    token: "",
    userInfo: {},
    userMenus: [],
  }),
  actions: {},
});

export default useLonginStore;
