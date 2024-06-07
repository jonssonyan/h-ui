export interface LogDto {
  numLine: number;
}

export interface LogExportDto {
  option: number;
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

export interface LogHysteria2Vo {
  level: string;
  msg: string;
  time: string;
}
