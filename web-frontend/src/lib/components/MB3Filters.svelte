<script lang="ts">
    import FilterButton from "$component/FilterBox.svelte";
    import {Filters, getFilterValues} from "$lib/common/FilterFunctions.js"

    export let filters: Filters;
    export let baseUrl: string;
</script>

{#await getFilterValues(baseUrl)}
    <div class="info">loading filters...</div>
{:then filterValues}
    <div class>
        <h2>Filters</h2>
        <h3>Instrument Type</h3>
        <FilterButton bind:result={filters.instrumentType} group="itype" values={filterValues.instrument_type}></FilterButton>
        <h3>MS Type</h3>
        <FilterButton bind:result={filters.msType} group="mstype" values={filterValues.ms_type}></FilterButton>
        <h3>Ion Mode</h3>
        <FilterButton bind:result={filters.ionMode} group="imode" values={filterValues.ion_mode}></FilterButton>
        <h3>Contributor</h3>
        <FilterButton bind:result={filters.contributors} group="cont" values={filterValues.contributor}></FilterButton>
    </div>
{:catch error}
    <div class="error">Error during Filter loading</div>
{/await}
