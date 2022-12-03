<template>
  <el-row :gutter="15" class="main" justify="center">
    <el-col>
      <h3>风险地区</h3>
      <el-col class="tips">
        <span>* 本服务由国家卫生健康委提供</span>
        <el-divider direction="vertical"/>
        <span>更新时间：{{ updateTime }}</span>
      </el-col>
    </el-col>
    <el-col>
      <el-tabs
          v-model="activeTabs"
          class="main-tabs"
          tab-position="top"
          @tab-click="tabsClick"
      >
        <el-tab-pane class="main-tabs-pane" label="归属地风险地区" name="location">
          <!--START-->
          <h2>重庆</h2>

          <h3 v-if="dangerAreaList.h_count !== 0" class="hei">高风险 x{{ dangerAreaList.h_count }}</h3>
          <el-card v-for="(item,index) in dangerAreaList.h_list" v-if="dangerAreaList.h_count !== 0" :key="index"
                   class="main-tabs-pane-hei"
                   shadow="always">
            <template #header>
              <span style="font-weight: bold">{{ item.area_name }}</span>
            </template>
            <p v-for="(item2,index2) in item.communitys" :key="index2">{{ item2 }}</p>
          </el-card>

          <h3 v-if="dangerAreaList.m_count !== 0" class="mid">中风险 x{{}}</h3>
          <el-card v-for="(item,index) in dangerAreaList.m_list" v-if="dangerAreaList.m_count !== 0" :key="index"
                   class="main-tabs-pane-mid"
                   shadow="always">
            <template #header>
              <span style="font-weight: bold">{{ item.area_name }}</span>
            </template>
            <p v-for="(item2,index2) in item.communitys" :key="index2">{{ item2 }}</p>
          </el-card>

          <h3 v-if="dangerAreaList.l_count !== 0" class="low">低风险 x{{ dangerAreaList.l_count }}</h3>
          <el-card v-for="(item,index) in dangerAreaList.l_list" v-if="dangerAreaList.l_count !== 0" :key="index"
                   class="main-tabs-pane-low"
                   shadow="always">
            <template #header>
              <span style="font-weight: bold">{{ item.area_name }}</span>
            </template>
            <p v-for="(item2,index2) in item.communitys" :key="index2">{{ item2 }}</p>
          </el-card>
          <!--END-->
        </el-tab-pane>
        <el-tab-pane class="main-tabs-pane" label="精确位置查询" name="accurate">
          <!--START-->
          <el-col class="cascader">
            <span class="hidden-xs-only">当前地区： </span>
            <el-cascader
                :options="cityData.list"
                :show-all-levels="false"
                clearable
                filterable
                placeholder="请选择查询地区"
                separator=" - "
                @change="changeCity"
            />
          </el-col>
          <el-col v-if="cityData.res.level.length === 0">
            <el-empty :image-size="250" description="空数据，请选择地区查询数据！"/>
          </el-col>
          <el-col v-else>
            <el-result style="width: 100%;">
              <template #icon>
                <i-ep-SuccessFilled v-if="cityData.res.level === '常态化防控区域'" style="color: #67C23A;"/>
                <i-ep-CircleCloseFilled v-else style="color: #F56C6C;"/>
              </template>
              <template #title>
                <p v-if="cityData.res.level === '常态化防控区域'">{{ cityData.res.level }}</p>
                <p v-else>{{ cityData.res.level }} x{{ cityData.res.info.communitys.length }}</p>
              </template>
              <template #sub-title>
                <p>{{ citySelect.province }} {{ citySelect.county }} {{ citySelect.city }}</p>
              </template>
              <template #extra>
                <el-col v-if="cityData.res.level !== '常态化防控区域'">
                  <el-card class="main-tabs-pane-hei" shadow="always">
                    <el-row v-for="(item,index) in cityData.res.info.communitys" :key="index" :span="24" align="middle">
                      <el-col :md="18"><p>{{ item.info }}</p></el-col>
                      <el-col :md="6" class="hidden-sm-and-down">{{ item.level }}</el-col>
                    </el-row>
                  </el-card>
                </el-col>
                <el-col style="padding-top: 20px">
                  <p class="status-desc">{{ updateTime }}</p>
                  <p class="status-desc">有关信息来自各地确定的疫情风险等级</p>
                </el-col>
              </template>
            </el-result>
          </el-col>
          <!--END-->
        </el-tab-pane>
        <el-tab-pane class="main-tabs-pane" label="全国风险地区" name="country">
          <!--START-->
          <h3 v-if="cityData.all.hcount !== 0" class="hei">高风险 x{{ cityData.all.hcount }}</h3>
          <el-card v-for="(item,index) in cityData.all.highlist" v-if="cityData.all.hcount !== 0" :key="index"
                   class="main-tabs-pane-hei"
                   shadow="always">
            <template #header>
              <span style="font-weight: bold">{{ item.area_name }}</span>
            </template>
            <p v-for="(item2,index2) in item.communitys" :key="index2">{{ item2 }}</p>
          </el-card>

          <h3 v-if="cityData.all.mcount !== 0" class="mid">中风险 x{{ cityData.all.mcount }}</h3>
          <el-card v-for="(item,index) in cityData.all.middlelist" v-if="cityData.all.mcount !== 0" :key="index"
                   class="main-tabs-pane-mid"
                   shadow="always">
            <template #header>
              <span style="font-weight: bold">{{ item.area_name }}</span>
            </template>
            <p v-for="(item2,index2) in item.communitys" :key="index2">{{ item2 }}</p>
          </el-card>

          <h3 v-if="cityData.all.lcount !== 0" class="low">低风险 x{{ cityData.all.lcount }}</h3>
          <el-card v-for="(item,index) in cityData.all.lowlist" v-if="cityData.all.lcount !== 0" :key="index"
                   class="main-tabs-pane-low"
                   shadow="always">
            <template #header>
              <span style="font-weight: bold">{{ item.area_name }}</span>
            </template>
            <p v-for="(item2,index2) in item.communitys" :key="index2">{{ item2 }}</p>
          </el-card>
          <!--END-->
        </el-tab-pane>
      </el-tabs>
    </el-col>
  </el-row>
</template>

<script lang="ts" setup>
// tabs 激活
import {
  ChinaCityList,
  DangerAreaLocation,
  DangerAreaLocationRes,
  QueryDanger,
  QueryDangerAll,
  QueryDangerReq,
} from "../../../api/query";

// 活动 Tab 下标
const activeTabs = ref("location")

// 地区数据
const cityData = reactive({
  list: [] as any,
  res: {
    info: {
      type: "",
      province: "",
      city: "",
      county: "",
      area_name: "",
      communitys: []
    },
    level: "",
    update_time: "",
  } as any,
  all: {
    end_update_time: "",
    hcount: 0,
    highlist: [],
    mcount: 0,
    middlelist: [],
    lcount: 0,
    lowlist: [],
  } as any,
})

// 地区选项
const citySelect = reactive<QueryDangerReq>({
  province: "",
  county: "",
  city: "",
})


// 数据更新时间
const updateTime = ref<string>()

// 归属地风险地区
const dangerAreaList = reactive<DangerAreaLocationRes>({
  update_time: "",
  h_count: 0,
  h_list: [],
  m_count: 0,
  m_list: [],
  l_count: 0,
  l_list: [],
})

// 获取归属地风险地区
const GetDangerAreaList = async () => {
  DangerAreaLocation().then(res => {
    const {data} = res
    dangerAreaList.update_time = data.update_time
    updateTime.value = data.update_time
    dangerAreaList.h_count = data.h_count
    dangerAreaList.h_list = data.h_list
    dangerAreaList.m_count = data.m_count
    dangerAreaList.m_list = data.m_list
    dangerAreaList.l_count = data.l_count
    dangerAreaList.l_list = data.l_list

  }).catch(err => {
    ElNotification({
      title: "请求发送错误",
      message: err.message,
    })
  })
}
GetDangerAreaList()

// 改变城市
const changeCity = (v: any) => {
  citySelect.province = v[0]
  citySelect.county = v[1]
  citySelect.city = v[2]

  QueryDangerInfo()
}

// 获取城市列表
const GetCityList = async () => {
  ChinaCityList().then(res => {
    cityData.list = res.data
  }).catch(err => {
    ElNotification({
      title: "请求发送错误",
      message: err.message,
    })
  })
}
GetCityList()

// 查询风险地区
const QueryDangerInfo = async () => {
  await QueryDanger(citySelect).then(res => {
    if (res.data.level === "n") {
      res.data.level = "常态化防控区域"
    } else if (res.data.level === "h") {
      res.data.level = "高风险"
    } else if (res.data.level === "m") {
      res.data.level = "中风险"
    } else if (res.data.level === "l") {
      res.data.level = "低风险"
    }
    cityData.res = res.data
    updateTime.value = res.data.update_time
  }).catch(err => {
    ElNotification({
      title: "请求发送错误",
      message: err.message,
    })
  })
}

// 获取全国风险地区
const GetDangerAll = async () => {
  await QueryDangerAll().then(res => {
    cityData.all = res.data
    updateTime.value = res.data.end_update_time
  }).catch(err => {
    ElNotification({
      title: "请求发送错误",
      message: err.message,
    })
  })
}
GetDangerAll()
</script>

<style lang="scss" scoped>
.main {
  background-color: rgba(0, 0, 0, 0.4);
  border-radius: 8px;
  box-shadow: 0 0 8px rgba(0, 0, 0, 0.3);
  padding-bottom: 30px;

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

  &-tabs {
    padding: 0 30px;
    overflow: hidden;

    &-pane {
      h2 {
        letter-spacing: 0.5em;
        text-align: center;
      }

      .hei {
        color: #d00909;
        padding: unset;
      }

      .mid {
        color: #ffff00;
        padding: unset;
      }

      .low {
        color: #ffffff;
        padding: unset;
      }

      &-hei {
        background-color: #d00909;
        margin-bottom: 15px;
      }

      &-mid {
        background-color: #262294;
        margin-bottom: 15px;
      }

      &-low {
        margin-bottom: 15px;
      }

      .status-desc {
        letter-spacing: 2px;
        font-size: 14px;
        color: #808080;
      }

      @media screen and (max-width: 768px) {
        .cascader {
          margin: 10px 0 40px;
          text-align: center;
        }
      }
    }
  }
}
</style>

<style lang="scss">
//覆写 element plus 样式
.el-result__extra {
  width: 80%;
}

@media screen and (max-width: 768px) {
  .el-result {
    padding: 0;

    .el-result__extra {
      width: 100%;
    }
  }
}
</style>