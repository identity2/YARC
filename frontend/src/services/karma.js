import axios from 'axios';
import {API_URL} from './const';
import article from './article';

export default {
  voteArticle(articleID, action, authHeader) {
    return axios.post(API_URL+"/karma/article/"+articleID, {}, {
      params: {
        action
      },
      headers: {
        Authorization: authHeader
      }
    });
  },
  voteComment(commentID, action, authHeader) {
    return axios.post(API_URL+"/karma/comment/"+commentID, {}, {
      params: {
        action
      },
      headers: {
        Authorization: authHeader
      }
    });
  }
};