import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";

export const Layout = () => import("@/layout/index.vue");

// 静态路由
export const constantRoutes: RouteRecordRaw[] = [
  {
    path: "/redirect",
    component: Layout,
    meta: { hidden: true },
    children: [
      {
        path: "/redirect/:path(.*)",
        component: () => import("@/views/redirect/index.vue"),
      },
    ],
  },

  {
    path: "/login",
    component: () => import("@/views/login/index.vue"),
    meta: { hidden: true },
  },

  {
    path: "/register",
    component: () => import("@/views/register/index.vue"),
    meta: { hidden: true },
  },

  {
    path: "/",
    component: Layout,
    redirect: "/dashboard",
    children: [
      {
        path: "dashboard",
        component: () => import("@/views/dashboard/index.vue"),
        name: "Dashboard",
        meta: { title: "dashboard", icon: "homepage", affix: true },
      },
      {
        path: "401",
        component: () => import("@/views/error-page/401.vue"),
        meta: { hidden: true },
      },
      {
        path: "404",
        component: () => import("@/views/error-page/404.vue"),
        meta: { hidden: true },
      },
    ],
  },
];

export const asyncRoutes: any[] = [
  {
    path: "/account",
    component: "Layout",
    redirect: "/list",
    name: "Account",
    meta: { title: "account", icon: "user", roles: ["admin"] },
    children: [
      {
        path: "list",
        component: "account/list/index",
        name: "AccountList",
        meta: {
          title: "accountList",
          icon: "user",
          roles: ["admin"],
        },
      },
    ],
  },
  {
    path: "/hysteria",
    component: "Layout",
    redirect: "/list",
    name: "Hysteria",
    meta: { title: "hysteria", icon: "hysteria", roles: ["admin"] },
    children: [
      {
        path: "list",
        component: "hysteria/list/index",
        name: "HysteriaList",
        meta: {
          title: "hysteriaList",
          icon: "hysteria",
          roles: ["admin"],
        },
      },
    ],
  },
  {
    path: "/config",
    component: "Layout",
    redirect: "/list",
    name: "Config",
    meta: { title: "config", icon: "config", roles: ["admin"] },
    children: [
      {
        path: "list",
        component: "config/list/index",
        name: "ConfigList",
        meta: {
          title: "configList",
          icon: "config",
          roles: ["admin"],
        },
      },
    ],
  },
  {
    path: "/monitor",
    component: "Layout",
    redirect: "/monitor",
    name: "Monitor",
    meta: { title: "monitor", icon: "monitor", roles: ["admin"] },
    children: [
      {
        path: "system",
        component: "monitor/system/index",
        name: "MonitorSystem",
        meta: {
          title: "monitorSystem",
          icon: "monitor",
          roles: ["admin"],
        },
      },
    ],
  },
  {
    path: "/log",
    component: "Layout",
    redirect: "/system",
    name: "Log",
    meta: { title: "log", icon: "log-system", roles: ["admin"] },
    children: [
      {
        path: "system",
        component: "log/system/index",
        name: "LogSystem",
        meta: {
          title: "logSystem",
          icon: "log-system",
          roles: ["admin"],
        },
      },
    ],
  },
  {
    path: "/log",
    component: "Layout",
    redirect: "/hysteria",
    name: "Log",
    meta: { title: "log", icon: "log-hysteria", roles: ["admin"] },
    children: [
      {
        path: "hysteria",
        component: "log/hysteria/index",
        name: "LogHysteria",
        meta: {
          title: "logHysteria",
          icon: "log-hysteria",
          roles: ["admin"],
        },
      },
    ],
  },
  {
    path: "/info",
    component: "Layout",
    redirect: "/account",
    name: "Info",
    meta: { title: "info", icon: "info-account", roles: ["user", "admin"] },
    children: [
      {
        path: "account",
        component: "info/account/index",
        name: "AccountInfo",
        meta: {
          title: "infoAccount",
          icon: "info-account",
          roles: ["user", "admin"],
        },
      },
    ],
  },
];

/**
 * 创建路由
 */
const router = createRouter({
  history: createWebHashHistory(),
  routes: constantRoutes as RouteRecordRaw[],
  // 刷新时，滚动条位置还原
  scrollBehavior: () => ({ left: 0, top: 0 }),
});

/**
 * 重置路由
 */
export function resetRouter() {
  router.replace({ path: "/login" });
  location.reload();
}

export default router;
