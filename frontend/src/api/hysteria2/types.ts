import { hysteria2UrlApi } from "@/api/hysteria2/index";

export interface Hysteria2KickDto {
  ids: number[];
  kickUtilTime: number;
}

export interface Hysteria2VersionDto {
  version: string;
}

export interface Hysteria2ReleaseVo {
  tagName: string;
  browserDownloadURL: string;
}

export interface Hysteria2UrlDto {
  accountId: number;
  hostname: string;
}
