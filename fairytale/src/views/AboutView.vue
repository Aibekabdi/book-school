<template>
  <div class="about">
    <img src="../assets/wave.png" class="wave" alt="" />
    <div class="img"><img src="../assets/kitap-r.png" alt="" /></div>
    <h2 class="heading-2">{{ t(currentLocalization, 'ABOUT_HELLO') }}</h2>
    <h3 class="title">{{ t(currentLocalization, 'ABOUT_US') }}</h3>
    <div class="aboutus">
      <div class="aboutus__title">
        {{ t(currentLocalization, 'ABOUT_PROJECT') }}
        <div class="line"></div>
      </div>
      <div class="aboutus__text">
        <p>
         {{ t(currentLocalization, 'ABOUT_PROJECT_TEXT') }}
        </p>
      </div>
    </div>
    <div class="aboutus">
      <div class="aboutus__title">
        {{ t(currentLocalization, 'ABOUT_MISSION') }}
        <div class="line"></div>
      </div>
      <div class="aboutus__text">
        <p>
         {{ t(currentLocalization, 'ABOUT_MISSION_TEXT') }}
        </p>
      </div>
    </div>
    <div class="aim">
      <div class="aim__img"><img src="../assets/aim-robot.png" alt="" /></div>
      <div class="aim__text">
        <div class="aim__title">{{ t(currentLocalization, 'ABOUT_AIM') }}</div>
        <ol>
          <li>
            {{ t(currentLocalization, 'OM1') }}
          </li>
          <li>
            {{ t(currentLocalization, 'OM2') }}
          </li>
          <li>
            {{ t(currentLocalization, 'OM3') }}
          </li>
          <li>
            {{ t(currentLocalization, 'OM4') }}
          </li>
          <li>
            {{ t(currentLocalization, 'OM5') }}
          </li>
        </ol>
      </div>
    </div>
    <div class="aboutus">
      <div class="aboutus__title">
        {{ t(currentLocalization, 'ABOUT_METHODS') }}
        <div class="line"></div>
      </div>
      <div class="aboutus__text">
        <p>
         {{ t(currentLocalization, 'ABOUT_METHODS_TEXT') }}
        </p>
      </div>
    </div>
    <div class="aboutus">
      <div class="aboutus__title">
        {{ t(currentLocalization, 'ABOUT_CREATION_AIM') }}
        <div class="line"></div>
      </div>
      <div class="aboutus__text">
        <p>
         {{ t(currentLocalization, 'ABOUT_CREATION_AIM_TEXT') }}
        </p>
      </div>
    </div>
    <div class="aboutus">
      <div class="aboutus__title">
        {{ t(currentLocalization, 'ABOUT_KIDS_LIKE') }}
        <div class="line"></div>
      </div>
      <div class="aboutus__text">
        <p>
         {{ t(currentLocalization, 'ABOUT_KIDS_LIKE_TEXT') }}
        </p>
      </div>
    </div>
    <div class="aboutus">
      <div class="aboutus__title">
        {{ t(currentLocalization, 'ABOUT_HELP_PARRENTS') }}
        <div class="line"></div>
      </div>
      <div class="aboutus__text">
        <p>
         {{ t(currentLocalization, 'ABOUT_HELP_PARRENTS_TEXT') }}
        </p>
      </div>
    </div>
    <div class="aboutus">
      <div class="aboutus__title">
        {{ t(currentLocalization, 'ABOUT_DEVELOPMENT') }}
        <div class="line"></div>
      </div>
      <div class="aboutus__text">
        <p>
         {{ t(currentLocalization, 'ABOUT_DEVELOPMENT_TEXT') }}
        </p>
      </div>
    </div>
  </div>
</template>
<script setup>
import { currentLocalization } from '@/App.vue';
import { t } from '@/utils/i18n.js';
</script>
<style scoped>
.aboutus,
.aim {
  z-index: 50;
}
.wave {
  position: absolute;
  bottom: 0;
  left: 0;
  width: 100%;
}
.aim__title {
  font-weight: 600;
  font-size: 24px;
  line-height: 150%;
  color: #000000;
}

.aim__text {
  font-size: 16px;
  line-height: 150%;
  color: #000000;
  border: 2px solid #000000;
  border-left: none;
  padding-top: 10px;
}
.aim__text li {
  margin-bottom: 20px;
}
.aim {
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.aim__img {
  width: 554px;
  height: 685px;
}
.aim__img img {
  width: 100%;
  height: 100%;
  object-fit: contain;
}
.aim__img,
.aim__text {
  width: 50%;
  display: flex;
  flex-direction: column;
}
.heading-2 {
  font-weight: 600;
  font-size: 42px;
  color: #000000;
  margin-bottom: 80px;
}
.title {
  font-weight: 600;
  font-size: 34px;
  color: #000000;
  position: relative;
  padding-bottom: 30px;
}
.title::before {
  position: absolute;
  content: '';
  bottom: 0;
  left: 0;
  width: 100%;
  height: 3px;
  background: #000000;
}

.aboutus {
  width: 80%;
  max-width: 2400px;
  margin-bottom: 80px;
}
.aboutus__title {
  font-weight: 600;
  font-size: 24px;
  color: #000000;
  display: flex;
  align-items: center;
  white-space: nowrap;
}
.aboutus__title .line {
  width: 100%;
  height: 2px;
  background: #000000;
  margin-left: 50px;
}
.aboutus__text {
  font-size: 16px;
  line-height: 150%;
  text-align: justify;
}
.about {
  background: #fff;
  margin: 0 -30px;
  padding: 0 30px;
  display: flex;
  flex-direction: column;
  align-items: center;
  position: relative;
}
.img {
  position: absolute;
  right: 0;
  top: 10%;
  width: 276px;
  height: 244px;
}
.img img {
  width: 100%;
  height: 100%;
  object-fit: contain;
}

@media only screen and (max-width: 1560px) {
  .aim__img {
    height: 400px;
  }
  .heading-2 {
    font-size: 36px;
    margin-bottom: 60px;
  }
  .title {
    font-size: 28px;
    padding-bottom: 20px;
  }
  .aboutus__title {
    font-size: 20px;
  }
  .aboutus__text[data-v-d4948df9] {
    font-size: 14px;
  }
  .img[data-v-d4948df9] {
    top: 0;
  }
}
@media only screen and (max-width: 1000px) {
  .aboutus__title .line {
    margin-left: 20px;
  }
  .img {
    display: none;
  }
  .heading-2 {
    font-size: 32px;
    margin-bottom: 20px;
  }
  .title {
    font-size: 24px;
    padding-bottom: 15px;
  }
  .aboutus {
    margin-bottom: 50px;
  }
  .aim {
    flex-direction: column;
  }
  .aim__text {
    width: 100%;
  }
  .aim__text {
    font-size: 14px;
  }
}
@media only screen and (max-width: 500px) {
  .aboutus {
    margin-bottom: 30px;
  }
  .heading-2 {
    font-size: 25px;
  }
  .title {
    font-size: 22px;
    padding-bottom: 10px;
  }
  .aim__img {
    height: 200px;
  }
}
</style>
