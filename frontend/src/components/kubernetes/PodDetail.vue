<template>
    <div class="pod-detail" v-loading="loading">
        <el-page-header @back="goBack" :title="pod.name || 'Pod Details'">
            <template #content>
                <span class="page-header-title">{{ pod.name }}</span>
                <el-tag class="ml-10" :type="getPodStatusType(pod.status)" v-if="pod.status">
                    {{ pod.status }}
                </el-tag>
            </template>
        </el-page-header>

        <div class="pod-content" v-if="!loading && pod.name">
            <!-- Basic Information Card -->
            <el-card class="detail-card">
                <template #header>
                    <div class="card-header">
                        <span>Basic Information</span>
                    </div>
                </template>
                <el-descriptions :column="2" border>
                    <el-descriptions-item label="Name">{{ pod.name }}</el-descriptions-item>
                    <el-descriptions-item label="Namespace">{{ pod.namespace }}</el-descriptions-item>
                    <el-descriptions-item label="Status">
                        <el-tag :type="getPodStatusType(pod.status)">{{ pod.status }}</el-tag>
                    </el-descriptions-item>
                    <el-descriptions-item label="IP Address">{{ pod.podIP || 'N/A' }}</el-descriptions-item>
                    <el-descriptions-item label="Node">{{ pod.nodeName || 'N/A' }}</el-descriptions-item>
                    <el-descriptions-item label="Creation Time">{{ formatDate(pod.creationTime) }}</el-descriptions-item>
                </el-descriptions>
            </el-card>

            <!-- Labels and Annotations -->
            <el-card class="detail-card">
                <template #header>
                    <div class="card-header">
                        <span>Labels</span>
                    </div>
                </template>
                <div v-if="Object.keys(pod.labels || {}).length === 0" class="no-data">No Labels</div>
                <div v-else class="labels-container">
                    <el-tag 
                        v-for="(value, key) in pod.labels" 
                        :key="key"
                        class="label-tag"
                    >
                        {{ key }}: {{ value }}
                    </el-tag>
                </div>
            </el-card>

            <!-- Container Information -->
            <el-card class="detail-card">
                <template #header>
                    <div class="card-header">
                        <span>Containers</span>
                    </div>
                </template>
                <div v-if="!pod.containers || pod.containers.length === 0" class="no-data">No Container Information</div>
                <el-table v-else :data="pod.containers" style="width: 100%">
                    <el-table-column prop="name" label="Name" min-width="120" />
                    <el-table-column prop="image" label="Image" min-width="200" />
                    <el-table-column prop="state" label="State" min-width="100">
                        <template #default="scope">
                            <el-tag :type="getContainerStatusType(scope.row.state)">
                                {{ scope.row.state || 'Unknown' }}
                            </el-tag>
                        </template>
                    </el-table-column>
                    <el-table-column prop="ready" label="Ready" min-width="80">
                        <template #default="scope">
                            <el-tag type="success" v-if="scope.row.ready">Yes</el-tag>
                            <el-tag type="danger" v-else>No</el-tag>
                        </template>
                    </el-table-column>
                    <el-table-column prop="restartCount" label="Restart Count" min-width="100" />
                </el-table>
            </el-card>
        </div>

        <el-empty v-if="!loading && !pod.name" description="Pod Information Not Found" />
    </div>
</template>

<script>
import { ref, onMounted, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import kubernetesApi from '@/api/kubernetes'

export default {
    name: 'PodDetail',
    props: {
        namespace: {
            type: String,
            default: ''
        },
        name: {
            type: String,
            default: ''
        }
    },
    setup(props) {
        const route = useRoute()
        const router = useRouter()
        const pod = ref({})
        const loading = ref(true)

        const fetchPodDetail = async (namespace, name) => {
            if (!namespace || !name) {
                loading.value = false
                return
            }

            loading.value = true
            try {
                const podData = await kubernetesApi.getPodDetail(namespace, name)
                pod.value = podData || {}
            } catch (error) {
                console.error('Failed to fetch pod details:', error)
                ElMessage.error(`Failed to fetch pod details: ${error.message || 'Unknown error'}`)
                pod.value = {}
            } finally {
                loading.value = false
            }
        }

        // Navigate back to the previous page
        const goBack = () => {
            router.go(-1)
        }

        // Return different tag types based on Pod status
        const getPodStatusType = (status) => {
            const types = {
                'Running': 'success',
                'Pending': 'warning',
                'Failed': 'danger',
                'Succeeded': 'info',
                'Unknown': 'info'
            }
            return types[status] || 'info'
        }

        // Return different tag types based on container state
        const getContainerStatusType = (state) => {
            const types = {
                'Running': 'success',
                'Waiting': 'warning',
                'Terminated': 'info'
            }
            return types[state] || 'info'
        }

        // Format date
        const formatDate = (dateString) => {
            if (!dateString) return '-'
            const date = new Date(dateString)
            return date.toLocaleString()
        }

        // Get namespace and Pod name from route or props
        const getParams = () => {
            const namespace = props.namespace || route.params.namespace
            const name = props.name || route.params.name
            return { namespace, name }
        }

        // Load data when the component is mounted or parameters change
        onMounted(() => {
            const { namespace, name } = getParams()
            fetchPodDetail(namespace, name)
        })

        // Watch for route parameter changes
        watch(() => [route.params.namespace, route.params.name], ([newNamespace, newName]) => {
            if (newNamespace && newName) {
                fetchPodDetail(newNamespace, newName)
            }
        })

        // Watch for props changes
        watch(() => [props.namespace, props.name], ([newNamespace, newName]) => {
            if (newNamespace && newName) {
                fetchPodDetail(newNamespace, newName)
            }
        })

        return {
            pod,
            loading,
            goBack,
            getPodStatusType,
            getContainerStatusType,
            formatDate
        }
    }
}
</script>

<style scoped>
.pod-detail {
    padding: 20px;
    background-color: white;
    border-radius: 4px;
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.page-header-title {
    font-size: 18px;
    font-weight: bold;
    margin-right: 10px;
}

.pod-content {
    margin-top: 20px;
}

.detail-card {
    margin-bottom: 20px;
}

.card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.no-data {
    color: #909399;
    font-style: italic;
    text-align: center;
    padding: 20px;
}

.labels-container {
    display: flex;
    flex-wrap: wrap;
    gap: 10px;
}

.label-tag {
    margin-right: 0;
}

.ml-10 {
    margin-left: 10px;
}
</style>