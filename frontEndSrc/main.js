import Vue from "vue";
import App from './App.vue';
import './assets/style/base.css';
import 'element-ui/lib/theme-chalk/index.css';
import router from "./router";
import ElementUI from 'element-ui';
Vue.use(ElementUI);

router.beforeEach((to, from, next) => {
    if (to.meta.title) {
        document.title = to.meta.title;
    }
    
    if (to.path == '/login' || sessionStorage.getItem("token")) {
        next();
    } else {
        next('/login');
    }
});

Vue.config.productionTip = false;

new Vue({
    router,
    render: h => h(App)
}).$mount("#app");
