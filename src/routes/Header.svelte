<script>
    export let user = "ralf"

    function doLogout() {
    }
</script>

<header>
    <h1>Wingfoil News 0.2</h1>
    <ul class="nav">
        <li><a class="nav-link" href="/">News</a></li>
        <li><a class="nav-link" href="/publishlink">Publish</a></li>
        <li><a class="nav-link" href="/blogs">Blogs</a></li>
        <li><a class="nav-link" href="/groups">Groups</a></li>
        <li><a class="nav-link" href="/sites">Sites</a></li>
        <li><a class="nav-link" href="/about">About</a></li>
        {#if user !== ''}
            <li><a class="nav-link" on:click={doLogout} href="/">Logout</a></li>
        {:else}
            <li><a class="nav-link" href="/login">Login</a></li>
        {/if}
    </ul>
    <hr>
</header>