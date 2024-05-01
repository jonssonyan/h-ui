import axios, { InternalAxiosRequestConfig, AxiosResponse } from "axios";
import { useAccountStoreHook } from "@/store/modules/account";

// 创建 axios 实例
const service = axios.create({
  baseURL: import.meta.env.VITE_APP_BASE_API,
  timeout: 50000,
  headers: { "Content-Type": "application/json;charset=utf-8" },
});

// 请求拦截器
service.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    const accountStore = useAccountStoreHook();
    if (accountStore.token) {
      config.headers.Authorization = accountStore.token;
    }
    return config;
  },
  (error: any) => {
    return Promise.reject(error);
  }
);

// 响应拦截器
service.interceptors.response.use(
  (response: AxiosResponse) => {
    const { code, message } = response.data;
    if (code === 20000) {
      return response.data;
    }
    // 响应数据为二进制流处理(文件导出)
    if (response.data instanceof ArrayBuffer || response.data instanceof Blob) {
      return response;
    }

    ElMessage.error(message || "系统出错");
    return Promise.reject(new Error(message || "Error"));
  },
  (error: any) => {
    if (error.response.data) {
      const { code, msg } = error.response.data;
      // token 过期,重新登录
      if (code === "A0230") {
        ElMessageBox.confirm("当前页面已失效，请重新登录", "提示", {
          confirmButtonText: "确定",
          type: "warning",
        }).then(() => {
          localStorage.clear();
          window.location.href = "/";
        });
      } else {
        ElMessage.error(msg || "系统出错");
      }
    }
    return Promise.reject(error.message);
  }
);

// 导出 axios 实例
export default service;
