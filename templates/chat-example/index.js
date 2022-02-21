var submit_btn = document.getElementById("form-submit")
var chat_inp = document.getElementById("chat_inp")
var socket = new WebSocket("ws://localhost:3001/v1/ws");

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