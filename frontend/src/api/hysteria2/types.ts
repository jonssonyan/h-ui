export interface Hysteria2KickDto {
  ids: number[];
  kickUtilTime: number;
}

export interface Hysteria2VersionDto {
  version: string;
}

export interface Hysteria2SubscribeUrlDto {
  accountId: number;
  protocol: string;
  host: string;
}

export interface Hysteria2UrlDto {
  accountId: number;
  hostname: string;
}

export interface Hysteria2SubscribeVo {
	url: string;
	qrCode: string;
}


export interface Hysteria2UrlVo {
  url: string;
  qrCode: string;
}
