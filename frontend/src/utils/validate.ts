export const validateNumberStr = (rule: any, value: any, callback: any) => {
  if (!value) {
    callback(new Error("field is required"));
  } else if (!/^\d+$/.test(value)) {
    callback(new Error("field must be a number"));
  } else {
    callback();
  }
};
