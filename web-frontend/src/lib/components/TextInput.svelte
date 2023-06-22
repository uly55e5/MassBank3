<script lang="ts">
    import LabeledLine from "$component/LabeledLine.svelte";

    export let value:string
    export let label:string
</script>

<LabeledLine label={label}>
<input bind:value size="30" placeholder={label} on:change>
</LabeledLine>
