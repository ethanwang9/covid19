<template>
  <el-row align="middle" class="main" justify="center">
    <el-col>
      <el-result
          icon="error"
          sub-title="您当前的用户权限不足, 无法访问该页面"
          title="权限不足"
      >
        <template #extra>
          <el-button size="large" type="primary" @click="backUp">返回上一页</el-button>
        </template>
      </el-result>
    </el-col>
  </el-row>
</template>

<script lang="ts" setup>
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