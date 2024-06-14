/**
 * 浅拷贝 忽略 null 和 undefined
 * @param target
 * @param sources
 */
export function assignIgnoringNull<T extends object>(
  target: T,
  ...sources: (Partial<T> | null | undefined)[]
): T {
  sources.forEach((source) => {
    if (source) {
      Object.keys(source).forEach((key) => {
        const value = source[key as keyof T];
        if (value !== null && value !== undefined) {
          target[key as keyof T] = value;
        }
      });
    }
  });
  return target;
}

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
