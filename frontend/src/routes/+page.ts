import type { PageLoad } from './$types';
import type {Sport} from "$lib/types";
import {building} from "$app/environment";

export const prerender = true;
export const load: PageLoad = ({ fetch }) => {
    return {
        sports: building ? Promise.resolve([]) : fetch("/api/sports", {
            headers : {
                'Content-Type': 'application/json',
                'Accept': 'application/json'
            }

        }).then(value => value.json()).then(value => value as Sport[])
    };
};