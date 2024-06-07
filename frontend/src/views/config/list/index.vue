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
              导入
            </el-button>
          </el-upload>
        </el-form-item>
        <el-form-item>
          <el-button @click="handleExport">
            <template #icon>
              <i-ep-download />
            </template>
            导出
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
import {
  exportConfigApi,
  importConfigApi,
  importHysteria2ConfigApi,
  listConfigApi,
  updateConfigsApi,
} from "@/api/config";
import { ConfigsUpdateDto } from "@/api/config/types";
import type { FormInstance, FormRules } from "element-plus";
import {validateIntegerStr, validateNumberStr} from "@/utils/validate";
import {
  UploadFile,
  UploadRawFile,
  UploadRequestOptions,
} from "element-plus/lib/components";

const huiWebPortKey = "H_UI_WEB_PORT";
const hysteria2TrafficTimeKey = "HYSTERIA2_TRAFFIC_TIME";

const formDataRef = ref<FormInstance>();
const rules = reactive<FormRules>({
  huiWebPort: [
    {
      validator: validateIntegerStr,
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

const state = reactive({
  formData: {
    huiWebPort: "8081",
    hysteria2TrafficTime: "1",
  },
  fileList: [] as UploadFile[],
});

const { formData, fileList } = toRefs(state);

const submitForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.validate((valid) => {
    if (valid) {
      let configs: Array<ConfigsUpdateDto> = [
        {
          key: huiWebPortKey,
          value: state.formData.huiWebPort,
        },
        {
          key: hysteria2TrafficTimeKey,
          value: state.formData.hysteria2TrafficTime,
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
        state.formData.huiWebPort = configVo.value;
      } else if (configVo.key === hysteria2TrafficTimeKey) {
        state.formData.hysteria2TrafficTime = configVo.value;
      }
    });
  });
};

const handleImport = (params: UploadRequestOptions) => {
  if (state.fileList.length > 0) {
    let formData = new FormData();
    formData.append("file", params.file);
    importConfigApi(formData).then(() => {
      ElMessage.success("导入成功");
    });
    state.fileList = [];
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
  let response = await exportConfigApi();
  try {
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
    ElMessage.success("导出成功");
  } catch (err) {
    ElMessage.error("导出失败");
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
