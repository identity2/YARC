import axios from 'axios';
import {IMGUR_AUTH_HEADER} from './const';

export default {
  upload(title, image) {
    let formData = new FormData();
    formData.append('title', title);
    formData.append('image', image);

    return axios.post('https://api.imgur.com/3/image', formData, {
      headers: {
        Authorization: IMGUR_AUTH_HEADER
      },
      mimeType: 'multipart/form-data'
    }).then(response => response.data.data.link);
  }
};