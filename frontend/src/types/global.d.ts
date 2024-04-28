declare global {
  interface IdDto {
    id: number;
  }

  interface BaseDto {
    pageNum: number;
    pageSize: number;
    startTime?: string;
    endTime?: string;
  }

  interface PageVo<T> {
    records: T[];
    total: number;
  }
}
export {};
