<script>
    import { base } from '$app/paths';
</script>

<nav class="navbar navbar-expand-md bg-dark navbar-dark">
    <div class="container-fluid">
        <span class="navbar-brand">Tickets | Unauthorized</span>
        <div class="collapse navbar-collapse">
            <ul class="navbar-nav me-auto">
                <li class="nav-item">
                    <a class="nav-link" href="/">User Management</a>
                </li>
                <li class="nav-item">
                    <a class="nav-link active" href="{base}">Main</a>
                </li>
            </ul>
        </div>
    </div>
</nav>

<div class="Home container mt-5">
    <div class="row">
        <div class="col-12">
            <div class="text-center text-danger">
                <p class="specialParagraph">You are not authorized to view tickets.</p>
                <p class="specialParagraph">Please log in to using the User Management.</p>
            </div>
        </div>
    </div>
</div>

<style>
    .Home {
      width: 90em;
    }
    .specialParagraph {
        font-size: 1.5em;
    }
</style>