import { AxiosPromise } from "axios";
import { Hysteria2ServerConfig } from "@/api/config/types";
import request from "@/utils/request";
import {
  Hysteria2KickDto,
  Hysteria2SubscribeUrlDto,
  Hysteria2UrlDto,
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

export function hysteria2ChangeVersionApi(
  data: Hysteria2VersionDto
): AxiosPromise {
  return request({
    url: "/hysteria2/hysteria2ChangeVersion",
    method: "post",
    data: data,
  });
}

export function listReleaseApi(): AxiosPromise<string[]> {
  return request({
    url: "/hysteria2/listRelease",
    method: "get",
  });
}

export function hysteria2SubscribeUrlApi(dto: Hysteria2SubscribeUrlDto): AxiosPromise<string> {
	return request({
		url: "/hysteria2/hysteria2SubscribeUrl",
		method: "get",
		params: dto,
	});
}

export function hysteria2UrlApi(dto: Hysteria2UrlDto): AxiosPromise<string> {
  return request({
    url: "/hysteria2/hysteria2Url",
    method: "get",
    params: dto,
  });
}
