/**
 * 格式化字节大小
 * @param bytes 字节数
 * @param decimals 小数位数，默认为 2
 * @returns 格式化后的字节大小字符串
 */
export const formatBytes = (bytes: number, decimals = 2): string => {
  // 检查是否为特殊值
  if (bytes === -1) {
    return "Unlimited";
  }
  if (bytes === 0) {
    return "0 Bytes";
  }

  // 计算单位和大小
  const k = 1024;
  const dm = decimals < 0 ? 0 : decimals;
  const sizes = ["Bytes", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"];
  const i = Math.floor(Math.log(bytes) / Math.log(k));

  // 返回格式化后的字符串
  return parseFloat((bytes / Math.pow(k, i)).toFixed(dm)) + " " + sizes[i];
};

export const calculateBytes = (value = 0, unit = "Bytes"): number => {
  // 将单位转换为大写，并去除空格
  const formattedUnit = unit.toUpperCase().trim();

  // 定义存储单位和对应的字节数的映射关系
  const unitToBytes: Record<string, number> = {
    BYTES: 1,
    KB: 1024 ** 1,
    MB: 1024 ** 2,
    GB: 1024 ** 3,
    TB: 1024 ** 4,
    PB: 1024 ** 5,
    EB: 1024 ** 6,
    ZB: 1024 ** 7,
    YB: 1024 ** 8,
  };

  // 检查传入的单位是否存在于映射关系中
  if (!Object.prototype.hasOwnProperty.call(unitToBytes, formattedUnit)) {
    throw new Error("Invalid unit");
  }

  if (value == -1) {
    return -1;
  }

  // 计算并返回字节数
  return value * unitToBytes[formattedUnit];
};

/**
 * 格式化存储容量单位
 * @param bytes 存储容量（字节数）
 * @param decimals 小数位数，默认为 2
 * @returns 格式化后的存储容量值
 */
export const formatStorageCapacity = (bytes: number, decimals = 2): number => {
  // 检查输入是否有效
  if (!bytes || bytes <= 0) {
    return bytes;
  }

  // 计算存储单位
  const k = 1024;
  const dm = decimals < 0 ? 0 : decimals;
  const i = Math.floor(Math.log(bytes) / Math.log(k));

  // 格式化存储容量值并返回
  return parseFloat((bytes / Math.pow(k, i)).toFixed(dm));
};

/**
 * 格式化存储容量单位
 * @param bytes 存储容量（字节数）
 * @returns 格式化后的存储单位
 */
export const formatStorageUnit = (bytes: number): string => {
  // 检查输入是否有效
  if (!bytes || bytes <= 0) {
    return "Bytes";
  }

  // 计算存储单位
  const k = 1024;
  const sizes = ["Bytes", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"];
  const i = Math.floor(Math.log(bytes) / Math.log(k));

  // 返回格式化后的存储单位
  return sizes[i];
};
