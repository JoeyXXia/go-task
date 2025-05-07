import { createRouter, createWebHistory } from "vue-router";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: "/",
      redirect: "/main",
    },
    {
      path: "/login",
      component: () => import("../pages/login/Login.vue"),
    },
    {
      path: "/main",
      component: () => import("../pages/main/Main.vue"),
    },
    {
      path: "/:pathMatch(.*)*",
      component: () => import("../pages/not-found/NotFound.vue"),
    },
  ],
});

export default router;
