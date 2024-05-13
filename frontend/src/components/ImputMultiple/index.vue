<template>
  <div>
    <el-tag :key="tag" v-for="tag in tags" closable @close="handleClose(tag)">
      {{ item }}
    </el-tag>
    <el-input
      v-model="item"
      @keyup.enter="handleConfirm"
      @blur="handleConfirm"
    />
  </div>
</template>

<script setup lang="ts">
defineProps({
  tags: {
    required: false,
    type: Array,
    default: () => [],
  },
  handleConfirm: {
    required: false,
    type: Function,
  },
  handleClose: {
    required: false,
    type: Function,
  },
});

const state = reactive({
  item: "",
});

const { item } = toRefs(state);
</script>

<style scoped></style>
