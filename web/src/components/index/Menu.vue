<template>
  <el-row class="menu">
    <el-col class="menu-left hidden-sm-and-down" :md="10">
      <img src="@/assets/icon/du.png" alt="logo">
      <h3>COVID-19 大数据可视化系统</h3>
    </el-col>
    <el-col class="menu-right" :sm="24" :md="10">
      <el-col class="menu-right-btns" :xs="16" :sm="10" :md="14">
        <el-button text dark @click="goto('/')">首页</el-button>
        <el-button text dark @click="goto('/publicity')">防疫宣传</el-button>
      </el-col>
      <el-col class="menu-right-otns" :xs="6" :sm="10" :md="6">
        <el-button class="login" color="#ffffff" dark @click="goto('/login')">进入系统</el-button>
      </el-col>
    </el-col>
  </el-row>
</template>

<script setup lang="ts">
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