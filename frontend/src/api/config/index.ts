import { AxiosPromise } from "axios";
import request from "@/utils/request";
import { ConfigDto, ConfigVo, Hysteria2ServerConfig } from "@/api/config/types";

export function getHysteria2Config(): AxiosPromise<Hysteria2ServerConfig> {
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

export function getConfig(data: ConfigDto): AxiosPromise<ConfigVo> {
  return request({
    url: "/config/getConfig",
    method: "get",
    params: data,
  });
}
