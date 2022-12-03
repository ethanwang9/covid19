<template>
  <el-row :gutter="15" justify="center">
    <el-col :lg="14" :md="18" class="main">
      <el-col>
        <h3>个人资料</h3>
      </el-col>
      <el-col>
        <el-form
            label-width="120px"
        >
          <el-form-item label="用户UID">
            <el-input v-model="userInfo.uid" disabled readonly/>
          </el-form-item>
          <el-form-item label="用户昵称">
            <el-input v-model="userInfo.nickname" disabled readonly/>
          </el-form-item>
          <el-form-item label="用户头像">
            <el-avatar
                :size="100"
                :src="userInfo.avatar"
                alt="用户头像"
                fit="cover"
                shape="square"
            ></el-avatar>
          </el-form-item>
        </el-form>
      </el-col>
    </el-col>
  </el-row>
</template>

<script lang="ts" setup>
// 获取用户信息
import {UserBaseInfo} from "../../../api/user";

const userInfo = reactive({
  nickname: "无名氏",
  avatar: "",
  uid: "0",
})

async function init() {
  await UserBaseInfo().then(res => {
    userInfo.uid = res.data.uid
    userInfo.avatar = res.data.avatar
    userInfo.nickname = res.data.nickname
  }).catch(err => {
    ElNotification({
      title: '请求失败',
      message: "查询信息失败, " + err.message,
    })
  })
}

init()
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
</style>