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
    data,
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
    data,
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
    data,
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
    data,
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
