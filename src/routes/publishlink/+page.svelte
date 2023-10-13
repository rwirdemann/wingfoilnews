<script context="module">
    import {goto} from "$app/navigation";

    let title = "";
    let uri = "";
    let tags = "";

    function isValidUrl(uriString) {
        try {
            return Boolean(new URL(uriString));
        } catch (e) {
            return false;
        }
    }

    function toArray(tags) {
        if (!tags) {
            return []
        }

        return tags.split(",").map(t => {
            return t.trim()
        })
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
            tags: toArray(tags)
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

<form>
    <div class="mb-3">
        <label for="title" class="form-label">Title</label>
        <input class="form-control" bind:value={title} id="title" placeholder="Cold Hawaii Games" required autofocus/>
    </div>

    <div class="mb-3">
        <label for="url" class="form-label">URL</label>
        <input class="form-control" bind:value={uri} id="url" placeholder="https://coldhawaiigames.com" required/>
    </div>

    <div class="mb-3">
        <label for="tags" class="form-label">Tags</label>
        <input class="form-control" bind:value={tags} id="tags" placeholder="Event, Test, Report" required/>
    </div>

    <button on:click={doPost} type="submit" class="btn btn-primary">Publish</button>
</form>