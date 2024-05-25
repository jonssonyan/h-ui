<template>
  <div>
    <div class="flex gap-2">
      <el-tag
        :key="item"
        v-for="item in outbounds"
        @close="handleClose(item)"
        @click="handleUpdate(item)"
        size="large"
        closable
      >
        {{ item.name }}
      </el-tag>

      <el-button @click="handleAdd">+</el-button>

      <el-dialog
        :title="dialog.title"
        v-model="dialog.visible"
        width="600px"
        append-to-body
        @close="closeDialog"
      >
        <el-form ref="dataFormRef" label-position="top" :model="formData">
          <el-tooltip
            :content="$t('hysteria.config.outbounds.name')"
            placement="bottom"
          >
            <el-form-item label="name" prop="name">
              <el-input v-model="formData.name" clearable />
            </el-form-item>
          </el-tooltip>
          <el-tooltip
            :content="$t('hysteria.config.outbounds.type')"
            placement="bottom"
          >
            <el-form-item label="type" prop="type">
              <el-input v-model="formData.type" clearable />
            </el-form-item>
          </el-tooltip>
        </el-form>
        <template #footer>
          <div class="dialog-footer">
            <el-button type="primary" @click="submitForm"
              >{{ $t("common.confirm") }}
            </el-button>
            <el-button @click="closeDialog">{{
								$t("common.cancel")
                }}</el-button>
          </div>
        </template>
      </el-dialog>
    </div>
  </div>
</template>

<script setup lang="ts">
import { Hysteria2ServerConfigOutbound } from "@/api/config/types";

const props = defineProps({
  outbounds: {
    required: true,
    type: Array,
    default: () => [],
  },
});

const emit = defineEmits<{
  (
    event: "update:outbounds",
    value: Array<Hysteria2ServerConfigOutbound>
  ): void;
}>();

const outbounds = useVModel(props, "outbounds", emit);

const dataFormRef = ref(ElForm);

const state = reactive({
  formData: {
    name: "",
    type: "",
    socks5: undefined,
    http: undefined,
    direct: undefined,
  } as Hysteria2ServerConfigOutbound,
  dialog: {
    visible: false,
  } as DialogType,
});

const { formData, dialog } = toRefs(state);

function resetFormData() {
  Object.assign(state.formData, {
    name: "",
    type: "",
    socks5: undefined,
    http: undefined,
    direct: undefined,
  });
}

const handleAdd = () => {
  state.dialog = {
    title: "Add Outbound",
    visible: true,
  };
};

const handleClose = (outbound: Hysteria2ServerConfigOutbound): void => {
  const index = outbounds.value?.indexOf(outbound);
  if (index !== -1) {
    outbounds.value?.splice(index, 1);
  }
};

const handleUpdate = (outbound: Hysteria2ServerConfigOutbound) => {
  Object.assign(formData.value, outbound);

  dialog.value = {
    title: "Update Outbound",
    visible: true,
  };
};

const submitForm = () => {
  dataFormRef.value.validate((valid: any) => {
    if (valid) {
      const title = state.dialog.title;
      const outbound = { ...state.formData };
      if (title == "Update Outbound") {
        console.log(state.formData);
      } else {
        outbounds.value?.push(outbound);
      }
      console.log(outbounds.value);
      closeDialog();
    }
  });
};
const closeDialog = (): void => {
  state.dialog.visible = false;
  dataFormRef.value.resetFields();
  dataFormRef.value.clearValidate();

  if (state.dialog.title == "Update Outbound") {
    resetFormData();
  }
};
</script>

<style lang="scss" scoped>
.flex.gap-2 {
  flex-wrap: wrap;
}
</style>
