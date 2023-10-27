<script>
    import Link from "../components/Link.svelte";

    export let data = []
    const excluded = "site,group,blog"
</script>

<ol>
    {#if data.links !== null && data.links.length > 0}
        {#each data.links as l}
            {#if !excluded.includes(l.tags)}
                <Link title={l.title} href={l.uri} tags={l.tags} created={l.created}/>
            {/if}
        {/each}
    {/if}
</ol>