// jquery 
$(function() {
  var socket = null;
  var msgBox = $("#chatbox textarea"); // element of the fomr
  var label = $("#chatbox label");
  var messages = $("#messages");
  
  $("#chatbox").submit(function() {
    if (!msgBox.val()) return false; // if empty return fale 
    if (!socket) { // check the socket
      alert("Error: There is no socket connection.");
      return false;
    }
    // send the message by the socket
    socket.send(label.text() + " " + msgBox.val() + "\n");
    msgBox.val(""); // clean the message box
    return false;
  });

  // when charge the js check if the browser support the web sockets
  if (!window["WebSocket"]) {
    alert("Error: Your browser does not support web sockets.")
  } else {
    // check the protocol of the socket
    if (window.location.protocol == "https:"){
      socket = new WebSocket("wss://localhost:8062/ChatRoom/");// web socket secure
    } else {
      socket = new WebSocket("ws://localhost:8061/ChatRoom/");// web socket
    }
    socket.onclose = function() {
      alert("Connection has been closed.");
    }
    socket.onmessage = function(e) {
      messages.append(
        $("<li>").append(
          e.data
        ));
    }
  }
});
