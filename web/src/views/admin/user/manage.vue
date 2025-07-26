<template>
  <el-row :class="isAdmin?'main':''" :gutter="15" justify="center">
    <template v-if="isAdmin">
      <el-col>
        <h3>用户管理</h3>
      </el-col>
      <el-col>
        <el-form ref="formRef" :inline="true" :model="form" :rules="rules" class="form">
          <el-form-item label="UID" prop="uid">
            <el-input v-model.trim="form.uid" clearable placeholder="请输入用户UID"/>
          </el-form-item>
          <el-form-item label="昵称" prop="nickname">
            <el-input v-model.trim="form.nickname" clearable placeholder="请输入用户昵称"/>
          </el-form-item>
          <el-form-item label="权限" prop="level">
            <el-select v-model="form.level" clearable placeholder="请选择用户权限">
              <el-option label="用户" value="user"/>
              <el-option label="管理员" value="admin"/>
              <el-option label="禁用" value="stop"/>
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="Query">查询</el-button>
          </el-form-item>
        </el-form>
      </el-col>
      <el-col>
        <el-table
            :data="data.list"
            :highlight-current-row="true"
            empty-text="没有数据"
            style="width: 100%;padding: 0 30px"
        >
          <el-table-column fixed label="用户UID" prop="uid" sortable/>
          <el-table-column label="昵称" prop="nickname"/>
          <el-table-column label="头像" prop="avatar">
            <template #default="scope">
              <el-avatar :size="50" :src="scope.row.avatar" alt="用户头像" fit="cover" shape="square"></el-avatar>
            </template>
          </el-table-column>
          <el-table-column label="权限" prop="level" sortable>
            <template #default="scope">
              <el-tag v-if="scope.row.level === 'admin'" effect="dark" type="info">管理员</el-tag>
              <el-tag v-else-if="scope.row.level === 'user'" effect="dark">用户</el-tag>
              <el-tag v-else-if="scope.row.level === 'stop'" effect="dark" type="danger">禁用</el-tag>
            </template>
          </el-table-column>
          <el-table-column label="归属地" prop="location" sortable/>
          <el-table-column label="创建时间" prop="create_at" sortable/>
          <el-table-column label="更新时间" prop="updated_at" sortable/>
          <el-table-column fixed="right" label="操作" width="80">
            <template #default="scope">
              <el-button size="small" @click="openSide(scope.row.uid)">编辑</el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-col>
      <el-col class="pagination">
        <el-pagination
            :current-page="pagination.current"
            :hide-on-single-page="true"
            :page-size="pagination.limit"
            :pager-count="4"
            :total="data.total"
            background
            layout="prev, pager, next"
            @current-change="currentChange"
        />
      </el-col>
    </template>
    <el-col v-else :sm="8" class="error">
      <Auth></Auth>
    </el-col>
  </el-row>

  <!--  侧边栏内容-->
  <el-drawer
      v-model="isOpenSide"
      size="40%"
      title="编辑用户信息"
  >
    <el-row :gutter="15">
      <el-col class="side">
        <el-avatar :size="100" :src="sideData.avatar"/>
        <p class="side-nickname">{{ sideData.nickname }}</p>
      </el-col>
      <el-col style="text-align: center;padding-top: 30px;padding-bottom: 100px;">
        <el-select v-model="sideData.level" @change="changeUserLevel">
          <el-option v-show="sideData.level!=='admin'" label="管理员" value="admin"/>
          <el-option v-show="sideData.level!=='user'" label="用户" value="user"/>
          <el-option v-show="sideData.level!=='stop'" label="禁用" value="stop"/>
        </el-select>
      </el-col>
    </el-row>
  </el-drawer>
</template>

<script lang="ts" setup>
import Auth from "@/components/status/auth.vue"
import {UpdateUserLevel, UserList, UserListRes, UserQuery, UserQueryReq} from "../../../api/user"
import {FormInstance} from "element-plus";

// 权限
const isAdmin = ref<Boolean>(true)

// 页显示条数
const pagination = reactive({
  // 页数据条数
  limit: 10,
// 当前页码
  current: 1,
})

// 数据
const data = reactive<UserListRes>({
  list: [],
  total: 0,
})

// 是否打开侧边栏
const isOpenSide = ref<Boolean>(false)

// 侧边栏数据
const sideData = reactive({
  uid: "",
  avatar: "",
  nickname: "",
  level: "",
})

// 表单元素
const formRef = ref<FormInstance>()

// 表单数据
const form = reactive<UserQueryReq>({
  uid: "",
  nickname: "",
  level: "",
  limit: 10,
  page: 1,
})

// 表单规则
const rules = {
  uid: [
    {min: 1, max: 32, message: '用户ID长度是1-32位之间', trigger: 'blur'},
  ],
  nickname: [
    {min: 1, max: 50, message: '昵称长度在1-50位之间', trigger: 'blur'},
  ],
}

// 获取用户列表
const GetUserList = async () => {
  await UserList({
    limit: pagination.limit,
    page: pagination.current,
  }).then(res => {
    data.list = res.data.list
    data.total = res.data.total
  }).catch(err => {
    if (err.code === 306) {
      isAdmin.value = false
    }
    ElNotification({
      title: '错误',
      message: err.message,
    })
  })
}

GetUserList()

// 页码改变
const currentChange = (val: number) => {
  pagination.current = val
  let isF = false

  if (form.uid.length != 0 || form.level.length != 0 || form.nickname.length != 0) {
    isF = true
  }

  if (isF) {
    GetUserQuery()
  } else {
    GetUserList()
  }
}

// 复合条件查询用户信息
const GetUserQuery = async () => {
  await UserQuery({
    ...form,
    limit: pagination.limit,
    page: pagination.current,
  }).then(res => {
    data.list = res.data.list
    data.total = res.data.total
    ElNotification({
      message: "查询成功",
    })
  }).catch(err => {
    ElNotification({
      title: '错误',
      message: err.message,
    })
  })
}

// 查询用户
const Query = async () => {
  if (!formRef.value) return

  await formRef.value.validate((valid) => {
    if (valid) {
      if (form.uid.length === 0 && form.nickname.length === 0 && form.level.length === 0) {
        ElNotification({
          title: "发现错误",
          message: "请设置任意一样查询参数",
        })
        return
      }
      GetUserQuery()
    } else {
      ElNotification({
        title: "发现错误",
        message: "提交参数不正确,请检查表单",
      })
    }
  })
}

// 打开侧边栏
const openSide = (uid: string) => {
  isOpenSide.value = true
  data.list.map(v => {
    if (v.uid === uid) {
      sideData.uid = v.uid
      sideData.nickname = v.nickname
      sideData.avatar = v.avatar
      sideData.level = v.level
    }
  })
}

// 修改用户权限
const changeUserLevel = async (v: string) => {
  await UpdateUserLevel({
    uid: sideData.uid,
    level: sideData.level,
  }).then(() => {
    data.list.map(v => {
      if (v.uid === sideData.uid) {
        v.level = sideData.level
      }
    })
    ElNotification({
      message: "修改成功",
    })
  }).catch(err => {
    ElNotification({
      title: "错误",
      message: err.message,
    })
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

  .form {
    padding: 0 30px;
  }

  .pagination {
    display: flex;
    justify-content: center;
    padding: 15px 30px 30px;
    overflow: hidden;
  }
}

.side {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;

  p {
    font-size: 20px;
    font-weight: bold;
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