import axios from 'axios';
import {API_URL} from './const';

export default {
  getList() {
    return axios.get(API_URL+"/subreddit").then(response => response.data.subreddits);
  },
  get(name) {
    return axios.get(API_URL+"/subreddit/"+name).then(response => response.data);
  },
  create(name, description, authHeader) {
    return axios.post(API_URL+"/subreddit", {
      name,
      description
    }, {
      headers: {
        Authorization: authHeader
      }
    });
  },
  getTrending(limit) {
    return axios.get(API_URL+"/trending", {
      params: {
        limit
      }
    }).then(response => response.data.subreddits);
  }
};