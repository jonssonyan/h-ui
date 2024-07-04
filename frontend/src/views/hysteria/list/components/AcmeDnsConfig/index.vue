<template>
  <div class="flex gap-2">
    <el-tag
      :key="key"
      v-for="[key] in acmeDnsConfig"
      @close="handleClose(key)"
      @click="handleInfo(key)"
      size="large"
      closable
    >
      {{ key }}
    </el-tag>

    <el-button @click="handleAdd">+</el-button>

    <el-dialog
      :title="dialog.title"
      v-model="dialog.visible"
      width="600px"
      append-to-body
      @close="closeDialog"
    >
      <el-form
        ref="dataFormRef"
        :rules="dataFormRules"
        label-position="top"
        :model="dataForm"
      >
        <el-form-item label="key" prop="key">
          <el-input v-model="dataForm.key" clearable />
        </el-form-item>
        <el-form-item label="value" prop="value">
          <el-input v-model="dataForm.value" clearable />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="submitForm"
            >{{ $t("common.confirm") }}
          </el-button>
          <el-button @click="closeDialog">{{ $t("common.cancel") }}</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog
      :title="acmeDnsConfigInfoDialog.title"
      v-model="acmeDnsConfigInfoDialog.visible"
      width="600px"
      append-to-body
      @close="acmeDnsConfigInfoDialog.visible = false"
    >
      <el-form label-position="top">
        <el-form-item label="key" prop="key">
          <el-tag>{{ acmeDnsConfigInfo.key }}</el-tag>
        </el-form-item>
        <el-form-item label="value" prop="value">
          <el-tag>{{ acmeDnsConfigInfo.value }}</el-tag>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="acmeDnsConfigInfoDialog.visible = false"
            >{{ $t("common.cancel") }}
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts">
export default {
  name: "acmeDnsConfig",
};
</script>

<script setup lang="ts">
import { PropType } from "vue";

const props = defineProps({
  acmeDnsConfig: {
    required: true,
    type: {} as PropType<{ [key: string]: string }>,
    default: () => {},
  },
});

const emit = defineEmits<{
  (event: "update:acmeDnsConfig", value: Map<string, string>): void;
}>();

const acmeDnsConfig = useVModel(props, "acmeDnsConfig", emit);

const dataFormRef = ref(ElForm);

const dataFormRules = {
  key: [
    {
      required: true,
      message: "Required",
      trigger: ["change", "blur"],
    },
  ],
  value: [
    {
      required: true,
      message: "Required",
      trigger: ["change", "blur"],
    },
  ],
};

interface Form {
  key: string;
  value: string;
}

const state = reactive({
  dataForm: {
    key: "",
    value: "",
  } as Form,
  dialog: {
    title: "Add ACME DNS Config",
    visible: false,
  } as DialogType,
  acmeDnsConfigInfoDialog: {
    title: "ACME DNS Config Info",
    visible: false,
  },
  acmeDnsConfigInfo: {
    key: "",
    value: "",
  } as Form,
});

const { dataForm, dialog, acmeDnsConfigInfoDialog, acmeDnsConfigInfo } =
  toRefs(state);

const handleAdd = () => {
  state.dialog.visible = true;
};

const handleClose = (key: string): void => {
  delete acmeDnsConfig.value[key];
};

const handleInfo = (key: string) => {
  const value = acmeDnsConfig.value[key];
  state.acmeDnsConfigInfo = {
    key: key,
    value: value ? value : "",
  };
  state.acmeDnsConfigInfoDialog.visible = true;
};

const submitForm = () => {
  dataFormRef.value.validate((valid: any) => {
    if (valid) {
      if (acmeDnsConfig.value[state.dataForm.key]) {
        ElMessage.error("key cannot be repeated");
        return;
      }
      acmeDnsConfig.value[state.dataForm.key] = state.dataForm.value;
      closeDialog();
    }
  });
};

const closeDialog = (): void => {
  state.dialog.visible = false;
  dataFormRef.value.resetFields();
  dataFormRef.value.clearValidate();
};
</script>

<style lang="scss" scoped>
.flex.gap-2 {
  flex-wrap: wrap;
}
</style>
