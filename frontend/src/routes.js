import Vue from 'vue';
import VueRouter from 'vue-router';
import ListOfArticles from './components/page/ListOfArticles';
import CreateArticle from './components/page/CreateArticle';

Vue.use(VueRouter);

export const router = new VueRouter({
  mode: 'history',
  routes: [
    {
      name: 'home',
      path: '/',
      component: ListOfArticles
    },
    {
      name: 'submit',
      path: '/submit',
      component: CreateArticle
    }
  ]
});
