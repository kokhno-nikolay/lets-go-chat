var serverURL = 'http://localhost:3001/v1/user/login'

function LoginRequest() {
    var username = document.getElementById("username")
    var password = document.getElementById("password")

    fetch(serverURL, {
        method: 'post',
        body: JSON.stringify({"username": username.value, "password": password.value}),
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        }
    }).then((res) => {
        if (res.status === 200) {
            console.log("login successfully")
        }
        return res.json();
    }).then((data) => {
        console.log(data)
        const params = new URL(data.url).searchParams;
        location.href = "http://localhost:63342/lets-go-chat/templates/chat-example/pages/chat.html?token=" + params.get('token');
        window.localStorage.setItem('token', params.get('token'));
    }).catch((error) => {
        console.log(error)
    })
}