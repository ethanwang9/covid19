<template>
  <el-row :gutter="15" class="main" justify="center">
    <el-col>
      <h3>出行政策</h3>
      <el-col class="tips">
        <span>* 本服务由腾讯健康提供</span>
      </el-col>
    </el-col>
    <el-col class="main-select">
      <el-cascader
          :options="areaList"
          clearable
          filterable
          placeholder="出发地"
          @change="changeArea($event, 'from')"
      />
      <i-ep-Switch/>
      <el-cascader
          :options="areaList"
          clearable
          filterable
          placeholder="目的地"
          @change="changeArea($event, 'to')"
      />
    </el-col>
    <el-col v-show="travelData.to.cityName.length !== 0" :span="24" class="main-content">
      <el-row justify="center" style="margin-bottom: 20px">
        <el-col :md="14" :sm="16" :xs="20">
          <el-alert
              :title="travelData.importantNotice"
              closable
              show-icon
              type="success"
          />
        </el-col>
      </el-row>
      <el-col :sm="11" :xs="20" class="main-content-item">
        <el-col class="main-content-item-title">
          <h2>出{{ travelData.from.cityName }}</h2>
          <span v-if="travelData.from.isOutUpdate">离开政策有更新</span>
          <span v-if="travelData.from.isInUpdate">进入政策有更新</span>
        </el-col>
        <el-col class="main-content-item-tag">
          <el-tag v-for="(item,index) in travelData.from.labels" :key="index" effect="dark" type="danger">{{
              item.label
            }}
          </el-tag>
        </el-col>
        <el-col class="main-content-item-text">
          <p>{{ travelData.from.outDesc }}</p>
        </el-col>
        <el-col class="main-content-item-phone">
          <h3>防疫热线</h3>
          <!--防疫热线 START-->
          <el-col v-for="(item, index) in travelData.from.diseaseControlTel.diseaseControlTel"
                  :key="index"
                  class="main-content-item-phone-btns"
          >
            <p>{{ item.name }}</p>
            <el-button auto-insert-space bg size="large" text>
              <a v-for="(item2, index2) in item.title" :key="index2" :href="'tel:'+item2">{{ item2 }}</a>
            </el-button>
          </el-col>
          <!--防疫热线 END-->
        </el-col>
      </el-col>
      <el-col :sm="11" :xs="20" class="main-content-item">
        <el-col class="main-content-item-title">
          <h2>到{{ travelData.to.cityName }}</h2>
          <span v-if="travelData.to.isOutUpdate">离开政策有更新</span>
          <span v-if="travelData.to.isInUpdate">进入政策有更新</span>
        </el-col>
        <el-col class="main-content-item-tag">
          <el-tag v-for="(item,index) in travelData.to.labels" :key="index" effect="dark" type="danger">{{
              item.label
            }}
          </el-tag>
        </el-col>
        <el-col class="main-content-item-text">
          <p>{{ travelData.to.highInDesc }}</p>
        </el-col>
        <el-col class="main-content-item-phone">
          <h3>防疫热线</h3>
          <!--防疫热线 START-->
          <el-col v-for="(item, index) in travelData.to.diseaseControlTel.diseaseControlTel"
                  :key="index"
                  class="main-content-item-phone-btns"
          >
            <p>{{ item.name }}</p>
            <el-button auto-insert-space bg size="large" text>
              <a v-for="(item2, index2) in item.title" :key="index2" :href="'tel:'+item2">{{ item2 }}</a>
            </el-button>
          </el-col>
          <!--防疫热线 END-->
        </el-col>
      </el-col>
    </el-col>
    <el-col v-show="travelData.to.cityName.length === 0">
      <el-empty description="还没有内容,快去选择地区查询出行政策吧！"/>
    </el-col>
  </el-row>
</template>

<script lang="ts" setup>
import {CityList, Travel, TravelRes} from "../../../api/query";

// 接口结构
interface CityListItem {
  value: string
  label: string
  children: CityListChildren[]
}

interface CityListChildren {
  value: string
  label: string
}

// 选项卡中数据
const selectData = reactive({
  from: "",
  to: "",
})

// 出行数据
const travelData = reactive<TravelRes>({
  "code": 0,
  "from": {
    "cityCode": "",
    "cityName": "",
    "code": 0,
    "diseaseControlTel": "",
    "highInDesc": "",
    "isInUpdate": false,
    "isOutUpdate": false,
    "labels": [],
    "lowInDesc": "",
    "medianInDesc": "",
    "message": "",
    "news": {
      "context": "",
      "isShow": false,
      "jumpLink": {
        "appId": "",
        "appVer": "",
        "type": 0,
        "url": ""
      }
    },
    "outDesc": "",
    "provinceName": "",
    "riskLevel": 0,
    "sourceContext": "",
    "sourceLink": {
      "appId": "",
      "appVer": "",
      "type": 0,
      "url": ""
    }
  },
  "importantDict": {
    "isShow": false,
    "type": "",
    "value": ""
  },
  "importantNotice": "",
  "inKeywordDict": {
    "isShow": false,
    "type": "",
    "value": ""
  },
  "message": "",
  "outKeywordDict": {
    "isShow": false,
    "type": "",
    "value": ""
  },
  "to": {
    "cityCode": "",
    "cityName": "",
    "code": 0,
    "diseaseControlTel": "",
    "highInDesc": "",
    "isInUpdate": false,
    "isOutUpdate": false,
    "labels": [],
    "lowInDesc": "",
    "medianInDesc": "",
    "message": "",
    "news": {
      "context": "",
      "isShow": false,
      "jumpLink": {
        "appId": "",
        "appVer": "",
        "type": 0,
        "url": ""
      }
    },
    "outDesc": "",
    "provinceName": "",
    "riskLevel": 0,
    "sourceContext": "",
    "sourceLink": {
      "appId": "",
      "appVer": "",
      "type": 0,
      "url": ""
    }
  }
})

// 地区列表
const areaList: CityListItem[] = reactive([])

// 获取地区列表
const GetCityList = async () => {
  await CityList().then(res => {
    let {data} = res

    // 清洗数据
    data.forEach(v1 => {
      let tempArr: CityListItem = {
        value: v1.cityCode,
        label: v1.label,
        children: [],
      }

      v1.children.forEach(v2 => {
        tempArr.children.push({
          value: v2.cityCode,
          label: v2.label,
        })
      })

      areaList.push(tempArr)
    })
  }).catch(err => {
    ElNotification({
      title: "请求发送错误",
      message: err.message,
    })
  })
}
GetCityList()

// 改变地区
const changeArea = (event: any, where: string) => {
  let t = event[1]
  switch (where) {
    case "from":
      selectData.from = t
      break
    case "to":
      selectData.to = t
      break
  }

  if (selectData.from.length !== 0 && selectData.to.length !== 0) {
    GetTravel()
  }
}

// 获取出行政策
const GetTravel = async () => {
  Travel({
    to: selectData.to,
    from: selectData.from,
  }).then(res => {
    let {data} = res
    travelData.code = data.code
    travelData.from = data.from
    travelData.importantDict = data.importantDict
    travelData.importantNotice = data.importantNotice
    travelData.inKeywordDict = data.inKeywordDict
    travelData.message = data.message
    travelData.outKeywordDict = data.outKeywordDict
    travelData.to = data.to

    travelData.from.diseaseControlTel = JSON.parse(travelData.from.diseaseControlTel)
    travelData.to.diseaseControlTel = JSON.parse(travelData.to.diseaseControlTel)
  }).catch(err => {
    ElNotification({
      title: "请求发送错误",
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
    padding: 15px 30px 0;
  }

  .tips {
    font-size: 12px;
    margin: 0 30px 5px;
    color: #808080;
    line-height: 20px;
    letter-spacing: 1px;
  }

  &-select {
    display: flex;
    justify-content: center;
    align-items: center;
    margin-top: 10px;

    * {
      padding: 20px 20px;
    }

    @media screen and (max-width: 768px) {
      flex-direction: column;
    }
  }

  &-content {
    display: flex;
    justify-content: center;
    align-items: flex-start;
    flex-wrap: wrap;
    padding-bottom: 30px;
    margin-top: 25px;

    &-item {
      display: flex;
      flex-direction: column;
      align-items: flex-start;
      justify-content: flex-start;


      &-title {
        display: flex;
        align-items: center;
        letter-spacing: .5em;

        span {
          padding-left: 5px;
          font-size: 14px;
          letter-spacing: 0;
          color: #e1e1e1;
        }
      }

      &-tag {
        * {
          margin: 6px;
        }
      }

      &-text {
        letter-spacing: 1px;
        line-height: 25px;
        white-space: pre-wrap;
      }

      &-phone {
        margin-bottom: 40px;

        h3 {
          padding: 0;
          margin: 20px 0;
        }

        &-btns {

          p {
            margin: 15px 0;
            padding: 0;
            font-weight: bold;
            letter-spacing: 2px;
          }

          * {
            margin: 4px;
          }
        }
      }
    }
  }
}

//覆盖样式
a {
  text-decoration: none;
  color: var(--el-button-text-color);
}
</style>