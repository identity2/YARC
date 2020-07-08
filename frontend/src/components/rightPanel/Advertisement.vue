<template>
  <q-card dark flat bordered class="bg-grey-10 q-mb-md">
    <q-card-section>
      <div class="q-mb-md">
        Advertisement
      </div>
      <a v-if="currAdvert" :href="currAdvert.link" target="_blank"><img :src="currAdvert.img" style="width: 100%" /></a>
    </q-card-section>
  </q-card>
</template>

<script>
const adverts = [
  {img: 'https://i.imgur.com/meeL3Z4.png', link: 'https://www.youtube.com/channel/UCx_wmkyTwDVuYWULSTBlYyw'},
  {img: 'https://i.imgur.com/p97Item.png', link: 'http://shinerightstudio.com/apnea-app/'},
  {img: 'https://i.imgur.com/bpwXei3.png', link: 'https://www.instagram.com/yuchao.jpg/'},
  {img: 'https://i.imgur.com/XqnbrvP.png', link: 'https://www.instagram.com/tilted_food_poorn/'},
  {img: 'https://i.imgur.com/f0R26lJ.jpg', link: 'https://www.amazon.com/Inter-dimensional-Time-Yu-Chao-ebook/dp/B07HDSTV5F'},
  {img: 'https://i.imgur.com/GnP4qaD.jpg', link: 'https://store.steampowered.com/app/980440/Knights_Rubbish'},
  {img: 'https://i.imgur.com/LKuKIAL.png', link: 'http://shinerightstudio.com/kqq/'},
  {img: 'https://i.imgur.com/DH2GDcL.png', link: 'https://store.line.me/stickershop/author/122259/'},
  {img: 'https://i.imgur.com/bMjuASH.png', link: 'http://shinerightstudio.com/'},
  {img: 'https://i.imgur.com/0uqhnrV.png', link: 'http://shinerightstudio.com/recordable-launchpad/'},
  {img: 'https://i.imgur.com/Gjxqqzy.png', link: 'http://shinerightstudio.com/tall-as-the-sky/'}
];

const SWITCH_INTERVAL = 30;

export default {
  data() {
    return {
      currAdvert: null
    };
  },
  async mounted() {
    let index = Math.floor(Math.random() * adverts.length);
    this.currAdvert = adverts[index];

    for(;;) {
      // Wait for some time and switch to the next ad.
      await new Promise(resolve => setTimeout(resolve, SWITCH_INTERVAL * 1000));
      index = (index + 1) % adverts.length;
      this.currAdvert = adverts[index];
    }
  }
}
</script>