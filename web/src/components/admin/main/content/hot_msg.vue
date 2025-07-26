<template>
  <div class="MSGMain">
    <dv-border-box7 class="MSGMain-content">
      <p class="MSGMain-content-title">热点动态</p>
      <el-carousel :interval="interval" height="185px">
        <!--START-->
        <el-carousel-item v-for="(item,index) in data" :key="index" class="MSGMain-content-item"
                          @click="goUrl(item.jumpLink.url)">
          <h3 class="MSGMain-content-item-title">{{ item.title }}</h3>
          <p class="MSGMain-content-item-desc">{{ item.desc }}</p>
          <div class="MSGMain-content-item-info">
            <p>{{ item.from }}</p>
            <p>{{ item.publicTime }}</p>
          </div>
        </el-carousel-item>
        <!--END-->
      </el-carousel>
    </dv-border-box7>
  </div>
</template>

<script lang="ts" setup>
import {HotMessage} from "../../../../api/panel";
import useStore from "../../../../store";

// store
const store = useStore().app

// 数据
const data = ref()

// 热点切换时间
const interval = 5000

// 热点动态跳转
function goUrl(url: string) {
  window.open(url, '_blank')
}

// 获取热点消息
const GetHotMessage = async () => {
  await HotMessage({
    province: store.userInfo.location
  }).then(res => {
    data.value = res.data
  }).catch(err => {
    ElNotification({
      title: "错误",
      message: err.message,
    })
  })
}
GetHotMessage()
</script>

<style lang="scss" scoped>
.MSGMain {
  width: 900px;
  height: 230px;

  &-content {
    &-title {
      color: #a5d5f6;
      font-weight: bold;
      font-size: 20px;
      margin: 0;
      padding: 15px;
    }

    &-item {
      padding: 10px;
      box-sizing: border-box;

      &-title {
        margin: 0;
        width: 100%;
        font-size: 26px;
        display: inline-block;
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
      }

      &-desc {
        letter-spacing: 2px;
        overflow: hidden;
        text-overflow: ellipsis;
        display: -webkit-box;
        -webkit-line-clamp: 2;
        -webkit-box-orient: vertical;
        font-size: 20px;
        margin: 10px 0 0 0;
      }

      &-info {
        display: flex;
        justify-content: space-between;
        align-items: center;
        font-size: 20px;
      }
    }
  }
}
</style>