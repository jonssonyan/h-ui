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
              <el-select v-model="formData.type" style="width: 100%">
                <el-option
                  v-for="item in outboundTypes"
                  :key="item"
                  :label="item"
                  :value="item"
                />
              </el-select>
            </el-form-item>
          </el-tooltip>
          <el-tooltip
            :content="$t('hysteria.config.outbounds.socks5.addr')"
            placement="bottom"
          >
            <el-form-item
              label="socks5.addr"
              prop="socks5.addr"
              v-if="formData.type === 'socks5'"
            >
              <el-input v-model="formData.socks5.addr" clearable />
            </el-form-item>
          </el-tooltip>
          <el-tooltip
            :content="$t('hysteria.config.outbounds.socks5.username')"
            placement="bottom"
          >
            <el-form-item
              label="socks5.username"
              prop="socks5.username"
              v-if="formData.type === 'socks5'"
            >
              <el-input v-model="formData.socks5.username" clearable />
            </el-form-item>
          </el-tooltip>
          <el-tooltip
            :content="$t('hysteria.config.outbounds.socks5.password')"
            placement="bottom"
          >
            <el-form-item
              label="socks5.password"
              prop="socks5.password"
              v-if="formData.type === 'socks5'"
            >
              <el-input v-model="formData.socks5.password" clearable />
            </el-form-item>
          </el-tooltip>
          <el-tooltip
            :content="$t('hysteria.config.outbounds.http.url')"
            placement="bottom"
          >
            <el-form-item
              label="http.url"
              prop="http.url"
              v-if="formData.type === 'http'"
            >
              <el-input v-model="formData.http.url" clearable />
            </el-form-item>
          </el-tooltip>
          <el-tooltip
            :content="$t('hysteria.config.outbounds.http.insecure')"
            placement="bottom"
          >
            <el-form-item
              label="http.insecure"
              prop="http.insecure"
              v-if="formData.type === 'http'"
            >
              <el-switch v-model="formData.http.insecure" />
            </el-form-item>
          </el-tooltip>
          <el-tooltip
            :content="$t('hysteria.config.outbounds.direct.mode')"
            placement="bottom"
          >
            <el-form-item
              label="direct.mode"
              prop="direct.mode"
              v-if="formData.type === 'direct'"
            >
              <el-select v-model="formData.direct.mode" style="width: 100%">
                <el-option
                  v-for="item in outboundDirectModes"
                  :key="item"
                  :label="item"
                  :value="item"
                />
              </el-select>
            </el-form-item>
          </el-tooltip>
          <el-tooltip
            :content="$t('hysteria.config.outbounds.direct.bindIPv4')"
            placement="bottom"
          >
            <el-form-item
              label="direct.bindIPv4"
              prop="direct.bindIPv4"
              v-if="formData.type === 'direct'"
            >
              <el-input v-model="formData.direct.bindIPv4" clearable />
            </el-form-item>
          </el-tooltip>
          <el-tooltip
            :content="$t('hysteria.config.outbounds.direct.bindIPv6')"
            placement="bottom"
          >
            <el-form-item
              label="direct.bindIPv6"
              prop="direct.bindIPv6"
              v-if="formData.type === 'direct'"
            >
              <el-input v-model="formData.direct.bindIPv6" clearable />
            </el-form-item>
          </el-tooltip>
          <el-tooltip
            :content="$t('hysteria.config.outbounds.direct.bindDevice')"
            placement="bottom"
          >
            <el-form-item
              label="direct.bindDevice"
              prop="direct.bindDevice"
              v-if="formData.type === 'direct'"
            >
              <el-input v-model="formData.direct.bindDevice" clearable />
            </el-form-item>
          </el-tooltip>
        </el-form>
        <template #footer>
          <div class="dialog-footer">
            <el-button type="primary" @click="submitForm"
              >{{ $t("common.confirm") }}
            </el-button>
            <el-button @click="closeDialog"
              >{{ $t("common.cancel") }}
            </el-button>
          </div>
        </template>
      </el-dialog>
    </div>
  </div>
</template>

<script setup lang="ts">
import {
  defaultHysteria2ServerConfigOutbound,
  Hysteria2ServerConfigOutbound,
} from "@/api/config/types";

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
  formData:
    defaultHysteria2ServerConfigOutbound as Hysteria2ServerConfigOutbound,
  dialog: {
    visible: false,
  } as DialogType,
});

const { formData, dialog } = toRefs(state);

const outboundTypes = ["socks5", "http", "direct"];

const outboundDirectModes = ["auto", "64", "46", "6", "4"];

function resetFormData() {
  Object.assign(state.formData, defaultHysteria2ServerConfigOutbound);
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
  Object.assign(state.formData, outbound);

  dialog.value = {
    title: "Update Outbound",
    visible: true,
  };
};

const submitForm = () => {
  dataFormRef.value.validate((valid: any) => {
    if (valid) {
      const title = state.dialog.title;

      // 精简 美化 最终配置文件
      if (state.formData.type === "socks5") {
        state.formData.http = undefined;
        state.formData.direct = undefined;
      } else if (state.formData.type === "http") {
        state.formData.socks5 = undefined;
        state.formData.direct = undefined;
      } else if (state.formData.type === "direct") {
        state.formData.socks5 = undefined;
        state.formData.http = undefined;
      }

      const outbound = { ...state.formData };

      console.log(state.formData);
      console.log(outbound);

      if (title == "Update Outbound") {
        console.log("update outbound");
      } else {
        outbounds.value?.push(outbound);
      }
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
