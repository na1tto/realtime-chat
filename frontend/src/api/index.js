// This is the WebSocket connection code wich will be used by our fronted

var socket = new WebSocket("ws://localhost:8080/ws");

let connect = () => {
  console.log("Attempting Connection...");

  socket.onopen = () => {
    console.log("Successfully connected!");
  };

  socket.onmessage = (msg) => {
    console.log(msg);
  };

  socket.onclose = (event) => {
    console.log("Socket Closed Connection: ", event);
  };

  socket.onerror = (error) => {
    console.log("Socket error: ", error);
  };
};

let sendMsg = (msg) => {
  console.log("sending message: ", msg);
  socket.send(msg);
};

export { connect, sendMsg };
