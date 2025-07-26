<template>
  <el-container>
    <!--    菜单-->
    <el-header v-show="isShowHeader" class="header" height="60px">
      <el-row align="middle" justify="center" style="height: 100%">
        <el-col :md="8" class="header-logo hidden-sm-and-down">
          <h3>COVID-19 大数据可视化系统</h3>
        </el-col>
        <el-col :md="15" class="header-bar">
          <el-row align="middle" justify="center">
            <el-col :sm="18" :xs="14" class="header-bar-menu">
              <el-menu
                  :default-active="menuActive"
                  mode="horizontal"
                  router
              >
                <el-menu-item index="/admin/main">数据看板</el-menu-item>
                <el-sub-menu index="/admin/query">
                  <template #title>在线查询</template>
                  <el-menu-item index="/admin/query/danger">风险地区</el-menu-item>
                  <el-menu-item index="/admin/query/travel">出行政策</el-menu-item>
                </el-sub-menu>
                <el-sub-menu index="/admin/user">
                  <template #title>用户管理</template>
                  <el-menu-item index="/admin/user/self">个人资料</el-menu-item>
                  <el-menu-item index="/admin/user/manage">管理用户</el-menu-item>
                </el-sub-menu>
                <el-sub-menu index="/admin/sys">
                  <template #title>系统管理</template>
                  <el-menu-item index="/admin/sys/index">系统设置</el-menu-item>
                  <el-menu-item index="/admin/sys/mp">公众号设置</el-menu-item>
                </el-sub-menu>
              </el-menu>
            </el-col>
            <el-col :sm="6" :xs="10" class="header-bar-user">
              <el-dropdown>
                <span class="el-dropdown-link header-bar-user-box">
                  <el-avatar
                      :src="userInfo.avatar"
                      class="hidden-xs-only"
                  />
                  <p>{{ userInfo.nickname }}</p>
                  <i-ep-arrow-down></i-ep-arrow-down>
                </span>
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item @click="goUserSelf">个人资料</el-dropdown-item>
                    <el-dropdown-item divided @click="logout">退出登录</el-dropdown-item>
                  </el-dropdown-menu>
                </template>
              </el-dropdown>
            </el-col>
          </el-row>
        </el-col>
      </el-row>
    </el-header>
    <!--    内容-->
    <el-main :style="adminMainCustom" class="content">
      <router-view></router-view>
    </el-main>
  </el-container>

  <!--  返回顶部-->
  <el-backtop :bottom="100" target=".content">
    <div class="TOTOP">UP</div>
  </el-backtop>
</template>

<script lang="ts" setup>
import {useRoute, useRouter} from "vue-router";
import useStore from "../../store";
import {storeToRefs} from "pinia";

// 路由
const router = useRouter()
const route = useRoute()

// store
const store = useStore().admin
const {isShowHeader, adminMainCustom} = storeToRefs(store)

// 获取用户信息
const userInfo = useStore().app.userInfo

// 菜单默认激活下标
const menuActive = router.currentRoute.value.path

// 查看用户信息
function goUserSelf() {
  router.push({name: 'adminUserSelf'})
}

// 退出登录
function logout() {
  useStore().app.clear()
  router.push({name: "login"})
}

// 设置可视化界面
function adminPageCfg() {
  if (route.name === "adminMain") {
    store.changeShowHeader(false)
  } else {
    store.changeShowHeader(true)
  }
}

adminPageCfg()

onUpdated(() => {
  adminPageCfg()
})
</script>

<style lang="scss" scoped>
.header {
  overflow: hidden;
  border-bottom: 2px solid #5C47FF;

  &-logo {
    h3 {
      margin: 0;
      letter-spacing: 2px;
    }
  }

  &-bar {
    &-menu {
    }

    &-user {
      text-align: right;

      &-box {
        display: flex;
        justify-content: space-between;
        align-items: center;
        flex-wrap: nowrap;
        overflow: hidden;

        p {
          width: 100px;
          overflow: hidden;
          text-overflow: ellipsis;
          white-space: nowrap;
          margin: 0 8px;
          text-align: left;
        }
      }

    }
  }
}

.content {
  width: 100%;
}
</style>