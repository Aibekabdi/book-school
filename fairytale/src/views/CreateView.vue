<template>
  <Notification class="notification" v-if="onError" :message="onError" />
  <div class="create max-w">
    <Button
      type="Router"
      link="/admin"
      :label="t(currentLocalization, 'TO_ADMIN_PANEL')"
      color="primary"
      icon="admin_panel_settings"
      width="fit-content"
    />
    <h2 class="heading-2">{{ t(currentLocalization, 'CREATE_BOOK') }}</h2>

    <form class="editor" name="document" id="formElem" enctype="multipart/form-data">
      <div>
        <div
          class="imagePreviewWrapper"
          :style="{ 'background-image': `url(${previewImage})` }"
          @click="selectImage"
        >
          <img
            v-if="!previewImage"
            src="https://cdn-icons-png.flaticon.com/512/2899/2899181.png"
            alt=""
          />
        </div>

        <label for="file-upload" class="file-upload">
          <span class="material-symbols-outlined"> cloud_upload </span>
          <span>{{ t(currentLocalization, 'UPLOAD_PIC') }}</span>
        </label>
        <input
          ref="fileInput"
          type="file"
          name="preview"
          accept="image/*"
          id="file-upload"
          @input="pickFile"
        />
      </div>
      <Tabs
        :names="tTabs(currentLocalization, tabs)"
        :selectedTab="selectedTab"
        @changeTab="changeTab"
      >
        <div v-if="selectedTab === 'kindergarden'">
          <div class="dropdowns">
            <div>
              <span class="heading-3">{{ t(currentLocalization, 'SELECT_LANG') }}:</span>
              <v-select :options="lang" v-model="bookData.lang"></v-select>
            </div>
            <div>
              <span class="heading-3">{{ t(currentLocalization, 'SELECT_CATEGORY') }}:</span>
              <v-select
                :options="categories_kinder.map((e) => t(currentLocalization, e))"
                v-model="bookData.category"
              ></v-select>
            </div>
            <div>
              <span class="heading-3">{{ t(currentLocalization, 'SELECT_AGE') }}:</span>
              <v-select
                :options="classes_kinder.map((e) => t(currentLocalization, e))"
                v-model="bookData.class"
              ></v-select>
            </div>
          </div>
        </div>
        <div v-if="selectedTab === 'pupil'">
          <div class="dropdowns">
            <div>
              <span class="heading-3">{{ t(currentLocalization, 'SELECT_LANG') }}:</span>
              <v-select :options="lang" v-model="bookData.lang"></v-select>
            </div>
            <div>
              <span class="heading-3">{{ t(currentLocalization, 'SELECT_CATEGORY') }}:</span>
              <v-select
                :options="categories_pupil.map((e) => t(currentLocalization, e))"
                v-model="bookData.category"
              ></v-select>
            </div>
            <div>
              <span class="heading-3">{{ t(currentLocalization, 'SELECT_CLASS') }}:</span>
              <v-select
                :options="classes_pupil.map((e) => t(currentLocalization, e))"
                v-model="bookData.class"
              ></v-select>
            </div>
          </div>
        </div>
      </Tabs>
      <Input
        :label="t(currentLocalization, 'TITLE')"
        name="headline"
        :placeholder="t(currentLocalization, 'ENTER_WORK_NAME')"
        v-model:value="bookData.headline"
        width="100%"
      />
      <QuillEditor
        theme="snow"
        toolbar="#custom-toolbar"
        v-model:content="bookData.mainText"
        contentType="html"
        :options="{
          ...editorOptions,
          placeholder: t(currentLocalization, editorOptions.placeholder),
        }"
      >
        <template #toolbar>
          <div class="custom-toolbar" id="custom-toolbar">
            <!-- <span class="ql-formats">
              <button class="ql-bold"></button>
              <button class="ql-italic"></button>
              <button class="ql-underline"></button>
              <button class="ql-strike"></button>
            </span> -->

            <!-- <span class="ql-formats">
              <button class="ql-blockquote"></button>
              <button class="ql-code-block"></button>
            </span> -->

            <!-- <span class="ql-formats">
              <button class="ql-header" value="1"></button>
              <button class="ql-header" value="2"></button>
            </span> -->

            <!-- <span class="ql-formats">
              <button class="ql-list" value="ordered"></button>
              <button class="ql-list" value="bullet"></button>
            </span> -->

            <!-- <span class="ql-formats">
              <button class="ql-indent" value="-1"></button>
              <button class="ql-indent" value="+1"></button>
            </span> -->

            <!-- <span class="ql-formats">
              <select class="ql-align">
                <option selected></option>
                <option value="center"></option>
                <option value="right"></option>
                <option value="justify"></option>
              </select>
            </span> -->

            <!-- <div class="ql-formats">
              <select class="ql-header">
                <option selected></option>
                <option value="1"></option>
                <option value="2"></option>
                <option value="3"></option>
                <option value="4"></option>
                <option value="5"></option>
                <option value="6"></option>
              </select>
            </div> -->

            <!-- <div class="ql-formats">
              <select class="ql-size">
                <option value="small"></option>
                <option selected></option>
                <option value="large"></option>
                <option value="huge"></option>
              </select>
            </div> -->

            <!-- <div class="ql-formats">
              <select class="ql-font">
                <option selected="selected"></option>
                <option value="serif"></option>
                <option value="monospace"></option>
              </select>
            </div> -->

            <!-- <div class="ql-formats">
              <select class="ql-background color">
                <option selected></option>
                <option v-for="item in colors" :value="item"></option>
              </select>
            </div> -->

            <!-- <div class="ql-formats">
              <select class="ql-color color">
                <option selected></option>
                <option v-for="item in colors" :value="item"></option>
              </select>
            </div> -->

            <span class="ql-formats">
              <button type="button" class="ql-image"></button>
            </span>

            <span class="ql-formats">
              <button @click.prevent="createSplit">
                <span class="material-symbols-outlined"> splitscreen </span>
              </button>
            </span>
          </div>
        </template>
      </QuillEditor>
      <div class="line"></div>
      <button class="submit-btn" type="submit" @click.prevent="createBook">
        <span class="material-symbols-outlined"> add_circle </span>&nbsp;&nbsp;<b>{{
          t(currentLocalization, 'CB')
        }}</b>
      </button>
    </form>
    <span v-if="isSuc" class="success-message">{{ t(currentLocalization, 'CB_SUC') }}</span>
    <span v-else-if="isErr" class="error-message">{{ t(currentLocalization, 'REG_ERR') }}</span>
  </div>
</template>
<script setup>
import Tabs from '@/components/Tab-component.vue';
import Input from '@/components/Input-component.vue';
import Button from '@/components/Button-component.vue';
import Notification from '@/components/layout/Notification-component.vue';
import vSelect from 'vue-select';
import 'vue-select/dist/vue-select.css';
import '@vueup/vue-quill/dist/vue-quill.snow.css';
import { QuillEditor } from '@vueup/vue-quill';
import { ref, reactive } from 'vue';
import { user, userToken, onError, currentLocalization } from '@/App.vue';
import host from '@/main';
import { t, tTabs } from '@/utils/i18n.js';
</script>
<script>
const bookData = reactive({
  mainText: '',
  headline: '',
  category: '',
  class: '',
  lang: '',
});

const splittedContent = ref(null);

const tabs = [
  { name: 'kindergarden', label: 'KINDER' },
  { name: 'pupil', label: 'PUPIL' },
];
const selectedTab = ref('kindergarden');
const changeTab = (tabName) => {
  selectedTab.value = tabName;
};

const categories_pupil = ['WBW', 'JIT', 'WIGAWIB', 'WN', 'POETRY'];

const categories_kinder = ['WOT', 'AAN', 'RAF', 'AOB', 'BOG'];

const lang = ['🇷🇺 RU', '🇰🇿 KZ'];

const classes_pupil = ['CLASS_1', 'CLASS_2', 'CLASS_3', 'CLASS_4'];

const classes_kinder = ['YEAR_2', 'YEAR_3', 'YEAR_4', 'YEAR_5'];

const isSuc = ref(false);
const isErr = ref(false);
const pages = ref(0);

const colors = [
  '#333333',
  '#e60000',
  '#ff9900',
  '#ffff00',
  '#008a00',
  '#0066cc',
  '#9933ff',
  '#facccc',
  '#ffebcc',
  '#ffffcc',
  '#cce8cc',
  '#cce0f5',
  '#ebd6ff',
  '#bbbbbb',
  '#f06666',
  '#ffc266',
  '#ffff66',
  '#66b966',
  '#66a3e0',
  '#c285ff',
  '#888888',
  '#a10000',
  '#b26b00',
  '#b2b200',
  '#006100',
  '#0047b2',
  '#6b24b2',
  '#444444',
  '#5c0000',
  '#663d00',
  '#666600',
  '#003700',
  '#002966',
  '#3d1466',
];

let file = null;

const url = ref(null);
const divider = `<pre class=\"ql-syntax ql-align-center\" spellcheck=\"false\">НОВАЯ СТРАНИЦА\n</pre>`;

export default {
  name: 'CreateView',
  components: {
    QuillEditor,
    Button,
    Input,
    vSelect,
    Tabs,
  },
  methods: {
    selectImage() {
      this.$refs.fileInput.click();
    },
    pickFile() {
      let input = this.$refs.fileInput;
      file = input.files;
      console.log(input.files);
      if (input.files.size > 2097152) {
        alert(t(currentLocalization.value, 'TOO_BIG_PIC'));
        file = '';
        return;
      }
      if (file && file[0]) {
        let reader = new FileReader();
        reader.onload = (e) => {
          this.previewImage = e.target.result;
        };
        reader.readAsDataURL(file[0]);
        this.$emit('input', file[0]);
      }
    },
    createSplit() {
      // pages.value++;
      if (bookData.mainText.includes(divider)) {
        pages.value = bookData.mainText.split(divider).length;
      } else {
        pages.value++;
      }
      bookData.mainText += '<br>' + divider + '<br>';

      console.log(pages.value);
    },
    async createBook() {
      isSuc.value = false;
      isErr.value = false;

      splittedContent.value = bookData.mainText.split(divider);

      const json = {
        name: bookData.headline,
        pages: splittedContent.value,
        category: bookData.category,
        class: bookData.class,
        language: bookData.lang.slice(-2),
      };

      console.log(host);

      console.log(json);

      let form = new FormData();

      if (file !== null) {
        form.append('preview', file[0]);
      }

      form.append('document', JSON.stringify(json));

      let response = await fetch(`${host}/api/books/create`, {
        method: 'POST',
        body: form,
        headers: {
          // 'Content-type': 'multipart/form-data; charset=UTF-8',
          Accept: '*/*',
          Authorization: `Bearer ${userToken.value}`,
          'Access-Control-Allow-Origin': `${host}/`,
          'Access-Control-Allow-Methods': 'GET, POST, OPTIONS, PUT, PATCH, DELETE',
          'Access-Control-Allow-Headers': 'origin,X-Requested-With,content-type,accept',
          'Access-Control-Allow-Credentials': 'true',
        },
      });

      let result = await response.json();

      if (result.status != 'OK') {
        isErr.value = true;

        onError.value = 'ITSD';
        console.log('NOT OK');

        setTimeout(() => {
          onError.value = null;
        }, 5000);
      } else {
        isSuc.value = true;
      }
    },
  },
  beforeRouteLeave(to, from, next) {
    if (bookData.headline || bookData.mainText || bookData.category || bookData.class.classroom) {
      const answer = window.confirm(t(currentLocalization.value, 'AREYOUSURE'));
      if (answer) {
        next();
      } else {
        next(false);
      }
    } else {
      next();
    }
  },
  mounted() {
    user.value = localStorage.getItem('admin') ? 'admin' : localStorage.getItem('userType') || '';
    userToken.value = localStorage.getItem('admin') || localStorage.getItem('userToken') || '';
  },
  data() {
    return {
      url,
      model: '',
      fileInput: '',
      previewImage: null,
      selectedFile: '',
      editorContent: '',
      editorOptions: {
        theme: 'snow',
        placeholder: 'WTTOBH',
      },
    };
  },
};
</script>
<style scoped lang="scss">
.editor {
  padding: 20px 0;
  display: flex;
  flex-flow: column;
  gap: 10px;
  text-align: center;

  & .dropdowns {
    align-self: center;
    display: flex;
    flex-direction: column;
    gap: 20px;
    max-width: 400px;
    width: 400px;
    margin: 20px 0;
  }

  & .dropdowns > div {
    text-align: left;
    display: flex;
    flex-direction: column;
    width: 100%;
  }
}

.custom-toolbar {
  position: sticky;
  top: 80px;
  background-color: #d1d5db;
  padding: 20px 0;
  z-index: 5;
  border-radius: 12px 12px 0 0;
}

.ql-formats span {
  font-size: 18px;
}

.ql-snow .ql-picker.ql-header,
.ql-snow .ql-picker.ql-font {
  width: 132px;
}

.success-message,
.error-message {
  margin-top: 20px;
  display: inline-block;
  width: 100%;
  text-align: center;
}

.submit-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0 10px;
  height: 40px;
  color: #fff;
  border-radius: 12px;
  border: none;
  cursor: pointer;
  font-size: 14px;
  transition: 0.4s;
  background-color: var(--primary);
  border: 1px solid var(--primary);
  &:hover {
    background-color: var(--primary-hover);
  }
}

.imagePreviewWrapper {
  min-width: 250px;
  width: 250px;
  min-height: 250px;
  height: 250px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  margin: 0 auto 30px;
  background-size: cover;
  background-position: center center;
  border: 2px solid var(--primary);
  background-color: rgba(0, 0, 0, 0.8);
  z-index: 2;

  & > img {
    z-index: 0;
    width: 70%;
    height: 70%;
  }
}
input[type='file'] {
  display: none;
}
.file-upload {
  display: flex;
  justify-content: center;
  width: fit-content;
  gap: 4px;
  align-items: center;
  cursor: pointer;
  border: 2px solid var(--primary);
  padding: 6px 12px;
  color: var(--black);
  border-radius: 12px;
  margin: 0 auto;
  transition: all 0.4s;

  &:hover {
    color: var(--black-hover);
    border-color: var(--primary-hover);
  }
}
</style>
