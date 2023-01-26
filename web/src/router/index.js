import { createRouter, createWebHistory } from "vue-router";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "Dashboard Statistic",
      component: () => import("@/views/DashboardStatisticView.vue"),
    },
    {
      path: "/instruments",
      name: "Instruments",
      component: () => import("@/views/DashboardInstrumentsView.vue"),
    },
    {
      path: "/logs",
      name: "Logs",
      component: () => import("@/views/DashboardLogsView.vue"),
    },
    {
      path: "/:pathMatch(.*)*",
      component: () => import("@/views/PageNotFoundView.vue"),
    },
  ],
});

export default router;
