package main

import (
	"net/http"
	"html/template"
	"qiniupkg.com/x/log.v7"
	"golang.org/x/net/websocket"
)

const tpl = `<html>
<head></head>
<body>
    <script type="text/javascript">
        var sock = null;
        var wsuri = "ws://127.0.0.1:1234";
        window.onload = function() {
            console.log("onload");
            sock = new WebSocket(wsuri);
            sock.onopen = function() {
                console.log("connected to " + wsuri);
            }
            sock.onclose = function(e) {
                console.log("connection closed (" + e.code + ")");
            }
            sock.onmessage = function(e) {
                console.log("message received: " + e.data);
            }
        };
        function send() {
            var msg = document.getElementById('message').value;
            sock.send(msg);
        };
    </script>
    <h1>WebSocket Echo Test</h1>
    <form>
        <p>
            Message: <input id="message" type="text" value="Hello, world!">
        </p>
    </form>
    <button onclick="send();">Send Message</button>
</body>
</html>`

func myHander(w http.ResponseWriter, r *http.Request) {
	temp, err := template.New("tpl").Parse(tpl)
	if err != nil {
		http.Error(w, "no resource", http.StatusNotFound)
	}
	temp.Execute(w, nil)
}

func echo(ws *websocket.Conn) {
	var str string

	for {
		if err := websocket.Message.Receive(ws, &str); err != nil {
			log.Println("can't receive")
			break
		}

		log.Println("Received back from client: " + str)

		if err := websocket.Message.Send(ws, str); err != nil {
			log.Println("can't send")
			break
		}
		log.Println("sended message to client: " + str)

	}

}

func main() {
	http.HandleFunc("/send", myHander)
	http.Handle("/", websocket.Handler(echo))

	log.Print("serve is starting...")
	log.Fatal(http.ListenAndServe(":1234", nil))
}
