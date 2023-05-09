class WebSocketService {
    static instance = null;
    callbacks = {};

    static getInstance() {
        if (!WebSocketService.instance) {
            WebSocketService.instance = new WebSocketService();
        }
        return WebSocketService.instance;
    }

    constructor() {
        this.socketRef = null;
    }

    connect(url) {
        this.socketRef = new WebSocket(url);
        this.socketRef.onopen = () => {
            console.log("WebSocket open");
        };
        this.socketRef.onmessage = (e) => {
            this.handleMessage(e);
        };
        this.socketRef.onerror = (e) => {
            console.log("WebSocket error:", e);
        };
        this.socketRef.onclose = () => {
            console.log("WebSocket closed, reconnecting...");
            this.connect(url);
        };
    }

    handleMessage(e) {
        const parsedData = JSON.parse(e.data);
        const command = parsedData.command;
        if (this.callbacks[command]) {
            this.callbacks[command](parsedData);
        }
    }

    sendCommand(command, data) {
        try {
            this.socketRef.send(JSON.stringify({ ...data, command }));
        } catch (err) {
            console.log("WebSocket error:", err);
        }
    }

    addCallback(command, callback) {
        this.callbacks[command] = callback;
    }
}

const WebSocketInstance = WebSocketService.getInstance();

export default WebSocketInstance;