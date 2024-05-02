/**
 * 将时间戳转换为格式化日期时间字符串（YYYY-MM-DD HH:mm:ss）
 * @param timestamp 时间戳
 * @returns 格式化日期时间字符串
 */
export const timestampToDateTime = (timestamp: number): string => {
  const date = new Date(timestamp);
  const year = date.getFullYear();
  const month = (date.getMonth() + 1).toString().padStart(2, "0");
  const day = date.getDate().toString().padStart(2, "0");
  const hours = date.getHours().toString().padStart(2, "0");
  const minutes = date.getMinutes().toString().padStart(2, "0");
  const seconds = date.getSeconds().toString().padStart(2, "0");

  return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
};

/**
 * 获取一周后的时间戳
 * @returns 一周后的时间戳
 */
export const getWeekLater = (): number => {
  const date = new Date();
  date.setDate(date.getDate() + 7);
  return date.getTime();
};

/**
 * 获取一个月后的时间戳
 * @returns 一个月后的时间戳
 */
export const getMonthLater = (): number => {
  const date = new Date();
  date.setMonth(date.getMonth() + 1);
  return date.getTime();
};

/**
 * 获取一年后的时间戳
 * @returns 一年后的时间戳
 */
export const getYearLater = (): number => {
  const date = new Date();
  date.setFullYear(date.getFullYear() + 1);
  return date.getTime();
};
