<template>
  <div class="SMain">
    <div class="SMain-item">
      <dv-border-box8 ref="border1" :dur="4">
        <h4 v-if="db.currentConfirmedCount>99999">999999+</h4>
        <h4 v-else>{{ db.currentConfirmedCount }}</h4>
        <p>现存确诊</p>
      </dv-border-box8>
    </div>
    <div class="SMain-item">
      <dv-border-box8 ref="border2" :dur="6">
        <h4 v-if="db.confirmedCount>99999">{{ db.confirmedCount }}</h4>
        <h4 v-else>{{ db.confirmedCount }}</h4>
        <p>累计确诊</p>
      </dv-border-box8>
    </div>
    <div class="SMain-item">
      <dv-border-box8 ref="border3" :dur="8">
        <h4 v-if="db.deadCount>99999">999999+</h4>
        <h4 v-else>{{ db.deadCount }}</h4>
        <p>累计死亡</p>
      </dv-border-box8>
    </div>
    <div class="SMain-item">
      <dv-border-box8 ref="border4" :dur="10">
        <h4 v-if="db.curedCount>99999">999999+</h4>
        <h4 v-else>{{ db.curedCount }}</h4>
        <p>累计治愈</p>
      </dv-border-box8>
    </div>
  </div>
</template>

<script lang="ts" setup>
import useStore from "../../../../store";
import {AreaStat, AreaStatRes} from "../../../../api/panel";

// store
const store = useStore().app

// 数据
const db = reactive<AreaStatRes>({
  currentConfirmedCount: 0,
  confirmedCount: 0,
  deadCount: 0,
  curedCount: 0,
})

// 获取统计数据
const GetAreaStat = async () => {
  await AreaStat({
    province: store.userInfo.location,
  }).then(res => {
    db.currentConfirmedCount = res.data.currentConfirmedCount
    db.confirmedCount = res.data.confirmedCount
    db.deadCount = res.data.deadCount
    db.curedCount = res.data.curedCount
  }).catch(err => {
    ElNotification({
      title: "错误",
      message: err.message,
    })
  })
}
GetAreaStat()
</script>

<style lang="scss" scoped>
.SMain {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 900px;

  &-item {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    text-align: center;
    width: 220px;
    height: 120px;

    h4 {
      font-size: 40px;
      margin: 0 0;
      padding-top: 20px;
    }

    p {
      font-size: 20px;
      letter-spacing: 8px;
      font-weight: bold;
      margin-top: 10px;
    }
  }

  &-item:nth-child(1) {
    h4 {
      color: rgb(247, 76, 49);

    }
  }

  &-item:nth-child(2) {
    h4 {
      color: rgb(174, 33, 44);;
    }
  }

  &-item:nth-child(3) {
    h4 {
      color: rgb(93, 112, 146);
    }
  }

  &-item:nth-child(4) {
    h4 {
      color: rgb(40, 183, 163);
    }
  }
}
</style>