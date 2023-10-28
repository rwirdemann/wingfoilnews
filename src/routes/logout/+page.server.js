/** @type {import('./$types').Actions} */
import {fail, redirect} from "@sveltejs/kit";

export const actions = {

        default: async (event) => {
            event.cookies.delete('AuthorizationToken',{ path: '/' });
            throw redirect(302, '/');
        }
    }
;