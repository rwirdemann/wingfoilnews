/** @type {import('@sveltejs/kit').Handle} */
export async function handle({event, resolve}) {
    const cookies = event.cookies
    if (cookies && cookies.get('AuthorizationToken')) {
        const token = cookies.get('AuthorizationToken').split(" ")[1];
        console.log("Token: " + token)
    }

    const response = await resolve(event);
    return response;
}