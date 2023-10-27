/** @type {import('@sveltejs/kit').Handle} */
export async function handle({event, resolve}) {
    console.log("hook.server.js#handle")
    const cookies = event.cookies
    let all = cookies.getAll()
    if (cookies && cookies.get('AuthorizationToken')) {
        console.log("hook.server.js#handle => token found")
        const token = cookies.get('AuthorizationToken').split(" ")[1];
        event.locals.user = 'ralf'
    } else {
        console.log("hook.server.js#handle => no token found")
    }

    const response = await resolve(event);
    return response;
}