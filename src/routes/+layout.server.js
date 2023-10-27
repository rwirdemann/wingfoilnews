
export async function load(event) {
    console.log("layout.server.js#load")
    let user = event.locals.user;
    console.log("layout.server.js#load user = " + user)

    return {
        user: user
    }
}