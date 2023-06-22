<script lang="ts">
    import Card from "$component/Card.svelte";
    export let record: any;
</script>

<Card>
                <div class="record-name">{record.name}</div>
                <div>{record.formula}</div>
    <svg width="30%">
            {@html record.svg}
        </svg>


    <div>{record.mass}</div>
                    <div class="spectra">{record.spectra.length} Spectra :</div>
                        {#each record.spectra as spectrum}
                            <a href="/record/{spectrum.id}"><div>{spectrum.id} {spectrum.title}</div></a>
                    {/each}
                </Card>

<style>
    .card {
        border: solid #a9a9a9;
        margin: 0.5em;
    }

    .record-name {
        font-weight: bold;
        font-size: 1.5em;
    }
</style>
