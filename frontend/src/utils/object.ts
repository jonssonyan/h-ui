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
