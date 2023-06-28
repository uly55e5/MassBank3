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
let fullText
$: base = data.baseurl
$: filters.fullText = fullText

</script>

<div class="pure-g">
    <div class="pure-u-1-5">
        <Card>
            <MB3Filters bind:filters bind:baseUrl={base}></MB3Filters>
        </Card>
    </div>
    <div class="pure-u-3-5">
        <AccordionItem headline="Fulltext search" ><FullTextSearch bind:value={filters.fullText}/></AccordionItem>
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
                        <div><div class="key">{k.replace(/([A-Z]+)/g, " $1").replace(/([A-Z][a-z])/g, " $1")}</div>:
                        {#each v as vv}
                            <div class="cloud">{vv}</div>&nbsp;
                        {/each}
                        </div>
                    {/if}
                {:else}
                    {#if v !== null && typeof v !== 'undefined' && v !== ""}
                        <div><div class="key">{k.replace(/([A-Z]+)/g, " $1").replace(/([A-Z][a-z])/g, " $1")}</div>: <div class="cloud">{v}</div></div>
                    {/if}
                {/if}
            {/each}
        </Card>
    </div>
</div>

<style>
    .cloud {
        border: solid #aaaaaa;
        border-radius: 1em;
        display: inline-block;
        padding: 0.2em 0.5em;
        background-color: #dddddd;
    }

    .key {
        font-weight: bold;
        display: inline-block;
    }
    .key::first-letter {
        text-transform: capitalize;
    }
</style>