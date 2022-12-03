<template>
  <div class="YQMain">
    <dv-border-box12 class="YQMain-content">
      <p class="content-title">疫情数据</p>
      <div class="content-box">
        <dv-scroll-board :config="config" class="content-box-table"/>
      </div>
    </dv-border-box12>
  </div>
</template>

<script lang="ts" setup>
import {AreaStatDetails} from "../../../../api/panel";
import useStore from "../../../../store";

const config = reactive({
  header: ['地区', '现存', '累计', '死亡', '治愈'],
  data: [] as any,
  index: true,
  columnWidth: [50, 120],
  rowNum: 5,
  headerHeight: 45,
  headerBGC: "#4445c8",
  oddRowBGC: "#060e35",
  evenRowBGC: "#061a51",

})

// store
const store = useStore().app

// 获取数据
const GetData = async () => {
  await AreaStatDetails({
    province: store.userInfo.location,
  }).then(res => {
    config.data = res.data
  }).catch(err => {
    ElNotification({
      title: "错误",
      message: err.message,
    })
  })
}
GetData()
</script>

<style lang="scss" scoped>
.YQMain {
  width: 450px;
  height: 365px;

  &-content {
    padding: 12px;
    box-sizing: border-box;
    display: flex;
    flex-direction: column;

    .content {
      &-title {
        color: #a5d5f6;
        font-weight: bold;
        font-size: 20px;
        margin: 0;
        padding: 10px;
      }

      &-box {
        width: 450px;

        &-table {
          width: 426px;
          height: 300px
        }
      }
    }
  }
}
</style>