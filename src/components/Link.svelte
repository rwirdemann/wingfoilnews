<script>
    import Time from "svelte-time";

    export let title = "";
    export let href = "";
    export let tags = ""
    export let created = new Date();
    export let details = true

    function extractDomain(uri) {
        let url = new URL(uri);
        return url.hostname;
    }
</script>

<li>
    <a {href} target="_blank">{title}</a>
    {#if details}
        <small style="color: gray;">via {extractDomain(href)}</small><br/>
        <small>1 vote by rawi
            <Time relative timestamp={created}/>
        </small>
        {#if tags}
            <small>({tags})</small>
        {/if}
    {/if}
</li>
