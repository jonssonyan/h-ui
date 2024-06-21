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
  importConfigApi,
  listConfigApi,
  updateConfigsApi,
} from "@/api/config";
import { ConfigsUpdateDto } from "@/api/config/types";
import {
  UploadFile,
  UploadRawFile,
  UploadRequestOptions,
} from "element-plus/lib/components";
import { useI18n } from "vue-i18n";

const huiWebPortKey = "H_UI_WEB_PORT";
const hysteria2TrafficTimeKey = "HYSTERIA2_TRAFFIC_TIME";

const { t } = useI18n();

const dataFormRef = ref(ElForm);

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
  },
  fileList: [] as UploadFile[],
});

const { dataForm, fileList } = toRefs(state);

const submitForm = () => {
  dataFormRef.value.validate((valid: boolean) => {
    if (valid) {
      let configs: ConfigsUpdateDto[] = [
        {
          key: huiWebPortKey,
          value: state.dataForm.huiWebPort,
        },
        {
          key: hysteria2TrafficTimeKey,
          value: state.dataForm.hysteria2TrafficTime,
        },
      ];
      updateConfigsApi({ configUpdateDtos: configs }).then(() => {
        ElMessage.success(t("common.saveSuccess"));
      });
    }
  });
};

const setConfig = async () => {
  const { data } = await listConfigApi({
    keys: [huiWebPortKey, hysteria2TrafficTimeKey],
  });

  data.forEach((configVo) => {
    if (configVo.key === huiWebPortKey) {
      state.dataForm.huiWebPort = configVo.value;
    } else if (configVo.key === hysteria2TrafficTimeKey) {
      state.dataForm.hysteria2TrafficTime = configVo.value;
    }
  });
};

const handleImport = async (params: UploadRequestOptions) => {
  if (state.fileList.length > 0) {
    try {
      let formData = new FormData();
      formData.append("file", params.file);
      await importConfigApi(formData);
      ElMessage.success(t("common.exportSuccess"));
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
    ElMessage.success(t("common.exportSuccess"));
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
