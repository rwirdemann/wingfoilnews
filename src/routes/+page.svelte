<script>
    import Link from "../components/Link.svelte";

    export let data = {
        links: [],
        user: ""
    }
    const excluded = "site,group,blog"
</script>

<ol>
    {#if data.links.links !== null && data.links.links.length > 0}
        {#each data.links.links as l}
            {#if !excluded.includes(l.tags)}
                <Link id={l.id} title={l.title} href={l.uri} tags={l.tags} created={l.created} user={data.user}/>
            {/if}
        {/each}
    {/if}
</ol>