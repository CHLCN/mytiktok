import request from "../utils/request";

export default {
  follow(params, data) {
    return request({url: '/relation/follow/list/', method: 'get', params, data})
  },
  follower(params, data) {
    return request({url: '/relation/follower/list/', method: 'get', params, data})
  },
  author(params, data) {
    return request({url: '/user/author', method: 'get', params, data})
  },
  profile(params, data) {
    return request({url: '/user/profile/other', method: 'get', params, data})
  },
}