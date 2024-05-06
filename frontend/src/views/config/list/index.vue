<template>
  <div class="app-container">
    <div class="search">
      <el-form inline>
        <el-form-item>
          <el-button type="success" @click="submitForm" :icon="Select"
            >保存
          </el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-card shadow="never">
      <div class="config-container">
        <el-form ref="dataFormRef" label-position="top">
          <el-form-item :label="$t('config.huiWebPort')" prop="huiWebPort">
            <el-input
              v-model="huiWebPortVal"
              :placeholder="$t('config.huiWebPort')"
              clearable
            />
          </el-form-item>
          <el-form-item
            :label="$t('config.hysteria2TrafficTime')"
            prop="hysteria2TrafficTime"
          >
            <el-input
              v-model="hysteria2TrafficTimeVal"
              :placeholder="$t('config.hysteria2TrafficTime')"
              clearable
            />
          </el-form-item>
        </el-form>
      </div>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { Select } from "@element-plus/icons-vue";
import { listConfig, updateConfigs } from "@/api/config";
import { ConfigsUpdateDto } from "@/api/config/types";

const state = reactive({
  huiWebPortVal: "8081",
  hysteria2TrafficTimeVal: "1",
});

const { huiWebPortVal, hysteria2TrafficTimeVal } = toRefs(state);

const huiWebPortKey = "H_UI_WEB_PORT";
const hysteria2TrafficTimeKey = "HYSTERIA2_TRAFFIC_TIME";

const submitForm = () => {
  let configs: Array<ConfigsUpdateDto> = [
    {
      key: huiWebPortKey,
      value: state.huiWebPortVal,
    },
    {
      key: hysteria2TrafficTimeKey,
      value: state.hysteria2TrafficTimeVal,
    },
  ];
  updateConfigs({ configUpdateDtos: configs }).then(() => {
    ElMessage.success("修改配置成功");
  });
};

onMounted(() => {
  listConfig({
    keys: [huiWebPortKey, hysteria2TrafficTimeKey],
  }).then((response) => {
    const configVos = response.data;
    configVos.forEach((configVo) => {
      if (configVo.key === huiWebPortKey) {
        state.huiWebPortVal = configVo.value;
      } else if (configVo.key === hysteria2TrafficTimeKey) {
        state.hysteria2TrafficTimeVal = configVo.value;
      }
    });
  });
});
</script>

<style lang="scss" scoped>
.config-container {
  max-width: 800px;
  margin: 0 auto;
}
</style>
