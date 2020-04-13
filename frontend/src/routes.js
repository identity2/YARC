import Vue from 'vue';
import VueRouter from 'vue-router';
import ListOfArticles from './components/page/ListOfArticles';

Vue.use(VueRouter);

export const router = new VueRouter({
  mode: 'history',
  routes: [
    {path: '/', component: ListOfArticles}
  ]
});
