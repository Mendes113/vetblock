package network

import (
	"bytes"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"encoding/json"
	"vetblock/internal/blockchain"

)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

type Connection struct {
	Conn *websocket.Conn
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Erro ao estabelecer conexão WebSocket:", err)
		return
	}
	c := &Connection{Conn: conn}
	go handleMessages(c)
}

func handleMessages(conn *Connection) {
	for {
		_, msg, err := conn.Conn.ReadMessage()
		if err != nil {
			log.Println("Erro ao ler mensagem:", err)
			break
		}
		// Aqui você pode processar a mensagem recebida, como um bloco ou transação.
		log.Printf("Recebida mensagem: %s", msg)
		
		// Exemplo de como processar a mensagem
		var block blockchain.Block
		if err := json.Unmarshal(msg, &block); err != nil {
			log.Println("Erro ao decodificar mensagem como bloco:", err)
			continue
		}
		
		// Adicione o bloco à blockchain, valide-o, etc.
		// blockchain.AddBlock(block)
	}
}
func StartServer() {
	http.HandleFunc("/ws", handleConnection)
	go func() {
		log.Println("Servidor WebSocket iniciado na porta 8081...")
		log.Fatal(http.ListenAndServe(":8081", nil))
	}()
}

func propagateBlock(block blockchain.Block) {
    nodes := []string{"http://localhost:8081"} // Adicione a lista de nós
    for _, node := range nodes {
        resp, err := http.Post(node+"/blocks", "application/json", bytes.NewBuffer(blockToJSON(block)))
        if err != nil {
            log.Printf("Erro ao enviar bloco para %s: %v", node, err)
            continue
        }
        resp.Body.Close()
    }
}

func blockToJSON(block blockchain.Block) []byte {
    jsonData, _ := json.Marshal(block)
    return jsonData
}


func StartClient() {
    conn, err := ConnectToNode("ws://localhost:8081/ws")
    if err != nil {
        log.Fatalf("Erro ao conectar-se ao nó: %v", err)
    }
    defer conn.Close()

    message := []byte("Hello from new node!")
    err = conn.WriteMessage(websocket.TextMessage, message)
    if err != nil {
        log.Printf("Erro ao enviar mensagem: %v", err)
    }

    _, msg, err := conn.ReadMessage()
    if err != nil {
        log.Printf("Erro ao ler mensagem: %v", err)
    } else {
        log.Printf("Mensagem recebida: %s", msg)
    }
}


func ConnectToNode(url string) (*websocket.Conn, error) {
    conn, _, err := websocket.DefaultDialer.Dial(url, nil)
    if err != nil {
        return nil, err
    }
    return conn, nil
}