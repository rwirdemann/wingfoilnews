<script context="module">
  import { goto } from "$app/navigation";
  let title = "";

  async function doPost() {
    let link = {
      title: "New Link",
      uri: "https://wingbuddies.de",
    };

    let body = JSON.stringify(link);
    console.log(body);
    const res = await fetch("https://news.wingbuddies.de:8087/links", {
      method: "POST",
      body: body,
    });
    if (res.ok) {
      goto("/success");
    } else {
      console.log("error");
    }
  }
</script>

<h3>Publish Link</h3>
<form>
  <label for="title">Title</label>
  <input
    size="40"
    bind:value={title}
    id="title"
    placeholder="Cold Hawaii Games"
  />
  <label for="url">URL</label>
  <input size="40" id="url" placeholder="https://coldhawaiigames.com" />
  <input on:click={doPost} type="submit" value="Publish" />
</form>
