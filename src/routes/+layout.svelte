<script>
    export let data = {
        user: "",
    };

    function doLogout() {
    }
</script>

<div class="container mt-2">
    <header>
        <h1>Wingfoil News 0.2</h1>
        <ul class="nav">
            <li><a class="nav-link" href="/">News</a></li>
            <li><a class="nav-link" href="/publishlink">Publish</a></li>
            <li><a class="nav-link" href="/blogs">Blogs</a></li>
            <li><a class="nav-link" href="/groups">Groups</a></li>
            <li><a class="nav-link" href="/sites">Sites</a></li>
            <li><a class="nav-link" href="/about">About</a></li>
            {#if data.user === undefined}
                <li><a class="nav-link" href="/login">Login</a></li>
            {:else}
                <li><a class="nav-link" on:click={doLogout} href="/logout">Logout ({data.user})</a></li>
            {/if}
        </ul>
        <hr>
    </header>
    <main>
        <slot/>
    </main>

    <footer>
        <hr>
        <p>Wingfoil News powered by wingbuddies.de 2023</p>
    </footer>
</div>