<template>
  <div class="DENMain">
    <dv-border-box13 class="DENMain-content">
      <p class="content-title">风险地区</p>
      <div class="content-danger">
        <div class="content-danger-item">
          <p v-if="dangerData.h_count > 99999">99999+</p>
          <p v-else>{{ dangerData.h_count }}</p>
          <p>高风险</p>
        </div>
        <div class="content-danger-item">
          <p v-if="dangerData.m_count > 99999">99999+</p>
          <p v-else>{{ dangerData.m_count }}</p>
          <p>中风险</p>
        </div>
        <div class="content-danger-item">
          <p v-if="dangerData.l_count > 99999">99999+</p>
          <p v-else>{{ dangerData.l_count }}</p>
          <p>低风险</p>
        </div>
      </div>
      <div class="content-box">
        <dv-scroll-board :config="config" class="content-box-table"/>
      </div>
    </dv-border-box13>
  </div>
</template>

<script lang="ts" setup>
import {DangerArea} from "../../../../api/panel";
import useStore from "../../../../store";

const config = reactive({
  data: [] as any,
  rowNum: 7,
  columnWidth: [430, 70],
  headerBGC: "#4445c8",
  oddRowBGC: "#060e35",
  evenRowBGC: "#061a51",
})

// 风险地区数据
const dangerData = reactive({
  update_time: "",
  h_count: 0,
  h_list: [] as any,
  m_count: 0,
  m_list: [] as any,
  l_count: 0,
  l_list: [] as any,
})

const store = useStore().app

// 获取用户归属地风险地区
const GetUserDangerLocation = async () => {
  await DangerArea({
    province: store.userInfo.location
  }).then(res => {
    dangerData.update_time = res.data.update_time
    dangerData.h_count = res.data.h_count
    dangerData.h_list = res.data.h_list
    dangerData.m_count = res.data.m_count
    dangerData.m_list = res.data.m_list
    dangerData.l_count = res.data.l_count
    dangerData.l_list = res.data.l_list

    res.data.h_list.forEach(v => {
      let t = v.province + v.city + v.county
      v.communitys.forEach(v2 => {
        config.data.push([t + v2, "高风险"])
      })
    })
    res.data.m_list.forEach(v => {
      let t = v.province + v.city + v.county
      v.communitys.forEach(v2 => {
        config.data.push([t + v2, "中风险"])
      })
    })
    res.data.l_list.forEach(v => {
      let t = v.province + v.city + v.county
      v.communitys.forEach(v2 => {
        config.data.push([t + v2, "低风险"])
      })
    })
  }).catch(err => {
    ElNotification({
      title: "错误",
      message: err.message,
    })
  })
}
GetUserDangerLocation()
</script>

<style lang="scss" scoped>
.DENMain {
  width: 450px;
  height: 570px;

  &-content {
    padding: 10px;
    box-sizing: border-box;

    .content {
      &-title {
        color: #a5d5f6;
        font-weight: bold;
        font-size: 20px;
        margin: 0;
        padding: 10px;
      }

      &-danger {
        display: flex;
        justify-content: space-evenly;

        &-item {
          height: 90px;
          width: 130px;
          border-radius: 8px;
          overflow: hidden;
          display: flex;
          flex-direction: column;
          align-items: center;
          justify-content: center;

          & > p {
            margin: 0;
          }

          & > p:first-child {
            font-size: 26px;
            font-weight: bold;
            padding-bottom: 5px;
          }

          & > p:last-child {
            font-size: 20px;
            letter-spacing: 6px;
          }
        }

        &-item:nth-child(1) {
          background-color: #ea2027;
        }

        &-item:nth-child(2) {
          background-color: #e30097;
        }

        &-item:nth-child(3) {
          background-color: #005be3;
        }
      }

      &-box {
        height: 405px;
        margin-top: 10px;

        &-table {

        }
      }
    }
  }
}
</style>