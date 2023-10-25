<form method="POST">
    <div class="mb-3">
        <label for="username" class="form-label">Username</label>

        <!-- svelte-ignore a11y-autofocus -->
        <input class="form-control" name="username" id="username" required autofocus/>
    </div>

    <div class="mb-3">
        <label for="password" class="form-label">Password</label>
        <input type="password" class="form-control" name="password" id="password" required/>
    </div>

    <button type="submit" class="btn btn-primary">Login</button>
</form>