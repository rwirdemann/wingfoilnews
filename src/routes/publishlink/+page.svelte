<script context="module">
  import { goto } from "$app/navigation";
  let title = "";
  let uri = "";

  function isValidUrl(uriString) {
    try {
      return Boolean(new URL(uriString));
    } catch (e) {
      return false;
    }
  }

  async function doPost() {
    if (title.length == 0 || uri.length == 0) {
      return;
    }

    if (!isValidUrl(uri)) {
      alert("URL has invalid format")
      return;
    }

    let link = {
      title: title,
      uri: uri,
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
      console.log("error posting new link");
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
    required
  />
  <label for="url">URL</label>
  <input
    size="40"
    id="uri"
    bind:value={uri}
    placeholder="https://coldhawaiigames.com"
    required
  />
  <input on:click={doPost} type="submit" value="Publish" />
</form>
