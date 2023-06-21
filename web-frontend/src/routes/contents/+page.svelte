<script lang="ts">
    import MB3Filters from "$component/MB3Filters.svelte";
    import MB3CharInfoCharts from "$component/MB3CharInfoCharts.svelte";
    import MB3Results from "$component/MB3Results.svelte";
    import type {Filters} from "$lib/common/FilterFunctions";

    /** @type {import('./$typesunknownPageData} */
    export let data: any;
    let base: string;

    let filters: Filters  = {
        contributors: [],
        msType: [],
        instrumentType: [],
        ionMode : []
    }

    $: base = data.baseurl
    $: (filters.contributors || filters.ionMode || filters.msType || filters.instrumentType) && (filters = filters)
</script>

<div class="pure-g">
    <div class="pure-u-1-5">
        <div class="card">
            <MB3Filters bind:filters bind:baseUrl={base}></MB3Filters>
        </div>
    </div>
    <div class="pure-u-3-5">
        <h2>Results</h2>
            <MB3Results bind:baseURL={base} filters={filters}></MB3Results>
    </div>
    <div class="pure-u-1-5">
        <div class="card">
            <MB3CharInfoCharts bind:baseURL={base}></MB3CharInfoCharts>
        </div>
    </div>
</div>


<style>
    h2 {
        margin: 0;
        padding: 0.2em;
        font-weight: 600;
    }

    h3 {
        padding: 0.3em;
        margin: 0;
    }

    .card {
        border: solid #a9a9a9;
        margin: 0.5em;
    }

    .record-name {
        font-weight: bold;
        font-size: 1.5em;
    }

</style>
