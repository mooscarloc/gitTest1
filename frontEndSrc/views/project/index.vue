<template>
  <div class="page">
    <div class="page-header">
      <el-breadcrumb separator-class="el-icon-arrow-right">
        <el-breadcrumb-item :to="{ path: '/' }">Dataset Display</el-breadcrumb-item>
        <el-breadcrumb-item>Project Info: ({{ $route.query.name }})</el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <div class="page-body">
      <el-tabs
        v-model="activeName"
        @tab-click="handleClick"
        v-loading="loading"
        type="card"
        element-loading-spinner="el-icon-loading"
      >
        <el-tab-pane label="Details" name="first">
          <el-divider content-position="left">Project Details</el-divider>
          <el-form
            :model="projectDetails"
            class="demo-ruleForm2"
          >
            <el-form-item label="Project ID: ">
              <span>{{ projectDetails.id }}</span>
            </el-form-item>
            <el-form-item label="Project Name: ">
              <span>{{ projectDetails.name }}</span>
            </el-form-item>
            <el-form-item label="MN: ">
              <span>{{ projectDetails.mn }}</span>
            </el-form-item>
            <el-form-item label="Project Info: ">
              <span>{{ projectDetails.desc }}</span>
            </el-form-item>
            <el-form-item label="Industry: ">
              <span v-if="projectDetails.industry">{{ projectDetails.industry | getIndustry }}</span>
              <span v-else>N/A</span>
            </el-form-item>
          </el-form>
        </el-tab-pane>
        <el-tab-pane label="Test Report" name="second">
          <el-divider content-position="left">Test Report</el-divider>
          <el-form
            :model="projectDetails"
            class="demo-ruleForm2"
          >
          <el-row>
            <el-col span='800'>
              <el-form-item label="Project Name">
                <span>{{ projectDetails.name }}</span>
              </el-form-item>
            </el-col>
            <el-col span='1000'>
              <el-form-item label="Report Date">
                <span>{{ new Date().toISOString().split('T')[0] }}</span>
              </el-form-item>
            </el-col>
            <el-col span='400'>
              <el-form-item label="Tester">
                <span>
                  <input
                    v-model="tester"
                    style="width:150px"
                  />
                </span>
              </el-form-item>
            </el-col>
          </el-row>
          <ReportRow :metricVal="zeronoise" metricId="zeroNoise" :deviceName="projectDetails.name" />
          <ReportRow :metricVal="scopenoise" metricId="scopeNoise" :deviceName="projectDetails.name" />
          <ReportRow :metricVal="uptime" metricId="upTime" :deviceName="projectDetails.name" />
          <ReportRow :metricVal="downtime" metricId="downTime" :deviceName="projectDetails.name" />
          </el-form>
        </el-tab-pane>
      </el-tabs>
    </div>
  </div>
</template>

<script>
import api from "../../api/base.js";
import ReportRow from "./ReportRow.vue";
export default {
  name: "Project",
  data() {
    return {
      tester: "",
      projectDetails: {},
      option: {},
      zeronoise: 0,
      scopenoise: 0,
      uptime: 0,
      downtime: 0
    };
  },
  filters: {
    getIndustry(tem) {
      let industryList = [
        {
          label: "SO2",
          value: 1,
        },
        {
          label: "NO",
          value: 2,
        },
        {
          label: "O3",
          value: 3,
        },
        {
          label: "CO",
          value: 4,
        },
      ];

      for (let i = 0; i < industryList.length; i++) {
        if (industryList[i].value == tem) {
          return industryList[i].label;
        }
      }
    },
  },
  watch: {},
  created() {
    this.projectId = this.$route.query.id;
    this.getTimeFn();
    this.handleClick();
  },
  updated() {
    if (this.projectId !== this.$route.query.id) {
      this.projectId = this.$route.query.id;
      this.handleClick();
    }
  },
  methods: {
    handleClick() {
      switch (this.activeName) {
        case "first":
          this.getInfo();
          break;
        case "second":
          this.getInfo();
          break;
      }
    },
    getInfo() {
      try {
        this.loading = true;
        api
          .getprojectDetails({ id: this.projectId })
          .then((res) => {
            if (res.status) {
              this.projectDetails = res.data;
              this.loading = false;
            } else {
              this.$message.error(res.message);
              this.loading = false;
            }
          })
          .catch(() => {
            this.loading = false;
          });
      } catch (e) {
        throw new Error(e);
      }
    },
    getTimeFn() {
      const that = this;
      const end = new Date();
      const start = new Date();
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 1);
      end.setTime(start.getTime() + 3600 * 1000 * 24 * 1);
      that.ruleForm.time[0] = that.formatDate(start);
      that.ruleForm.time[1] = that.formatDate(end);
    },
    formatDate(date) {
      var myyear = date.getFullYear();
      var mymonth = date.getMonth() + 1;
      var myweekday = date.getDate();

      if (mymonth < 10) {
        mymonth = "0" + mymonth;
      }
      if (myweekday < 10) {
        myweekday = "0" + myweekday;
      }
      return myyear + "-" + mymonth + "-" + myweekday + " " + "00:00:00";
    }
  },
  components: {
    ReportRow
  }
};
</script>

<style lang="scss">
.page {
  width: calc(100% - 180px);
  background-color: #f5f6f9;
  margin-left: 180px;

  .demo-ruleForm {
    display: flex;
  }

  .el-tabs {
    height: calc(100% - 40px);
    overflow-y: scroll;
  }

  .flex_line {
    display: flex;
    align-items: center;
    margin-bottom: 8px;
    b {
      width: 80px;
      padding-right: 8px;
    }
  }
}

.res_wrap {
  padding: 12px 24px;
  line-height: 30px;
  border-top: 1px solid #f7f7f7;
}
</style>
