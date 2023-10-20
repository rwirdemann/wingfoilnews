/** @type {import('@sveltejs/kit').Handle} */
export async function handle({event, resolve}) {


    const { headers } = event.request;
    const cookies = headers.get("cookie") ?? ""
    if (cookies.AuthorizationToken) {
        const token = cookies.AuthorizationToken.split(" ")[1];
        console.log("Token: " + token)
    }

    const response = await resolve(event);
    return response;
}