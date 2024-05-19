import { AxiosPromise } from "axios";
import request from "@/utils/request";
import { LogSystemDto, LogSystemVo } from "@/api/log/types";

export function logSystemApi(
  params: LogSystemDto
): AxiosPromise<PageVo<LogSystemVo>> {
  return request({
    url: "/log/system",
    method: "get",
    params: params,
  });
}

export function logHysteria2Api(): AxiosPromise {
  return request({
    url: "/log/hysteria2",
    method: "get",
  });
}
