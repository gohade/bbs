import axios from 'axios'
import { Message } from 'element-ui'


// 创建一个axios
const service = axios.create({
    withCredentials: true, // send cookies when cross-domain requests
    timeout: 10000 // request timeout
})

// 请求的配置
service.interceptors.request.use(
    config => {
        return config
    },
    error => {
        // 如果request 有错误，打印信息
        console.log(error) // for debug
        return Promise.reject(error)
    }
)

// response中统一做处理
service.interceptors.response.use(
    response => {
        // Any status code that lie within the range of 2xx cause this function to trigger
        // Do something with response data
        // 判断http status是否为200
        if (response.status !== 200) {
            const data = response.data
            if (typeof data == 'string') {
                Message({
                    message: data,
                    type: 'error',
                    duration: 5 * 1000
                })
            }
        }
        return response
    },
    error => {
        // Any status codes that falls outside the range of 2xx cause this function to trigger
        // Do something with response error
        console.log('err: ' + error) // for debug
        // 打印Message消息
        Message({
            message: error.response.data,
            type: 'error',
            duration: 5 * 1000
        })
        return Promise.reject(error)
    }
)

export default service
