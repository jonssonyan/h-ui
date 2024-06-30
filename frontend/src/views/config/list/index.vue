<template>
  <div class="app-container">
    <div class="search">
      <el-form inline>
        <el-form-item>
          <el-button type="primary" @click="submitForm" :icon="Select">
            {{ $t("common.save") }}
          </el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="handleRestartServer">
            <template #icon>
              <i-ep-refreshRight />
            </template>
            {{ $t("config.restartServer") }}
          </el-button>
        </el-form-item>
        <el-form-item>
          <el-upload
            v-model:file-list="fileList"
            :http-request="handleImport"
            :show-file-list="false"
            accept=".json"
            :limit="1"
            :before-upload="beforeImport"
          >
            <el-button>
              <template #icon>
                <i-ep-upload />
              </template>
              {{ $t("common.import") }}
            </el-button>
          </el-upload>
        </el-form-item>
        <el-form-item>
          <el-button @click="handleExport">
            <template #icon>
              <i-ep-download />
            </template>
            {{ $t("common.export") }}
          </el-button>
        </el-form-item>
      </el-form>
    </div>

    <el-card shadow="never">
      <el-form
        ref="dataFormRef"
        :rules="dataFormRules"
        :model="dataForm"
        label-position="top"
      >
        <el-form-item :label="$t('config.huiWebPort')" prop="huiWebPort">
          <el-input
            v-model="dataForm.huiWebPort"
            :placeholder="$t('config.huiWebPort')"
            clearable
          />
        </el-form-item>
        <el-form-item
          :label="$t('config.hysteria2TrafficTime')"
          prop="hysteria2TrafficTime"
        >
          <el-input
            v-model="dataForm.hysteria2TrafficTime"
            :placeholder="$t('config.hysteria2TrafficTime')"
            clearable
          />
        </el-form-item>
        <el-form-item :label="$t('config.huiHttps')" prop="huiHttps">
          <el-select v-model="huiHttps" style="width: 50%">
            <el-option
              v-for="item in huiHttpsList"
              :key="item.key"
              :label="item.key"
              :value="item.value"
            />
          </el-select>
          <el-button v-if="huiHttps" @click="setCertPath"
            >{{ t("config.useHysteria2Cert") }}
          </el-button>
        </el-form-item>
        <el-form-item
          v-if="huiHttps"
          :label="$t('config.huiCrtPath')"
          prop="huiCrtPath"
        >
          <el-input
            v-model="dataForm.huiCrtPath"
            :placeholder="$t('config.huiCrtPath')"
            clearable
          />
        </el-form-item>
        <el-form-item
          v-if="huiHttps"
          :label="$t('config.huiKeyPath')"
          prop="huiKeyPath"
        >
          <el-input
            v-model="dataForm.huiKeyPath"
            :placeholder="$t('config.huiKeyPath')"
            clearable
          />
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script lang="ts">
export default {
  name: "index",
};
</script>

<script setup lang="ts">
import { Select } from "@element-plus/icons-vue";
import {
  exportConfigApi,
  hysteria2AcmePathApi,
  importConfigApi,
  listConfigApi,
  restartServerApi,
  updateConfigsApi,
} from "@/api/config";
import { ConfigsUpdateDto } from "@/api/config/types";
import {
  UploadFile,
  UploadRawFile,
  UploadRequestOptions,
} from "element-plus/lib/components";
import { useI18n } from "vue-i18n";

const { t } = useI18n();

const dataFormRef = ref(ElForm);

const huiWebPortKey = "H_UI_WEB_PORT";
const hysteria2TrafficTimeKey = "HYSTERIA2_TRAFFIC_TIME";
const huiCrtPath = "H_UI_CRT_PATH";
const huiKeyPath = "H_UI_KEY_PATH";

const huiHttpsList = [
  { key: t("common.yes"), value: 1 },
  { key: t("common.no"), value: 0 },
];

const dataFormRules = {
  huiWebPort: [
    {
      required: true,
      message: "Required",
      trigger: ["change", "blur"],
    },
    {
      pattern: /^\d+$/,
      message: "field must be a integer",
      trigger: ["change", "blur"],
    },
  ],
  hysteria2TrafficTime: [
    {
      required: true,
      message: "Required",
      trigger: ["change", "blur"],
    },
    {
      pattern: /^\d+(\.\d)?$/,
      message: "field must be a number with up to one decimal place",
      trigger: ["change", "blur"],
    },
  ],
};

const state = reactive({
  dataForm: {
    huiWebPort: "8081",
    hysteria2TrafficTime: "1",
    huiCrtPath: "",
    huiKeyPath: "",
  },
  huiHttps: 0,
  fileList: [] as UploadFile[],
});

const { dataForm, huiHttps, fileList } = toRefs(state);

const submitForm = () => {
  dataFormRef.value.validate((valid: boolean) => {
    if (valid) {
      if (state.huiHttps) {
        if (!state.dataForm.huiCrtPath || !state.dataForm.huiKeyPath) {
          ElMessage.error("crt and key required");
          return;
        }
      }

      if (!state.huiHttps) {
        state.dataForm.huiCrtPath = "";
        state.dataForm.huiKeyPath = "";
      }

      let configs: ConfigsUpdateDto[] = [
        {
          key: huiWebPortKey,
          value: state.dataForm.huiWebPort,
        },
        {
          key: hysteria2TrafficTimeKey,
          value: state.dataForm.hysteria2TrafficTime,
        },
        {
          key: huiCrtPath,
          value: state.dataForm.huiCrtPath,
        },
        {
          key: huiKeyPath,
          value: state.dataForm.huiKeyPath,
        },
      ];

      updateConfigsApi({ configUpdateDtos: configs }).then(() => {
        ElMessage.success(t("common.success"));
      });
    }
  });
};

const setConfig = async () => {
  const { data } = await listConfigApi({
    keys: [huiCrtPath, huiKeyPath, huiWebPortKey, hysteria2TrafficTimeKey],
  });

  data.forEach((configVo) => {
    if (configVo.key === huiWebPortKey) {
      state.dataForm.huiWebPort = configVo.value;
    } else if (configVo.key === hysteria2TrafficTimeKey) {
      state.dataForm.hysteria2TrafficTime = configVo.value;
    } else if (configVo.key === huiCrtPath) {
      state.dataForm.huiCrtPath = configVo.value;
    } else if (configVo.key === huiKeyPath) {
      state.dataForm.huiKeyPath = configVo.value;
    }
  });

  if (state.dataForm.huiCrtPath != "" && state.dataForm.huiKeyPath != "") {
    state.huiHttps = 1;
  }
};

const handleImport = async (params: UploadRequestOptions) => {
  if (state.fileList.length > 0) {
    try {
      let formData = new FormData();
      formData.append("file", params.file);
      await importConfigApi(formData);
      ElMessage.success(t("common.success"));
      state.fileList = [];
    } catch (e) {
      /* empty */
    } finally {
      await setConfig();
    }
  }
};
const beforeImport = (file: UploadRawFile) => {
  if (!file.name.endsWith(".json")) {
    ElMessage.error("file format not supported");
    return false;
  }
  if (file.size / 1024 / 1024 > 2) {
    ElMessage.error("the file is too big, less than 2 MB");
    return false;
  }
};

const handleExport = async () => {
  try {
    let response = await exportConfigApi();
    const blob = new Blob([response.data], {
      type: "application/octet-stream",
    });
    let url = window.URL.createObjectURL(blob);
    let a = document.createElement("a");
    document.body.appendChild(a);
    a.href = url;
    let dis = response.headers["content-disposition"];
    a.download = dis.split("attachment; filename=")[1];
    // 模拟点击下载
    a.click();
    window.URL.revokeObjectURL(url);
    ElMessage.success(t("common.success"));
  } catch (e) {
    /* empty */
  }
};

const setCertPath = async () => {
  try {
    const { data } = await hysteria2AcmePathApi();
    const { crtPath, keyPath } = data;
    state.dataForm.huiCrtPath = crtPath;
    state.dataForm.huiKeyPath = keyPath;
  } catch (e) {
    /* empty */
  }
};

const handleRestartServer = async () => {
  try {
    ElMessageBox.confirm("Are you sure to restart panel?", "Warning", {
      confirmButtonText: t("common.confirm"),
      cancelButtonText: t("common.cancel"),
      type: "warning",
    }).then(() => {
      restartServerApi();
      ElMessage.success(t("config.restartTip"));
    });
  } catch (e) {
    /* empty */
  }
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
