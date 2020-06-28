import axios from 'axios';
import {API_URL} from './const';

export default {
  getArticleVote(articleID, authHeader) {
    return axios.get(API_URL+"/me/karma/article/"+articleID, {
      headers: {
        Authorization: authHeader
      }
    }).then(response => response.data.action);
  },
  getCommentVote(commentID, authHeader) {
    return axios.get(API_URL+"/me/karma/comment/"+commentID, {
      headers: {
        Authorization: authHeader
      }
    }).then(response => response.data.action);
  },
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