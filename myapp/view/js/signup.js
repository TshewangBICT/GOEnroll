function signUp() {
    // retrieve form data
    var _data = {
        firstname: document.getElementById("fname").value,
        lastname: document.getElementById("lname").value,
        email: document.getElementById("email").value,
        password: document.getElementById("pw1").value,
        pw: document.getElementById("pw2").value
    }
    if (_data.password !== _data.pw) {
        alert("Passwords doesn't match!")
        return
    }
    fetch("/signup", {
        method: "POST",
        body: JSON.stringify(_data),
        headers: {"Content-type": "application/json; charset=UTF-8"}
    }).then(response => {
        if(response.status == 201) {
            window.open("index.html", "_self")
        } else {
            throw new Error(response.statusText)
        }
    }).catch(e => {
        alert(e)
    });
}