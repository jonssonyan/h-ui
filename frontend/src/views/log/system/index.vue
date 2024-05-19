<template>
  <div class="app-container">
    <div class="search">
      <el-form ref="queryFormRef" :model="queryParams" :inline="true">
        <el-form-item :label="$t('log.numLine')" prop="numLine">
          <el-select
            v-model="queryParams.numLine"
            :placeholder="$t('log.numLine')"
            style="width: 200px"
            @change="setRecords"
          >
            <el-option label="100" value="100" />
            <el-option label="200" value="200" />
            <el-option label="300" value="300" />
          </el-select>
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

<script setup lang="ts">
import { LogSystemDto, LogSystemVo } from "@/api/log/types";
import { logSystemApi } from "@/api/log";

const state = reactive({
  loading: true,
  total: 0,
  records: [] as LogSystemVo[],
  queryParams: {
    numLine: 100,
  } as LogSystemDto,
});

const { loading, records, queryParams } = toRefs(state);

const setRecords = () => {
  state.loading = true;
  logSystemApi(state.queryParams).then((response) => {
    const { data } = response;
    state.records = data.records;
    state.total = data.total;
    state.loading = false;
  });
};

onMounted(() => {
  setRecords();
});
</script>

<style lang="scss" scoped></style>
