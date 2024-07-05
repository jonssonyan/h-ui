<template>
  <div class="flex gap-2">
    <el-tag
      :key="item"
      v-for="item in tags"
      closable
      @close="handleClose(item)"
      size="large"
    >
      {{ item }}
    </el-tag>
    <el-input
      v-if="inputVisible"
      ref="inputRef"
      v-model="tag"
      class="w-50"
      @keyup.enter="handleConfirm"
      @blur="handleConfirm"
    />
    <el-button v-else @click="showInput">+</el-button>
  </div>
</template>

<script setup lang="ts">
import { ElInput } from "element-plus";
import { PropType } from "vue";

const props = defineProps({
  tags: {
    required: false,
    type: Array as PropType<string[]>,
    default: () => [],
  },
});

const emit = defineEmits<{
  (event: "update:tags", value: string[]): void;
}>();

const tags = useVModel(props, "tags", emit);

const inputRef = ref(ElInput);

const state = reactive({
  tag: "",
  inputVisible: false,
});

const { tag, inputVisible } = toRefs(state);

const showInput = () => {
  state.inputVisible = true;
  nextTick(() => {
    inputRef.value!.input!.focus();
  });
};
const handleConfirm = (): void => {
  const newTag = state.tag.trim();
  if (newTag && !tags.value?.includes(newTag)) {
    tags.value?.push(newTag);
    state.tag = "";
  }
  state.inputVisible = false;
};

const handleClose = (tag: string): void => {
  const index = tags.value?.indexOf(tag.trim());
  if (index !== -1) {
    tags.value?.splice(index, 1);
  }
};
</script>

<style lang="scss" scoped>
.flex.gap-2 {
  flex-wrap: wrap;
}
</style>
