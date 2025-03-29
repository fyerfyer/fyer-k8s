<template>
    <div class="header">
        <el-menu class="el-menu-demo" mode="horizontal" :ellipsis="false" background-color="#409EFF" text-color="#fff"
            active-text-color="#ffd04b">
            <el-menu-item index="0">
                <div class="logo-container">
                    <img class="logo" src="../../assets/logo.png" alt="Logo" />
                    <span class="title">Kubernetes Dashboard</span>
                </div>
            </el-menu-item>
            <div class="flex-grow" />
            <el-menu-item index="1">
                <el-dropdown>
                    <span class="el-dropdown-link">
                        <el-icon>
                            <setting />
                        </el-icon>
                        Settings
                        <el-icon class="el-icon--right"><arrow-down /></el-icon>
                    </span>
                    <template #dropdown>
                        <el-dropdown-menu>
                            <el-dropdown-item @click="pingServer">Ping Server</el-dropdown-item>
                            <el-dropdown-item>About</el-dropdown-item>
                        </el-dropdown-menu>
                    </template>
                </el-dropdown>
            </el-menu-item>
        </el-menu>
    </div>
</template>

<script>
import { Setting, ArrowDown } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'
import kubernetesApi from '@/api/kubernetes'

export default {
    name: 'AppHeader',
    components: {
        Setting,
        ArrowDown
    },
    methods: {
        async pingServer() {
            try {
                const response = await kubernetesApi.pingServer()
                ElMessage({
                    type: 'success',
                    message: `Server response: ${response.message}`,
                    duration: 3000
                })
            } catch (error) {
                ElMessage.error('Failed to ping server')
            }
        }
    }
}
</script>

<style scoped>
.header {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    z-index: 1000;
}

.logo-container {
    display: flex;
    align-items: center;
}

.logo {
    height: 32px;
    margin-right: 10px;
}

.title {
    font-size: 1.3rem;
    font-weight: bold;
    color: white;
}

.flex-grow {
    flex-grow: 1;
}

.el-dropdown-link {
    display: flex;
    align-items: center;
    color: white;
    cursor: pointer;
}
</style>