<template>
  <Notification class="notification" v-if="onError" :message="onError" />
  <div class="paint">
    <canvas
      width="1000"
      height="400"
      ref="canvas"
      class="paint__canvas"
      @mousedown="startDrawing"
      @mousemove="draw"
      @mouseup="stopDrawing"
    ></canvas>
    <div class="paint__color" v-if="showPalette">
      <div
        v-for="(item, i) in colors"
        @click="setColor(item), (showPalette = false)"
        class="color-box"
        :class="color === item ? 'color-picked' : ''"
        :key="item"
        :style="{ background: item }"
      >
        <span v-if="color === item" class="material-symbols-outlined"> done_outline </span>
      </div>
    </div>
    <div class="paint__controls">
      <div class="paint__control row">
        <img @click="undo" class="tool" src="../assets/undo-1.png" alt="" />
        <img @click="redo" class="tool" src="../assets/undo-2.png" alt="" />
      </div>
      <div class="paint__control row">
        <div @click="lineWidth = 1" class="size size-1"></div>
        <div @click="lineWidth = 3" class="size size-3"></div>
        <div @click="lineWidth = 5" class="size size-5"></div>
        <div @click="lineWidth = 7" class="size size-7"></div>
        <div @click="lineWidth = 10" class="size size-10"></div>
      </div>
      <div class="paint__control row">
        <img @click="showPalette = !showPalette" class="tool" src="../assets/colors.png" alt="" />
      </div>
      <div class="paint__control">
        <img
          v-if="!this.eraser"
          @click="toggleEraser"
          class="tool"
          src="../assets/eraser.png"
          alt=""
        />
        <img
          v-else-if="this.eraser"
          @click="toggleEraser"
          class="tool"
          src="../assets/pencil.png"
          alt=""
        />
      </div>
      <div class="paint__control row">
        <img @click="clear" class="tool" src="../assets/paint.png" alt="" />
      </div>
    </div>
    <div class="btn">
      <Button
        type="Button"
        @click="saveImage"
        :label="t(currentLocalization, 'SEND')"
        icon="save"
        color="success"
        size="large"
      />
    </div>
  </div>
</template>
<script setup>
import Button from './Button-component.vue';
import Notification from '@/components/layout/Notification-component.vue';
import { defineProps } from 'vue';
import { userToken, onError, currentLocalization } from '@/App.vue';
import host from '@/main';
import { t } from '@/utils/i18n.js';
defineProps({
  BookId: {
    type: Number,
    required: true,
  },
  QuestionId: {
    type: Number,
    required: true,
  },
});
</script>
<script>
export default {
  components: {
    Button,
  },
  data() {
    return {
      lineWidth: 3,
      color: 'black',
      colors: [
        'red',
        'orange',
        'yellow',
        'green',
        'blue',
        'purple',
        'pink',
        'brown',
        'gray',
        'black',
        'white',
        'teal',
        'maroon',
        'navy',
        'olive',
        'coral',
        'lavender',
        'turquoise',
      ],
      eraser: false,
      drawing: false,
      showPalette: false,
      lastX: 0,
      lastY: 0,
      undoStack: [],
      redoStack: [],
    };
  },
  methods: {
    getMicAccess() {
      /* navigator.mediaDevices
        .getUserMedia({video:false, audio:true})
        .then((stream)=>{
          window.localStream = stream
          window.localAudio.srcObject = stream
        })
        .catch((err) => {
          console.error(`you got an error: ${err}`);
        }); */
    },
    getVidAccess() {},
    startDrawing(event) {
      this.drawing = true;
      this.lastX = event.offsetX;
      this.lastY = event.offsetY;
    },
    draw(event) {
      if (!this.drawing) return;
      const canvas = this.$refs.canvas;
      const ctx = canvas.getContext('2d');
      ctx.lineWidth = this.lineWidth;
      ctx.strokeStyle = this.eraser ? 'white' : this.color;
      ctx.beginPath();
      ctx.moveTo(this.lastX, this.lastY);
      ctx.lineTo(event.offsetX, event.offsetY);
      ctx.stroke();
      this.lastX = event.offsetX;
      this.lastY = event.offsetY;
    },
    stopDrawing() {
      this.drawing = false;
      this.undoStack.push(this.$refs.canvas.toDataURL());
    },
    toggleEraser() {
      this.eraser = !this.eraser;
    },
    undo() {
      if (this.undoStack.length > 0) {
        this.redoStack.push(this.undoStack.pop());
        const previousState = this.undoStack[this.undoStack.length - 1];
        const canvas = this.$refs.canvas;
        const ctx = canvas.getContext('2d');
        ctx.clearRect(0, 0, canvas.width, canvas.height);
        const img = new Image();
        img.src = previousState;
        img.onload = () => {
          ctx.drawImage(img, 0, 0);
        };
      }
    },
    redo() {
      if (this.redoStack.length > 0) {
        this.undoStack.push(this.redoStack.pop());
        const canvas = this.$refs.canvas;
        const ctx = canvas.getContext('2d');
        const img = new Image();
        img.src = this.undoStack[this.undoStack.length - 1];
        img.onload = () => {
          ctx.clearRect(0, 0, canvas.width, canvas.height);
          ctx.drawImage(img, 0, 0);
        };
      }
    },
    setColor(color) {
      this.color = color;
    },
    clear() {
      this.undoStack = [];
      this.redoStack = [];
      const canvas = this.$refs.canvas;
      const ctx = canvas.getContext('2d');
      ctx.clearRect(0, 0, canvas.width, canvas.height);
    },
    dataURLtoFile(dataurl, filename) {
      let arr = dataurl.split(','),
        mime = arr[0].match(/:(.*?);/)[1],
        bstr = atob(arr[1]),
        n = bstr.length,
        u8arr = new Uint8Array(n);
      while (n--) {
        u8arr[n] = bstr.charCodeAt(n);
      }
      return new File([u8arr], filename, { type: mime });
    },
    async saveImage() {
      const canvas = this.$refs.canvas;
      const dataUrl = canvas.toDataURL('image/png').replace('image/png', 'image/octet-stream');
      const file = this.dataURLtoFile(dataUrl, 'my-art.png');

      const json = {
        book_id: +this.BookId,
        question_id: +this.QuestionId,
        is_art: true,
        answer: '',
      };

      console.log(json);

      const form = new FormData();
      if (file !== null) {
        form.append('art', file);
      }
      form.append('document', JSON.stringify(json));

      let response = await fetch(`${host}/api/creative/pass/create`, {
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
        onError.value = 'ITCU';
        console.log('NOT OK');

        setTimeout(() => {
          onError.value = null;
        }, 5000);
      } else {
        console.log('OK');
      }
    },

    mounted() {
      const canvas = this.$refs.canvas;
      canvas.width = 1000;
      canvas.height = 500;
    },
  },
};
</script>
<style scoped lang="scss">
.paint {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  flex-wrap: wrap;
}
.paint__canvas {
  /* width: 100%; */
  background-color: white;
  border: 2px solid black;
}
.paint__controls {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  justify-content: center;
  margin-top: 20px;
  gap: 10px;
  height: 100px;
  background-color: #84eee8;
}
.paint__control {
  display: flex;
  align-items: center;
}
.paint__color-picker {
  width: 30px;
  height: 30px;
  border-radius: 50%;
  cursor: pointer;
  margin-right: 10px;
}
.paint__color-picker.active {
  border: 2px solid #000;
}
.btn {
  margin-top: 10px;
}

.paint__color {
  position: fixed;
  left: 50%;
  top: 50%;

  width: 720px;
  height: 380px;
  padding: 24px;
  border-radius: 20px;

  display: grid;
  grid-template-rows: repeat(3, 1fr);
  grid-template-columns: repeat(6, 1fr);
  gap: 20px;

  background-color: var(--black-hover);

  transform: translate(-50%, -50%);

  & .color-box {
    display: flex;
    align-items: center;
    justify-content: center;

    cursor: pointer;
    width: 100%;
    height: 100%;

    border-radius: 14px;
    border: 2px solid var(--white);

    align-self: center;
    justify-self: center;

    &.color-picked {
      border-color: var(--primary);
    }

    & span {
      color: var(--white);
    }
  }
}

.size {
  cursor: pointer;
  border: 3px solid var(--white);
  border-radius: 50%;
  padding: 10px;
}

.size-1 {
  width: 40px;
  height: 40px;
}

.size-3 {
  width: 50px;
  height: 50px;
}

.size-5 {
  width: 60px;
  height: 60px;
}

.size-7 {
  width: 70px;
  height: 70px;
}

.size-10 {
  width: 80px;
  height: 80px;
}

.tool {
  width: 100px;
  height: 100px;
  cursor: pointer;
  padding: 10px;
  background-color: #fff;
  border-radius: 10px;
}
</style>
