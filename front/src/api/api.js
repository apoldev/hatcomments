import axios from 'axios'
import {API_URL} from "@/api/config";

const $apiAxios = axios.create({
  baseURL: API_URL
})

class Api {

  token = ''

  setToken = (token) => {
    this.token = token
  }

  getConfig = () => {
    var o = {}

    if(this.token !== ""){
      o["headers"] = {
        Authorization: `Bearer ${this.token}`
      }
    }

    return o
  }

  post = async(method, data = {}, config) => {
    const response = await $apiAxios.post(method, data, this.getConfig())
    return response
  }

  put = async(method, data = {}, config) => {
    const response = await $apiAxios.put(method, data, this.getConfig())
    return response
  }

  get = async(method, config = {}) => {
    const response = await $apiAxios.get(method, this.getConfig())
    return response
  }

  delete = async(method) => {
    const response = await $apiAxios.delete(method, this.getConfig())
    return response
  }
}

const $api = new Api()

const $adminApi = new Api()

export { $api, $adminApi }

