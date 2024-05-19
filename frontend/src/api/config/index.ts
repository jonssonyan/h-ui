import { AxiosPromise } from "axios";
import request from "@/utils/request";
import {
  ConfigDto,
  ConfigsDto,
  ConfigUpdateDto,
  ConfigVo,
  Hysteria2ServerConfig,
} from "@/api/config/types";

export function getHysteria2ConfigApi(): AxiosPromise<Hysteria2ServerConfig> {
  return request({
    url: "/config/getHysteria2Config",
    method: "get",
  });
}

export function updateHysteria2ConfigApi(
  data: Hysteria2ServerConfig
): AxiosPromise {
  return request({
    url: "/config/updateHysteria2Config",
    method: "post",
    data,
  });
}

export function getConfigApi(data: ConfigDto): AxiosPromise<ConfigVo> {
  return request({
    url: "/config/getConfig",
    method: "get",
    params: data,
  });
}

export function listConfigApi(data: ConfigsDto): AxiosPromise<Array<ConfigVo>> {
  return request({
    url: "/config/listConfig",
    method: "post",
    data,
  });
}

export function updateConfigsApi(data: ConfigUpdateDto): AxiosPromise {
  return request({
    url: "/config/updateConfigs",
    method: "post",
    data,
  });
}
