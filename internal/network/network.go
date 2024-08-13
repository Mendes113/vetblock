package network

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"vetblock/internal/blockchain"

	"github.com/gorilla/websocket"
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

	// Envia a blockchain completa para o cliente assim que ele se conecta
	err = sendBlockchain(c)
	if err != nil {
		log.Println("Erro ao enviar a blockchain para o cliente:", err)
	}

	// Mantém a conexão aberta e escuta por mensagens do cliente
	go handleMessages(c)
}

func sendBlockchain(conn *Connection) error {
	// Aqui você recupera a blockchain completa
	chain := blockchain.GetBlockchain() // Supondo que GetBlockchain() retorna a blockchain completa

	// Converte a blockchain para JSON
	chainData, err := json.Marshal(chain)
	if err != nil {
		return err
	}

	// Envia a blockchain para o cliente
	err = conn.Conn.WriteMessage(websocket.TextMessage, chainData)
	if err != nil {
		return err
	}

	return nil
}

func handleMessages(conn *Connection) {
	defer conn.Conn.Close() // Fecha a conexão quando a função termina

	for {
		_, msg, err := conn.Conn.ReadMessage()
		if err != nil {
			log.Println("Erro ao ler mensagem:", err)
			break
		}

		// Processa a mensagem recebida
		log.Printf("Recebida mensagem: %s", msg)

		// Exemplo de processamento de um bloco recebido
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

	// Mantém a conexão aberta para continuar recebendo mensagens
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Erro ao ler mensagem: %v", err)
			break
		}

		log.Printf("Mensagem recebida: %s", msg)

		// Exemplo de como processar a blockchain recebida
		var chain []blockchain.Block
		if err := json.Unmarshal(msg, &chain); err != nil {
			log.Println("Erro ao decodificar a blockchain:", err)
			continue
		}

		// Processar a blockchain, por exemplo, armazenar ou validar
	}
}

func ConnectToNode(url string) (*websocket.Conn, error) {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
