import Vue from 'vue';
import App from './App.vue';
import {router} from './routes';
import {store} from './store';
import './quasar';
import moment from 'moment';

Vue.config.productionTip = false;

Vue.filter('formatDate', value => {
  if (value) {
    return moment(String(value)).format('YYYY-MM-DD hh:mm:ss a');
  }
});

new Vue({
  render: h => h(App),
  router,
  store,
}).$mount('#app');
