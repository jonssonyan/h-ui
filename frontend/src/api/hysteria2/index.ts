import { AxiosPromise } from "axios";
import { Hysteria2ServerConfig } from "@/api/config/types";
import request from "@/utils/request";
import { Hysteria2KickDto } from "@/api/hysteria2/types";

export function startHysteria2Api(): AxiosPromise<Hysteria2ServerConfig> {
  return request({
    url: "/hysteria2/startHysteria2",
    method: "post",
  });
}

export function stopHysteria2Api(): AxiosPromise<Hysteria2ServerConfig> {
  return request({
    url: "/hysteria2/stopHysteria2",
    method: "post",
  });
}

export function restartHysteria2Api(): AxiosPromise<Hysteria2ServerConfig> {
  return request({
    url: "/hysteria2/restartHysteria2",
    method: "post",
  });
}

export function countOnlineApi(): AxiosPromise<Hysteria2ServerConfig> {
  return request({
    url: "/hysteria2/countOnline",
    method: "get",
  });
}

export function hysteria2KickApi(
  data: Hysteria2KickDto
): AxiosPromise<Hysteria2ServerConfig> {
  return request({
    url: "/hysteria2/hysteria2Kick",
    method: "post",
    data,
  });
}
