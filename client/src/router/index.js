import { createRouter, createWebHistory } from "vue-router";
//import Home from "../views/Home.vue";
import Code from "../views/CodeAsk.vue";

const routes = [
  {
    path: "/",
    name: "Code",
    component: Code
  },
  {
    path: "/info",
    name: "Information",
    // route level code-splitting
    // this generates a separate chunk (info.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: function() {
      return import(/* webpackChunkName: "info" */ "../views/Info.vue");
    }
  }
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
});

export default router;
