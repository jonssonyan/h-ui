<template>
  <div style="display: flex; align-items: center">
    <el-input-number
      v-model="capacity"
      placeholder="Please enter a value"
      :min="-1"
      :controls="false"
      :precision="0"
      clearable
      style="width: 220px"
    />
    <el-select
      v-model="unit"
      :placeholder="$t('account.unit')"
      style="width: 100px"
    >
      <el-option
        v-for="item in units"
        :key="item"
        :label="item"
        :value="item"
      />
    </el-select>
  </div>
</template>

<script setup lang="ts">
import { PropType } from "vue";
import {
  calculateBytes,
  formatStorageCapacity,
  formatStorageUnit,
} from "@/utils/byte";

const units = ["Bytes", "KB", "MB", "GB", "TB", "PB"];

const props = defineProps({
  valueTmp: {
    type: Number as PropType<number>,
    required: true,
  },
  setValue: {
    type: Function as PropType<(newValue: number) => void>,
    required: true,
  },
});

const state = reactive({
  capacity: 0,
  unit: "GB",
});

const { capacity, unit } = toRefs(state);

watch(
  [capacity, unit],
  ([newC, newU]) => {
    const newValue = calculateBytes(newC, newU);
    props.setValue(newValue);
  },
  { immediate: true }
);

watch(
  () => props.valueTmp,
  (newValue) => {
    state.capacity = formatStorageCapacity(newValue);
    state.unit = formatStorageUnit(newValue);
  },
  { immediate: true }
);
</script>

<style lang="scss" scoped></style>
