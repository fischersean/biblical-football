<template>
  <div class="page-wrapper has-background-light" id="app">
    <div class="content-wrapper" data-app>
      <h1 class="title mx-6 pb-5">
        I wonder what the man upstairs thinks about these games?
      </h1>
      <Selection />

      <p class="is-italic py-4">
        All verses are chosen based on the home team's score = chapter and the visiting team's score = verse
      </p>
    </div>
<footer class="footer has-background-white py-6">
  <p><a href=https://sean-fischer.mit-license.org target="_blank">MIT</a> - <a href=# target="_blank">Source</a></p> 
  <p><b>Made with ❤️  by <a href=https://github.com/fischersean target="_blank">Sean Fischer </a> </b></p> 
  <p>This is a joke</p>
</footer>
  </div>
</template>

<script>
import Selection from './components/Selection';

export default {
  name: 'App',
  components: {
    Selection,
  },
  data: function() {
    return {
      t: null,
    };
  },
};
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Poppins&display=swap');

#app {
  font-family: 'Poppins', sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  /*color: #2c3e50;*/
  padding-top: 60px;
  /*padding-bottom: 10em;*/
}
.page-wrapper {
	display: flex;
	min-height: 100vh;
	flex-direction: column;
}

.content-wrapper {
	flex: 1;
}
:root {
  /* Color Defs */
  /* CSS HEX */
  --space-cadet: #283044ff;
  --air-superiority-blue: #78a1bbff;
  --wheat: #f5e0b7ff;
  --vegas-gold: #d6ba73ff;
  --cinnabar: #dc493aff;

  /* CSS HSL */
  --space-cadet: hsla(223, 26%, 21%, 1);
  --air-superiority-blue: hsla(203, 33%, 60%, 1);
  --wheat: hsla(40, 76%, 84%, 1);
  --vegas-gold: hsla(43, 55%, 65%, 1);
  --cinnabar: hsla(6, 70%, 55%, 1);

  /* SCSS HEX */
  $space-cadet: #283044ff;
  $air-superiority-blue: #78a1bbff;
  $wheat: #f5e0b7ff;
  $vegas-gold: #d6ba73ff;
  $cinnabar: #dc493aff;

  /* SCSS HSL */
  $space-cadet: hsla(223, 26%, 21%, 1);
  $air-superiority-blue: hsla(203, 33%, 60%, 1);
  $wheat: hsla(40, 76%, 84%, 1);
  $vegas-gold: hsla(43, 55%, 65%, 1);
  $cinnabar: hsla(6, 70%, 55%, 1);

  /* SCSS RGB */
  $space-cadet: rgba(40, 48, 68, 1);
  $air-superiority-blue: rgba(120, 161, 187, 1);
  $wheat: rgba(245, 224, 183, 1);
  $vegas-gold: rgba(214, 186, 115, 1);
  $cinnabar: rgba(220, 73, 58, 1);

  /* SCSS Gradient */
  $gradient-top: linear-gradient(
    0deg,
    #283044ff,
    #78a1bbff,
    #f5e0b7ff,
    #d6ba73ff,
    #dc493aff
  );
  $gradient-right: linear-gradient(
    90deg,
    #283044ff,
    #78a1bbff,
    #f5e0b7ff,
    #d6ba73ff,
    #dc493aff
  );
  $gradient-bottom: linear-gradient(
    180deg,
    #283044ff,
    #78a1bbff,
    #f5e0b7ff,
    #d6ba73ff,
    #dc493aff
  );
  $gradient-left: linear-gradient(
    270deg,
    #283044ff,
    #78a1bbff,
    #f5e0b7ff,
    #d6ba73ff,
    #dc493aff
  );
  $gradient-top-right: linear-gradient(
    45deg,
    #283044ff,
    #78a1bbff,
    #f5e0b7ff,
    #d6ba73ff,
    #dc493aff
  );
  $gradient-bottom-right: linear-gradient(
    135deg,
    #283044ff,
    #78a1bbff,
    #f5e0b7ff,
    #d6ba73ff,
    #dc493aff
  );
  $gradient-top-left: linear-gradient(
    225deg,
    #283044ff,
    #78a1bbff,
    #f5e0b7ff,
    #d6ba73ff,
    #dc493aff
  );
  $gradient-bottom-left: linear-gradient(
    315deg,
    #283044ff,
    #78a1bbff,
    #f5e0b7ff,
    #d6ba73ff,
    #dc493aff
  );
  $gradient-radial: radial-gradient(
    #283044ff,
    #78a1bbff,
    #f5e0b7ff,
    #d6ba73ff,
    #dc493aff
  );
}
</style>
