<template>
  <div class="app-container">
    <el-card shadow="never">
      <template #header>
        <div class="flex justify-between">
          <div></div>
          <div>
            <el-button type="success" @click="submitForm" :icon="Select"
              >保存
            </el-button>
          </div>
        </div>
      </template>
      <el-form ref="dataFormRef" :model="formData" label-width="150px">
        <el-row :gutter="20">
          <el-col :span="8">
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
          </el-col>
          <el-col :span="8"></el-col>
          <el-col :span="8"></el-col>
        </el-row>
      </el-form>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { Select } from "@element-plus/icons-vue";
import { listConfig, updateConfigs } from "@/api/config";
import { ConfigsUpdateDto } from "@/api/config/types";

const state = reactive({
  huiWebPort: "H_UI_WEB_PORT",
  hysteria2TrafficTime: "HYSTERIA2_TRAFFIC_TIME",
  formData: {
    huiWebPort: "8081",
    hysteria2TrafficTime: "1",
  },
});

const { formData } = toRefs(state);

const submitForm = () => {
  let configs: Array<ConfigsUpdateDto> = [
    {
      key: state.huiWebPort,
      value: state.formData.huiWebPort,
    },
    {
      key: state.hysteria2TrafficTime,
      value: state.formData.hysteria2TrafficTime,
    },
  ];
  updateConfigs({ configUpdateDtos: configs }).then(() => {
    ElMessage.success("修改配置成功");
  });
};

onMounted(() => {
  listConfig({ keys: [state.huiWebPort, state.hysteria2TrafficTime] }).then(
    (response) => {
      const configVos = response.data;
      configVos.forEach((configVo) => {
        if (configVo.key === state.huiWebPort) {
          state.formData.huiWebPort = configVo.value;
        } else if (configVo.key === state.hysteria2TrafficTime) {
          state.formData.hysteria2TrafficTime = configVo.value;
        }
      });
    }
  );
});
</script>

<style lang="sass" scoped></style>
