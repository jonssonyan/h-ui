import { AxiosPromise } from "axios";
import request from "@/utils/request";
import {
  LogExportDto,
  LogHysteria2Vo,
  LogSystemDto,
  LogSystemVo,
} from "@/api/log/types";

export function logSystemApi(
  params: LogSystemDto
): AxiosPromise<PageVo<LogSystemVo>> {
  return request({
    url: "/log/logSystem",
    method: "get",
    params: params,
  });
}

export function logHysteria2Api(
  params: LogSystemDto
): AxiosPromise<PageVo<LogHysteria2Vo>> {
  return request({
    url: "/log/logHysteria2",
    method: "get",
    params: params,
  });
}

export function exportLogApi(params: LogExportDto): AxiosPromise {
  return request({
    url: "/log/exportLog",
    method: "post",
    params,
    responseType: "blob",
  });
}
