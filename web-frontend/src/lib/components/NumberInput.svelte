<script lang="ts">
    import LabeledLine from "$component/LabeledLine.svelte";

    export let value:number
    export let label:string
</script>

<LabeledLine label={label}>
    <input bind:value type="number" step="0.001" placeholder=100.000 on:change>
</LabeledLine>
