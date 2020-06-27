import axios from 'axios';
import {API_URL} from './const';

export default {
  getList(sortBy, after, limit, criterion, key, authHeader) {
    return axios.get(API_URL+"/article", {
      params: {
        sort: sortBy,
        after,
        limit,
        criterion,
        key
      },
      headers: {
        Authorization: authHeader
      }
    }).then(response => response.data.articles);
  },
  get(id) {
    return axios.get(API_URL+"/article/"+id).then(response => response.data);
  },
  create(subreddit, type, body, title, authHeader) {
    return axios.post(API_URL+"/article",{
      subreddit,
      type,
      body,
      title
    }, {
      headers: {
        Authorization: authHeader
      }
    }).then(response => response.data.articleID);
  },
  modify(id, body, authHeader) {
    return axios.put(API_URL+"/article/"+id, {
      body
    }, {
      headers: {
        Authorization: authHeader
      }
    });
  },
  remove(id, authHeader) {
    return axios.delete(API_URL+"/article/"+id, {
      headers: {
        Authorization: authHeader
      }
    });
  }
};