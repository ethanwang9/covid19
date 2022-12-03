<template>
  <el-row class="main" justify="center" align="middle">
    <el-col :xs="22" :sm="12" :md="18" :lg="14" class="login">
      <el-col class="login-left hidden-sm-and-down" :md="14">
        <img src="@/assets/images/index/login/login-bg.png" alt="login_bg">
      </el-col>
      <el-col class="login-right" :md="10">
        <el-col class="login-right-title">
          <h1>COVID-19 大数据可视化系统</h1>
          <p>可视化网络爬虫测绘网络空间疫情数据</p>
        </el-col>
        <el-col class="login-right-qr" v-if="!isWx">
          <qrcode-vue :value="qrInfo.content" :size="qrInfo.size" class="login-right-qr-content"></qrcode-vue>
          <p>微信扫码登录</p>
        </el-col>
        <el-col v-else style="margin-top: 50px;">
          <p>请在微信外打开该页面</p>
          <p><b>微信扫一扫</b> 或 <b>截图后微信二维码识别</b> 登录</p>
        </el-col>
      </el-col>
    </el-col>
  </el-row>
</template>

<script setup lang="ts">
import {LoginStatus, LoginUrl} from "../../api/login";
import QrcodeVue from 'qrcode.vue'
import useStore from "../../store";
import {useRoute, useRouter} from "vue-router";
import axios from "axios/index";
import request from "../../api";

// 二维码信息
const qrInfo = reactive({
  // 内容
  content: "",
  // 大小
  size: 200,
})

// store
const store = useStore().app

// 路由
const router = useRouter()

// 获取请求链接
async function getUrl() {
  let back:string = ""
  if(import.meta.env.DEV) {
    back = import.meta.env.VITE_URL
  }else {
    back = import.meta.env.VITE_SERVER_ADDR
  }
  await LoginUrl({
    back: back + "/api/auth/wx/login/token",
  }).then(res => {
    let {code, url} = res.data
    // 设置QR
    qrInfo.content = url
    // 保存code
    store.uuid = code
  }).catch(err => {
    ElNotification({
      title: '请求失败',
      message: "获取登录链接失败, " + +err.message,
    })
  })
}

// 获取登录状态
async function getStatus() {
  await LoginStatus({
    query: store.uuid,
  }).then(res => {
    let {status, token, uid, avatar, nickname, location} = res.data
    if (status) {
      // 设置数据
      store.token = token
      store.setUserInfo(uid, nickname, avatar, location)
      // 清楚轮询
      clearInterval(query)
      clearInterval(timeoutSet)
      ElNotification({
        title: '登录成功',
        message: "尊敬的用户，欢迎回来！",
      })
      router.push({name: "admin"})
    }
  }).catch(err => {
    if (err.code == 305) {
      // 用户已被禁用
      getUrl()
      ElNotification({
        title: '警告',
        message: err.message,
      })
    } else {
      ElNotification({
        title: '请求失败',
        message: "查询登录状态失败, " + err.message,
      })
    }
  })
}

// 更新请求链接
let timeoutSet = setInterval(() => {
  getUrl()
}, 2.5 * 60 * 1000)
// 轮询
let query = setInterval(() => {
  getStatus()
}, 2 * 1000)

// 判断是否是微信浏览器
const isWx = ref(false)
const isWxFun = () => {
  let ua = window.navigator.userAgent.toLowerCase()
  let isWX = ua.indexOf('micromessenger') != -1;
  if (isWX) {
    isWx.value = true
  }
}
isWxFun()

if (!isWx.value) {
  getUrl()
} else {
  clearInterval(timeoutSet)
  clearInterval(query)
}

</script>

<style lang="scss" scoped>
.main {
  width: 100vw;
  height: 100vh;
}

// 登录表单样式
.login {
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 0 8px rgba(0, 0, 0, 0.3);
  padding: 50px;
  color: #000;
  display: flex;
  justify-content: space-around;
  align-items: center;
  box-sizing: content-box;

  &-left {
    text-align: center;

    img {
      width: 80%;
    }
  }

  &-right {
    text-align: center;

    &-title {
      h1 {
        font-size: 20px;
        margin: 0;
        letter-spacing: 1px;
        padding-bottom: 10px;
      }

      p {
        margin: 0;
        letter-spacing: 4px;
        font-size: 14px;
        color: gray;
      }
    }

    &-qr {
      margin-top: 50px;

      &-content {
        width: 200px;
        height: 200px;
      }

      p {
        margin: 0;
        letter-spacing: 1px;
      }
    }
  }
}

// 适应小屏幕
@media screen and (max-width: 992px) {
  .login {
    padding: 30px;
    box-sizing: border-box;
  }
}
</style>