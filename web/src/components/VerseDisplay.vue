<template>
  <div id="verse">
    <div class="card">
      <h1 class="title">
        <span>{{ this.homeTeam }} vs {{ this.visitingTeam }}</span>
      </h1>
      <div class="has-text-weight-bold">
        <span>{{ this.homeScore }} - {{ this.visitingScore }}</span>
      </div>
      <div class="is-italic card-content">
        <div>
          <p>
            {{ this.verseData.Text }}
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'VerseDisplay',
  components: {},
  data: function() {
    return {
      homeTeam: '',
      visitingTeam: '',
      homeScore: 0,
      visitingScore: 0,
      gameDate: 0,
      verseData: {},
    };
  },
  props: ['gameData'],
  methods: {},
  created: function() {
    this.homeTeam = this.gameData.Home.split(',')[1].trim();
    this.visitingTeam = this.gameData.Visitor.split(',')[1].trim();
    this.homeScore = this.gameData.HScore;
    this.visitingScore = this.gameData.VScore;
    this.gameDate = this.gameData.Date;
    this.verseData = this.gameData.VerseData;
  },
};
</script>

<style></style>
