/**
 * 浅拷贝 忽略 null 和 undefined
 * @param target
 * @param sources
 */
export const assignIgnoringNull = <T extends object>(
  target: T,
  ...sources: (Partial<T> | null | undefined)[]
): T => {
  sources.forEach((source) => {
    if (source) {
      Object.keys(source).forEach((key) => {
        const value = source[key as keyof T];
        if (value !== null && value !== undefined) {
          if (typeof value === "object" && !Array.isArray(value)) {
            // 如果值是对象，则递归调用 assignIgnoringNull
            target[key as keyof T] = assignIgnoringNull(
              {} as any, // 使用 {} as any 来避免类型错误，实际运行时应为 T[keyof T]
              value as Partial<any>
            );
          } else {
            target[key as keyof T] = value as any;
          }
        }
      });
    }
  });
  return target;
};

/**
 * 深拷贝
 * @param obj
 */
export const deepCopy = <T>(obj: T): T => {
  if (obj === null || typeof obj !== "object") {
    return obj;
  }

  if (Array.isArray(obj)) {
    const arrCopy = [] as any[];
    obj.forEach((item, index) => {
      arrCopy[index] = deepCopy(item);
    });
    return arrCopy as any;
  }

  const objCopy = {} as { [key: string]: any };
  Object.keys(obj).forEach((key) => {
    objCopy[key] = deepCopy((obj as { [key: string]: any })[key]);
  });

  return objCopy as T;
};
