<template>
  <el-row :gutter="15" justify="center">
    <el-col v-if="isAdmin" :lg="14" :md="18" class="main">
      <el-col>
        <h3>系统设置</h3>
      </el-col>
      <el-col>
        <el-form
            ref="form"
            :model="data"
            :rules="rules"
            label-width="120px"
        >
          <el-form-item label="版权主体" prop="copyright">
            <el-input v-model.trim="data.copyright"/>
          </el-form-item>
          <el-form-item label="工信部备案号" prop="gov_no1">
            <el-input v-model.trim="data.gov_no1"/>
          </el-form-item>
          <el-form-item label="公安部备案号" prop="gov_no2">
            <el-input v-model.trim="data.gov_no2"/>
          </el-form-item>
          <el-form-item label="博客地址" prop="blog">
            <el-input v-model.trim="data.blog"/>
          </el-form-item>
          <el-form-item label="公众号地址" prop="mp_url">
            <el-input v-model.trim="data.mp_url"/>
          </el-form-item>
          <el-form-item label="公众号图片" prop="mp_img">
            <el-input v-model.trim="data.mp_img" placeholder="请填写远程图片地址"/>
          </el-form-item>
          <el-form-item label="联系邮箱" prop="mail">
            <el-input v-model.trim="data.mail" type="email"/>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="save">保存</el-button>
            <el-button @click="reset">重置</el-button>
          </el-form-item>
        </el-form>
      </el-col>
    </el-col>
    <el-col v-else :sm="8" class="error">
      <Auth></Auth>
    </el-col>
  </el-row>
</template>

<script lang="ts" setup>
// 公众号表单
import {FormInstance, FormRules} from "element-plus";
import {SysInfo, SysInfoRes, SysInfoSet} from "../../../api/sys";
import Auth from "@/components/status/auth.vue"

// 系统基本信息数据
const data = reactive<SysInfoRes>({
  copyright: "权限不足，无法查看",
  gov_no1: "权限不足，无法查看",
  gov_no2: "权限不足，无法查看",
  mp_url: "权限不足，无法查看",
  mp_img: "权限不足，无法查看",
  mail: "权限不足，无法查看",
  blog: "权限不足，无法查看",
})

// 权限
const isAdmin = ref<Boolean>(true)

// 表单DOM
const form = ref<FormInstance>()

// 表单规则
const rules: FormRules = reactive<FormRules>({
  copyright: [
    {required: true, message: '请输入版权主体', trigger: 'blur'},
  ],
})

// 获取系统信息
const getSysInfo = async function () {
  await SysInfo().then(res => {
    let {copyright, gov_no1, gov_no2, mp_url, mp_img, mail, blog} = res.data
    data.copyright = copyright
    data.gov_no1 = gov_no1
    data.gov_no2 = gov_no2
    data.mp_url = mp_url
    data.mp_img = mp_img
    data.mail = mail
    data.blog = blog
  }).catch(err => {
    if (err.code === 306) {
      isAdmin.value = false
    }
    ElNotification({
      title: '请求失败',
      message: err.message,
    })
  })
}

getSysInfo()

// 设置系统信息
const SetSysInfo = async function () {
  for (let v in data) {
    if (data[v as keyof SysInfoRes].length == 0) {
      data[v as keyof SysInfoRes] = "EMPTY"
    }
  }

  await SysInfoSet({
    ...data
  }).then(res => {
    ElNotification({
      message: "保存成功",
    })
  }).catch(err => {
    ElNotification({
      title: '请求失败',
      message: err.message,
    })
  })

  for (let v in data) {
    if (data[v as keyof SysInfoRes] == "EMPTY") {
      data[v as keyof SysInfoRes] = ""
    }
  }
}

// 保存
const save = async function () {
  if (!form.value) return
  await form.value.validate((valid) => {
    if (valid) {
      SetSysInfo()
    } else {
      ElNotification({
        title: "发现错误",
        message: "提交参数不正确,请检查表单",
      })
      return false
    }
  })
}

// 重置
const reset = async function () {
  await getSysInfo()
}
</script>

<style lang="scss" scoped>
.main {
  background-color: rgba(0, 0, 0, 0.4);
  border-radius: 8px;
  box-shadow: 0 0 8px rgba(0, 0, 0, 0.3);

  h3 {
    padding: 15px 30px;
  }
}

.error {
  margin: auto 0;
  background-color: rgba(0, 0, 0, 0.4);
  border-radius: 8px;
  box-shadow: 0 0 8px rgba(0, 0, 0, 0.3);
  overflow: hidden;
}
</style>