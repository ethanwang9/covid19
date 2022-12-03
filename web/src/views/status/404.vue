<template>
  <el-row align="middle" class="main" justify="center">
    <el-col>
      <el-result
          icon="error"
          sub-title="当前页面不存在"
          title="404"
      >
        <template #extra>
          <el-button size="large" type="primary" @click="backUp">返回首页</el-button>
        </template>
      </el-result>
    </el-col>
  </el-row>
</template>

<script lang="ts" setup>
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