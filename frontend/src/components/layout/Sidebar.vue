<template>
    <el-menu :default-active="activeIndex" class="sidebar-menu" :collapse="isCollapse" background-color="#304156"
        text-color="#bfcbd9" active-text-color="#409EFF" router>
        <el-menu-item index="/" class="menu-item">
            <el-icon>
                <monitor />
            </el-icon>
            <template #title>Dashboard</template>
        </el-menu-item>

        <el-sub-menu index="kubernetes" class="menu-item">
            <template #title>
                <el-icon>
                    <connection />
                </el-icon>
                <span>Kubernetes</span>
            </template>
            <el-menu-item index="/pods" class="submenu-item">
                <el-icon>
                    <cpu />
                </el-icon>
                <span>Pods</span>
            </el-menu-item>
            <!-- More Kubernetes resources can be added here -->
        </el-sub-menu>

        <div class="sidebar-footer">
            <el-tooltip :content="isCollapse ? 'Expand' : 'Collapse'" placement="right">
                <div class="collapse-btn" @click="toggleCollapse">
                    <el-icon v-if="isCollapse">
                        <right />
                    </el-icon>
                    <el-icon v-else>
                        <left />
                    </el-icon>
                </div>
            </el-tooltip>
        </div>
    </el-menu>
</template>

<script>
import { ref, computed } from 'vue'
import { useRoute } from 'vue-router'
import { Monitor, Connection, Cpu, Right, ArrowLeft as Left } from '@element-plus/icons-vue'

export default {
    name: 'Sidebar',
    components: {
        Monitor,
        Connection,
        Cpu,
        Right,
        Left
    },
    setup() {
        const isCollapse = ref(false)
        const route = useRoute()

        const activeIndex = computed(() => {
            return route.path
        })

        const toggleCollapse = () => {
            isCollapse.value = !isCollapse.value
        }

        return {
            isCollapse,
            activeIndex,
            toggleCollapse
        }
    }
}
</script>

<style scoped>
.sidebar-menu {
    height: 100%;
    border-right: none;
}

.sidebar-menu:not(.el-menu--collapse) {
    width: 200px;
}

.menu-item {
    text-align: left;
}

.submenu-item {
    background-color: #1f2d3d !important;
}

.sidebar-footer {
    position: absolute;
    bottom: 0;
    width: 100%;
    background-color: #304156;
    display: flex;
    justify-content: center;
    padding: 10px 0;
}

.collapse-btn {
    color: #bfcbd9;
    cursor: pointer;
    width: 100%;
    text-align: center;
}

.collapse-btn:hover {
    color: #409EFF;
}
</style>