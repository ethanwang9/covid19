<template>
  <el-row class="menu">
    <el-col :md="10" class="menu-left hidden-sm-and-down">
      <img alt="logo" src="@/assets/icon/du.png">
      <h3>COVID-19 大数据可视化系统</h3>
    </el-col>
    <el-col :md="10" :sm="24" class="menu-right">
      <el-col :md="14" :sm="10" :xs="16" class="menu-right-btns">
        <el-button dark text @click="goto('/')">首页</el-button>
        <el-button dark text @click="goto('/publicity')">防疫宣传</el-button>
      </el-col>
      <el-col :md="6" :sm="10" :xs="6" class="menu-right-otns">
        <el-button class="login" color="#ffffff" dark @click="goto('/login')">进入系统</el-button>
      </el-col>
    </el-col>
  </el-row>
</template>

<script lang="ts" setup>
import {useRouter} from "vue-router";

const router = useRouter()

// 调整页面
function goto(path: string) {
  router.push(path)
}
</script>

<style lang="scss" scoped>
.menu {
  height: 70px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 20px;
  overflow: hidden;
  box-sizing: border-box;

  &-left {
    display: flex;
    align-items: center;

    img {
      width: 30px;
      height: 30px;
      margin-right: 10px;
    }
  }

  &-right {
    display: flex;
    justify-content: space-between;
    align-items: center;


    &-btns {
      display: flex;
      justify-content: flex-start;

      // 修改element按钮样式
      .el-button.is-text {
        color: #ffffff;
        font-size: 18px;
        letter-spacing: 1px;
        font-weight: bold;
      }
    }

    &-otns {
      text-align: right;

      .login {
        height: 38px;
        padding: 15px 20px;
        font-weight: bold;
        color: #000;
        font-size: 16px;
        letter-spacing: 1px;
      }
    }
  }
}
</style>