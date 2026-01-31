<template>
  <div class="space-y-8">
    <!-- Header -->
    <header class="flex items-end justify-between">
      <div>
        <h1 class="text-3xl font-bold tracking-tight text-q-heading">Dashboard</h1>
        <p class="mt-2 text-sm text-slate-500">Project overview and current status</p>
      </div>
      <div>
        <button class="flex items-center gap-2 rounded-lg bg-q-primary px-4 py-2 text-sm font-medium text-white shadow-sm transition-colors hover:bg-q-primary-hover focus:outline-none focus:ring-2 focus:ring-q-primary focus:ring-offset-2">
            <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2.5">
                <path stroke-linecap="round" stroke-linejoin="round" d="M12 4v16m8-8H4" />
            </svg>
            New Project
        </button>
      </div>
    </header>

    <!-- KPI Cards -->
    <div class="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-4">
      <div v-for="kpi in kpis" :key="kpi.label" class="overflow-hidden rounded-xl bg-white p-6 shadow-sm ring-1 ring-slate-100 transition-shadow hover:shadow-md">
        <div class="flex items-center gap-4">
          <div :class="['flex h-12 w-12 items-center justify-center rounded-lg', kpi.bgClass, kpi.textClass]">
            <component :is="kpi.icon" class="h-6 w-6" />
          </div>
          <div>
            <p class="text-sm font-medium text-slate-500">{{ kpi.label }}</p>
            <p class="mt-1 text-2xl font-bold text-q-heading">{{ kpi.value }}</p>
          </div>
        </div>
        <div class="mt-4 flex items-center gap-2">
            <span :class="['text-xs font-medium', kpi.trend > 0 ? 'text-emerald-600' : 'text-rose-600']">
                 {{ kpi.trend > 0 ? '+' : ''}}{{ kpi.trend }}%
            </span>
            <span class="text-xs text-slate-400">from last month</span>
        </div>
      </div>
    </div>

    <!-- Main Content Grid -->
    <div class="grid grid-cols-1 gap-8 lg:grid-cols-3">
      <!-- Left Column: Project Health Overview (Updates list) -->
      <div class="lg:col-span-2 space-y-6">
        <div class="flex items-center justify-between">
            <h2 class="text-lg font-bold text-q-heading">Project Health Overview</h2>
            <a href="#" class="text-sm font-medium text-q-primary hover:text-q-primary-hover">View all projects</a>
        </div>
        
        <div class="overflow-hidden rounded-xl bg-white shadow-sm ring-1 ring-slate-100">
            <div class="border-b border-slate-100 bg-slate-50/50 px-6 py-4">
                <div class="grid grid-cols-12 gap-4 text-xs font-semibold uppercase tracking-wider text-slate-500">
                    <div class="col-span-5">Project Name</div>
                    <div class="col-span-3">Leader</div>
                    <div class="col-span-2">Deadline</div>
                    <div class="col-span-2 text-right">Status</div>
                </div>
            </div>
            <ul role="list" class="divide-y divide-slate-100">
                <li v-for="project in projects" :key="project.id" class="px-6 py-4 transition-colors hover:bg-slate-50/50">
                    <div class="grid grid-cols-12 gap-4 items-center">
                        <div class="col-span-5">
                            <p class="text-sm font-medium text-q-heading">{{ project.name }}</p>
                            <p class="text-xs text-slate-400 mt-0.5">{{ project.client }}</p>
                        </div>
                        <div class="col-span-3">
                             <div class="flex items-center gap-2">
                                <img :src="project.leader.avatar" class="h-6 w-6 rounded-full object-cover" alt="" />
                                <span class="text-sm text-slate-600">{{ project.leader.name }}</span>
                             </div>
                        </div>
                        <div class="col-span-2">
                            <span class="text-sm text-slate-500">{{ project.deadline }}</span>
                        </div>
                        <div class="col-span-2 text-right">
                            <span :class="['inline-flex items-center rounded-full px-2.5 py-0.5 text-xs font-medium', statusStyles[project.status]]">
                                {{ project.status }}
                            </span>
                        </div>
                    </div>
                </li>
            </ul>
        </div>
      </div>

      <!-- Right Column: Risk Summary -->
      <div class="space-y-6">
        <div class="flex items-center justify-between">
            <h2 class="text-lg font-bold text-q-heading">Risk Summary</h2>
            <a href="#" class="text-sm font-medium text-q-primary hover:text-q-primary-hover">View all risks</a>
        </div>

        <div class="rounded-xl bg-white p-6 shadow-sm ring-1 ring-slate-100">
            <div class="space-y-6">
                 <div v-for="risk in risks" :key="risk.id" class="relative pl-4">
                    <!-- Severity Indicator Line -->
                    <div :class="['absolute left-0 top-1 h-full w-1 rounded-full', risk.severityColor]"></div>
                    
                    <div class="flex justify-between items-start">
                        <div>
                             <h3 class="text-sm font-medium text-q-heading">{{ risk.title }}</h3>
                             <p class="mt-1 text-xs text-slate-500 line-clamp-2">{{ risk.description }}</p>
                        </div>
                    </div>
                     <div class="mt-3 flex items-center gap-3">
                        <span :class="['inline-flex items-center rounded-md px-2 py-1 text-xs font-medium', risk.badgeBg, risk.badgeText]">
                            {{ risk.severity }} Severity
                        </span>
                        <span class="text-xs text-slate-400">{{ risk.project }}</span>
                    </div>
                </div>
            </div>
            
            <button class="mt-6 w-full rounded-lg border border-slate-200 py-2.5 text-sm font-medium text-slate-600 hover:bg-slate-50 hover:text-slate-900 transition-colors">
                Run Risk Analysis
            </button>
        </div>

        <!-- Mini Widget: Recent Activity or Messages? Let's do a simple summary widget -->
         <div class="rounded-xl bg-q-primary/5 p-6 ring-1 ring-q-primary/10">
            <h3 class="text-base font-semibold text-q-heading">Weekly Report</h3>
            <p class="mt-2 text-xs text-slate-600">Your team has completed 12 tasks this week. Keep up the momentum!</p>
            <div class="mt-4 h-1.5 w-full rounded-full bg-slate-200">
                <div class="h-1.5 w-3/4 rounded-full bg-q-primary"></div>
            </div>
             <p class="mt-2 text-right text-xs font-medium text-q-primary">75% Goal Reached</p>
         </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { h } from 'vue';

definePageMeta({
    middleware: 'auth',
    layout: 'app'
});

// Mock Icons
const IconProjects = h('svg', { fill: 'none', viewBox: '0 0 24 24', stroke: 'currentColor', 'stroke-width': 2 }, [
    h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', d: 'M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10' })
])
const IconClock = h('svg', { fill: 'none', viewBox: '0 0 24 24', stroke: 'currentColor', 'stroke-width': 2 }, [
    h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', d: 'M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z' })
])
const IconCurrency = h('svg', { fill: 'none', viewBox: '0 0 24 24', stroke: 'currentColor', 'stroke-width': 2 }, [
    h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', d: 'M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z' })
])
const IconAlert = h('svg', { fill: 'none', viewBox: '0 0 24 24', stroke: 'currentColor', 'stroke-width': 2 }, [
    h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', d: 'M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z' })
])


const kpis = [
  { label: 'Active Projects', value: '12', trend: 8, icon: IconProjects, bgClass: 'bg-blue-50', textClass: 'text-blue-600' },
  { label: 'Upcoming Deadlines', value: '5', trend: -2, icon: IconClock, bgClass: 'bg-amber-50', textClass: 'text-amber-600' },
  { label: 'Budget Usage', value: '$42.5k', trend: 12, icon: IconCurrency, bgClass: 'bg-emerald-50', textClass: 'text-emerald-600' },
  { label: 'Open Risks', value: '3', trend: -5, icon: IconAlert, bgClass: 'bg-rose-50', textClass: 'text-rose-600' },
];

const statusStyles: Record<string, string> = {
    'On Track': 'bg-emerald-100 text-emerald-800',
    'At Risk': 'bg-amber-100 text-amber-800',
    'Delayed': 'bg-rose-100 text-rose-800',
};

const projects = [
    { 
        id: 1, 
        name: 'Website Redesign', 
        client: 'Acme Corp',
        leader: { name: 'Sarah J.', avatar: 'https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80' },
        deadline: 'Nov 14, 2026',
        status: 'On Track'
    },
    { 
        id: 2, 
        name: 'Mobile App Launch', 
        client: 'Globex',
        leader: { name: 'Mike T.', avatar: 'https://images.unsplash.com/photo-1500648767791-00dcc994a43e?ixlib=rb-1.2.1&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80' },
        deadline: 'Dec 02, 2026',
        status: 'At Risk'
    },
    { 
        id: 3, 
        name: 'Internal Audit System', 
        client: 'Quanta Inc',
        leader: { name: 'Alex M.', avatar: 'https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80' },
        deadline: 'Jan 20, 2027',
        status: 'On Track'
    },
    { 
        id: 4, 
        name: 'Marketing Campaign', 
        client: 'Umbrella Corp',
        leader: { name: 'Emily R.', avatar: 'https://images.unsplash.com/photo-1438761681033-6461ffad8d80?ixlib=rb-1.2.1&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80' },
        deadline: 'Oct 30, 2026',
        status: 'Delayed'
    },
];

const risks = [
    {
        id: 1,
        title: 'Budget Overrun',
        description: 'Marketing campaign spending is projected to exceed Q4 allocation by 15% due to ad costs.',
        severity: 'High',
        badgeBg: 'bg-rose-100',
        badgeText: 'text-rose-800',
        severityColor: 'bg-rose-500',
        project: 'Marketing Campaign'
    },
    {
        id: 2,
        title: 'Resource Shortage',
        description: 'Lead developer is out on medical leave for 2 weeks during critical sprint phase.',
        severity: 'Medium',
        badgeBg: 'bg-amber-100',
        badgeText: 'text-amber-800',
        severityColor: 'bg-amber-500',
        project: 'Mobile App'
    },
    {
        id: 3,
        title: 'API Rate Limits',
        description: 'Third-party integration is hitting legacy API limits during peak hours.',
        severity: 'Low',
        badgeBg: 'bg-blue-100',
        badgeText: 'text-blue-800',
        severityColor: 'bg-blue-500',
        project: 'Internal Audit'
    }
];

</script>
