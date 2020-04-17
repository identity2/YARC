import Vue from 'vue';
import VueRouter from 'vue-router';
import ListOfArticles from './components/page/ListOfArticles';
import CreateArticle from './components/page/CreateArticle';
import ViewArticle from './components/page/ViewArticle';

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
    },
    {
      name: 'subreddit',
      path: '/r/:subreddit',
      component: ListOfArticles
    },
    {
      name: 'article',
      path: '/r/:subreddit/comments/:pid',
      component: ViewArticle
    }
  ]
});
