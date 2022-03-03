var serverURL = 'http://localhost:3001/v1/user'

function RegisterRequest() {
    var username = document.getElementById("username")
    var password = document.getElementById("password")

    fetch(serverURL, {
        method: 'post',
        body: JSON.stringify({"username": username.value, "password": password.value}),
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        }
    }).then((response) => {
        console.log(response)
        if (response.status === 201) {
            console.log("register successfully")
            location.href = "http://localhost:63342/lets-go-chat/templates/chat-example/pages/chat.html"
        }
    }).catch((error) => {
        console.log(error)
    })
}