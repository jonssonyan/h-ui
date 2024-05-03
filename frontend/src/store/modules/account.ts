import { defineStore } from "pinia";

import { getAccountInfoApi, loginApi } from "@/api/account";
import { resetRouter } from "@/router";
import { store } from "@/store";

import { AccountInfo, AccountLoginDto } from "@/api/account/types";

import { useStorage } from "@vueuse/core";

export const useAccountStore = defineStore("account", () => {
  // state
  const token = useStorage("accessToken", "");
  const id = ref(0);
  const username = ref("");
  const roles = ref<Array<string>>([]); // 用户角色编码集合 → 判断路由权限

  /**
   * 登录
   *
   * @returns
   */
  function login(accountLoginDto: AccountLoginDto) {
    return new Promise<void>((resolve, reject) => {
      loginApi(accountLoginDto)
        .then((response) => {
          const { tokenType, accessToken } = response.data;
          token.value = tokenType + " " + accessToken; // Bearer eyJhbGciOiJIUzI1NiJ9.xxx.xxx
          resolve();
        })
        .catch((error) => {
          reject(error);
        });
    });
  }

  // 查询当前
  function getAccountInfo() {
    return new Promise<AccountInfo>((resolve, reject) => {
      getAccountInfoApi()
        .then(({ data }) => {
          if (!data) {
            return reject("Verification failed, please Login again.");
          }
          if (!data.roles || data.roles.length <= 0) {
            reject("getAccountInfoApi: roles must be a non-null array!");
          }
          id.value = data.id;
          username.value = data.username;
          roles.value = data.roles;
          resolve(data);
        })
        .catch((error) => {
          reject(error);
        });
    });
  }

  // 注销
  function logout() {
    return new Promise<void>((resolve, reject) => {
      resetRouter();
      resetToken();
      resolve();
    });
  }

  // 重置
  function resetToken() {
    token.value = "";
    id.value = 0;
    username.value = "";
    roles.value = [];
  }

  return {
    token,
    id,
    username,
    roles,
    login,
    getAccountInfo,
    logout,
    resetToken,
  };
});

// 非setup
export function useAccountStoreHook() {
  return useAccountStore(store);
}
