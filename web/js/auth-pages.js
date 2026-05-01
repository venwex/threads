const loginForm = document.getElementById("loginForm");
const registerForm = document.getElementById("registerForm");
const errorBox = document.getElementById("errorBox");

redirectIfAuthenticated();

const params = new URLSearchParams(window.location.search);

if (params.get("registered") === "true" && errorBox) {
    errorBox.textContent = "Account created successfully. Now log in.";
    errorBox.classList.add("visible", "success");
}

function showError(message) {
    if (!errorBox) return;

    errorBox.textContent = message;
    errorBox.classList.add("visible");
}

function hideError() {
    if (!errorBox) return;

    errorBox.textContent = "";
    errorBox.classList.remove("visible");
}

function saveAuthData(data) {
    const accessToken = extractAccessToken(data);
    const refreshToken = extractRefreshToken(data);

    if (!accessToken) {
        throw new Error("Backend did not return access token");
    }

    const user = extractUser(data, accessToken);

    setTokens(accessToken, refreshToken);
    setCurrentUser(user);
}

if (loginForm) {
    loginForm.addEventListener("submit", async (event) => {
        event.preventDefault();
        hideError();

        const formData = new FormData(loginForm);

        const payload = {
            login: formData.get("login"),
            password: formData.get("password"),
        };

        try {
            const data = await apiFetch(ROUTES.signIn, {
                method: "POST",
                body: JSON.stringify(payload),
            });

            saveAuthData(data);
            window.location.href = "/app.html";
        } catch (err) {
            showError(err.message);
        }
    });
}

if (registerForm) {
    registerForm.addEventListener("submit", async (event) => {
        event.preventDefault();
        hideError();

        const formData = new FormData(registerForm);

        const payload = {
            username: formData.get("username"),
            email: formData.get("email"),
            password: formData.get("password"),
        };

        try {
            await apiFetch(ROUTES.signUp, {
                method: "POST",
                body: JSON.stringify(payload),
            });

            window.location.href = "/login.html?registered=true";
        } catch (err) {
            showError(err.message);
        }
    });
}