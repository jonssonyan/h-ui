/**
 * 浅拷贝，忽略 null，支持嵌套对象
 * @param target
 * @param source
 */
export const assignWith = <T>(target: T, source: Partial<T>): void => {
  if (source === null || typeof source !== "object") {
    return;
  }

  for (const key in source) {
    if (source[key] !== null) {
      if (typeof source[key] === "object") {
        if (!target[key]) {
          target[key] = (Array.isArray(source[key]) ? [] : {}) as T[Extract<
            keyof T,
            string
          >];
        }
        assignWith(target[key] as any, source[key] as any);
      } else {
        target[key] = source[key] as T[Extract<keyof T, string>];
      }
    }
  }
};

/**
 * 深拷贝，忽略 null，支持嵌套对象
 * @param source
 */
export const deepCopy = <T>(source: Partial<T>): T => {
  if (source === null || typeof source !== "object") {
    return source;
  }

  if (Array.isArray(source)) {
    const arrCopy = [] as any[];
    source.forEach((item, index) => {
      arrCopy[index] = deepCopy(item);
    });
    return arrCopy as any;
  }

  const objCopy = {} as { [key: string]: any };
  Object.keys(source).forEach((key) => {
    objCopy[key] = deepCopy((source as { [key: string]: any })[key]);
  });

  return objCopy as T;
};
