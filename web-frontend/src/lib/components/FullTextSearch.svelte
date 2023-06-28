<script lang="ts">
    import TextInput from "$component/TextInput.svelte";
    import Card from "$component/Card.svelte";

    export let value:string
</script>
<Card>
<TextInput label="Search text" bind:value></TextInput>
</Card>
