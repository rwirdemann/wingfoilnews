export async function load(event) {
    const res = await fetch("https://news.wingbuddies.de:8087/links", {
        method: "GET",
    });
    const links = await res.json();
    return links
}