<template>
  <div class="min-h-full">
    <!-- Page Header -->
    <ProjectHeader
      v-model:searchQuery="searchQuery"
      v-model:viewMode="viewMode"
      v-model:sortBy="sortBy"
      :projectCount="filteredProjects.length"
      @new-project="handleNewProject"
    />

    <!-- Main Content Area -->
    <div class="flex gap-6">
      <!-- Sidebar Quick Filters -->
      <ProjectFilters
        v-model:activeFilter="activeFilter"
      />

      <!-- Content Area -->
      <div class="min-w-0 flex-1">
        <!-- Empty State -->
        <ProjectEmptyState 
          v-if="filteredProjects.length === 0" 
          @create="handleNewProject"
        />

        <!-- Grid View -->
        <div v-else-if="viewMode === 'grid'" class="grid gap-6 sm:grid-cols-2 xl:grid-cols-3">
          <ProjectCard
            v-for="project in filteredProjects"
            :key="project.id"
            :project="project"
            @open="handleOpen"
            @settings="handleSettings"
            @duplicate="handleDuplicate"
            @archive="handleArchive"
            @toggle-favorite="handleToggleFavorite"
          />
        </div>

        <!-- List View -->
        <div v-else class="overflow-hidden rounded-xl bg-white shadow-sm ring-1 ring-slate-100">
          <!-- Table Header -->
          <div class="border-b border-slate-100 bg-slate-50/50 px-6 py-4">
            <div class="grid grid-cols-12 gap-4 text-xs font-semibold uppercase tracking-wider text-slate-500">
              <div class="col-span-4">Project Name</div>
              <div class="col-span-1">Status</div>
              <div class="col-span-2">Workflow State</div>
              <div class="col-span-2">Progress</div>
              <div class="col-span-2">Last Activity</div>
              <div class="col-span-1 text-right">Actions</div>
            </div>
          </div>

          <!-- Table Body -->
          <ul role="list" class="divide-y divide-slate-100">
            <ProjectListRow
              v-for="project in filteredProjects"
              :key="project.id"
              :project="project"
              @open="handleOpen"
              @settings="handleSettings"
              @duplicate="handleDuplicate"
              @archive="handleArchive"
              @toggle-favorite="handleToggleFavorite"
            />
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';
import type { Project } from '~/components/app/projects/types';
import ProjectHeader from '~/components/app/projects/ProjectHeader.vue';
import ProjectFilters from '~/components/app/projects/ProjectFilters.vue';
import ProjectCard from '~/components/app/projects/ProjectCard.vue';
import ProjectListRow from '~/components/app/projects/ProjectListRow.vue';
import ProjectEmptyState from '~/components/app/projects/ProjectEmptyState.vue';

definePageMeta({
  middleware: 'auth',
  layout: 'app'
});

// State
const searchQuery = ref('');
const viewMode = ref<'grid' | 'list'>('grid');
const activeFilter = ref('all');
const sortBy = ref('updated');

// Mock Data
const projects = ref<Project[]>([
  {
    id: 1,
    name: 'Website Redesign',
    description: 'Complete overhaul of the company website with modern design patterns and improved UX.',
    status: 'Active',
    priority: 'high',
    favorite: true,
    workflow: {
      automations: 3,
      state: 'Running',
      lastRun: '2 min ago',
      lastResult: 'Success',
    },
    tasksCompleted: 12,
    tasksTotal: 30,
    team: [
      { name: 'Sarah J.', avatar: 'https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80' },
      { name: 'Mike T.', avatar: 'https://images.unsplash.com/photo-1500648767791-00dcc994a43e?ixlib=rb-1.2.1&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80' },
      { name: 'Alex M.', avatar: 'https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80' },
    ],
    updatedAt: '2h ago',
  },
  {
    id: 2,
    name: 'Mobile App Launch',
    description: 'Cross-platform mobile application for customer engagement and loyalty programs.',
    status: 'Active',
    priority: 'high',
    favorite: false,
    workflow: {
      automations: 5,
      state: 'Error',
      lastRun: '15 min ago',
      lastResult: 'Error',
    },
    tasksCompleted: 45,
    tasksTotal: 60,
    team: [
      { name: 'Emily R.', avatar: 'https://images.unsplash.com/photo-1438761681033-6461ffad8d80?ixlib=rb-1.2.1&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80' },
      { name: 'David K.', avatar: 'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?ixlib=rb-1.2.1&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80' },
    ],
    updatedAt: '30m ago',
  },
  {
    id: 3,
    name: 'Internal Audit System',
    description: 'Automated compliance and audit tracking system with real-time reporting dashboards.',
    status: 'Active',
    priority: 'normal',
    favorite: true,
    workflow: {
      automations: 2,
      state: 'Idle',
      lastRun: '1 day ago',
      lastResult: 'Success',
    },
    tasksCompleted: 8,
    tasksTotal: 25,
    team: [
      { name: 'Lisa N.', avatar: 'https://images.unsplash.com/photo-1517841905240-472988babdf9?ixlib=rb-1.2.1&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80' },
      { name: 'James W.', avatar: 'https://images.unsplash.com/photo-1519345182560-3f2917c472ef?ixlib=rb-1.2.1&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80' },
      { name: 'Anna B.', avatar: 'https://images.unsplash.com/photo-1534528741775-53994a69daeb?ixlib=rb-1.2.1&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80' },
      { name: 'Tom H.', avatar: 'https://images.unsplash.com/photo-1506794778202-cad84cf45f1d?ixlib=rb-1.2.1&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80' },
      { name: 'Chris P.', avatar: 'https://images.unsplash.com/photo-1500648767791-00dcc994a43e?ixlib=rb-1.2.1&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80' },
    ],
    updatedAt: '1d ago',
  },
  {
    id: 4,
    name: 'Marketing Campaign Q1',
    description: 'Multi-channel marketing campaign targeting enterprise customers in the technology sector.',
    status: 'Draft',
    priority: 'normal',
    favorite: false,
    workflow: {
      automations: 0,
      state: 'Idle',
      lastRun: 'Never',
      lastResult: 'Success',
    },
    tasksCompleted: 0,
    tasksTotal: 15,
    team: [
      { name: 'Rachel G.', avatar: 'https://images.unsplash.com/photo-1487412720507-e7ab37603c6f?ixlib=rb-1.2.1&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80' },
    ],
    updatedAt: '3d ago',
  },
  {
    id: 5,
    name: 'API Integration Platform',
    description: 'Unified API gateway for third-party service integrations with monitoring and analytics.',
    status: 'Active',
    priority: 'normal',
    favorite: true,
    workflow: {
      automations: 8,
      state: 'Running',
      lastRun: '5 min ago',
      lastResult: 'Success',
    },
    tasksCompleted: 32,
    tasksTotal: 40,
    team: [
      { name: 'Mike T.', avatar: 'https://images.unsplash.com/photo-1500648767791-00dcc994a43e?ixlib=rb-1.2.1&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80' },
      { name: 'Sarah J.', avatar: 'https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80' },
    ],
    updatedAt: '5m ago',
  },
  {
    id: 6,
    name: 'Customer Portal v2',
    description: 'Next generation self-service portal with AI-powered support and personalized dashboards.',
    status: 'Paused',
    priority: 'normal',
    favorite: false,
    workflow: {
      automations: 4,
      state: 'Idle',
      lastRun: '1 week ago',
      lastResult: 'Success',
    },
    tasksCompleted: 18,
    tasksTotal: 50,
    team: [
      { name: 'Alex M.', avatar: 'https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80' },
      { name: 'Lisa N.', avatar: 'https://images.unsplash.com/photo-1517841905240-472988babdf9?ixlib=rb-1.2.1&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80' },
      { name: 'David K.', avatar: 'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?ixlib=rb-1.2.1&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80' },
    ],
    updatedAt: '1w ago',
  },
]);

// Computed
const filteredProjects = computed(() => {
  let result = projects.value;

  // Filter by search query
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase();
    result = result.filter(p => 
      p.name.toLowerCase().includes(query) || 
      p.description.toLowerCase().includes(query)
    );
  }

  // Filter by active filter
  if (activeFilter.value === 'active') {
    result = result.filter(p => p.status === 'Active');
  } else if (activeFilter.value === 'draft') {
    result = result.filter(p => p.status === 'Draft');
  } else if (activeFilter.value === 'archived') {
    result = result.filter(p => p.status === 'Completed');
  } else if (activeFilter.value === 'favorites') {
    result = result.filter(p => p.favorite);
  }

  return result;
});

// Event handlers
function handleNewProject() {
  console.log('Create new project');
}

function handleOpen(id: number) {
  console.log('Open project', id);
}

function handleSettings(id: number) {
  console.log('Open settings for project', id);
}

function handleDuplicate(id: number) {
  console.log('Duplicate project', id);
}

function handleArchive(id: number) {
  console.log('Archive project', id);
}

function handleToggleFavorite(id: number) {
  const project = projects.value.find(p => p.id === id);
  if (project) {
    project.favorite = !project.favorite;
  }
}
</script>
