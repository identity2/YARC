import axios from 'axios';
import {API_URL} from './const';

export default {
  getUser(username) {
    return axios.get(API_URL+"/user/"+username).then(response => response.data);
  },
  modify(bio, authHeader) {
    return axios.put(API_URL+"/me/bio", {
      bio
    },{
      headers: {
        Authorization: authHeader
      }
    });
  },
  saveArticle(articleID, authHeader) {
    return axios.post(API_URL+"/me/save/"+articleID, {}, {
      headers: {
        Authorization: authHeader
      }
    });
  },
  unsaveArticle(articleID, authHeader) {
    return axios.delete(API_URL+"/me/save/"+articleID, {
      headers: {
        Authorization: authHeader
      }
    });
  },
  joinSubreddit(subreddit, authHeader) {
    return axios.post(API_URL+"/me/join/"+subreddit, {}, {
      headers: {
        Authorization: authHeader
      }
    });
  },
  leaveSubreddit(subreddit, authHeader) {
    return axios.delete(API_URL+"/me/join/"+subreddit, {
      headers: {
        Authorization: authHeader
      }
    });
  }
};