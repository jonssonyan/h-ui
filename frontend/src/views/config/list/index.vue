<template>
  <div class="app-container">
    <div class="search">
      <el-form inline>
        <el-form-item>
          <el-button
            type="primary"
            @click="submitForm(formDataRef)"
            :icon="Select"
            >保存
          </el-button>
        </el-form-item>
      </el-form>
    </div>

    <el-card shadow="never">
      <el-form
        ref="formDataRef"
        :rules="rules"
        :model="formData"
        label-position="top"
      >
        <el-form-item :label="$t('config.huiWebPort')" prop="huiWebPort">
          <el-input
            v-model="formData.huiWebPort"
            :placeholder="$t('config.huiWebPort')"
            clearable
          />
        </el-form-item>
        <el-form-item
          :label="$t('config.hysteria2TrafficTime')"
          prop="hysteria2TrafficTime"
        >
          <el-input
            v-model="formData.hysteria2TrafficTime"
            :placeholder="$t('config.hysteria2TrafficTime')"
            clearable
          />
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { Select } from "@element-plus/icons-vue";
import { listConfigApi, updateConfigsApi } from "@/api/config";
import { ConfigsUpdateDto } from "@/api/config/types";
import type { FormInstance, FormRules } from "element-plus";
import { validateNumberStr } from "@/utils/validate";

const huiWebPortKey = "H_UI_WEB_PORT";
const hysteria2TrafficTimeKey = "HYSTERIA2_TRAFFIC_TIME";

const formDataRef = ref<FormInstance>();
const rules = reactive<FormRules>({
  huiWebPort: [
    {
      validator: validateNumberStr,
      trigger: ["change", "blur"],
    },
  ],
  hysteria2TrafficTime: [
    {
      validator: validateNumberStr,
      trigger: ["change", "blur"],
    },
  ],
});

const formData = reactive({
  huiWebPort: "8081",
  hysteria2TrafficTime: "1",
});

const submitForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.validate((valid) => {
    if (valid) {
      let configs: Array<ConfigsUpdateDto> = [
        {
          key: huiWebPortKey,
          value: formData.huiWebPort,
        },
        {
          key: hysteria2TrafficTimeKey,
          value: formData.hysteria2TrafficTime,
        },
      ];
      updateConfigsApi({ configUpdateDtos: configs }).then(() => {
        ElMessage.success("修改配置成功");
      });
    }
  });
};

const setConfig = () => {
  listConfigApi({
    keys: [huiWebPortKey, hysteria2TrafficTimeKey],
  }).then((response) => {
    const configVos = response.data;
    configVos.forEach((configVo) => {
      if (configVo.key === huiWebPortKey) {
        formData.huiWebPort = configVo.value;
      } else if (configVo.key === hysteria2TrafficTimeKey) {
        formData.hysteria2TrafficTime = configVo.value;
      }
    });
  });
};

onMounted(() => {
  setConfig();
});
</script>

<style lang="scss" scoped>
.el-card .el-form {
  max-width: 800px;
  margin: 0 auto;
}
</style>
