import axios from 'axios';
import {API_URL} from './const';

export default {
  search(queryStr) {
    return axios.get(API_URL+"/search", {
      params: {
        q: queryStr
      }
    }).then(response => response.data);
  }
};