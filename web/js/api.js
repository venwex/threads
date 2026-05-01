const API_BASE = "";

const ROUTES = {
    signUp: "/sign-up",
    signIn: "/sign-in",
    refresh: "/refresh",
    posts: "/posts",
    ws: "/ws",
};

function getAccessToken() {
    return localStorage.getItem("access_token");
}

function getRefreshToken() {
    return localStorage.getItem("refresh_token");
}

function setTokens(accessToken, refreshToken) {
    if (accessToken) {
        localStorage.setItem("access_token", accessToken);
    }

    if (refreshToken) {
        localStorage.setItem("refresh_token", refreshToken);
    }
}

function setCurrentUser(user) {
    if (!user) return;

    const hasUsefulData =
        user.id ||
        user.username ||
        user.email ||
        user.role;

    if (!hasUsefulData) return;

    localStorage.setItem("user", JSON.stringify(user));
}

function getCurrentUser() {
    const raw = localStorage.getItem("user");

    if (!raw) return null;

    try {
        return JSON.parse(raw);
    } catch {
        return null;
    }
}

function clearAuth() {
    localStorage.removeItem("access_token");
    localStorage.removeItem("refresh_token");
    localStorage.removeItem("user");
}

function extractAccessToken(data) {
    return data?.access_token || data?.accessToken || data?.token || data?.access;
}

function extractRefreshToken(data) {
    return data?.refresh_token || data?.refreshToken || data?.refresh;
}

function extractUser(data, accessToken) {
    if (data?.user) return data.user;

    const claims = decodeJwtPayload(accessToken);

    return {
        id:
            data?.user_id ||
            data?.userID ||
            data?.id ||
            claims?.user_id ||
            claims?.userID ||
            claims?.sub,

        username:
            data?.username ||
            claims?.username ||
            claims?.name,

        email:
            data?.email ||
            claims?.email,

        role:
            data?.role ||
            claims?.role,
    };
}

async function apiFetch(path, options = {}) {
    const token = getAccessToken();

    const headers = {
        "Content-Type": "application/json",
        ...(options.headers || {}),
    };

    if (token) {
        headers.Authorization = `Bearer ${token}`;
    }

    const response = await fetch(`${API_BASE}${path}`, {
        ...options,
        headers,
    });

    let data = null;

    try {
        data = await response.json();
    } catch {
        data = null;
    }

    if (!response.ok) {
        const message =
            data?.error ||
            data?.message ||
            data?.detail ||
            `Request failed with status ${response.status}`;

        throw new Error(message);
    }

    return data;
}

function requireAuthPage() {
    if (!getAccessToken()) {
        window.location.href = "/login.html";
    }
}

function redirectIfAuthenticated() {
    if (getAccessToken()) {
        window.location.href = "/app.html";
    }
}

function decodeJwtPayload(token) {
    if (!token) return null;

    try {
        const payload = token.split(".")[1];

        if (!payload) return null;

        const normalizedPayload = payload
            .replace(/-/g, "+")
            .replace(/_/g, "/");

        const decoded = atob(normalizedPayload);

        return JSON.parse(decoded);
    } catch {
        return null;
    }
}