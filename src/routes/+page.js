export async function load() {
  const res = await fetch("http://news.wingbuddies.de:8087/links", {
    method: "GET",
  });
  return await res.json();
}