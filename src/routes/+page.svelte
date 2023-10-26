<script>
    import Link from "../components/Link.svelte";

    export let data = {
        user: "",
        content: []
    };
    const excluded = "site,group,blog"
</script>

<h1>{data.user}</h1>
<ol>
    {#if data.content.links !== null && data.content.links.length > 0}
        {#each data.content.links as l}
            {#if !excluded.includes(l.tags)}
                <Link title={l.title} href={l.uri} tags={l.tags} created={l.created}/>
            {/if}
        {/each}
    {/if}
</ol>