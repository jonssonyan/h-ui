<script lang="ts">
export default { name: "Dashboard" };
</script>

<script setup lang="ts">
import { useTransition, TransitionPresets } from "@vueuse/core";

import SvgIcon from "@/components/SvgIcon/index.vue";

const duration = 5000;

// 收入金额
const amount = ref(0);
const amountOutput = useTransition(amount, {
  duration: duration,
  transition: TransitionPresets.easeOutExpo,
});
amount.value = 2000;

// 访问数
const visitCount = ref(0);
const visitCountOutput = useTransition(visitCount, {
  duration: duration,
  transition: TransitionPresets.easeOutExpo,
});
visitCount.value = 2000;

//消息数
const messageCount = ref(0);
const messageCountOutput = useTransition(messageCount, {
  duration: duration,
  transition: TransitionPresets.easeOutExpo,
});
messageCount.value = 2000;

// 订单数
const orderCount = ref(0);
const orderCountOutput = useTransition(orderCount, {
  duration: duration,
  transition: TransitionPresets.easeOutExpo,
});
orderCount.value = 2000;
</script>

<template>
  <div class="dashboard-container"></div>
</template>

<style lang="scss" scoped>
.dashboard-container {
  position: relative;
  padding: 24px;

  .user-avatar {
    width: 40px;
    height: 40px;
    border-radius: 50%;
  }

  .github-corner {
    position: absolute;
    top: 0;
    right: 0;
    z-index: 99;
    border: 0;
  }

  .data-box {
    display: flex;
    justify-content: space-between;
    padding: 20px;
    font-weight: bold;
    color: var(--el-text-color-regular);
    background: var(--el-bg-color-overlay);
    border-color: var(--el-border-color);
    box-shadow: var(--el-box-shadow-dark);
  }

  .svg-icon {
    fill: currentcolor !important;
  }
}
</style>
