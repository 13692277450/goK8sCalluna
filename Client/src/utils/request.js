import axios from 'axios'
import { severalUrl } from './tools.js'

const instanceAxios = axios.create({
  baseURL: severalUrl,
  timeout: 1000,
})

export const get = (url, params = {} ) => instanceAxios.get(url,{params})
//   return new Promise((resolve, reject) => {
//     instanceAxios.get(url, params).then(res => {
//       if (res.status === 200) {
//         resolve(res.data)
//       } else {
//         reject('error')
//       }
//     })
//   })
// }