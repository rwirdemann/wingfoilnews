export async function load() {
  const res = await fetch("https://news.wingbuddies.de:8087/links", {
    method: "GET",
  });
  return await res.json();
}