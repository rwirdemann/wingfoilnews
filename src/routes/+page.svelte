<script>
  // @ts-nocheck

  import Link from "../components/Link.svelte";
  export let data = [];
</script>

<ul>
  {#if data !== null && data.links.length > 0}
    {#each data.links as l}
      <Link title={l.title} href={l.uri} created={new Date()} />
    {/each}
  {/if}
</ul>
