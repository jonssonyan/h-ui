export interface AccountPageDto extends BaseDto {
  username: string;
}

export interface AccountUpdateProfileDto {
  username: string;
  newPass: string;
  oldPass: string;
  email: string;
}

export interface AccountCreateDto {
  username: string;
  pass: string;
  roleId: number;
  email: string;
  expireTime: number;
  deleted: number;
}

export interface AccountUpdateDto extends RequiredIdDto {
  username: string;
  pass: string;
  roleId: number;
  email: string;
  expireTime: number;
  deleted: number;
}

export interface AccountLoginDto extends CaptureDto {
  username: string;
  pass: string;
}

export interface AccountRegisterDto extends CaptureDto {
  username: string;
  pass: string;
}

export interface CaptureDto {
  captchaId?: string;
  captchaCode?: string;
}

export interface AccountVo {
  id: number | undefined;
  username: string;
  email: string;
  roleId: number;
  expireTime: number;
  deleted: number;
  createTime: string;
  roles: string[];
}

export interface AccountPageVo extends BaseVoPage {
  accounts: AccountVo[];
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

export interface AccountForm extends RequiredIdDto {
  username: string;
  pass: string;
  roleId: number;
  email: string;
  expireTime: number;
  deleted: number;
}
