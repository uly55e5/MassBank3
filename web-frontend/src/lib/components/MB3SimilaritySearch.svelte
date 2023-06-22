<script lang="ts">
    import LabeledLine from "$component/LabeledLine.svelte";
    import Card from "$component/Card.svelte";
    import NumberInput from "$component/NumberInput.svelte";
    import SearchButton from "$component/SearchButton.svelte";
    let inputMethods=["Paste list","Manual input"]
    let input = 0
    interface mzInt{
        mz: number | null
        intensity: number | null
    }
    let values: mzInt[] = []
    function deleteValue(i:number) {

    }

    function addValue(mz:number,int:number) {

    }

    let newValue:mzInt =  {
       mz: null,
        intensity: null

    }
    let intThreshold: number

</script>

Input Method:
{#each inputMethods as method,i}
    <div class="select">
   <input type="radio" bind:group={input} name="input" value={i}> {method}
    </div>
{/each}

<Card>
{#if input===0}
    <textarea></textarea>
    {:else }
    <div class="pure-g">
        <div class="pure-u-1-2">
            <Card>
                <div class="headline">Peak m/z</div>
                {#each values as val,i}
                    <div class="line">
                        <input bind:value={val.mz} type="number" step="0.001" placeholder=100.000 on:change>
                        <input bind:value={val.intensity} type="number" step="0.001" placeholder=100.000 on:change>
                        <button on:click={(i) => deleteValue(i)}>Delete</button></div>
                {/each}
                <div class="line"><LabeledLine label="Add value">
                    <input bind:value={newValue.mz} type="number" step="0.001" placeholder=100.000 on:change>
                    <input bind:value={newValue.intensity} type="number" step="0.001" placeholder=100.000 on:change>
                </LabeledLine></div>
            </Card>
        </div>
        <block class="pure-u-1-2">
            <Card>
                <div class="headline">Parameters</div>
                <div class="line"><NumberInput label="Intensity threshold" bind:value={intThreshold}></NumberInput> </div>

            </Card>
        </block>
    </div>
    <SearchButton></SearchButton>
    {/if}
</Card>
<style>
    .select {
        display: inline-block;
        padding: 0.5em;
    }
    textarea {
        width: 100%;
        height: 8em;
    }
</style>
