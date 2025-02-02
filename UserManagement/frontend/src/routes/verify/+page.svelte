<script>
    import { apiFetch } from "../../lib/api";

    let verificationMessage = "";

    async function verifyToken() {
        try {
            await apiFetch("/verify", {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem("token")}`,
                },
            });
            verificationMessage = "Token is valid!";
        } catch {
            verificationMessage = "Invalid or expired token.";
        }
    }

    verifyToken();
</script>

<h1>{verificationMessage}</h1>