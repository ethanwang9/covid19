<template>
  <v-scale-screen :wrapperStyle="contentStyle" height="1080" width="1920">
    <el-row>
      <!--      菜单-->
      <el-col :span="24">
        <Header @changeArea="openDatabase"></Header>
      </el-col>
      <!--      内容-->
      <el-col :span="24">
        <Content></Content>
      </el-col>
    </el-row>
  </v-scale-screen>
  <!--  改变地区抽屉-->
  <el-drawer
      v-model="isOpen"
      title=切换区域
  >
    <el-select v-model="area" clearable placeholder="请选择地区" @change="changeArea">
      <el-option
          v-for="item in cityInfo"
          :key="item.value"
          :label="item.label"
          :value="item.value"
      />
    </el-select>
  </el-drawer>
</template>

<script lang="ts" setup>
import {CSSProperties} from 'vue'
import VScaleScreen from 'v-scale-screen'
import Header from '@/components/admin/main/header.vue'
import Content from '@/components/admin/main/content.vue'
import {CityList, CityListRes} from "../../../api/panel";
import useStore from "../../../store";

// 内容样式
const contentStyle: CSSProperties = reactive<CSSProperties>({
  backgroundColor: '#010959',
})

// store
const store = useStore().app

// 当前地区
const area = ref("")
area.value = store.userInfo.location

// 城市信息
const cityInfo = ref<CityListRes[]>([])

// 是否打开抽屉
const isOpen = ref<boolean>(false)

// 打开抽屉
function openDatabase() {
  isOpen.value = true
}

// 获取城市信息
const GetCityInfo = async () => {
  await CityList({
    province: store.userInfo.location
  }).then(res => {
    cityInfo.value = res.data
  }).catch(err => {
    ElNotification({
      title: "错误",
      message: err.message,
    })
  })
}

// 改变区域
const changeArea = (value: string) => {
  if (value.length == 0) return
  if (value !== store.userInfo.location) {
    if (value === "香港特别行政区" || value === "澳门特别行政区" || value === "台湾省") {
      ElNotification({
        message: "该区域暂未开发查询",
      })
    } else {
      isOpen.value = false
      store.setUserLocation(value)
      if (value === "新疆生产建设兵团") {
        ElNotification({
          message: "该区域仅可显示风险地区，更多信息未公示！",
        })
        setTimeout(() => {
        }, 2000)
      } else {
        ElNotification({
          message: "切换成功",
        })
      }
      setTimeout(() => {
        window.location.reload()
      }, 2000)
    }
  }
}

// 定时刷新
setTimeout(() => {
  window.location.reload()
}, 1000 * 60 * 25)
GetCityInfo()
</script>

<style lang="scss" scoped>
</style>
