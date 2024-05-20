<template>
  <div class="dashboard-container">
    <el-card shadow="never">
      <el-row justify="space-between">
        <el-col :span="8" :xs="24">
          <div class="flex h-full items-center">
            <img
              class="w-20 h-20 mr-5 rounded-full"
              src="/src/assets/logo.png"
            />
            <div>
              <p>{{ greetings }}</p>
              <p class="text-sm text-gray">
                {{ $t("account.createTime") }}:
                {{ timestampToDateTime(account.createTime) }}
              </p>
            </div>
          </div>
        </el-col>

        <el-col :span="16" :xs="24">
          <div class="flex h-full items-center">
            <el-dropdown>
              <el-button size="large" type="primary" :icon="Share">
                订 阅 链 接
              </el-button>
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item>Shadowrocket</el-dropdown-item>
                  <el-dropdown-item>v2rayN</el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown>

            <el-button size="large" type="primary" :icon="Share">
              节 点 URL
            </el-button>
          </div>
        </el-col>
      </el-row>
    </el-card>

    <el-row :gutter="10" class="mt-3">
      <el-col :xs="24" :sm="12" :lg="6">
        <el-card shadow="never">
          <template #header>
            <div class="flex items-center justify-between">
              <span class="text-[var(--el-text-color-secondary)]">
                {{ $t("account.quota") }}
              </span>
              <el-tag type="success"
                >{{ formatStorageUnit(account.quota) }}
              </el-tag>
            </div>
          </template>
          <div class="flex items-center justify-between mt-5">
            <div class="text-lg text-right">
              {{ formatBytes(account.quota) }}
            </div>
            <svg-icon icon-class="quota" size="2em" />
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :lg="6">
        <el-card shadow="never">
          <template #header>
            <div class="flex items-center justify-between">
              <span class="text-[var(--el-text-color-secondary)]">
                {{ $t("account.download") }}
              </span>
              <el-tag type="success"
                >{{ formatStorageUnit(account.download) }}
              </el-tag>
            </div>
          </template>
          <div class="flex items-center justify-between mt-5">
            <div class="text-lg text-right">
              {{ formatBytes(account.download) }}
            </div>
            <svg-icon icon-class="download" size="2em" />
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :lg="6">
        <el-card shadow="never">
          <template #header>
            <div class="flex items-center justify-between">
              <span class="text-[var(--el-text-color-secondary)]">
                {{ $t("account.upload") }}
              </span>
              <el-tag type="success"
                >{{ formatStorageUnit(account.upload) }}
              </el-tag>
            </div>
          </template>
          <div class="flex items-center justify-between mt-5">
            <div class="text-lg text-right">
              {{ formatBytes(account.upload) }}
            </div>
            <svg-icon icon-class="upload" size="2em" />
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :lg="6">
        <el-card shadow="never">
          <template #header>
            <div class="flex items-center justify-between">
              <span class="text-[var(--el-text-color-secondary)]">
                {{ $t("account.expireTime") }}
              </span>
              <el-tag type="success">{{ $t("info.expireTime") }}</el-tag>
            </div>
          </template>
          <div class="flex items-center justify-between mt-5">
            <div class="text-lg text-right">
              {{ timestampToDateTime(account.expireTime) }}
            </div>
            <svg-icon icon-class="expire-time" size="2em" />
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { getAccountApi } from "@/api/account";
import { AccountVo } from "@/api/account/types";
import { useAccountStore } from "@/store/modules/account";
import { timestampToDateTime } from "@/utils/time";
import { formatBytes, formatStorageUnit } from "@/utils/byte";
import { Share } from "@element-plus/icons-vue";

const accountStore = useAccountStore();

const date: Date = new Date();

const greetings = computed(() => {
  const hours = date.getHours();
  if (hours >= 6 && hours < 8) {
    return "微凉扑面，清新的空气，唤醒一天的活力🌅！";
  } else if (hours >= 8 && hours < 12) {
    return "上午好，" + state.account.username + "！";
  } else if (hours >= 12 && hours < 18) {
    return "下午好，" + state.account.username + "！";
  } else if (hours >= 18 && hours < 24) {
    return "晚上好，" + state.account.username + "！";
  } else if (hours >= 0 && hours < 6) {
    return "我愿成为流星，划破黑夜，只为照亮你的梦境，晚安🌛！";
  }
});

const state = reactive({
  account: {} as AccountVo,
});

const { account } = toRefs(state);

onMounted(() => {
  getAccountApi({ id: accountStore.id }).then((response) => {
    Object.assign(state.account, response.data);
  });
});
</script>

<style lang="scss" scoped>
.dashboard-container {
  position: relative;
  padding: 24px;

  .user-avatar {
    width: 40px;
    height: 40px;
    border-radius: 50%;
  }

  .github-corner {
    position: absolute;
    top: 0;
    right: 0;
    z-index: 1;
    border: 0;
  }

  .data-box {
    display: flex;
    justify-content: space-between;
    padding: 20px;
    font-weight: bold;
    color: var(--el-text-color-regular);
    background: var(--el-bg-color-overlay);
    border-color: var(--el-border-color);
    box-shadow: var(--el-box-shadow-dark);
  }

  .svg-icon {
    fill: currentcolor !important;
  }
}

.flex.h-full.items-center {
  .el-button {
    margin: 10px;
  }
}

@media (max-width: 768px) {
  .flex.h-full.items-center {
    justify-content: center;
  }
}

@media (max-width: 375px) {
  .flex.h-full.items-center {
    flex-direction: column;
  }
}
</style>
