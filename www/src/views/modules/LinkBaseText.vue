<template>
  <a href="javascript:;" class="user-link" @click="toRightPage">{{ text }}</a>
</template>

<script>
export default {
  name: 'LinkBaseText',
  props: {
    text: {
      required: true,
      type: String,
      default: ''
    },
    type: {
      required: true,
      type: String,
      default: ''
    },
    groupId: {
      type: [String, Number],
      default: ''
    },
    buildingId: {
      type: [String, Number],
      default: ''
    },
    unitId: {
      type: [String, Number],
      default: ''
    }
  },
  methods: {
    toRightPage () {
      const query = {
        name: this.text
      }
      if (this.type === 'project') {
        this.$router.push({ name: 'BaseProject', query })
      } else if (this.type === 'group') {
        this.$router.push({ name: 'BaseGroup', query })
      } else if (this.type === 'building') {
        query.group_id = this.groupId
        this.$router.push({ name: 'BaseBuilding', query })
      } else if (this.type === 'unit') {
        this.$router.push({ name: 'BaseUnitList',
          query,
          params: {
            id: this.buildingId
          }
        })
      } else if (this.type === 'room') {
        query.unit_id = this.unitId
        this.$router.push({ name: 'BaseRoom', query })
      }
    }
  }
}
</script>
