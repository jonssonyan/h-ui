import { AxiosPromise } from "axios";
import { Hysteria2ServerConfig } from "@/api/config/types";
import request from "@/utils/request";
import {
  Hysteria2KickDto,
  Hysteria2ReleaseVo,
  Hysteria2VersionDto,
} from "@/api/hysteria2/types";

export function hysteria2KickApi(
  data: Hysteria2KickDto
): AxiosPromise<Hysteria2ServerConfig> {
  return request({
    url: "/hysteria2/hysteria2Kick",
    method: "post",
    data: data,
  });
}

export function changeHysteria2VersionApi(
  data: Hysteria2VersionDto
): AxiosPromise {
  return request({
    url: "/hysteria2/changeHysteria2Version",
    method: "post",
    data: data,
  });
}

export function listReleaseApi(): AxiosPromise<Hysteria2ReleaseVo[]> {
  return request({
    url: "/hysteria2/listRelease",
    method: "get",
  });
}
