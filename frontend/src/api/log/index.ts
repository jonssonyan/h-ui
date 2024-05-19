import { AxiosPromise } from "axios";
import { SystemMonitorVo } from "@/api/monitor/types";
import request from "@/utils/request";

export function logSystemApi(): AxiosPromise<SystemMonitorVo> {
  return request({
    url: "/log/system",
    method: "get",
  });
}

export function logHysteria2Api(): AxiosPromise<SystemMonitorVo> {
  return request({
    url: "/log/hysteria2",
    method: "get",
  });
}
