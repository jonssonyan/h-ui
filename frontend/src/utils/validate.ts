export const validateIntegerStr = (rule: any, value: any, callback: any) => {
  if (!value) {
    callback(new Error("field is required"));
  } else if (!/^\d+$/.test(value)) {
    callback(new Error("field must be a integer"));
  } else {
    callback();
  }
};

export const validateNumberStr = (
  rule: any,
  value: any,
  callback: any
): void => {
  if (!value) {
    callback(new Error("field is required"));
  } else if (!/^\d+(\.\d)?$/.test(value)) {
    callback(new Error("field must be a number with up to one decimal place"));
  } else {
    callback();
  }
};
