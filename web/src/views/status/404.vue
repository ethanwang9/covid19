<template>
  <el-row justify="center" align="middle" class="main">
    <el-col>
      <el-result
          icon="error"
          title="404"
          sub-title="当前页面不存在"
      >
        <template #extra>
          <el-button type="primary" size="large" @click="backUp">返回首页</el-button>
        </template>
      </el-result>
    </el-col>
  </el-row>
</template>

<script setup lang="ts">
import {useRouter} from "vue-router";
// 路由
const router = useRouter()

// 返回首页
function backUp() {
  router.replace({path: '/'})
}
</script>

<style lang="scss" scoped>
.main {
  width: 100vw;
  height: 100vh;
}
</style>