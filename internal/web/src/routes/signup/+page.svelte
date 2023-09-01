<script lang="ts">
    let email: string
    let password: string

    function submit() {
        console.log(email, password)
    }
</script>

<section class="container">
    <form on:submit={submit}>
        <label for="email">email:</label>
        <input id="email" name="email" bind:value={email}>

        <label for="password">password:</label>
        <input id="password" bind:value={password}>

        <button on:click={submit}>Signup</button>
    </form>
</section>

<style>
    .container {
        display: flex;
        justify-content: center;
        align-items: center;
        width: 100%;
        height: 100%;
    }
    form {
        border: 1px solid gray;
        border-radius: 6px;
        padding: 5px;
    }
</style>
