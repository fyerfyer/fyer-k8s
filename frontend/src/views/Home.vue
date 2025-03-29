<template>
    <div class="home">
        <div class="dashboard-header">
            <h1>Kubernetes Dashboard</h1>
            <el-button type="primary" @click="pingServer" :loading="pingLoading">
                <el-icon>
                    <connection />
                </el-icon>
                Test Connection
            </el-button>
        </div>

        <el-row :gutter="20">
            <el-col :xs="24" :sm="12" :md="8" :lg="6">
                <el-card class="dashboard-card">
                    <template #header>
                        <div class="card-header">
                            <el-icon class="icon">
                                <cpu />
                            </el-icon>
                            <span>Pods</span>
                        </div>
                    </template>
                    <div class="card-content">
                        <div class="metric">3</div>
                        <el-button type="primary" @click="navigateTo('/pods')">
                            View Pods
                        </el-button>
                    </div>
                </el-card>
            </el-col>

            <el-col :xs="24" :sm="12" :md="8" :lg="6">
                <el-card class="dashboard-card">
                    <template #header>
                        <div class="card-header">
                            <el-icon class="icon">
                                <service />
                            </el-icon>
                            <span>Services</span>
                        </div>
                    </template>
                    <div class="card-content">
                        <div class="metric">2</div>
                        <el-button type="primary" disabled>
                            View Services
                        </el-button>
                    </div>
                </el-card>
            </el-col>

            <el-col :xs="24" :sm="12" :md="8" :lg="6">
                <el-card class="dashboard-card">
                    <template #header>
                        <div class="card-header">
                            <el-icon class="icon">
                                <files />
                            </el-icon>
                            <span>Deployments</span>
                        </div>
                    </template>
                    <div class="card-content">
                        <div class="metric">1</div>
                        <el-button type="primary" disabled>
                            View Deployments
                        </el-button>
                    </div>
                </el-card>
            </el-col>

            <el-col :xs="24" :sm="12" :md="8" :lg="6">
                <el-card class="dashboard-card">
                    <template #header>
                        <div class="card-header">
                            <el-icon class="icon">
                                <lock />
                            </el-icon>
                            <span>Namespaces</span>
                        </div>
                    </template>
                    <div class="card-content">
                        <div class="metric">4</div>
                        <el-button type="primary" disabled>
                            View Namespaces
                        </el-button>
                    </div>
                </el-card>
            </el-col>
        </el-row>
    </div>
</template>

<script>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { Connection, Cpu, Files, Service, Lock } from '@element-plus/icons-vue'
import kubernetesApi from '@/api/kubernetes'

export default {
    name: 'HomeView',
    components: {
        Connection,
        Cpu,
        Files,
        Service,
        Lock
    },
    setup() {
        const router = useRouter()
        const pingLoading = ref(false)

        const navigateTo = (path) => {
            router.push(path)
        }

        const pingServer = async () => {
            pingLoading.value = true
            try {
                const response = await kubernetesApi.pingServer()
                ElMessage({
                    type: 'success',
                    message: `Server response: ${response.message}`,
                    duration: 3000
                })
            } catch (error) {
                ElMessage.error('Failed to ping server')
            } finally {
                pingLoading.value = false
            }
        }

        return {
            navigateTo,
            pingServer,
            pingLoading
        }
    }
}
</script>

<style scoped>
.home {
    padding: 20px;
}

.dashboard-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 30px;
}

.dashboard-card {
    margin-bottom: 20px;
    transition: all 0.3s;
}

.dashboard-card:hover {
    transform: translateY(-5px);
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.card-header {
    display: flex;
    align-items: center;
}

.icon {
    margin-right: 8px;
    font-size: 20px;
}

.card-content {
    display: flex;
    flex-direction: column;
    align-items: center;
}

.metric {
    font-size: 3rem;
    font-weight: bold;
    margin: 20px 0;
    color: #409EFF;
}
</style>