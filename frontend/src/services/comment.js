import axios from 'axios';
import {API_URL} from './const';

export default {
  getList(after, limit, ofArticle, postedBy) {
    return axios.get(API_URL+"/comment", {
      params: {
        after,
        limit,
        ofArticle,
        postedBy
      }
    }).then(response => response.data.comments);
  },
  get(id) {
    return axios.get(API_URL+"/comment/"+id).then(response => response.data);
  },
  create(articleID, body, authHeader) {
    return axios.post(API_URL+"/comment", {
      articleID,
      body
    }, {
      headers: {
        Authorization: authHeader
      }
    }).then(response => response.data.commentID);
  },
  modify(id, body, authHeader) {
    return axios.put(API_URL+"/comment/"+id, {
      body
    }, {
      headers: {
        Authorization: authHeader
      }
    });
  },
  remove(id, authHeader) {
    return axios.delete(API_URL+"/comment/"+id, {
      headers: {
        Authorization: authHeader
      }
    });
  }
};