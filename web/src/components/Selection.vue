<template>
  <div class="container is-widescreen">
    <div>
      <div class="control columns mx-6">
        <dropdown
          class="column"
          id="seasonDropdown"
          v-model="seasonSelect"
          :items="seasonList"
          @input="goJuice"
        ></dropdown>
        <dropdown
          class="column"
          id="weekDropdown"
          v-model="weekSelect"
          :items="weekList"
          @input="goJuice"
        ></dropdown>
        <dropdown
          class="column"
          id="bookDropdown"
          v-model="bookSelect"
          :items="bookOptions"
          @input="goJuice"
        ></dropdown>
      </div>
    </div>
    <div id="versedisplay" class="">
      <div class="my-5" v-for="game in scoreVerses" :key="game.id">
        <VerseDisplay :gameData="game" />
      </div>
    </div>
  </div>
</template>

<script>
import VerseDisplay from './VerseDisplay';
import Dropdown from './Dropdown';
export default {
  name: 'Selection',
  components: {
    VerseDisplay,
    Dropdown,
  },
  data: function() {
    return {
      seasonsMetaData: {},
      /*bookOptions: [],*/
      seasonSelect: '2019',
      weekSelect: 'Week 1',
      bookSelect: 'Numbers',
      scoreVerses: {},
    };
  },
  created: function() {
    // Convert the raw response from the API into something more useful for the frontend.
    // I dont want to touch the API's response format on the backend. That's more effort
    // than it is worth.
    const flattenMeta = function(d) {
      let rDict = {};
      for (const season of d) {
        rDict[season.Season] = (function(w) {
          let wDict = {};

          for (const week of w) {
            wDict[week.Label] = (function(b) {
              let bArr = [];
              for (const book of b) {
                bArr.push(book);
              }

              return bArr;
            })(week.ValidBooks);
          }
          return wDict;
        })(season.Weeks);
      }
      return rDict;
    };
    fetch('https://' + this.apiPath + '/seasonsmeta')
      .then(response => response.json())
      .then(data => {
        this.seasonsMetaData = flattenMeta(data);
      });

    this.goJuice();
  },
  methods: {
    goJuice() {
      let url = new URL('https://' + this.apiPath + '/weekverses'),
        params = {
          book: this.bookSelect,
          season: this.seasonSelect,
          week: this.weekSelect,
        };
      url.search = new URLSearchParams(params).toString();

      fetch(url)
        .then(response => response.json())
        .then(data => {
          this.scoreVerses = data;
          if (!this.bookOptions.includes(this.bookSelect)) {
            this.bookSelect = this.bookOptions[0];
          }
        });
      console.log('Fetching score data');
    },
  },
  computed: {
    seasonList() {
      return Object.keys(this.seasonsMetaData);
    },
    weekList() {
      if (this.seasonList.length != 0) {
        return Object.keys(this.seasonsMetaData[this.seasonSelect]);
      } else {
        return [];
      }
    },
    bookOptions() {
      if (this.seasonList.length != 0) {
        return this.seasonsMetaData[this.seasonSelect][this.weekSelect];
      } else {
        return [];
      }
    },
    apiPath() {
      var base = window.location.host;
      var domain = base.split('www')[1];
      return 'api' + domain;
    },
  },
};
</script>

<style>
.card {
  padding: 1em 0em 1em 0;
}
#versedisplay {
  max-width: 65em;
  margin: auto;
}
</style>
