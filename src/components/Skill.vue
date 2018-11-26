<template>
  <span :class="'badge badge-pill badge-info font-weight-normal ' + this.cclass">{{ cskill }}</span>
</template>

<script>
export default {
  name: 'skill',
  props: ['cskill', 'cclass']
}
</script>

<style>
span {
  flex-shrink: 1;
  flex-grow: 1;
}


.badge {
  padding: 8px;
}

span.badge {
  font-size: 1em;
  background-color: rgba(52, 73, 94, 0.6);
  border-radius: 5px;
  margin: 5px;
}

span.badge.greenish {
  background-color: rgba(77, 175, 124, 0.8);
}

span.badge.greyish {
  background-color: rgba(149, 165, 166, 0.8);
}
</style>