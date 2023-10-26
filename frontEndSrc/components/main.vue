<template>
    <div class="main">
        <div class="main-header">
            <span>Management System</span>
            <div class="main-menu">
                <el-menu :default-active="activePath" class="el-menu-demo" mode="horizontal" @select="handleSelect" router>
                    <el-menu-itemindex="/home"> 
                        <template slot="title">
                            <i class="el-icon-s-marketing"></i>
                            <span>Illustration</span>
                        </template>
                    </el-menu-item>
                    <el-menu-item index="/info">
                        <template slot="title">
                            <i class="el-icon-message"></i>
                            <span>System Configuration</span>
                        </template>
                    </el-menu-item>
                </el-menu>
            </div>
            <div>
                <span>{{username}}</span>
                <i>|</i>
                <button @click="logout">Log out</button>
            </div>
        </div>

        <div class="main-content">
            <Menu v-if="'/'+this.$route.path.split('/')[1]=='/home'"></Menu>
            <router-view />
        </div>
    </div>
</template>

<script>
    import Menu from './menu.vue';
    export default {
        name: "Main",
        components: {
            Menu
        },
        data() {
            return {
                activePath:'',
                username:''
            }
        },
        created() {
            this.username = sessionStorage.getItem('user')
            this.activePath = "/"+this.$route.path.split('/')[1];
        },
        updated() {
            this.activePath = "/"+this.$route.path.split('/')[1];
        },
        methods: {
            handleSelect(key, keyPath) {
            },
            logout(){
                sessionStorage.clear();
                this.$router.push('/login');
            }
        }
    }
</script>

<style scoped lang="scss">
    .main {
        width: 100%;
        height: 100%;
        display: flex;
        .header {
            z-index: 100;
            width: 100%;
            height: 60px;
            background: #1d1e23;
            color: #f6ca9d;
            padding: 0 25px;
            display: flex;
            align-items: center;
            justify-content: space-between;
            position: fixed;
            left: 0;
            right: 0;
            top: 0;

            span {
                color: #f5f5f5;
                width: 180px;
            }

            div {
                color: #f5f5f5;

            }

            button {
                background: #1d1e23;
                border: none;
                color: #f5f5f5;
                font-size: 14px;
                cursor: pointer;
            }
        }
    }
</style>