<template>
    <div class="menu">
        <el-form :model="searchForm" ref="searchForm" class="demo-ruleForm">
            <el-form-item>
                <el-select
                    size="small"
                    v-model="searchForm.industry"
                    placeholder="Please select industry"
                >
                    <el-option
                        v-for="item in industryList"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value"
                    ></el-option>
                </el-select>
            </el-form-item>
            <el-form-item>
                <el-select
                    size="small"
                    v-model="searchForm.status"
                    placeholder="Please select status"
                >
                    <el-option
                        v-for="item in statusList"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value"
                    ></el-option>
                </el-select>
            </el-form-item>
            <el-form-item>
                <div style="display: flex;justify-content: space-between">
                    <el-autocomplete
                        class="inline-input"
                        v-model="searchForm.name"
                        size="small"
                        :fetch-suggestions="querySearch"
                        placeholder="Project Name"
                        :trigger-on-focus="false"
                        @select="handleSelect"
                    ></el-autocomplete>
                    <div>
                        <el-button size="small" type="primary" @click="getProject" :loading="loading">Inquiry</el-button>
                    </div>
                </div>
            </el-form-item>
        </el-form>

        <ul
            class="mainUL"
            v-loading="loading"
            element-loading-spinner="el-icon-loading"
            element-loading-background="rgba(0, 0, 0, 0.8)"
        >
            <li
                v-for="(menu,key) in menuList"
                :key="key"
                @click="showProject(menu)"
                :class="{'is-active': $route.query.id == menu.id}"
            >
                <i class="el-icon-office-building"></i>
                {{menu.value}}
            </li>
        </ul>
    </div>
</template>

<script>
import api from "../api/base.js";
export default {
    name: "menu",
    data() {
        return {
            loading: true,
            currentId: null,
            searchForm: {
                name: "",
                industry: "",
                status: ""
            },
            industryList: [
                 {
                    label: "All",
                    value: 0
                },
                {
                    label: "SO2",
                    value: 1
                },
                {
                    label: "NO",
                    value: 2
                },
                {
                    label: "O3",
                    value: 3
                },
                {
                    label: "CO",
                    value: 4
                }
            ],
            statusList: [
                 {
                    label: "All",
                    value: 0
                },
                {
                    value: 1,
                    label: "Online"
                },
                {
                    value: 2,
                    label: "Offline"
                }
            ],
            index: "",
            menuList: [],
            menuTree: []
        };
    },
    created() {
        this.getProject();
    },
    mounted() {},
    methods: {
        querySearch(queryString, callback) {
            var items = this.menuList;
            var results = queryString
                ? items.filter(this.createFilter(queryString))
                : items;
            callback(results);
        },
        handleSelect(item) {},
        showProject(menu) {
            this.currentId = menu.id;
            this.$router.push({
                path: "/home/project",
                query: {
                    id: menu.id,
                    name: menu.value
                }
            });
        },
        getProject() {
            try {
                this.loading = true;
                let data = {};
                if (this.searchForm.name) {
                    data.name = this.searchForm.name;
                }

                if (this.searchForm.industry) {
                    data.industry = Number(this.searchForm.industry);
                }

                if (this.searchForm.status) {
                    data.status = Number(this.searchForm.status);
                }

                api.getProjects(data).then(res => {
                    if (res.status) {
                        this.menuList = JSON.parse(
                            JSON.stringify(res.data)
                        );

                        this.loading = false;
                    }
                });
            } catch (e) {
                throw new Error(e);
            }
        }
    }
};
</script>

<style scoped lang="scss">
.is-active {
    background: #000;
    color: #fff;
}

.menu {
    height: 100%;
    width: 180px;
    background: #1d1e23;
    z-index: 10;
    overflow: hidden;
    position: fixed;
    top: 60px;
    left: 0;
    bottom: 0;

    .demo-ruleForm {
        padding: 10px;
        .el-form-item {
            margin-bottom: 0;
        }
    }

    .mainUL {
        height: calc(100% - 200px);
        background: #1e1e1e;
        overflow: auto;
        color: #fff;
        padding: 0px 0;
        line-height: 30px;
        font-size: 16px;

        li {
            padding: 10px 10px;
            padding-left: 25px;
            color: rgba(255, 255, 255, 0.7);

            i {
                padding: 0 3px;
            }
        }

        li:hover {
            cursor: pointer;
            background: #000;
            color: rgba(255, 255, 255, 1);
        }
    }
}
</style>
