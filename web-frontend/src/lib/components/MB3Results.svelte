<script lang="ts">
import Pagination from "$component/Pagination.svelte";
import ShortRecordSummary from "$component/ShortRecordSummary.svelte";
import type {Filters} from "$lib/common/FilterFunctions.js";

let curPage=1
let pages=1
export let filters: Filters;

let results: any  = null
let error: Error | null = null;

async function getResults(page) {
    error = null
    let url = new URL("/v1/records",baseURL)
    url.searchParams.append('page',page.toString())
    url.searchParams.append('instrument_type',filters.instrumentType.join())
    url.searchParams.append('ms_type',filters.msType.join())
    url.searchParams.append('ion-mode',filters.ionMode.join())
    url.searchParams.append('contributor',filters.contributors.join())
    let resp = await fetch(url)
    let jsonData = await  resp.json();
    if(resp.ok) {
        pages = Math.floor(Number(jsonData.metadata.result_count)/20+1)
        if (curPage>pages) {
            curPage=1
        }
        results = jsonData;
        console.log(jsonData)
        return

    } else {
        throw new Error("Could not get results")
    }
}

$: filters && typeof window !== 'undefined' && getResults(curPage)
export let baseURL
</script>

<h2>Results</h2>
{#if error === null && results === null}
    <div class="info">Loading results...</div>
{:else if (error === null)}

    <Pagination bind:currentPage={curPage} pages={pages}>(Found {results.metadata.result_count} compounds with {results.metadata.spectra_count} spectra)
        {#each results.data as record}
            <ShortRecordSummary record="{record}"></ShortRecordSummary>
        {/each}
    </Pagination>
{:else}
    <div class="error">Error while loading results</div>
{/if}
