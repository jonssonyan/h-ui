<template>
  <div class="app-container">
    <div class="search">
      <el-form ref="queryFormRef" :model="queryParams" :inline="true">
        <el-form-item :label="$t('log.numLine')" prop="numLine">
          <el-select
            v-model="queryParams.numLine"
            style="width: 200px"
            @change="setRecords"
          >
            <el-option label="100" value="100" />
            <el-option label="200" value="200" />
            <el-option label="300" value="300" />
          </el-select>
        </el-form-item>
        <el-form-item prop="export">
          <el-button @click="handleExport">
            <template #icon>
              <i-ep-download />
            </template>
            {{ $t("common.export") }}
          </el-button>
        </el-form-item>
        <el-form-item prop="refresh">
          <el-button @click="setRecords">
            <template #icon>
              <i-ep-refresh />
            </template>
            {{ $t("common.refresh") }}
          </el-button>
        </el-form-item>
      </el-form>
    </div>

    <el-card shadow="never">
      <el-table v-loading="loading" :data="records">
        <el-table-column
          key="level"
          label="level"
          align="center"
          prop="level"
        />
        <el-table-column key="msg" label="msg" align="center" prop="msg" />
        <el-table-column key="time" label="time" align="center" prop="time" />
      </el-table>
    </el-card>
  </div>
</template>

<script lang="ts">
export default {
  name: "index",
};
</script>

<script setup lang="ts">
import { LogDto, LogHysteria2Vo } from "@/api/log/types";
import { exportLogApi, logHysteria2Api } from "@/api/log";
import { useI18n } from "vue-i18n";

const { t } = useI18n();

const state = reactive({
  loading: true,
  total: 0,
  records: [] as LogHysteria2Vo[],
  queryParams: {
    numLine: 100,
  } as LogDto,
});

const { loading, records, queryParams } = toRefs(state);

const setRecords = async () => {
  try {
    state.loading = true;
    const { data } = await logHysteria2Api(state.queryParams);
    state.records = data.records;
    state.total = data.total;
  } finally {
    state.loading = false;
  }
};

const handleExport = async () => {
  let response = await exportLogApi({ option: 1 });
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
    ElMessage.success(t("common.success"));
  } catch (e) {
    /* empty */
  }
};

onMounted(() => {
  setRecords();
});
</script>

<style lang="scss" scoped></style>
