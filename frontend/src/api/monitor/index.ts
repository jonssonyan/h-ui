import { AxiosPromise } from "axios";
import request from "@/utils/request";
import { SystemMonitorVo } from "@/api/monitor/types";

export function monitorSystemApi(): AxiosPromise<SystemMonitorVo> {
  return request({
    url: "/monitor/system",
    method: "get",
  });
}
