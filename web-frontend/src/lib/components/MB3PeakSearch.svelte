<script lang="ts">
    import NumberInput from "$component/NumberInput.svelte";
    import Card from "$component/Card.svelte";
    import SearchButton from "$component/SearchButton.svelte";

    export let values: number[] = [];
    export let tolerance = 0.3;
    export let intensity = 100;
    let newValue: number | null = null;

    function addValue() {
        if (newValue !== null) {
            values.push(newValue)
            newValue = null
            values = values
        }
        console.log("addValue")
        console.log(values)
    }

    function deleteValue(i: number) {
       values.splice(i,1)
        values = values
        console.log("deleteValue")
        console.log(values)
    }
</script>

<div class="pure-g">
<div class="pure-u-1-2">
    <Card>
    <div class="headline">Peak m/z</div>
{#each values as val,i}
    <div class="line"><NumberInput label={i+1} bind:value={val}></NumberInput><button on:click={(i) => deleteValue(i)}>Delete</button></div>
{/each}
<div class="line"><NumberInput label="Add to list:" bind:value={newValue} on:change={addValue}></NumberInput></div>
    </Card>
</div>
<block class="pure-u-1-2">
    <Card>
        <div class="headline">Parameters</div>
        <div class="line"><NumberInput label="Mass tolerance" bind:value={tolerance}></NumberInput> </div>
        <div class="line"><NumberInput label="Intensity threshold" bind:value={intensity}></NumberInput> </div>

    </Card>
</block>
</div>
<SearchButton></SearchButton>
<style>
    .line {
        padding: 0.5em;
    }
    Card {
        vertical-align: top;
    }
</style>
