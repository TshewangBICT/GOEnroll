function login() {
    var data = {
        email: document.getElementById("email").value,
        password: document.getElementById("pw").value
    }

    fetch("/login", {
        method: "POST",
        body: JSON.stringify(data),
        headers: {"Content-type": "application/json; charset=UTF-8"},
        credentials: "include"
    }).then(response => {
        if (response.ok) {
            window.open("student.html", "_self")
        } else {
            throw new Error(response.statusText)
        }
    }).catch(e => {
        alert(e)
    });
}

function logout() {
    fetch("/logout", {
        method: "GET",
        credentials: "include"
    })
    .then(response => {
        if (response.ok) {
            // Clear all storage
            localStorage.clear();
            sessionStorage.clear();
            // Force redirect
            window.location.replace("/index.html");
        } else {
            throw new Error("Logout failed");
        }
    }).catch(e => {
        console.error(e);
        // Still redirect even if fetch fails
        window.location.replace("/index.html");
    });
}