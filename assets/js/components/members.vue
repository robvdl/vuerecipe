<template>
<div>
  <h1 class="page-header">{{band.name}}</h1>

  <blockquote>
    {{band.bio}}
  </blockquote>

  <ul class="list-unstyled">
    <li v-for="member in members">
      <h2>
        {{member.name}} - {{member.instrument}}
      </h2>
    </li>
  </ul>

</div>
</template>

<script charset="utf-8">
export default {
  data() {
    return {
      band: {},
      members: {}
    };
  },

  created() {
    this.fetchData();
  },

  watch: {
    $route: "fetchData"
  },

  methods: {
    fetchData() {
      let id = this.$route.params.id;

      // Don't use await on these fetch statements so both can run in parallel.
      // If you use await here then it can't fetch the members until it has
      // finished fetching the band first, we don't want that.
      fetch(`/api/bands/${id}`).then(async response => {
        this.band = await response.json();
      });

      fetch(`/api/bands/${id}/members`).then(async response => {
        this.members = await response.json();
      });
    }
  }
};
</script>
