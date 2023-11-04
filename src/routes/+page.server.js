export async function load(event) {
    const res = await fetch("https://news.wingbuddies.de:8087/links", {
        method: "GET",
    });
    let user = event.locals.user;
    const links = await res.json();
    return {
        links: links,
        user: user
    }
}