import { AxiosPromise } from "axios";
import request from "@/utils/request";
import {
  LogDto,
  LogExportDto,
  LogHysteria2Vo,
  LogSystemVo,
} from "@/api/log/types";

export function logSystemApi(data: LogDto): AxiosPromise<PageVo<LogSystemVo>> {
  return request({
    url: "/log/logSystem",
    method: "get",
    params: data,
  });
}

export function logHysteria2Api(
  data: LogDto
): AxiosPromise<PageVo<LogHysteria2Vo>> {
  return request({
    url: "/log/logHysteria2",
    method: "get",
    params: data,
  });
}

export function exportLogApi(data: LogExportDto): AxiosPromise {
  return request({
    url: "/log/exportLog",
    method: "post",
    data: data,
    responseType: "blob",
  });
}
