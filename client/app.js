const btn = document.getElementById("button")

let socket = new WebSocket("ws://127.0.0.1:8080/ws")

socket.addEventListener("open", ()=> {
    console.log("WS connection open")    
})

// Listen for messages
socket.addEventListener('message', function (event) {
    console.log('Message from server:', event.data);
});

// Connection closed
socket.addEventListener('close', function (event) {
    console.log('WebSocket connection closed:', event);
});

// Error handling
socket.addEventListener('error', function (error) {
    console.error('WebSocket Error:', error);
});

btn.addEventListener("click", ()=> {
     
})
