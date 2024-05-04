<template>
  <div class="app-container">
    <el-card shadow="never">
      <template #header>
        <div class="flex justify-between">
          <div>
            <el-button @click="handleImport">
              <template #icon>
                <i-ep-upload />
              </template>
              导入
            </el-button>
            <el-button @click="handleExport">
              <template #icon>
                <i-ep-download />
              </template>
              导出
            </el-button>
          </div>
          <div>
            <el-button type="success" @click="submitForm" :icon="Select"
              >保存
            </el-button>
          </div>
        </div>
      </template>
      <el-form ref="enableFormRef" inline>
        <el-form-item label="enable" prop="enable">
          <el-switch
            v-model="enableValue"
            active-value="1"
            inactive-value="0"
            @change="handleChange"
          />
        </el-form-item>
      </el-form>
      <el-form ref="dataFormRef" inline :model="formData">
        <el-form-item label="listen" prop="listen">
          <el-input v-model="formData.listen" placeholder="listen" clearable />
        </el-form-item>
        <el-form-item label="obfs.type" prop="obfs.type">
          <el-input
            v-model="formData.obfs.type"
            placeholder="obfs.type"
            clearable
          />
        </el-form-item>
        <el-form-item
          label="obfs.salamander.password"
          prop="obfs.salamander.password"
        >
          <el-input
            v-model="formData.obfs.salamander.password"
            placeholder="obfs.salamander.password"
            clearable
          />
        </el-form-item>
        <el-form-item label="tls.cert" prop="tls.cert">
          <el-input
            v-model="formData.tls.cert"
            placeholder="tls.cert"
            clearable
          />
        </el-form-item>
        <el-form-item label="tls.key" prop="tls.key">
          <el-input
            v-model="formData.tls.key"
            placeholder="tls.key"
            clearable
          />
        </el-form-item>
        <el-form-item label="acme.domains" prop="acme.domains">
          <el-input
            v-model="formData.acme.domains"
            placeholder="acme.domains"
            clearable
          />
        </el-form-item>
        <el-form-item label="acme.email" prop="acme.email">
          <el-input
            v-model="formData.acme.email"
            placeholder="acme.email"
            clearable
          />
        </el-form-item>
        <el-form-item label="acme.ca" prop="acme.ca">
          <el-input
            v-model="formData.acme.ca"
            placeholder="acme.ca"
            clearable
          />
        </el-form-item>
        <el-form-item label="acme.disableHTTP" prop="acme.disableHTTP">
          <el-input
            v-model="formData.acme.disableHTTP"
            placeholder="acme.disableHTTP"
            clearable
          />
        </el-form-item>
        <el-form-item label="acme.disableTLSALPN" prop="acme.disableTLSALPN">
          <el-input
            v-model="formData.acme.disableTLSALPN"
            placeholder="acme.disableTLSALPN"
            clearable
          />
        </el-form-item>
        <el-form-item label="acme.listenHost" prop="acme.listenHost">
          <el-input
            v-model="formData.acme.listenHost"
            placeholder="acme.listenHost"
            clearable
          />
        </el-form-item>
        <el-form-item label="acme.altHTTPPort" prop="acme.altHTTPPort">
          <el-input
            v-model="formData.acme.altHTTPPort"
            placeholder="acme.altHTTPPort"
            clearable
          />
        </el-form-item>
        <el-form-item label="acme.altTLSALPNPort" prop="acme.altTLSALPNPort">
          <el-input
            v-model="formData.acme.altTLSALPNPort"
            placeholder="acme.altTLSALPNPort"
            clearable
          />
        </el-form-item>
        <el-form-item label="acme.dir" prop="acme.dir">
          <el-input
            v-model="formData.acme.dir"
            placeholder="acme.dir"
            clearable
          />
        </el-form-item>
        <el-form-item
          label="quic.initStreamReceiveWindow"
          prop="quic.initStreamReceiveWindow"
        >
          <el-input
            v-model="formData.quic.initStreamReceiveWindow"
            placeholder="quic.initStreamReceiveWindow"
            clearable
          />
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import {
  defaultHysteria2ServerConfig,
  Hysteria2ServerConfig,
} from "@/api/config/types";
import { Select } from "@element-plus/icons-vue";
import { getConfig, getHysteria2Config } from "@/api/config";

const state = reactive({
  enableKey: "HYSTERIA2_ENABLE",
  enableValue: "0",
  formData: defaultHysteria2ServerConfig as Hysteria2ServerConfig,
});

const { enableValue, formData } = toRefs(state);

const handleImport = () => {};
const handleExport = () => {};

const handleChange = () => {};
const submitForm = () => {};

onMounted(() => {
  getConfig({ key: state.enableKey }).then((response) => {
    if (response.data) {
      const { value } = response.data;
      state.enableValue = value;
    }
  });
  getHysteria2Config().then((response) => {
    Object.assign(state.formData, response.data);
  });
});
</script>

<style lang="scss" scoped>
.el-switch {
  --el-switch-on-color: #00ff00;
  --el-switch-off-color: #cccccc;
}
</style>
