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
    fetchData: function() {
      let id = this.$route.params.id;

      fetch(`/api/bands/${id}`).then(response => {
        response.json().then(data => {
          this.band = data;
        })
      });

      fetch(`/api/bands/${id}/members`).then(response => {
        response.json().then(data => {
          this.members = data;
        })
      });
    }
  }
};
</script>
