$(function() {
	var socket = null;
	var msgBox = $("#chatbox textarea");
	var messages = $("#messages");
	$("#chatbox").submit(function(event) {
		event.preventDefault()

		if (!msgBox.val()) {
			return false;
		}
		if (!socket) {
			alert("Error: There is no socket connection.")
			return false;
		}
		socket.send(JSON.stringify({
			"Message": msgBox.val()
		}));
		msgBox.val("");
		return false
	})

	if (!window["WebSocket"]) {
		alert("Error: Your browser does not support web sockets.")
	} else {
		socket = new WebSocket("ws://" + host + "/room");
		socket.onclose = function() {
			alert("Connection has been closed.");
		}
		socket.onmessage = function(event) {
			var msg = JSON.parse(event.data);
			messages.append(
				$("<li>").append(
					$("<img>")
					.attr("title", msg.Name)
					.attr("src", msg.AvatarURL),
					$("<strong>").text(msg.Name + ": "),
					$("<span>").text(msg.Message)
				)
			);
		}
	}
})
