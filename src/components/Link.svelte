<script>
  import Time from "svelte-time";
  export let title = "";
  export let href = "";
  export let created = new Date();
</script>

<li>
  <a {href} target="_blank">{title}</a><br />
  <small>1 vote by rawi <Time relative timestamp={created} /></small>
</li>
