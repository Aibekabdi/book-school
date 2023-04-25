<template>
  <div class="home">
    <section class="banner__container">
      <div class="banner__img"><img src="../assets/Group 3245.png" alt="" /></div>
      <h1 class="banner__title">KITAP LAB</h1>

      <div class="landing-img"><img src="../assets/landing-img.jpg" alt="" /></div>
      <p>{{ t(currentLocalization, 'SLOGAN') }}</p>
      <div class="btn-container">
        <Button
          type="Router"
          link="/login"
          :label="t(currentLocalization, 'SIGN_IN')"
          color="info"
          size="large"
        />
        <Button
          type="Link"
          link="/payment"
          :label="t(currentLocalization, 'SIGN_UP')"
          color="success"
          icon="trending_flat"
          size="large"
        />
      </div>

      <div class="bgr__img">
        <img class="bgr__img-l" src="../assets/самолетик(1).png" alt="" />
        <img class="bgr__img-r" src="../assets/самолетик.png" alt="" />
      </div>
    </section>
    <!-- <div class="top-home">
      <div class="main-content">
        <h2 class="heading-1">Fairytale Platform - {{ t(currentLocalization, 'FE') }}</h2>
        <h3 class="heading-3">{{ t(currentLocalization, 'MFFE') }}</h3>
        <p>{{ t(currentLocalization, 'TB') }}</p>
        <div class="btn-container">
          <Button
            type="Router"
            link="/login"
            :label="t(currentLocalization, 'SIGN_IN')"
            color="info"
            size="large"
          />
          <Button
            type="Link"
            link="/payment"
            :label="t(currentLocalization, 'SIGN_UP')"
            color="success"
            icon="trending_flat"
            size="large"
          />
        </div>
      </div>
      <div class="picture">
        <img src="@/assets/undraw_Book_lover_re_rwjy.png" alt="" />
      </div>
    </div> -->

    <!-- <div class="categories-container">
      <h4 class="heading-2">{{ t(currentLocalization, 'CATEGORIES') }}</h4>
      <h6>{{ t(currentLocalization, 'MORE_THAN_1000') }}</h6>
      <div class="categories-grid">
        <div class="category-card" v-for="card in cards" :key="card.name">
          <p class="icon">{{ card.icon }}</p>
          <p class="name">{{ t(currentLocalization, card.name) }}</p>
        </div>
      </div>
    </div> -->

    <section class="about__company">
      <div class="about__bg-img">
        <img src="../assets/kitap.png" alt="" />
      </div>
      <h3 class="main__title">О компании</h3>
      <div class="about__company-text">
        <h4>
          {{ t(currentLocalization, 'MAIN_TEXT') }}
        </h4>
        <ul style="text-align: left">
          <li>{{ t(currentLocalization, 'FEATS1') }}</li>
          <li>{{ t(currentLocalization, 'FEATS2') }}</li>
          <li>{{ t(currentLocalization, 'FEATS3') }}</li>
          <li>{{ t(currentLocalization, 'FEATS4') }}</li>
          <li>{{ t(currentLocalization, 'FEATS5') }}</li>
          <li>{{ t(currentLocalization, 'FEATS6') }}</li>
          <li>{{ t(currentLocalization, 'FEATS7') }}</li>
        </ul>
      </div>
      <Button
        type="Link"
        link="/about"
        :label="t(currentLocalization, 'ABOUT_US')"
        color="yellow"
        size="large"
      />
      <div class="about__company-items">
        <div class="about__company-item blue">
          <div class="about__company-item-img">
            <img src="../assets/cup.png" alt="" />
          </div>
          <p>{{ t(currentLocalization, 'MAIN_BLUE') }}</p>
        </div>
        <div class="about__company-item green">
          <div class="about__company-item-img">
            <img src="../assets/teacher.png" alt="" />
          </div>
          <p>{{ t(currentLocalization, 'MAIN_GREEN') }}</p>
        </div>
        <div class="about__company-item red">
          <div class="about__company-item-img">
            <img src="../assets/calendar.png" alt="" />
          </div>
          <p>{{ t(currentLocalization, 'MAIN_RED') }}</p>
        </div>
        <div class="about__company-item orange">
          <div class="about__company-item-img">
            <img src="../assets/graduation.png" alt="" />
          </div>
          <p>{{ t(currentLocalization, 'MAIN_YELLOW') }}</p>
        </div>
      </div>
    </section>
    <section class="about__company instructions">
      <div class="instructions__bg">
        <img src="../assets/bg-g.png" alt="" />
      </div>
      <div class="about__bg-img right">
        <img src="../assets/kitap-r.png" alt="" />
      </div>
      <div class="about__bg-img">
        <img src="../assets/robot.png" alt="" />
      </div>
      <div class="himia">
        <img src="../assets/himia.png" alt="" />
      </div>
      <div class="robot"><img src="../assets/robot.png" alt="" /></div>
      <h3 class="main__title">{{ t(currentLocalization, 'RESULTS') }}</h3>
      <div class="about__company-text">
        <ol>
          <li>{{ t(currentLocalization, 'RES1') }}</li>
          <li>{{ t(currentLocalization, 'RES2') }}</li>
          <li>{{ t(currentLocalization, 'RES3') }}</li>
          <li>{{ t(currentLocalization, 'RES4') }}</li>
        </ol>
      </div>
      <h3 class="main__title">{{ t(currentLocalization, 'INSTRUCTIONS') }}</h3>
      <div class="about__company-text">
        <p>
          {{ t(currentLocalization, 'USECASE') }}
        </p>
        <p>
          {{ t(currentLocalization, 'RECOMEND_LIST') }}
        </p>
        <ol>
          <li>
            {{ t(currentLocalization, 'RECOMEND1') }}
            <ul>
              <li>{{ t(currentLocalization, 'AGE_TIME1') }}</li>
              <li>{{ t(currentLocalization, 'AGE_TIME2') }}</li>
              <li>{{ t(currentLocalization, 'AGE_TIME3') }}</li>
              <li>{{ t(currentLocalization, 'AGE_TIME4') }}</li>
              <li>{{ t(currentLocalization, 'AGE_TIME5') }}</li>
              <li>{{ t(currentLocalization, 'AGE_TIME6') }}</li>
              <li>{{ t(currentLocalization, 'AGE_TIME7') }}</li>
            </ul>
          </li>
          <li>{{ t(currentLocalization, 'RECOMEND2') }}</li>
          <li>{{ t(currentLocalization, 'RECOMEND3') }}</li>
          <li>{{ t(currentLocalization, 'RECOMEND4') }}</li>
          <li>{{ t(currentLocalization, 'RECOMEND5') }}</li>
        </ol>
      </div>
    </section>
  </div>
</template>
<script setup>
import { defineComponent } from 'vue';
import Button from '@/components/Button-component.vue';
import { currentLocalization } from '@/App.vue';
import { t } from '@/utils/i18n.js';
</script>
<script>
export default defineComponent({
  data() {
    return {
      cards: [
        { icon: '📚', name: 'BIG_COLLECTION' },
        { icon: '🇰🇿', name: 'KAZAKH_BOOKS' },
        { icon: '📘', name: 'TALES_FOR_KIDS' },
        { icon: '🧑‍🏫', name: 'WITH_TASK_CHECK' },
        { icon: '🎧', name: 'WITH_AUDIO' },
        { icon: '🔧', name: 'MBFE' },
        { icon: '🏆', name: 'WAAS' },
        { icon: '🖼️', name: 'TWI' },
        { icon: '👍', name: 'TWT' },
        { icon: '📈', name: 'AR' },
        { icon: '🎓', name: 'CWS' },
        { icon: '🍀', name: 'HWG' },
      ],
    };
  },
  name: 'HomeView',
  components: { Button },
});
</script>
<style scoped lang="scss">
@import url('https://fonts.googleapis.com/css2?family=Modak&display=swap');

// new banner
.landing-img {
  width: 100%;
  max-width: 800px;
  height: 350px;
}
.landing-img img {
  width: 100%;
  height: 100%;
  object-fit: contain;
}
.banner__img {
  width: 211px;
  height: 191px;
}
.banner__img img {
  width: 100%;
  height: 100%;
  object-fit: contain;
}
.bgr__img img {
  width: 170px;
  height: 90px;
}
.bgr__img-r {
  position: absolute;
  right: 150px;
  top: 300px;
}
.bgr__img-l {
  bottom: 150px;
  left: 150px;
  position: absolute;
}
.banner__container {
  margin: 0 -30px;
  padding-top: 30px;
  background: #84eee8;
  height: 100vh;
  min-height: 940px;
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
  & .banner__title {
    font-family: 'Modak';
    font-style: normal;
    font-weight: 400;
    font-size: 48px;
    color: #000000;
    position: relative;
    // padding-bottom: 30px;
  }
  & .banner__title::before {
    position: absolute;
    content: '';
    bottom: 0;
    left: 0;
    width: 100%;
    height: 2px;
    background: #000000;
  }
  & p {
    font-weight: 600;
    font-size: 28px;
    line-height: 86%;
    color: #000000;
    width: 645px;
  }
}
// new bannew ends

// about company
.about__company {
  background: #fff;
  margin: 0 -30px;
}
.about__company-items {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  grid-gap: 32px;
  margin-top: 100px;
  padding: 0 30px;

  & .about__company-item {
    border-radius: 10px;
    padding: 30px;
  }
  .about__company-item.orange {
    background: #ffed42;
  }
  .about__company-item.green {
    background: #76ff40;
  }
  .about__company-item.blue {
    background: #4285ff;
  }
  .about__company-item.red {
    background: #ff4440;
  }
  & .about__company-item p {
    margin-top: auto;
    font-weight: 600;
    font-size: 24px;
    line-height: 100%;
    text-align: center;
  }

  & .about__company-item-img {
    width: 130px;
    height: 130px;
    margin: 0 auto;
    margin-bottom: 20px;
  }
  & .about__company-item-img img {
    width: 100%;
    height: 100%;
    object-fit: contain;
  }
}
.about__bg-img {
  position: absolute;
  left: 0;
  top: 20%;
}
.about__bg-img.right {
  right: 0;
  left: unset;
  top: 0;
}
.about__company {
  position: relative;
  padding: 80px 0;
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
  & .about__company-text {
    font-size: 16px;
    line-height: 150%;
    width: 802px;
    margin-bottom: 30px;
    z-index: 50;
  }
}
.btn_yellow {
  display: inline-block;
  background: #ffeb00;
  box-shadow: 1.08px 0.76px 2.24px rgba(0, 0, 0, 0.56);
  font-weight: 500;
  text-align: center;
  letter-spacing: 0.02em;
  text-transform: uppercase;
  color: #000000;
  padding: 13px 70px;
  background: #ffeb00;
  box-shadow: 1.08px 0.76px 2.24px rgba(0, 0, 0, 0.56);
  border-radius: 10px;

  transition: all 0.4s;
}
.btn_yellow:hover {
  background: radial-gradient(
        96.87% 124.51% at 45.33% 23.44%,
        rgba(255, 255, 255, 0.49) 11.84%,
        rgba(255, 224, 224, 0) 63.33%,
        rgba(255, 224, 224, 0) 100%
      )
      /* warning: gradient uses a rotation that is not supported by CSS and may not behave as expected */,
    #ffeb00;
}
.main__title {
  font-weight: 600;
  font-size: 34px;
  line-height: 176%;
  color: #000000;
  position: relative;
  width: fit-content;
  margin-bottom: 30px;
  margin-top: 0;
}
.main__title::before {
  position: absolute;
  content: '';
  bottom: 0;
  left: 0;
  width: 100%;
  height: 2px;
  background: #000000;
}
// about company ends
// about__company instructions
.himia {
  position: absolute;
  bottom: 0;
  right: 10%;
}
.robot {
  position: absolute;
  top: 20%;
  left: 10%;
}
.about__company.instructions {
  position: relative;
  // overflow: hidden;
  height: 80vh;
  min-height: 100%;
}
.instructions__bg {
  position: absolute;
  top: 0px;
  width: 100%;
  height: 100%;
}
.instructions__bg img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}
// about__company instructions ends
.home {
  display: flex;
  flex-flow: column;
  width: 100%;
  // padding: 100px 0;
}

.top-home {
  display: flex;
  flex-flow: row nowrap;
  gap: 100px;
}

.heading {
  &-1 {
    color: var(--primary);
  }
  &-2 {
    color: var(--info);
  }
  &-3 {
    color: var(--second);
  }
}

.main-content {
  width: 50%;
}

.picture {
  width: 50%;
}

.picture img {
  width: 100%;
  height: auto;
  border-radius: 20px;
}

.categories-container {
  padding: 25px 0;
  width: 100%;
  text-align: center;
}

.categories-grid {
  display: flex;
  flex-flow: row wrap;
  justify-content: center;
  gap: 20px;
}

h6 {
  font-size: 16px;
  margin: 0 0 20px 0;
}

.category-card {
  display: flex;
  flex-flow: column;
  align-items: center;
  gap: 10px;
  padding: 10px;
  justify-content: center;
  width: 150px;
  height: 150px;
  color: var(--white);
  background: var(--black-hover);
  border-radius: 12px;
  transition: all 0.4s;

  & p {
    margin: 0;
  }

  &:hover {
    transform: translate(0, 2px);
    background: var(--black);
  }
}

@media only screen and (max-width: 1440px) {
  .about__bg-img {
    width: 300px;
    height: 250px;
  }
  .robot {
    width: 300px;
    height: 300px;
    left: 10px;
  }
  .himia {
    width: 300px;
    height: 130px;
    bottom: 0;
  }
  .robot img,
  .himia img,
  .about__bg-img img {
    width: 100%;
    height: 100%;
    object-fit: contain;
  }
  .about__company-items .about__company-item {
    border-radius: 10px;
  }
  .about__company-items .about__company-item-img {
    width: 90px;
    height: 90px;
    margin-bottom: 15px;
  }
  .about__company-items .about__company-item p {
    font-size: 20px;
  }
  .banner__container .banner__title {
    font-size: 35px;
  }
  .banner__container p {
    font-size: 20px;
    width: 450px;
  }
}
@media only screen and (max-width: 1260px) {
  .bgr__img {
    display: none;
  }
}
@media only screen and (max-width: 1000px) {
  .robot img,
  .himia img,
  .about__bg-img img {
    display: none;
  }
  .banner__container {
    width: 100vw;
  }
  .about__bg-img {
    display: none;
  }
  .about__company {
    padding: 40px 0;
  }
  .about__company-items {
    grid-template-columns: repeat(2, 1fr);
    grid-gap: 25px;
    margin-top: 80px;
  }
}
@media only screen and (max-width: 768px) {
  .about__company .about__company-text {
    font-size: 14px;
    width: 90%;
    margin-bottom: 20px;
  }
  .about__company-items .about__company-item p {
    font-size: 18px;
  }
  .home {
    // padding: 40px 0;
  }

  .top-home {
    flex-flow: column wrap;
    gap: 20px;
  }

  .main-content,
  .picture {
    width: 100%;
  }

  .banner__container .banner__title {
    font-size: 30px;
  }
  .banner__container p {
    font-size: 16px;
    width: 70%;
  }
  .btn-container {
    justify-content: center;
  }
  .banner__img {
    width: 150px;
    height: 100px;
  }
}
@media only screen and (max-width: 508px) {
  .about__company-items {
    grid-template-columns: repeat(1, 1fr);
    grid-gap: 20px;
    margin-top: 50px;
  }
}
</style>
