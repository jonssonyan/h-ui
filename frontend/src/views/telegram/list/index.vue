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
            <el-input v-model="dataForm.telegramToken" clearable />
          </el-form-item>
        </el-tooltip>
        <el-tooltip
          :content="$t('telegram.telegramUsername')"
          placement="bottom"
        >
          <el-form-item
            :label="$t('telegram.telegramUsername')"
            prop="telegramUsername"
          >
            <el-input v-model="dataForm.telegramUsername" clearable />
          </el-form-item>
        </el-tooltip>
        <el-tooltip :content="$t('telegram.telegramJob')" placement="bottom">
          <el-form-item :label="$t('telegram.telegramJob')" prop="telegramJob">
          </el-form-item>
        </el-tooltip>
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

const dataFormRules = {
  telegramToken: [
    {
      required: true,
      message: "Required",
      trigger: ["change", "blur"],
    },
  ],
  telegramUsername: [
    {
      required: true,
      message: "Required",
      trigger: ["change", "blur"],
    },
  ],
};

const state = reactive({
  dataForm: {
    telegramEnable: 0,
    telegramToken: "",
    telegramUsername: "",
    telegramLoginJobEnable: 0,
    telegramLoginJobText: "",
  },
});

const { dataForm } = toRefs(state);
</script>

<style lang="scss" scoped>
.el-card .el-form {
  max-width: 1000px;
  margin: 0 auto;
}
</style>
