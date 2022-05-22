import { createRouter, createWebHistory } from "vue-router";
import DashboardView from "../views/DashboardView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "Dashboard",
      component: DashboardView,
    },
    // {
    //   path: "/profile",
    //   name: "Profile",
    //   component: () => import("../views/ProfileView.vue"),
    // },
    {
      path: "/logs",
      name: "Logs",
      component: () => import("../views/LogsView.vue"),
    },
    {
      path: "/:pathMatch(.*)*",
      component: () => import("../views/PageNotFound.vue"),
    },
  ],
});

export default router;
