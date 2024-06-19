import request from "@/utils/request";
import { AxiosPromise } from "axios";
import {
  AccountSaveDto,
  AccountInfo,
  AccountLoginDto,
  AccountLoginVo,
  AccountPageDto,
  AccountUpdateDto,
  AccountVo,
} from "./types";

/**
 * 查询
 */
export function getAccountApi(data: IdDto): AxiosPromise<AccountVo> {
  return request({
    url: "/account/getAccount",
    method: "get",
    params: data,
  });
}

/**
 * 保存
 *
 * @param data
 */
export function saveAccountApi(data: AccountSaveDto): AxiosPromise {
  return request({
    url: "/account/saveAccount",
    method: "post",
    data: data,
  });
}

/**
 * 查询当前
 */
export function getAccountInfoApi(): AxiosPromise<AccountInfo> {
  return request({
    url: "/account/getAccountInfo",
    method: "get",
  });
}

/**
 * 分页
 * @param data
 */
export function pageAccountApi(
  data: AccountPageDto
): AxiosPromise<PageVo<AccountVo>> {
  return request({
    url: "/account/pageAccount",
    method: "get",
    params: data,
  });
}

/**
 * 删除
 *
 * @param data
 */
export function deleteAccountApi(data: IdDto): AxiosPromise {
  return request({
    url: "/account/deleteAccount",
    method: "post",
    data: data,
  });
}

/**
 * 修改
 * @param data
 */
export function updateAccountApi(data: AccountUpdateDto): AxiosPromise {
  return request({
    url: "/account/updateAccount",
    method: "post",
    data: data,
  });
}

/**
 * 重设流量
 * @param data
 */
export function resetTrafficApi(data: IdDto): AxiosPromise {
  return request({
    url: "/account/resetTraffic",
    method: "post",
    data: data,
  });
}

/**
 * 登录
 * @param data
 */
export function loginApi(data: AccountLoginDto): AxiosPromise<AccountLoginVo> {
  return request({
    url: "/auth/login",
    method: "post",
    data: data,
  });
}

/**
 * 导入
 */
export function importAccountApi(data: FormData): AxiosPromise {
  return request({
    url: "/account/importAccount",
    method: "post",
    headers: {
      "Content-Type": "multipart/form-data",
    },
    data: data,
  });
}

/**
 * 导出
 */
export function exportAccountApi(): AxiosPromise {
  return request({
    url: "/account/exportAccount",
    method: "post",
    responseType: "blob",
  });
}

/**
 * 解除下线状态
 */
export function releaseKickAccountApi(data: IdDto): AxiosPromise {
  return request({
    url: "/account/releaseKickAccount",
    method: "post",
    data: data,
  });
}
