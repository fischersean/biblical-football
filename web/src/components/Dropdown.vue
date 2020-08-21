<template>
  <div >
    <label class="label" :for="name"></label>
    <div class="select is-rounded">
      <select v-model="computedValue" :name="name" :id="name">
        <option v-for="item in items" :key="item.id"> {{ item }}</option>
      </select>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Dropdown',
  components: {},
  data: function() {
    return {
      selected: this.value,
    };
  },
  props: ['items', 'name', 'value'],
  methods: {},
  computed: {
    computedValue: {
      get() {
        return this.selected;
      },
      set(value) {
        this.selected = value;
        this.$emit('input', value);
      },
    },
  },
};
</script>
