export async function load(event) {
  console.log("EVENT: " + event.data.user)

  const res = await fetch("https://news.wingbuddies.de:8087/links", {
    method: "GET",
  });
  const links = await res.json();
  return {
    user: event.data.user,
    content: links
  }
}