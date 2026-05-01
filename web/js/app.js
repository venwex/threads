requireAuthPage();

const postsList = document.getElementById("postsList");
const postForm = document.getElementById("postForm");
const postInput = document.getElementById("postInput");
const postButton = document.getElementById("postButton");
const logoutBtn = document.getElementById("logoutBtn");

const profileName = document.getElementById("profileName");
const profileEmail = document.getElementById("profileEmail");
const composerAvatar = document.getElementById("composerAvatar");
const sidebarAvatar = document.getElementById("sidebarAvatar");

const user = getCurrentUser();

renderCurrentUser();
loadPosts();
connectWebSocket();

logoutBtn.addEventListener("click", () => {
    clearAuth();
    window.location.href = "/login.html";
});

postInput.addEventListener("input", () => {
    postButton.disabled = postInput.value.trim().length === 0;
});

postForm.addEventListener("submit", async (event) => {
    event.preventDefault();

    const content = postInput.value.trim();

    if (!content) return;

    postButton.disabled = true;

    try {
        await apiFetch(ROUTES.posts, {
            method: "POST",
            body: JSON.stringify({ content }),
        });

        postInput.value = "";
    } catch (err) {
        showFeedError(err.message);
    } finally {
        postButton.disabled = false;
    }
});

function renderCurrentUser() {
    const username = user?.username || "user";
    const email = user?.email || "no email";

    profileName.textContent = username;
    profileEmail.textContent = email;

    const letter = username.slice(0, 1).toUpperCase();

    composerAvatar.textContent = letter;
    sidebarAvatar.textContent = letter;
}

async function loadPosts() {
    postsList.innerHTML = createLoader();

    try {
        const data = await apiFetch(ROUTES.posts);
        const posts = normalizePosts(data);

        renderPosts(posts);
    } catch (err) {
        showFeedError(err.message);
    }
}

function normalizePosts(data) {
    if (Array.isArray(data)) return data;
    if (Array.isArray(data?.posts)) return data.posts;
    if (Array.isArray(data?.data)) return data.data;
    if (Array.isArray(data?.items)) return data.items;

    return [];
}

function renderPosts(posts) {
    postsList.innerHTML = "";

    if (!posts.length) {
        postsList.innerHTML = `
            <section class="empty-state">
                <div class="empty-icon">◎</div>
                <h3>No posts yet</h3>
                <p>Write the first post. Humanity has survived worse.</p>
            </section>
        `;
        return;
    }

    posts.forEach((post) => renderPost(post));
}

function renderPost(post, prepend = false) {
    removeEmptyState();

    const normalized = normalizePost(post);

    const article = document.createElement("article");
    article.className = "post-card";
    article.dataset.postId = normalized.id || "";

    article.innerHTML = `
        <div class="post-avatar">${escapeHtml(getInitial(normalized.author))}</div>

        <div class="post-main">
            <header class="post-header">
                <div class="post-author-block">
                    <strong>${escapeHtml(normalized.author)}</strong>
                    <span>@${escapeHtml(normalized.author)}</span>
                </div>

                <time>${escapeHtml(formatDate(normalized.createdAt))}</time>
            </header>

            <p class="post-content">${escapeHtml(normalized.content)}</p>

            <footer class="post-actions">
                <button type="button" aria-label="Like">♡</button>
                <button type="button" aria-label="Reply">💬</button>
                <button type="button" aria-label="Repost">↻</button>
                <button type="button" aria-label="Share">↗</button>
            </footer>
        </div>
    `;

    if (prepend) {
        postsList.prepend(article);
    } else {
        postsList.appendChild(article);
    }
}

function normalizePost(post) {
    const author =
        post?.author?.username ||
        post?.author_username ||
        post?.authorUsername ||
        post?.username ||
        post?.user?.username ||
        "unknown";

    return {
        id: post?.id,
        author,
        content: post?.content || "",
        createdAt: post?.created_at || post?.createdAt || post?.created || "",
    };
}

function connectWebSocket() {
    const token = getAccessToken();

    if (!token) return;

    const protocol = location.protocol === "https:" ? "wss" : "ws";
    const wsUrl = `${protocol}://${location.host}${ROUTES.ws}?token=${encodeURIComponent(token)}`;

    const socket = new WebSocket(wsUrl);

    socket.onopen = () => {
        console.log("websocket connected");
    };

    socket.onmessage = (event) => {
        try {
            const post = JSON.parse(event.data);
            renderPost(post, true);
        } catch {
            console.warn("Invalid websocket message:", event.data);
        }
    };

    socket.onerror = () => {
        console.warn("websocket error");
    };

    socket.onclose = () => {
        setTimeout(connectWebSocket, 3000);
    };
}

function removeEmptyState() {
    const emptyState = postsList.querySelector(".empty-state");
    if (emptyState) {
        emptyState.remove();
    }
}

function showFeedError(message) {
    postsList.innerHTML = `
        <section class="empty-state error-state">
            <div class="empty-icon">!</div>
            <h3>Something broke</h3>
            <p>${escapeHtml(message)}</p>
        </section>
    `;
}

function createLoader() {
    return `
        <section class="loader-state">
            <div class="loader"></div>
            <p>Loading posts...</p>
        </section>
    `;
}

function getInitial(value) {
    if (!value) return "?";
    return value.slice(0, 1).toUpperCase();
}

function formatDate(value) {
    if (!value) return "now";

    const date = new Date(value);

    if (Number.isNaN(date.getTime())) {
        return "now";
    }

    return date.toLocaleString("en", {
        month: "short",
        day: "numeric",
        hour: "2-digit",
        minute: "2-digit",
    });
}

function escapeHtml(value) {
    return String(value)
        .replaceAll("&", "&amp;")
        .replaceAll("<", "&lt;")
        .replaceAll(">", "&gt;")
        .replaceAll('"', "&quot;")
        .replaceAll("'", "&#039;");
}