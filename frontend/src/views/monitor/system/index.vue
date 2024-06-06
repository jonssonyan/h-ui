<template>
  <div class="dashboard-container">
    <el-row :gutter="10" class="mt-3">
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
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { monitorSystemApi } from "@/api/monitor";

const state = reactive({
  systemMonitor: {
    cpuPercent: 0,
    memPercent: 0,
    diskPercent: 0,
  },
});

const { systemMonitor } = toRefs(state);

const setSystemMonitor = () => {
  monitorSystemApi().then((response) => {
    const { data } = response;
    Object.assign(state.systemMonitor, data);
  });
};

onMounted(() => {
  setSystemMonitor();
});
</script>

<style lang="scss" scoped>
.dashboard-container {
  position: relative;
  padding: 24px;

  .svg-icon {
    fill: currentcolor !important;
  }
}
</style>
