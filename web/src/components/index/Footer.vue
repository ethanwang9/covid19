<template>
  <el-row class="copyright">
    <!--        社交媒体-->
    <el-col :lg="8" :sm="10" :xs="24" class="copyright-icons">
      <img alt="RSS" src="@/assets/icon/blog.png" @click="GO(data.blog)">
      <el-popover :width="225" placement="top-start" trigger="hover">
        <template #reference>
          <img alt="微信" src="@/assets/icon/wechat.png" @click="GO(data.mp_url)">
        </template>
        <img :src="data.mp_img" alt="微信公众号图片" style="width: 200px; height: auto;" @click="GO(data.mp_url)">
      </el-popover>
      <img alt="邮箱" src="@/assets/icon/mail.png" @click="GO('mailto:'+data.mail)">
    </el-col>
    <!--        版权信息-->
    <el-col class="copyright-text">© 2022-2023 {{ data.copyright }} All Rights Reserved.</el-col>
    <el-col class="copyright-text hidden-xs-only">
      <span v-if="data.gov_no1.length !== 0"
            @click="GO('https://beian.miit.gov.cn/#/Integrated/recordQuery')">{{ data.gov_no1 }}</span>
      <span v-show="data.gov_no2.length !== 0 && data.gov_no1.length !== 0" style="padding: 0 10px">|</span>
      <span v-if="data.gov_no2.length !== 0" @click="GO('https://www.beian.gov.cn/portal/registerSystemInfo')">
        <img alt="公安部备案图表" src="@/assets/icon/gongan.png" style="padding-right: 10px"/>{{ data.gov_no2 }}
      </span>
    </el-col>
    <el-col class="copyright-text hidden-sm-and-up">
      <el-col v-if="data.gov_no1.length !== 0" @click="GO('https://beian.miit.gov.cn/#/Integrated/recordQuery')">
        {{ data.gov_no1 }}
      </el-col>
      <el-col v-if="data.gov_no2.length !== 0" @click="GO('https://www.beian.gov.cn/portal/registerSystemInfo')">
        <img alt="公安部备案图表" src="@/assets/icon/gongan.png" style="padding-right: 10px"/>{{ data.gov_no2 }}
      </el-col>
    </el-col>
  </el-row>
</template>

<script lang="ts" setup>
// 获取页面信息
import {PublicInfo, PublicInfoRes} from "../../api/public";
import {ElNotification} from "element-plus";

// 内容
const data = reactive<PublicInfoRes>({
  create_at: "",
  updated_at: "",
  copyright: "",
  gov_no1: "",
  gov_no2: "",
  mp_url: "",
  mp_img: "",
  mail: "",
  blog: "",
})

const GetInfo = async () => {
  await PublicInfo().then(res => {
    data.create_at = res.data.create_at
    data.updated_at = res.data.updated_at
    data.copyright = res.data.copyright
    data.gov_no1 = res.data.gov_no1
    data.gov_no2 = res.data.gov_no2
    data.mp_url = res.data.mp_url
    data.mp_img = res.data.mp_img
    data.mail = res.data.mail
    data.blog = res.data.blog
  }).catch(err => {
    ElNotification({
      title: '错误',
      message: err.message,
    })
  })
}
GetInfo()

// 跳转页面
const GO = (url: string) => {
  window.open(url)
}
</script>

<style lang="scss" scoped>
//版权
.copyright {
  margin-top: 220px;
  margin-bottom: 100px;
  display: flex;
  justify-content: center;
  text-align: center;

  &-icons {
    display: flex;
    align-items: center;
    justify-content: space-around;
    margin: 100px 0;
  }

  &-text {
    line-height: 50px;
    font-size: 18px;
    text-align: center;
    letter-spacing: 0.1em;
    color: #9E9E9E;
  }
}
</style>