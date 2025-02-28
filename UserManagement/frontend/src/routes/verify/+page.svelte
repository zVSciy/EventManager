<script>
    let verificationMessage = "";

    async function verifyToken() {
        try {
            const response = await fetch("/api/token", {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem("token")}`,
                },
            });
            if (!response.ok) {
                throw new Error(response.statusText);
            }
            const result = await response.json();
            verificationMessage = result.message;
        } catch (error) {
            verificationMessage = "Invalid or expired token.";
        }
    }

    verifyToken();
</script>

<h1>{verificationMessage}</h1>