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
  getSaveState(articleID, authHeader) {
    return axios.get(API_URL+"/me/save/"+articleID, {
      headers: {
        Authorization: authHeader
      }
    }).then(response => response.data.saved);
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
  getJoinState(subreddit, authHeader) {
    return axios.get(API_URL+"/me/join/"+subreddit, {
      headers: {
        Authorization: authHeader
      }
    }).then(response => response.data.joined);
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