export interface TeamMember {
    name: string;
    avatar: string;
}

export interface Workflow {
    automations: number;
    state: 'Running' | 'Idle' | 'Error';
    lastRun: string;
    lastResult: 'Success' | 'Error';
}

export interface Project {
    id: number;
    name: string;
    description: string;
    status: 'Active' | 'Draft' | 'Paused' | 'Completed';
    priority: 'high' | 'normal';
    favorite: boolean;
    workflow: Workflow;
    tasksCompleted: number;
    tasksTotal: number;
    team: TeamMember[];
    updatedAt: string;
}

export interface FilterItem {
    id: string;
    label: string;
    icon: any;
    count?: number;
}
