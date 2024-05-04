<template>
  <div class="container">
    <el-input-number
      v-model="capacity"
      :placeholder="$t('account.quota')"
      :min="-1"
      :controls="false"
      :precision="0"
      clearable
      style="width: 214px"
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

const props = defineProps({
  quota: {
    type: Number as PropType<number>,
    required: true,
  },
  quotaTmp: {
    type: Number as PropType<number>,
    required: true,
  },
});

const state = reactive({
  capacity: 0,
  unit: "GB",
});

const { capacity, unit } = toRefs(state);

const units = ["Bytes", "KB", "MB", "GB", "TB", "PB"];

const emit = defineEmits(["update:quota"]);

const quota = useVModel(props, "quota", emit);

watch([capacity, unit], ([newC, newU]) => {
  quota.value = calculateBytes(newC, newU);
});

watch(
  () => props.quotaTmp,
  () => {
    capacity.value = formatStorageCapacity(props.quotaTmp);
    unit.value = formatStorageUnit(props.quotaTmp);
  }
);
</script>

<style lang="scss" scoped>
.container {
  display: flex;
}
</style>
