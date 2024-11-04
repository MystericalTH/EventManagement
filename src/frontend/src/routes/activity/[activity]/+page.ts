import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch, params }) => {
    return {
        post: await (await fetch(`/api/activities/${params.activity}`)).json()
    };
};