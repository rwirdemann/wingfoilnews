<script context="module">
    import {goto} from "$app/navigation";

    let username = "";
    let password = "";

    async function doPost() {
        let user = {
            username: username,
            password: password,
        };

        let body = JSON.stringify(user);
        console.log(body);
        const res = await fetch("https://news.wingbuddies.de:8087/login", {
            method: "POST",
            body: body,
        });
        if (res.ok) {
            console.log(res.json())
            goto("/");
        } else {
            console.log("error logging in");
        }
    }
</script>

<form>
    <div class="mb-3">
        <label for="username" class="form-label">Username</label>
        <input class="form-control" bind:value={username} id="username" required autofocus/>
    </div>

    <div class="mb-3">
        <label for="password" class="form-label">Password</label>
        <input type="password" class="form-control" bind:value={password} id="password" required/>
    </div>

    <button on:click={doPost} type="submit" class="btn btn-primary">Login</button>
</form>