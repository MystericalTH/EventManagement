import type { PageLoad } from './$types';
import type { ActivityData } from '$lib/types/activity';

export const load: PageLoad = async ({ fetch }) => {
    const response = await fetch('/api/activities');
    if (!response.ok) {
        throw new Error(`Failed to load activities: ${response.status} ${response.statusText}`);
    }
    const activities: ActivityData[] = await response.json();
    return { activities };
};
