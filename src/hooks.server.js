/** @type {import('@sveltejs/kit').Handle} */
import {JWT_KEY} from '$env/static/private';
import pkg from 'jsonwebtoken';

const {verify} = pkg;

export async function handle({event, resolve}) {
    const cookies = event.cookies
    if (cookies && cookies.get('AuthorizationToken')) {
        const token = cookies.get('AuthorizationToken').split(" ")[1];
        verify(token, JWT_KEY, function (err, decoded) {
            if (err) {
                console.log("hook.server.js#handle err " + err)
                event.locals.user = undefined
            } else {
                event.locals.user = decoded.user
            }
        });
    }

    const response = await resolve(event);
    return response;
}