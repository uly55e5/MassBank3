<script lang="ts">
    export let label:string
</script>

<block class="component">
    <label><div class="pure-g">
        <div class="pure-u-2-5 element">
            {label}
        </div>
        <div class="pure-u-3-5 element">
            <slot/>
        </div>
    </div>
    </label>
</block>

<style>
    .component {
        padding: 0.5em;
    }
    label {
        display: inline-block;
        width: 100%;
    }
</style>
