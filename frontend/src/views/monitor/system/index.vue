<template>
  <div class="dashboard-container">
    <el-row :gutter="10" class="mt-3">
      <el-col :xs="24" :sm="12" :lg="6">
        <el-card shadow="never">
          <template #header>
            <div class="flex items-center justify-between">
              <span class="text-[var(--el-text-color-secondary)]">
                {{ $t("monitor.huiVersion") }}
              </span>
              <el-tag type="success">version</el-tag>
            </div>
          </template>
          <div class="flex items-center justify-between mt-5">
            <div class="text-lg text-right">
              {{ systemMonitor.huiVersion ? systemMonitor.huiVersion : "-" }}
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :lg="6">
        <el-card shadow="never">
          <template #header>
            <div class="flex items-center justify-between">
              <span class="text-[var(--el-text-color-secondary)]">
                {{ $t("monitor.cpuPercent") }}
              </span>
              <el-tag type="success">%</el-tag>
            </div>
          </template>
          <div class="flex items-center justify-between mt-5">
            <div class="text-lg text-right">
              {{
                systemMonitor.cpuPercent ? systemMonitor.cpuPercent + "%" : "-"
              }}
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :lg="6">
        <el-card shadow="never">
          <template #header>
            <div class="flex items-center justify-between">
              <span class="text-[var(--el-text-color-secondary)]">
                {{ $t("monitor.memPercent") }}
              </span>
              <el-tag type="success">%</el-tag>
            </div>
          </template>
          <div class="flex items-center justify-between mt-5">
            <div class="text-lg text-right">
              {{
                systemMonitor.memPercent ? systemMonitor.memPercent + "%" : "-"
              }}
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :lg="6">
        <el-card shadow="never">
          <template #header>
            <div class="flex items-center justify-between">
              <span class="text-[var(--el-text-color-secondary)]">
                {{ $t("monitor.diskPercent") }}
              </span>
              <el-tag type="success">%</el-tag>
            </div>
          </template>
          <div class="flex items-center justify-between mt-5">
            <div class="text-lg text-right">
              {{
                systemMonitor.diskPercent
                  ? systemMonitor.diskPercent + "%"
                  : "-"
              }}
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :lg="6">
        <el-card shadow="never">
          <template #header>
            <div class="flex items-center justify-between">
              <span class="text-[var(--el-text-color-secondary)]">
                {{ $t("monitor.hysteria2Version") }}
              </span>
              <el-tag type="success">version</el-tag>
            </div>
          </template>
          <div class="flex items-center justify-between mt-5">
            <div class="text-lg text-right">
              {{ hysteria2Monitor.version ? hysteria2Monitor.version : "-" }}
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :lg="6">
        <el-card shadow="never">
          <template #header>
            <div class="flex items-center justify-between">
              <span class="text-[var(--el-text-color-secondary)]">
                {{ $t("monitor.hysteria2Running") }}
              </span>
              <el-tag type="success">running</el-tag>
            </div>
          </template>
          <div class="flex items-center justify-between mt-5">
            <div
              class="text-lg text-right"
              :style="
                hysteria2Monitor.running === undefined
                  ? '-'
                  : hysteria2Monitor.running
                  ? 'color: #2ecc71'
                  : 'color: #e74c3c'
              "
            >
              {{
                hysteria2Monitor.running === undefined
                  ? "-"
                  : hysteria2Monitor.running
                  ? $t("monitor.hysteria2RunningTrue")
                  : $t("monitor.hysteria2RunningFalse")
              }}
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :lg="6">
        <el-card shadow="never">
          <template #header>
            <div class="flex items-center justify-between">
              <span class="text-[var(--el-text-color-secondary)]">
                {{ $t("monitor.hysteria2UserTotal") }}
              </span>
              <el-tag type="success">account</el-tag>
            </div>
          </template>
          <div class="flex items-center justify-between mt-5">
            <div class="text-lg text-right">
              {{ hysteria2Monitor.userTotal }}
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :lg="6">
        <el-card shadow="never">
          <template #header>
            <div class="flex items-center justify-between">
              <span class="text-[var(--el-text-color-secondary)]">
                {{ $t("monitor.hysteria2DeviceTotal") }}
              </span>
              <el-tag type="success">device</el-tag>
            </div>
          </template>
          <div class="flex items-center justify-between mt-5">
            <div class="text-lg text-right">
              {{ hysteria2Monitor.deviceTotal }}
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts">
export default {
  name: "index",
};
</script>

<script setup lang="ts">
import { monitorHysteria2Api, monitorSystemApi } from "@/api/monitor";

const state = reactive({
  systemMonitor: {
    huiVersion: "",
    cpuPercent: 0,
    memPercent: 0,
    diskPercent: 0,
  },
  hysteria2Monitor: {
    userTotal: 0,
    deviceTotal: 0,
    version: undefined,
    running: undefined,
  },
});

const { systemMonitor, hysteria2Monitor } = toRefs(state);

const setMonitor = () => {
  monitorSystemApi().then((response) => {
    const { data } = response;
    Object.assign(state.systemMonitor, data);
  });
  monitorHysteria2Api().then((response) => {
    const { data } = response;
    Object.assign(state.hysteria2Monitor, data);
  });
};

onMounted(() => {
  setMonitor();
});
</script>

<style lang="scss" scoped>
.dashboard-container {
  position: relative;
  padding: 24px;

  .svg-icon {
    fill: currentcolor !important;
  }

  .el-col {
    margin-bottom: 10px;
  }
}
</style>
