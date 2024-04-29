export interface AccountPageDto extends BaseDto {
  username?: string;
  deleted?: number;
}

export interface AccountUpdateDto extends IdDto {
  username: string;
  pass: string;
  conPass: string;
  quota: number;
  expireTime: number;
  deleted: number;
}

export interface AccountSaveDto {
  username: string;
  pass: string;
  conPass: string;
  quota: number;
  expireTime: number;
  deleted: number;
}

export interface AccountLoginDto {
  username: string;
  pass: string;
}

export interface AccountVo extends IdDto {
  username: string;
  quota: number;
  download: number;
  upload: number;
  expireTime: number;
  role: string;
  deleted: number;
  createTime: string;
}

export interface AccountLoginVo {
  accessToken: string;
  tokenType: string;
}

export interface AccountInfo {
  id: number;
  username: string;
  roles: string[];
}

export interface AccountForm extends IdDto {
  username: string;
  pass: string;
  conPass: string;
  quota: number;
  expireTime: number;
  deleted: number;
}
