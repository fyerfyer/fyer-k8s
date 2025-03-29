<template>
    <div class="pod-list">
        <div class="table-header">
            <h2>Kubernetes Pods</h2>
            <div class="table-actions">
                <el-button type="primary" @click="fetchPods" :loading="loading" size="small">
                    <el-icon>
                        <refresh />
                    </el-icon>
                    Refresh
                </el-button>
            </div>
        </div>

        <el-table v-loading="loading" :data="pods" style="width: 100%" border>
            <el-table-column prop="name" label="Name" min-width="180" />
            <el-table-column prop="namespace" label="Namespace" min-width="120" />
            <el-table-column prop="status" label="Status" min-width="100">
                <template #default="scope">
                    <el-tag :type="getPodStatusType(scope.row.status)">
                        {{ scope.row.status }}
                    </el-tag>
                </template>
            </el-table-column>
            <el-table-column label="Actions" width="180">
                <template #default="scope">
                    <el-button type="primary" size="small" @click="viewPodDetails(scope.row)">
                        View
                    </el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { Refresh } from '@element-plus/icons-vue'
import kubernetesApi from '@/api/kubernetes'

export default {
    name: 'PodList',
    components: {
        Refresh
    },
    setup() {
        const pods = ref([])
        const loading = ref(false)

        const fetchPods = async () => {
            loading.value = true
            try {
                const response = await kubernetesApi.listPods()
                // Mocking pod data since the API currently only returns a message
                // In the future, you'll parse the actual pod data from the API response
                pods.value = [
                    { name: 'nginx-pod', namespace: 'default', status: 'Running' },
                    { name: 'redis-pod', namespace: 'kube-system', status: 'Pending' },
                    { name: 'mongodb-pod', namespace: 'database', status: 'Failed' }
                ]
                ElMessage({
                    type: 'success',
                    message: response.message || 'Pods loaded successfully',
                    duration: 2000
                })
            } catch (error) {
                ElMessage.error('Failed to fetch pods')
                console.error(error)
            } finally {
                loading.value = false
            }
        }

        const viewPodDetails = (pod) => {
            ElMessage({
                message: `Viewing details for pod: ${pod.name}`,
                type: 'info'
            })
            // Future implementation would navigate to a pod details page
        }

        const getPodStatusType = (status) => {
            const types = {
                'Running': 'success',
                'Pending': 'warning',
                'Failed': 'danger',
                'Succeeded': 'info'
            }
            return types[status] || 'info'
        }

        // Fetch pods on component mount
        fetchPods()

        return {
            pods,
            loading,
            fetchPods,
            viewPodDetails,
            getPodStatusType
        }
    }
}
</script>

<style scoped>
.pod-list {
    padding: 20px;
}

.table-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;
}

.table-actions {
    display: flex;
    gap: 10px;
}
</style>