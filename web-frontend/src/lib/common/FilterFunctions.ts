export interface Filters {
    instrumentType: string[]
    msType: string[]
    ionMode: string[]
    contributors: string[]
    fullText: string | null
}

export async function getFilterValues(baseURL: string, all = true, filters: Filters) {
    const url = new URL("/v1/filter/browse", baseURL);
    if (!all) {
        url.searchParams.append('instrument_type', filters.instrumentType.join())
        url.searchParams.append('ms_type', filters.msType.join())
        url.searchParams.append('ion-mode', filters.ionMode.join())
        url.searchParams.append('contributor', filters.contributors.join())
    }
    const resp = await fetch(url);
    const jsonData = await resp.json();
    if (resp.ok) {
        return jsonData
    } else {
        console.log(jsonData);
        throw new Error("Could not load filters.")
    }

}
