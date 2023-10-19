/** @type {import('@sveltejs/kit').Handle} */
export async function handle({event, resolve}) {


    const { headers } = event.request;
    const cookies = headers.get("cookie") ?? ""
    if (cookies.AuthorizationToken) {
        const token = cookies.AuthorizationToken.split(" ")[1];
        console.log("Token: " + token)
    } else {
        console.log("Request contains no AuthorizationToken")
    }

    const response = await resolve(event);
    return response;
}