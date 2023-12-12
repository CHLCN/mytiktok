import request from "../utils/request";

export default {
  me(params, data) {
    return request({url: '/me', method: 'get', params, data})
  },
  my(params, data) {
    // console.log('api - params',params)
    return request({url: '/me/my', method: 'get', params, data})
  },
  mypost(params, data) {
    return request({url: '/me/mypost', method: 'get', params, data})
  },
  like(params, data) {
    return request({url: '/me/like', method: 'get', params, data})
  },
  updatelike(params, data) {
    return request({url: '/favorite/action/', method: 'post', params, data})
  },
  collect(params, data) {
    return request({url: '/me/collect', method: 'get', params, data})
  },
  recommended(params, data) {
    return request({url: '/video/recommended', method: 'get', params, data})
  },
  historyVideo(params, data) {
    return request({url: '/video/historyVideo', method: 'get', params, data})
  },
  historyOther(params, data) {
    return request({url: '/video/historyOther', method: 'get', params, data})
  },
  post(params, data) {
    return request({url: '/aweme/post', method: 'get', params, data})
  },
}