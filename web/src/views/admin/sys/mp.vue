<template>
  <el-row :gutter="15" justify="center">
    <el-col v-if="isAdmin" :lg="14" :md="18" class="main">
      <el-col>
        <h3>公众号设置</h3>
      </el-col>
      <el-col>
        <el-form
            ref="form"
            :model="mp"
            :rules="rules"
            label-width="120px"
        >
          <el-form-item label="APPID" prop="appid">
            <el-input v-model.trim="mp.appid" clearable/>
          </el-form-item>
          <el-form-item label="SECRET" prop="secret">
            <el-input v-model.trim="mp.secret" clearable/>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="submitForm">保存</el-button>
            <el-button @click="reset()">重置</el-button>
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
import {SysMpInfo, SysMpInfoRes, SysMpInfoSet} from "../../../api/sys";
import {FormInstance, FormRules} from "element-plus";

// 公众号信息
const mp = reactive<SysMpInfoRes>({
  appid: "无",
  secret: "无",
})

// 公众号表单
const form = ref<FormInstance>()

// 表单规则
const rules: FormRules = reactive<FormRules>({
  appid: [
    {required: true, message: '请输入微信公众号APPID', trigger: 'blur'},
    {min: 10, max: 30, message: '长度在10-30', trigger: 'blur'},
  ],
  secret: [
    {required: true, message: '请输入微信公众号SECRET', trigger: 'blur'},
    {min: 20, max: 40, message: '长度在20-40', trigger: 'blur'},
  ],
})

// 权限
const isAdmin = ref<Boolean>(true)

// 获取公众号信息
async function getMpInfo() {
  SysMpInfo().then(res => {
    let {appid, secret} = res.data
    mp.appid = appid
    mp.secret = secret
  }).catch(err => {
    if (err.code === 306) {
      isAdmin.value = false
    }
    mp.appid = mp.secret = err.message
    ElNotification({
      title: '错误',
      message: err.message,
    })
  })
}

getMpInfo()

// 重置表单参数
async function reset() {
  await getMpInfo()
}

// 提交公众号信息
const submitForm = async () => {
  if (!form.value) return
  await form.value.validate((valid) => {
    if (valid) {
      // 设置公众号
      SysMpInfoSet({
        ...mp
      }).then(res => {
        ElNotification({
          message: res.message,
        })
      }).catch(err => {
        ElNotification({
          title: "请求失败",
          message: err.message,
        })
      })
    } else {
      ElNotification({
        title: "发现错误",
        message: "提交参数不正确,请检查表单",
      })
      return false
    }
  })
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