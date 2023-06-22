<script lang="ts">
    import TextInput from "$component/TextInput.svelte";
    import NumberInput from "$component/NumberInput.svelte";
    import SearchButton from "$component/SearchButton.svelte";
    import Card from "$component/Card.svelte";

    export let compoundName: string
    export let formula: string
    export let inchiKey: string
    export let smiles: string
    export let mass: number
    export let splash:string
    export let tolerance = 0.3
</script>

<div class="pure-g">
    <div class="pure-u-1-2">
        <Card>
<div class="line"><TextInput label="Compound name" bind:value={compoundName}></TextInput></div>
<div class="line"><NumberInput label="Exact Mass" bind:value={mass}></NumberInput></div>
<div class="line"><TextInput label="Formula" bind:value={formula}></TextInput></div>
<div class="line"><TextInput label="InChiKey" bind:value={inchiKey}></TextInput></div>
<div class="line"><TextInput label="SMILES" bind:value={smiles}></TextInput></div>
<div class="line"><TextInput label="SPLASH" bind:value={splash}></TextInput></div>
        </Card>
    </div>
    <div class="pure-u-1-2">
        <Card>
            <div>Parameters</div>
            <div class="line"><NumberInput label="Mass tolerance" bind:value={tolerance}></NumberInput></div>

        </Card>
    </div>
</div>
<SearchButton></SearchButton>
<style>
    .line {
        padding: 0.5em;
    }

</style>
