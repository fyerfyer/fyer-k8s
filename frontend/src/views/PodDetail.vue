<template>
    <div class="pod-detail-view">
        <pod-detail :namespace="$route.params.namespace" :name="$route.params.name" />
    </div>
</template>

<script>
import PodDetail from '@/components/kubernetes/PodDetail.vue'

export default {
    name: 'PodDetailView',
    components: {
        PodDetail
    }
}
</script>

<style scoped>
.pod-detail-view {
    width: 100%;
    height: 100%;
}
</style>