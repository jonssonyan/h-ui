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

export const calculateTimeDifference = (timestamp: number): string => {
  const now = Date.now();
  const diff = timestamp - now;

  if (diff <= 0) {
    return "-";
  }

  const seconds = Math.floor(diff / 1000);
  const minutes = Math.floor(seconds / 60);
  const hours = Math.floor(minutes / 60);
  const days = Math.floor(hours / 24);

  const remainingHours = hours % 24;
  const remainingMinutes = minutes % 60;
  const remainingSeconds = seconds % 60;

  const parts: string[] = [];

  if (days > 0) {
    parts.push(`${days}天`);
  }
  if (remainingHours > 0) {
    parts.push(`${remainingHours}小时`);
  }
  if (remainingMinutes > 0) {
    parts.push(`${remainingMinutes}分钟`);
  }
  if (remainingSeconds > 0) {
    parts.push(`${remainingSeconds}秒`);
  }

  return parts.join(" ");
};

/**
 * 获取一小时后的时间戳
 * @returns 一周后的时间戳
 */
export const getHourLater = (): number => {
  const date = new Date();
  date.setHours(date.getHours() + 1);
  return date.getTime();
};

/**
 * 获取一天后的时间戳
 * @returns 一周后的时间戳
 */
export const getDayLater = (): number => {
  const date = new Date();
  date.setDate(date.getDate() + 1);
  return date.getTime();
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
