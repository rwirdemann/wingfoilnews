<script>
    import {jwt} from "../store.js";

    let validJwt = false
    jwt.subscribe((value) => {
        validJwt = value !== ""
    })

    function doLogout() {
        localStorage.clear()
        jwt.set("")
    }
</script>

<header>
    <h1>Wingfoil News</h1>
    <ul class="nav">
        <li><a class="nav-link" href="/">News</a></li>
        <li><a class="nav-link" href="/publishlink">Publish</a></li>
        <li><a class="nav-link" href="/blogs">Blogs</a></li>
        <li><a class="nav-link" href="/groups">Groups</a></li>
        <li><a class="nav-link" href="/sites">Sites</a></li>
        <li><a class="nav-link" href="/about">About</a></li>
        {#if validJwt}
            <li><a class="nav-link" on:click={doLogout} href="/">Logout</a></li>
        {:else}
            <li><a class="nav-link" href="/login">Login</a></li>
        {/if}
    </ul>
    <hr>
</header>