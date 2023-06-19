<script lang="ts">
import Pagination from "$component/Pagination.svelte";
import ShortRecordSummary from "$component/ShortRecordSummary.svelte";
import {Filters} from "$lib/common/FilterFunctions.js";

let curPage=1
let pages=1
export let filters: Filters;

async function getResults(page) {
    let url = new URL("/v1/records",baseURL)
    url.searchParams.append('page',page.toString())
    url.searchParams.append('instrument_type',filters.instrumentType.join())
    url.searchParams.append('ms_type',filters.msType.join())
    url.searchParams.append('ion-mode',filters.ionMode.join())
    url.searchParams.append('contributor',filters.contributors.join())
    let resp = await fetch(url)
    let jsonData = await  resp.json();
    if(resp.ok) {
        console.log(JSON.stringify(jsonData))
        pages = Math.floor(Number(jsonData.metadata.result_count)/20+1)
        if (curPage>pages) {
            curPage=1
        }
        return jsonData

    } else {
        console.log(jsonData)
        throw new Error("Could not get results")
    }
}

export let baseURL
</script>

{#await getResults(curPage)}
    <div class="info">Loading results...</div>
{:then records}
    <Pagination bind:currentPage={curPage} pages={pages}>
        {#each records.data as record}
            <ShortRecordSummary record="{record}"></ShortRecordSummary>
        {/each}
    </Pagination>
{:catch error}
    <div class="error">Error while loading results</div>
{/await}
