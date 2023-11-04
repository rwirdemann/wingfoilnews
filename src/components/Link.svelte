<script>
    import Time from "svelte-time";
    import {goto} from "$app/navigation";

    export let id = 0;
    export let title = "";
    export let href = "";
    export let tags = ""
    export let created = new Date();
    export let details = true
    export let user = ""

    function extractDomain(uri) {
        let url = new URL(uri);
        return url.hostname;
    }

    function doPatch(id) {
        alert("Publish: " + id)

    }
</script>

<li>
    <a {href} target="_blank">{title}</a>
    {#if details}
        <small style="color: gray;">via {extractDomain(href)}</small>
        {#if user}
            <button on:click={() => doPatch(id)} class="btn btn-outline-primary btn-sm">
                Publish
            </button>
        {/if}
        <br/>
        <small>1 vote by rawi
            <Time relative timestamp={created}/>
        </small>
        {#if tags}
            <small>({tags})</small>
        {/if}
    {/if}
</li>
