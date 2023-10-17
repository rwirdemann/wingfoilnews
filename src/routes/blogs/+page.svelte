<script>
  import Link from "../../components/Link.svelte";
  export let data = [];
</script>

<ol>
  {#if data.links !== null && data.links.length > 0}
    {#each data.links as l}
      <Link title={l.title} href={l.uri} tags={l.tags} created={l.created} />
    {/each}
  {/if}
</ol>
