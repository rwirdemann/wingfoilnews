/** @type {import('./$types').Actions} */
import {fail, redirect} from "@sveltejs/kit";

export const actions = {

    default: async (event) => {
        const formData = Object.fromEntries(await event.request.formData());

        if (!formData.username || !formData.password) {
            return fail(400, {
                error: 'Missing username or password'
            });
        }

        let user = {
            username: formData.username,
            password: formData.password,
        };

        let body = JSON.stringify(user);
        const res = await fetch("https://news.wingbuddies.de:8087/login", {
            method: "POST",
            body: body,
        });
        if (res.ok) {
            console.log("logout/page.server.js#post => setting cookie")
            const response = await res.json();
            event.cookies.set('AuthorizationToken', `Bearer ${response.token}`, {
                httpOnly: true,
                path: '/',
                secure: false,
                sameSite: 'strict',
                maxAge: 60 * 60 * 24 // 1 day
            });
            throw redirect(302, '/');
        } else {
            console.log("error logging in");
        }
    }
};