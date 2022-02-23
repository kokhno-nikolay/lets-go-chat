var submit_btn = document.getElementById("form-submit")
var chat_inp = document.getElementById("chat_inp")
var socket = new WebSocket("ws://localhost:3001/ws?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1dWlkIjoiNjY1ZWZlNGYtZDdkYy00NDM0LTgzZTYtZmE2YzI3NDgxYzhkIiwiZXhwIjoxNjQ1NzgyNDg2LCJpYXQiOjE2NDU2MDk2ODZ9.edkubjIJjhlAoyR2jYX_bT0h1ltnL2oUceqd4_9IWsU");

socket.onopen = () => {
    console.log("Successfully connected from ws server");
};

socket.onclose = event => {
    console.log("Socket closed connection: ", event);
    socket.send("Client closed connection!")
};

socket.onerror = error => {
    console.log("Socket error: ", error);
};

socket.onmessage = function(event) {
    var out = document.getElementById('res');
    out.innerHTML = event.data;
}

submit_btn.addEventListener('click', function () {
    sendMessage(chat_inp.value)
});

function sendMessage(data) {
    socket.send(data)
}