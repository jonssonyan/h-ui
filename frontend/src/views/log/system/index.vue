<template>
  <div class="app-container">
    <el-card shadow="never">
      <el-form label-position="top">
        <el-form-item label="自动刷新" prop="autoRefresh">
          <el-switch
            v-model="autoRefresh"
            active-value="1"
            inactive-value="0"
            @change="handleChange"
          />
        </el-form-item>
        <el-form-item label="过滤" prop="filter" style="width: 300px">
          <el-input v-model="filter" placeholder="过滤" clearable />
        </el-form-item>
      </el-form>
      <el-card shadow="never"></el-card>
    </el-card>
  </div>
</template>

<script setup lang="ts">
const state = reactive({
  autoRefresh: true,
  filter: "",
});

const handleChange = () => {};
const { autoRefresh, filter } = toRefs(state);
</script>

<style scoped></style>
