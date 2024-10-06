<template>
  <div class="app-container">
    <div class="search">
      <el-form :inline="true">
        <el-form-item>
          <el-button type="primary" @click="submitForm" :icon="Select">
            {{ $t("common.save") }}
          </el-button>
        </el-form-item>
      </el-form>
    </div>

    <el-card shadow="never">
      <el-form
        ref="dataFormRef"
        :rules="dataFormRules"
        label-position="top"
        :model="dataForm"
      >
        <el-tooltip :content="$t('telegram.telegramEnable')" placement="bottom">
          <el-form-item prop="telegramEnable">
            <el-switch
              v-model="dataForm.telegramEnable"
              active-value="1"
              inactive-value="0"
              :active-text="$t('telegram.enable')"
              :inactive-text="$t('telegram.disable')"
              style="height: 32px"
            />
          </el-form-item>
        </el-tooltip>
        <el-tooltip :content="$t('telegram.telegramToken')" placement="bottom">
          <el-form-item
            :label="$t('telegram.telegramToken')"
            prop="telegramToken"
          >
            <el-input
              v-model="dataForm.telegramToken"
              type="password"
              clearable
            />
          </el-form-item>
        </el-tooltip>
        <el-tooltip :content="$t('telegram.telegramChatId')" placement="bottom">
          <el-form-item
            :label="$t('telegram.telegramChatId')"
            prop="telegramChatId"
          >
            <el-input v-model="dataForm.telegramChatId" clearable />
          </el-form-item>
        </el-tooltip>
        <el-form-item :label="$t('telegram.telegramJob')" prop="telegramJob">
          <el-tabs class="tabs">
            <el-tab-pane name="telegramLoginJob">
              <template #label>
                <span>
                  {{ $t("telegram.telegramLoginJob") }}
                </span>
                <el-badge
                  is-dot
                  v-show="
                    dataForm.telegramLoginJobEnable &&
                    dataForm.telegramLoginJobEnable === '1'
                  "
                />
              </template>
              <el-tooltip
                :content="$t('telegram.telegramLoginJobEnable')"
                placement="bottom"
              >
                <el-form-item prop="telegramLoginJobEnable">
                  <el-switch
                    v-model="dataForm.telegramLoginJobEnable"
                    active-value="1"
                    inactive-value="0"
                    :active-text="$t('telegram.enable')"
                    :inactive-text="$t('telegram.disable')"
                    style="height: 32px; margin-bottom: 8px"
                  />
                </el-form-item>
              </el-tooltip>
              <el-tooltip
                :content="$t('telegram.telegramLoginJobText')"
                placement="bottom"
              >
                <el-form-item
                  :label="$t('telegram.telegramLoginJobText')"
                  prop="telegramLoginJobText"
                >
                  <el-input
                    v-model="dataForm.telegramLoginJobText"
                    type="textarea"
                    clearable
                  />
                </el-form-item>
              </el-tooltip>
              <el-tooltip
                :content="$t('telegram.placeholder')"
                placement="bottom"
              >
                <el-form-item>
                  <el-alert type="info" :closable="false">
                    <p>[time] - 时间</p>
                    <p>[username] - 用户名</p>
                    <p>[ip] - IP地址</p>
                  </el-alert>
                </el-form-item>
              </el-tooltip>
            </el-tab-pane>
          </el-tabs>
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
import { listConfigApi, updateConfigsApi } from "@/api/config";
import { ConfigsUpdateDto } from "@/api/config/types";
import { useI18n } from "vue-i18n";

const { t } = useI18n();

const dataFormRef = ref(ElForm);

const telegramEnable = "TELEGRAM_ENABLE";
const telegramToken = "TELEGRAM_TOKEN";
const telegramChatId = "TELEGRAM_CHAT_ID";
const telegramLoginJobEnable = "TELEGRAM_LOGIN_JOB_ENABLE";
const telegramLoginJobText = "TELEGRAM_LOGIN_JOB_TEXT";

const dataFormRules = {
  telegramToken: [
    {
      required: true,
      message: "Required",
      trigger: ["change", "blur"],
    },
  ],
  telegramChatId: [
    {
      required: true,
      message: "Required",
      trigger: ["change", "blur"],
    },
  ],
};

const state = reactive({
  dataForm: {
    telegramEnable: "0",
    telegramToken: "",
    telegramChatId: "",
    telegramLoginJobEnable: "0",
    telegramLoginJobText: "",
  },
});

const { dataForm } = toRefs(state);

const submitForm = () => {
  dataFormRef.value.validate((valid: boolean) => {
    if (valid) {
      if (
        state.dataForm.telegramEnable &&
        state.dataForm.telegramEnable === "1"
      ) {
        if (!state.dataForm.telegramToken || !state.dataForm.telegramChatId) {
          ElMessage.error("token and chatId required");
          return;
        }
      }

      if (
        state.dataForm.telegramLoginJobEnable &&
        state.dataForm.telegramLoginJobEnable === "1"
      ) {
        if (!state.dataForm.telegramLoginJobText) {
          ElMessage.error("text required");
          return;
        }
      }

      let configs: ConfigsUpdateDto[] = [
        {
          key: telegramEnable,
          value: state.dataForm.telegramEnable,
        },
        {
          key: telegramToken,
          value: state.dataForm.telegramToken,
        },
        {
          key: telegramChatId,
          value: state.dataForm.telegramChatId,
        },
        {
          key: telegramLoginJobEnable,
          value: state.dataForm.telegramLoginJobEnable,
        },
        {
          key: telegramLoginJobText,
          value: state.dataForm.telegramLoginJobText,
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
    keys: [
      telegramEnable,
      telegramToken,
      telegramChatId,
      telegramLoginJobEnable,
      telegramLoginJobText,
    ],
  });

  data.forEach((configVo) => {
    if (configVo.key === telegramEnable) {
      state.dataForm.telegramEnable = configVo.value;
    } else if (configVo.key === telegramToken) {
      state.dataForm.telegramToken = configVo.value;
    } else if (configVo.key === telegramChatId) {
      state.dataForm.telegramChatId = configVo.value;
    } else if (configVo.key === telegramLoginJobEnable) {
      state.dataForm.telegramLoginJobEnable = configVo.value;
    } else if (configVo.key === telegramLoginJobText) {
      state.dataForm.telegramLoginJobText = configVo.value;
    }
  });
};
onMounted(() => {
  setConfig();
});
</script>

<style lang="scss" scoped>
.el-card .el-form {
  max-width: 1000px;
  margin: 0 auto;
}

.el-tabs {
  width: 100%;
}

.el-alert {
  margin: 20px 0 0;
}
</style>
