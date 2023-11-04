
export async function load(event) {
    let user = event.locals.user;

    return {
        user: user
    }
}