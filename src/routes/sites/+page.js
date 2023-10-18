export async function load() {
  const res = await fetch("https://news.wingbuddies.de:8087/links?tags=site", {
    method: "GET",
  });
  return await res.json();
}