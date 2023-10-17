export async function load() {
  const res = await fetch("https://news.wingbuddies.de:8087/links?tags=Blog", {
    method: "GET",
  });
  return await res.json();
}