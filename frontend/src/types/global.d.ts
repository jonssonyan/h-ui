declare global {
  interface RequiredIdDto {
    id: number;
  }

  interface BaseDto {
    pageNum: number;
    pageSize: number;
    startTime?: string;
    endTime?: string;
  }

  interface BaseVoPage {
    pageNum: number;
    pageSize: number;
    total: number;
  }
}
export {};
