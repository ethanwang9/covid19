<template>
  <el-row justify="center" align="middle" class="main">
    <el-col>
      <el-result
          icon="error"
          title="权限不足"
          sub-title="您当前的用户权限不足, 无法访问该页面"
      >
        <template #extra>
          <el-button type="primary" size="large" @click="backUp">返回上一页</el-button>
        </template>
      </el-result>
    </el-col>
  </el-row>
</template>

<script setup lang="ts">
import {useRouter} from "vue-router";
// 路由
const router = useRouter()

// 返回上一页
function backUp() {
  router.go(-2)
}
</script>

<style lang="scss" scoped>
.main {
  width: 100vw;
  height: 100vh;
}
</style>