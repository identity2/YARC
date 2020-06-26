import axios from 'axios';
import {API_URL} from './const';

export default {
  login(username, password) {
    return axios.post(API_URL+"/login",{
      username,
      password
    }).then((response) => {
      var user = {
        username,
        authHeader: "bearer " + response.data.token
      };
      
      // Save to the local storage, so the user don't need to log in again.
      localStorage.setItem("user", JSON.stringify(user));

      return user;
    });
  },
  logout() {
    localStorage.removeItem("user");
  },
  register(username, password, email) {
    return axios.post(API_URL+"/register",{
      username,
      password,
      email
    });
  }
};