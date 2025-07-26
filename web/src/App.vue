<template>
  <router-view></router-view>
</template>

<script lang="ts" setup>
import {useDark, useToggle} from '@vueuse/core'

// 切换暗黑主题
const isDark = useDark()
const toggleDark = useToggle(isDark)

// 默认为暗黑主题
if (isDark.value === false) {
  toggleDark()
}
</script>

<style lang="scss">
//导入全局公共样式
@use "@/styles/index.scss";

</style>