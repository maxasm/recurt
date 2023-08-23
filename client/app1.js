const btn_rewrite = document.getElementById("btn_rewrite")
const res_span = document.querySelector("#res")
const inp = document.querySelector("#inp")

// the websocket connection
let socket; 

btn_rewrite.addEventListener("click", async ()=>{
    // get the text from the textarea
    text = inp.value
    console.log(`text: ${text}`)    
    
    const resp = await fetch("/rewrite", 
    {
        method: "POST",
        body: text,
    })
    
    if (resp.status === 200) {
        console.log("submitted the text successfully ...") 
    } else {
        console.log("error submitting text to server") 
    } 
    
    // convert response to JSON
    const resp_text = await resp.json()

    console.log(`server sent the ode: ${resp_text}`)
    
    // create the websocket connection
    socket = new WebSocket(`ws://127.0.0.1:8080/ws/${resp_text.text}`)

    // Listen for open signal
    socket.addEventListener("open", ()=> {
        console.log("WS connection open")    
    })

    // Listen for messages
    socket.addEventListener('message', function (event) {
        console.log('Message from server:', event.data);
        res_span.textContent = event.data 
    });

    // Connection closed
    socket.addEventListener('close', function (event) {
        console.log('WebSocket connection closed:', event);
    });

    // Error handling
    socket.addEventListener('error', function (error) {
        console.error('WebSocket Error:', error);
    });

})

