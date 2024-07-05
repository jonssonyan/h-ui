import lodash from "lodash";

/**
 *  忽略 null 和 undefined
 * @param value
 */
const customizerCloneDeepWith = (value: any) => {
  if (value === null || value === undefined) {
    return undefined;
  }
};

/**
 * 忽略 null 和 undefined
 * @param objValue
 * @param srcValue
 */
const customizerAssignWith = (objValue: any, srcValue: any): any => {
  if (srcValue === null || srcValue === undefined) {
    return objValue;
  }
  return undefined;
};

/**
 * 浅拷贝
 * @param target
 * @param source
 */
export const assignWith = <T, U>(target: T, source: U): void => {
  lodash.assignWith(target, source, customizerAssignWith);
};

/**
 * 深拷贝
 * @param obj
 */
export const cloneDeepWith = <T>(obj: T): T => {
  return lodash.cloneDeepWith(obj, customizerCloneDeepWith);
};
