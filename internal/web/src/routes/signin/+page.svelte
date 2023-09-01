<script lang="ts">
  import { goto } from "$app/navigation"

  let email: string
  let password: string

  function submit() {
    fetch("/v1/users/signin", {
      method: "POST",
      body: JSON.stringify({email, password})
    }).then((data) => {
      data.json().then(d => console.log(d))
      goto("/")
    })
      .catch((err) => alert(err.name))
    // redirect()
  }
</script>

<section class="container">
    <form on:submit|preventDefault={submit}>
        <div class="form-element">
            <label for="email">Email</label>
            <input id="email" name="email" type="email" required bind:value={email}>
        </div>

        <div class="form-element">
            <label for="password">Password</label>
            <input id="password" required bind:value={password} type="password">
        </div>

        <button>Sign in</button>
    </form>
</section>

<style>
    .container {
        display: flex;
        justify-content: center;
        align-items: center;
        width: 100%;
        height: 100vh;
    }

    form {
        border: 1px solid gray;
        border-radius: 6px;
        padding: 15px 10px;
    }

    .form-element {
       margin: 15px 0;
    }
    .form-element label, input {
        display: block;
        font-size: 1.5rem;
    }

    button {
        background-color: aquamarine;
        outline: none;
        border: 1px solid gray;
        border-radius: 6px;
        font-size: 1.5rem;
        transition: opacity ease-in 100ms;
        margin-top: 5px;
    }

    button:hover {
        opacity: 0.8;
        cursor: pointer;
    }

</style>
