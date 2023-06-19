<script lang="ts">
    import {Pie} from "svelte-chartjs";
    import {getFilterValues} from "$lib/common/FilterFunctions";
    import {ArcElement, CategoryScale, Chart as ChartJS, Colors, Legend, Title, Tooltip} from "chart.js";

    let chartData = {
        contributorsData: {
            labels: [],
            datasets: [
                {
                    data: [],
                },
            ],
        },
        InstrumentType: {
            labels: [],
            datasets: [
                {
                    data: [],
                },
            ],
        },
        msType: {
            labels: [],
            datasets: [
                {
                    data: [],
                },
            ],
        },

    }

    export let baseURL: string;

    ChartJS.register(Title, Tooltip, Legend, ArcElement, CategoryScale, Colors);

    async function getChartData(base) {
        let jsonData = await getFilterValues(base)


            chartData.contributorsData.labels = []
            chartData.contributorsData.datasets[0].data = []
            for (const co of jsonData.contributor) {
                chartData.contributorsData.labels.push(co.value)
                chartData.contributorsData.datasets[0].data.push(co.count)
            }
            chartData.msType.labels = []
            chartData.msType.datasets[0].data = []
            for (const co of jsonData.ms_type) {
                chartData.msType.labels.push(co.value)
                chartData.msType.datasets[0].data.push(co.count)
            }
            chartData.InstrumentType.labels = []
            chartData.InstrumentType.datasets[0].data = []
            for (const co of jsonData.instrument_type) {
                chartData.InstrumentType.labels.push(co.value)
                chartData.InstrumentType.datasets[0].data.push(co.count)
            }
            chartData = chartData

        return jsonData

    }
</script>

{#await getChartData(baseURL)}
    <div class="info">loading charts...</div>
{:then filters}
    <div>
        <b>MassBank Version: </b>{filters.metadata.version}<br>
        <br>
        <b>Compounds: </b>{filters.metadata.compound_count}<br>
        <b>Isomers: </b>{filters.metadata.isomer_count}<br>
        <b>Spectra: </b>{filters.metadata.spectra_count}<br>
        <br>
        <Pie data={chartData.contributorsData} ></Pie>
        <Pie data={chartData.InstrumentType} ></Pie>
        <Pie data={chartData.msType} ></Pie>
    </div>
{/await}
