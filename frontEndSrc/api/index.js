import router from '../router';
import Axios from 'axios'; // Axios: promise based HTTP client for the browser and node.js
import { Message } from "element-ui";

let instance = Axios.create({
    baseURL: process.env.VUE_APP_BASEURL,
    headers: {
        'Content-Type': 'application/json; charset=UTF-8',
    },
    timeout: 30000
});

instance.interceptors.request.use(
    config => {
        let token = sessionStorage.getItem('token');
        if(token){
            config.headers.Authorization = token;
        }
        
        return config;
    },
    err => {
        router.replace({
            path: '/login' 
        }); 

        return Promise.reject(err);
    }
);

instance.interceptors.response.use(
    response => {
        if (response.data.code == 3) {
            sessionStorage.clear();
            router.replace({
                path: '/login', 
                query: {
                    redirect: router.currentRoute.fullPath
                }
            });
            
            Message.error(response.data.message);
        }

        if(response.data.code != 0 && response.data.code != 3){
            Message.error(response.data.message);
        }

        return response.data;
    },
    error => {
        return Promise.reject(error)
    }
);

export default instance;