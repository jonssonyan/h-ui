<template>
  <el-form>
    <el-form-item>
      <el-button type="primary" @click="handleAdd">
        {{ $t("hysteria.addOutbound") }}
        <el-icon class="el-icon--right">+</el-icon>
      </el-button>
    </el-form-item>
    <el-form-item>

		ssfds</el-form-item>
  </el-form>
</template>

<script setup lang="ts">
const props = defineProps({
  outbounds: {
    required: false,
    type: Array,
    default: () => [],
  },
});

const emit = defineEmits<{
  (event: "update:outbounds", value: Array<string>): void;
}>();

const outbounds = useVModel(props, "outbounds", emit);

const handleAdd = () => {};
</script>

<style lang="scss" scoped></style>
