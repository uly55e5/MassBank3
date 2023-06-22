<script>
    import { slide } from "svelte/transition";
    export let headline
    let isOpen = false
    const toggle = () => isOpen = !isOpen
</script>

<div class="box">
<button on:click={toggle} aria-expanded={isOpen}><svg style="tran"  width="1em" height="1em" fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" viewBox="0 0 20 20" stroke="currentColor"><path d="M9 5l7 7-7 7"></path></svg> {headline} </button>
{#if isOpen}
    <div transition:slide={{ duration: 300 }}>
        <slot />
    </div>
{/if}
</div>

<style>
    button {border: none; background: #eeeeee; width:100%; display:block; color: inherit; text-align: left; font-size: 120%; cursor: pointer; margin: 0; padding-bottom: 0.5em; padding-top: 0.5em}

    svg { transition: transform 0.2s ease-in;
    }

    [aria-expanded=true] svg { transform: rotate(0.25turn); }

    .box {
        margin: 1em 0;
        border: 0.1px solid;
    }
</style>
