<template>
  <div class="flex gap-2">
    <el-tag
      :key="key"
      v-for="(value, key) in mapObject"
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
      :title="infoDialog.title"
      v-model="infoDialog.visible"
      width="600px"
      append-to-body
      @close="infoDialog.visible = false"
    >
      <el-form label-position="top">
        <el-form-item label="key" prop="key">
          <el-tag>{{ dataInfo.key }}</el-tag>
        </el-form-item>
        <el-form-item label="value" prop="value">
          <el-tag>{{ dataInfo.value }}</el-tag>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="infoDialog.visible = false"
            >{{ $t("common.cancel") }}
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts">
export default {
  name: "mapObject",
};
</script>

<script setup lang="ts">
import { PropType } from "vue";

interface Form {
  key: string;
  value: string;
}

const props = defineProps({
  mapObject: {
    required: false,
    type: Object as PropType<{ [key: string]: string }>,
    default: () => ({}),
  },
});

const emit = defineEmits<{
  (event: "update:mapObject", value: { [key: string]: string }): void;
}>();

const mapObject = useVModel(props, "mapObject", emit);

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

const state = reactive({
  dataForm: {
    key: "",
    value: "",
  } as Form,
  dialog: {
    title: "Add",
    visible: false,
  } as DialogType,
  infoDialog: {
    title: "Info",
    visible: false,
  },
  dataInfo: {
    key: "",
    value: "",
  } as Form,
});

const { dataForm, dialog, infoDialog, dataInfo } = toRefs(state);

const handleAdd = () => {
  state.dialog.visible = true;
};

const handleClose = (key: string): void => {
  delete mapObject.value[key];
};

const handleInfo = (key: string) => {
  state.dataInfo = {
    key: key,
    value: mapObject.value[key] || "",
  };
  state.infoDialog.visible = true;
};

const submitForm = () => {
  dataFormRef.value.validate((valid: any) => {
    if (valid) {
      if (mapObject.value[state.dataForm.key]) {
        ElMessage.error("key cannot be repeated");
        return;
      }
      mapObject.value[state.dataForm.key] = state.dataForm.value;
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
