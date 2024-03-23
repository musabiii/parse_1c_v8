package parse_1c_v8

import (
	"bufio"
	"fmt"
	"os"
	"os/user" // for current user
	"strings"
)

type Connection struct {
	Connect                  string `json:"connect"`
	Name                     string `json:"name"`
	ID                       string `json:"id"`
	OrderInList              int    `json:"order_in_list"`
	Folder                   string `json:"folder"`
	OrderInTree              int    `json:"order_in_tree"`
	External                 int    `json:"external"`
	ClientConnectionSpeed    string `json:"client_connection_speed,omitempty"`
	App                      string `json:"app,omitempty"`
	WA                       int    `json:"wa,omitempty"`
	Version                  string `json:"version,omitempty"`
	DisableLocalSpeechToText int    `json:"disable_local_speech_to_text,omitempty"`
	DefaultVersion           string `json:"default_version,omitempty"`
	DefaultApp               string `json:"default_app,omitempty"`
}

func main() {

	connections := getConnections() // Get the connections
	foldersMap := getFoldersMap(connections)

	// Convert the unique folders map keys to a list

	for k, v := range foldersMap {
		fmt.Println("folder", k)
		for _, vv := range v {
			fmt.Println("\t" + vv.Name)
		}

	}

	// fmt.Printf("connections: %v\n", connections)

	// // Convert connections slice to JSON
	// jsonData, err := json.MarshalIndent(connections, "", "  ")
	// if err != nil {
	// 	fmt.Println("Error marshalling to JSON:", err)
	// 	return
	// }

	// Print JSON data
	// fmt.Println(string(jsonData))
	// fmt.Printf("uniqueFoldersList: %v\n", uniqueFoldersList)
}

func getFoldersMap(connections []Connection) map[string][]Connection {

	foldersMap := map[string][]Connection{}

	for _, connection := range connections {
		if connection.Folder != "" {
			foldersMap[connection.Folder] = append(foldersMap[connection.Folder], connection)
		}
	}
	return foldersMap
}

func getConnections() []Connection {
	foldersMap := make(map[string][]Connection)

	var currentConnection Connection
	var connections []Connection

	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error:", err)
		return connections
	}

	// Open the file
	file, err := os.Open(currentUser.HomeDir + "\\AppData\\Roaming\\1C\\1CEStart\\ibases.v8i")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return connections
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Create a map to store unique folders

	// Iterate over each line
	for scanner.Scan() {
		line := scanner.Text()

		// Check for section headers
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			if currentConnection.ID != "" {
				connections = append(connections, currentConnection)
				foldersMap[currentConnection.Folder] = append(foldersMap[currentConnection.Folder], currentConnection)
			}
			currentConnection = Connection{}
			line = "Name=" + line[1:len(line)-1]
			// continue
		}

		// Parse key-value pairs
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "Connect":
			currentConnection.Connect = value
		case "ID":
			currentConnection.ID = value
		case "Name":
			currentConnection.Name = value
		case "OrderInList":
			fmt.Sscanf(value, "%d", &currentConnection.OrderInList)
		case "Folder":
			currentConnection.Folder = value
		case "OrderInTree":
			fmt.Sscanf(value, "%d", &currentConnection.OrderInTree)
		case "External":
			fmt.Sscanf(value, "%d", &currentConnection.External)
		case "ClientConnectionSpeed":
			currentConnection.ClientConnectionSpeed = value
		case "App":
			currentConnection.App = value
		case "WA":
			fmt.Sscanf(value, "%d", &currentConnection.WA)
		case "Version":
			currentConnection.Version = value
		case "DisableLocalSpeechToText":
			fmt.Sscanf(value, "%d", &currentConnection.DisableLocalSpeechToText)
		case "DefaultVersion":
			currentConnection.DefaultVersion = value
		case "DefaultApp":
			currentConnection.DefaultApp = value
		}
	}

	if currentConnection.ID != "" {
		connections = append(connections, currentConnection)
		foldersMap[currentConnection.Folder] = append(foldersMap[currentConnection.Folder], currentConnection)
	}

	return connections

}

func SayHi() {
	fmt.Println("Hello, World!")
}
