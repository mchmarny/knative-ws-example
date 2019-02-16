window.onload = function () {

    console.log("Protocol: " + location.protocol);
    var wsURL = "ws://" + document.location.host + "/ws"
    // TODO: websocketUpgrade not set in Istio so errs on upgrade if WSS
    if (location.protocol == 'https:') {
        wsURL = "wss://" + document.location.host + "/ws"
    }
    console.log("WS URL: " + wsURL);

    var msg = document.getElementById("eventmsg");
    var log = document.getElementById("eventlog");

    function appendLog(item) {
        var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
        log.appendChild(item);
        if (doScroll) {
            log.scrollTop = log.scrollHeight - log.clientHeight;
        }
    }

    function setMsg(msf) {
        msg.innerHTML = msf;
    }

    if (log) {

        sock = new WebSocket(wsURL);


        sock.onopen = function () {
            console.log("connected to " + wsURL);
            setMsg("<b>Connection Opened</b>");
        };

        sock.onclose = function (e) {
            console.log("connection closed (" + e.code + ")");
            setMsg("<b>Connection closed</b>");
        };

        sock.onmessage = function (e) {
            console.log(e);
            var eventObj = JSON.parse(e.data);
            console.log(eventObj);
            var item = document.createElement("div");
            item.className = "item";
            item.innerHTML = JSON.stringify(eventObj, undefined, 2);
            appendLog(item);
        };

    } // if log

};