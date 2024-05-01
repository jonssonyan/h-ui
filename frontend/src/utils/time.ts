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

export const getWeekLater = (): number => {
  const date = new Date();
  date.setDate(date.getDate() + 7);
  return date.getTime();
};

export const getMonthLater = (): number => {
  const date = new Date();
  date.setMonth(date.getMonth() + 1);
  return date.getTime();
};

export const getYearLater = (): number => {
  const date = new Date();
  date.setFullYear(date.getFullYear() + 1);
  return date.getTime();
};
