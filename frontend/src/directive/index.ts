import type { App } from "vue";

import { hasRole } from "./permission";

// 全局注册 directive
export function setupDirective(app: App<Element>) {
  // 使 v-hasRole 在所有组件中都可用
  app.directive("hasRole", hasRole);
}
