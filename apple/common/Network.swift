import Foundation
import Starscream
import CocoaAsyncSocket

class Network: WebSocketDelegate {

    static var address = "ws://127.0.0.1:8000/ws"
    //static var address = "ws://192.168.2.172:8000/ws"

    static let shared = Network()

    private var websocket: WebSocket?

    func connect() {
        guard let url = URL(string: Network.address) else {
            print("could not create url from " + Network.address)
            return
        }
        websocket = WebSocket(url: url)
        websocket?.delegate = self
        websocket?.connect()
    }

    func send(_ data: Data) {
        websocket?.write(data: data)
    }

    // websocket delegate

    func websocketDidConnect(_ websocket: Starscream.WebSocket) {
        EventBus.post(.connected)
    }

    func websocketDidDisconnect(_ websocket: Starscream.WebSocket, error: NSError?) {
        EventBus.post(.disconnected)
    }

    func websocketDidReceiveData(_ websocket: Starscream.WebSocket, data: Data) {
        WireBackend.shared.didReceiveFromServer(data)
    }

    func websocketDidReceiveMessage(_ socket: WebSocket, text: String) {
        print("websocketDidReceiveMessage")
    }
}

class UDPNetwork: NSObject {
    
    static var host = "127.0.0.1"
    static var port: UInt16 = 8001
    
    private var udpSocket: GCDAsyncUdpSocket?
    
    func connect() {
        self.udpSocket = GCDAsyncUdpSocket(delegate: self, delegateQueue: DispatchQueue.main)
        
        do {
            try udpSocket?.connect(toHost: UDPNetwork.host, onPort: UDPNetwork.port)
        } catch  {
            print(error)
            self.udpSocket = nil
        }
    }
    
    func send(_ data: Data) {
        let tag = Int(Date.timeIntervalSinceReferenceDate)
        udpSocket?.send(data, withTimeout: 30, tag: tag)
    }
}

extension UDPNetwork: GCDAsyncUdpSocketDelegate {
    
    func udpSocket(_ sock: GCDAsyncUdpSocket, didConnectToAddress address: Data) {
        print("UDP didConnectToAddress \(UDPNetwork.host):\(UDPNetwork.port)")
    }
    
    func udpSocket(_ sock: GCDAsyncUdpSocket, didNotConnect error: Error?) {
        print("UDP didNotConnect \(UDPNetwork.host):\(UDPNetwork.port)")
    }
    
    func udpSocket(_ sock: GCDAsyncUdpSocket, didSendDataWithTag tag: Int) {
        print("UDP didSendDataWithTag \(tag)")
    }
    
    func udpSocket(_ sock: GCDAsyncUdpSocket, didNotSendDataWithTag tag: Int, dueToError error: Error?) {
        print("UDP didNotSendDataWithTag \(tag). Error: \(String(describing: error?.localizedDescription))")
    }
    
    func udpSocket(
        _ sock: GCDAsyncUdpSocket, didReceive data: Data, fromAddress address: Data, withFilterContext filterContext: Any?) {
        let saddr = String.init(data: address, encoding: String.Encoding.utf8)
        print("UDP didReceive from address \(String(describing: saddr))")
    }
}
