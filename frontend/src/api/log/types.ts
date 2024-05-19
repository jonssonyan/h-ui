export interface LogSystemDto {
  numLine: number;
}

export interface LogSystemVo {
  clientIp: string;
  latencyTime: string;
  level: string;
  msg: string;
  reqMethod: string;
  reqUri: string;
  statusCode: string;
  time: string;
}
