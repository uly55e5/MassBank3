<script>
    export let group;
    export let values;
    export let result= [];
</script>

{#each values as v}
<div class="filteroption">
<label>
    <input type=checkbox bind:group={result} name={group} value={v.value}>
    {v.value} ({v.count})
</label>
</div>
{/each}
<style>
    .filteroption {
        font-size: 0.8em;
        padding: 0 0.5em;
    }
</style>
