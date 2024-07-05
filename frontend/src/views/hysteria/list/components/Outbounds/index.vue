<template>
  <div class="flex gap-2">
    <el-tag
      :key="item"
      v-for="item in outbounds"
      @close="handleClose(item)"
      @click="handleInfo(item)"
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
      <el-form ref="dataFormRef" label-position="top" :model="dataForm">
        <el-tooltip
          :content="$t('hysteria.config.outbounds.name')"
          placement="bottom"
        >
          <el-form-item label="name" prop="name">
            <el-input v-model="dataForm.name" clearable />
          </el-form-item>
        </el-tooltip>
        <el-tooltip
          :content="$t('hysteria.config.outbounds.type')"
          placement="bottom"
        >
          <el-form-item label="type" prop="type">
            <el-select v-model="dataForm.type" style="width: 100%">
              <el-option
                v-for="item in outboundTypes"
                :key="item"
                :label="item"
                :value="item"
              />
            </el-select>
          </el-form-item>
        </el-tooltip>
        <template v-if="dataForm.type === 'socks5'">
          <el-tooltip
            :content="$t('hysteria.config.outbounds.socks5.addr')"
            placement="bottom"
          >
            <el-form-item label="socks5.addr" prop="socks5.addr">
              <el-input v-model="dataForm.socks5.addr" clearable />
            </el-form-item>
          </el-tooltip>
          <el-tooltip
            :content="$t('hysteria.config.outbounds.socks5.username')"
            placement="bottom"
          >
            <el-form-item label="socks5.username" prop="socks5.username">
              <el-input v-model="dataForm.socks5.username" clearable />
            </el-form-item>
          </el-tooltip>
          <el-tooltip
            :content="$t('hysteria.config.outbounds.socks5.password')"
            placement="bottom"
          >
            <el-form-item label="socks5.password" prop="socks5.password">
              <el-input v-model="dataForm.socks5.password" clearable />
            </el-form-item>
          </el-tooltip>
        </template>
        <template v-if="dataForm.type === 'http'">
          <el-tooltip
            :content="$t('hysteria.config.outbounds.http.url')"
            placement="bottom"
          >
            <el-form-item label="http.url" prop="http.url">
              <el-input v-model="dataForm.http.url" clearable />
            </el-form-item>
          </el-tooltip>
          <el-tooltip
            :content="$t('hysteria.config.outbounds.http.insecure')"
            placement="bottom"
          >
            <el-form-item label="http.insecure" prop="http.insecure">
              <el-switch v-model="dataForm.http.insecure" />
            </el-form-item>
          </el-tooltip>
        </template>
        <template v-if="dataForm.type === 'direct'">
          <el-tooltip
            :content="$t('hysteria.config.outbounds.direct.mode')"
            placement="bottom"
          >
            <el-form-item label="direct.mode" prop="direct.mode">
              <el-select v-model="dataForm.direct.mode" style="width: 100%">
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
            <el-form-item label="direct.bindIPv4" prop="direct.bindIPv4">
              <el-input v-model="dataForm.direct.bindIPv4" clearable />
            </el-form-item>
          </el-tooltip>
          <el-tooltip
            :content="$t('hysteria.config.outbounds.direct.bindIPv6')"
            placement="bottom"
          >
            <el-form-item label="direct.bindIPv6" prop="direct.bindIPv6">
              <el-input v-model="dataForm.direct.bindIPv6" clearable />
            </el-form-item>
          </el-tooltip>
          <el-tooltip
            :content="$t('hysteria.config.outbounds.direct.bindDevice')"
            placement="bottom"
          >
            <el-form-item label="direct.bindDevice" prop="direct.bindDevice">
              <el-input v-model="dataForm.direct.bindDevice" clearable />
            </el-form-item>
          </el-tooltip>
        </template>
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
      :title="outboundInfoDialog.title"
      v-model="outboundInfoDialog.visible"
      width="600px"
      append-to-body
      @close="outboundInfoDialog.visible = false"
    >
      <el-form label-position="top">
        <el-tooltip
          :content="$t('hysteria.config.outbounds.name')"
          placement="bottom"
        >
          <el-form-item label="name" prop="name">
            <el-tag>{{ outboundInfo.name }}</el-tag>
          </el-form-item>
        </el-tooltip>
        <el-tooltip
          :content="$t('hysteria.config.outbounds.type')"
          placement="bottom"
        >
          <el-form-item label="type" prop="type">
            <el-tag>{{ outboundInfo.type }}</el-tag>
          </el-form-item>
        </el-tooltip>
        <template v-if="outboundInfo.type === 'socks5'">
          <el-tooltip
            :content="$t('hysteria.config.outbounds.socks5.addr')"
            placement="bottom"
          >
            <el-form-item label="socks5.addr" prop="socks5.addr">
              <el-tag>{{ outboundInfo.socks5.addr }}</el-tag>
            </el-form-item>
          </el-tooltip>
          <el-tooltip
            :content="$t('hysteria.config.outbounds.socks5.username')"
            placement="bottom"
          >
            <el-form-item label="socks5.username" prop="outboundInfo.username">
              <el-tag>{{ outboundInfo.socks5.username }}</el-tag>
            </el-form-item>
          </el-tooltip>
          <el-tooltip
            :content="$t('hysteria.config.outbounds.socks5.password')"
            placement="bottom"
          >
            <el-form-item label="socks5.password" prop="socks5.password">
              <el-tag>{{ outboundInfo.socks5.password }}</el-tag>
            </el-form-item>
          </el-tooltip>
        </template>
        <template v-if="outboundInfo.type === 'http'">
          <el-tooltip
            :content="$t('hysteria.config.outbounds.http.url')"
            placement="bottom"
          >
            <el-form-item label="http.url" prop="http.url">
              <el-tag>{{ outboundInfo.http.url }}</el-tag>
            </el-form-item>
          </el-tooltip>
          <el-tooltip
            :content="$t('hysteria.config.outbounds.http.insecure')"
            placement="bottom"
          >
            <el-form-item label="http.insecure" prop="http.insecure">
              <el-tag>{{ outboundInfo.http.insecure }}</el-tag>
            </el-form-item>
          </el-tooltip>
        </template>
        <template v-if="outboundInfo.type === 'direct'">
          <el-tooltip
            :content="$t('hysteria.config.outbounds.direct.mode')"
            placement="bottom"
          >
            <el-form-item label="direct.mode" prop="direct.mode">
              <el-tag>{{ outboundInfo.direct.mode }}</el-tag>
            </el-form-item>
          </el-tooltip>
          <el-tooltip
            :content="$t('hysteria.config.outbounds.direct.bindIPv4')"
            placement="bottom"
          >
            <el-form-item label="direct.bindIPv4" prop="direct.bindIPv4">
              <el-tag>{{ outboundInfo.direct.bindIPv4 }}</el-tag>
            </el-form-item>
          </el-tooltip>
          <el-tooltip
            :content="$t('hysteria.config.outbounds.direct.bindIPv6')"
            placement="bottom"
          >
            <el-form-item label="direct.bindIPv6" prop="direct.bindIPv6">
              <el-tag>{{ outboundInfo.direct.bindIPv6 }}</el-tag>
            </el-form-item>
          </el-tooltip>
          <el-tooltip
            :content="$t('hysteria.config.outbounds.direct.bindDevice')"
            placement="bottom"
          >
            <el-form-item label="direct.bindDevice" prop="direct.bindDevice">
              <el-tag>{{ outboundInfo.direct.bindDevice }}</el-tag>
            </el-form-item>
          </el-tooltip>
        </template>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="outboundInfoDialog.visible = false"
            >{{ $t("common.cancel") }}
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts">
export default {
  name: "outbounds",
};
</script>

<script setup lang="ts">
import {
  defaultHysteria2ServerConfigOutbound,
  Hysteria2ServerConfigOutbound,
} from "@/api/config/types";
import { PropType } from "vue";
import { deepCopy } from "@/utils/copy";

const props = defineProps({
  outbounds: {
    required: true,
    type: Array as PropType<Array<Hysteria2ServerConfigOutbound>>,
    default: (): Array<Hysteria2ServerConfigOutbound> => [],
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
  dataForm: {
    ...defaultHysteria2ServerConfigOutbound,
  } as Hysteria2ServerConfigOutbound,
  dialog: {
    title: "Add Outbound",
    visible: false,
  } as DialogType,
  outboundInfoDialog: {
    title: "Outbound Info",
    visible: false,
  },
  outboundInfo: {} as Hysteria2ServerConfigOutbound,
});

const { dataForm, dialog, outboundInfoDialog, outboundInfo } = toRefs(state);

const outboundTypes = ["socks5", "http", "direct"];

const outboundDirectModes = ["auto", "64", "46", "6", "4"];

const handleAdd = () => {
  state.dialog.visible = true;
};

const handleClose = (outbound: Hysteria2ServerConfigOutbound): void => {
  const index = outbounds.value.indexOf(outbound);
  if (index !== -1) {
    outbounds.value.splice(index, 1);
  }
};

const handleInfo = (outbound: Hysteria2ServerConfigOutbound) => {
  state.outboundInfo = outbound;
  state.outboundInfoDialog.visible = true;
};

const submitForm = () => {
  dataFormRef.value.validate((valid: any) => {
    if (valid) {
      if (outbounds.value.some((item) => item.name === state.dataForm.name)) {
        ElMessage.error("name cannot be repeated");
        return;
      }
      if (state.dataForm.type === "socks5") {
        state.dataForm.http = undefined;
        state.dataForm.direct = undefined;
      } else if (state.dataForm.type === "http") {
        state.dataForm.socks5 = undefined;
        state.dataForm.direct = undefined;
      } else if (state.dataForm.type === "direct") {
        state.dataForm.socks5 = undefined;
        state.dataForm.http = undefined;
      }

      let outbound = deepCopy(state.dataForm);
      outbounds.value.push(outbound);
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
