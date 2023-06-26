<script lang="ts">
import MB3Filters from "$component/MB3Filters.svelte";
import MB3Results from "$component/MB3Results.svelte";
import type {Filters} from "$lib/common/FilterFunctions";
import AccordionItem from "$component/AccordionItem.svelte";
import Card from "$component/Card.svelte";
import MB3BasicSearch from "$component/MB3BasicSearch.svelte";
import MB3PeakSearch from "$component/MB3PeakSearch.svelte";
import FullTextSearch from "$component/FullTextSearch.svelte";
import MB3SimilaritySearch from "$component/MB3SimilaritySearch.svelte";

/** @type {import('./$typesunknownPageData} */
export let data: any;
    let base: string;

let filters: Filters  = {
    contributors: [],
    msType: [],
    instrumentType: [],
    ionMode : [],
    fullText: null
}

let activeFilters = filters

$: base = data.baseurl

</script>

<div class="pure-g">
    <div class="pure-u-1-5">
        <Card>
            <MB3Filters bind:filters bind:baseUrl={base}></MB3Filters>
        </Card>
    </div>
    <div class="pure-u-3-5">
        <AccordionItem headline="Fulltext search" bind:value={filters.fullText}><FullTextSearch/></AccordionItem>
        <AccordionItem headline="Basic search"><MB3BasicSearch/></AccordionItem>
        <AccordionItem headline="Peak search"><MB3PeakSearch></MB3PeakSearch></AccordionItem>
        <AccordionItem headline="Peak difference search"><MB3PeakSearch/></AccordionItem>
        <AccordionItem headline="Spectrum similarity search"><MB3SimilaritySearch></MB3SimilaritySearch></AccordionItem>

        <MB3Results bind:baseURL={base} filters={filters}></MB3Results>
    </div>
    <div class="pure-u-1-5">
        <Card>
        <h1>Active filters</h1>
            {#each Object.entries(filters) as [k,v] }
                {#if Array.isArray(v) }
                    {#if v.length > 0}
                    {k}:
                    {#each v as vv}
                        {vv}&nbsp;
                    {/each}
                    {/if}
                {:else }
                    {#if v !== null}
                        {k}.: {v} &nbsp;
                    {/if}
                {/if}

            {/each}
        </Card>
    </div>
</div>
