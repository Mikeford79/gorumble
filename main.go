package main

import (
    "flag"
    "fmt"
    "os"
    "os/signal"
    "syscall"

    "github.com/Mikeford79/gorumble/botgen"
    "github.com/eiannone/keyboard"
)

var (
    targetCount int
    realViewers int = 0 // Ensure realViewers is initialized to 0
    viewerIDs   map[string]string
    videoID     string
    verbose     bool
)

func banner() {
    fmt.Println(` ┳┓ ┓ ┣┫┓┏┏┳┓┣┓┏┓╋ ┛┗┗┻┛┗┗┗┛┗┛┗ V3 Forked from Ryuku ^_^ `)
}

func manageViewers(targetCount int) {
    currentCount := len(viewerIDs)
    diff := targetCount - currentCount
    fmt.Printf("Managing viewers. Target: %d, Current: %d, Diff: %d\n", targetCount, currentCount, diff)

    if diff > 0 {
        newUserAgents := botgen.GenerateUserAgents(diff)
        for _, ua := range newUserAgents {
            viewerIDs[ua] = ua // Using user agent as viewer ID for simplicity
            botgen.Viewbot(map[string]string{ua: ua}, videoID, verbose)
            fmt.Printf("Added bot: %s\n", ua)
        }
    } else if diff < 0 {
        for id := range viewerIDs {
            if targetCount <= 0 {
                break
            }
            delete(viewerIDs, id)
            fmt.Printf("Removed bot: %s\n", id)
            targetCount--
        }
    }
}

func updateStatus() {
    fmt.Printf("Total viewers: %d (Real: %d, Bot: %d)\n", targetCount+realViewers, realViewers, len(viewerIDs))
}

func main() {
    urlFlag := flag.String("u", "", "Video URL")
    botsFlag := flag.Int("b", 0, "Number of bots")
    verboseFlag := flag.Bool("v", false, "Verbose mode")
    flag.Parse()

    if *urlFlag == "" || *botsFlag == 0 {
        fmt.Println("usage: go run main.go -u <videoURL> -b <num> [-v]")
        fmt.Println("e.g: go run main.go -u <your-video-url-here> -b <number-of-bots>")
        return
    }

    banner()

    var err error
    videoID, err = botgen.ExtractVideoID(*urlFlag)
    if err != nil {
        fmt.Println(err)
        return
    }

    viewerIDs, _, _ = botgen.GetViewerIds(videoID, *botsFlag)
    botgen.Viewbot(viewerIDs, videoID, *verboseFlag)

    targetCount = *botsFlag
    verbose = *verboseFlag

    fmt.Println("(+) Viewbotting Channel")
    fmt.Println("(+) Click CTRL + C when you are done to exit.")

    // Handle graceful shutdown
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

    // Initialize keyboard input
    if err := keyboard.Open(); err != nil {
        fmt.Println("Failed to open keyboard: ", err)
        return
    }
    defer keyboard.Close()

    // Handle key press events
    for {
        select {
        case <-sigChan:
            fmt.Println("Shutting down gracefully...")
            return
        default:
            if key, _, err := keyboard.GetKey(); err == nil {
                fmt.Printf("Key pressed: %v\n", key) // Debugging: print key pressed
                switch key {
                case rune(keyboard.KeyArrowRight):
                    fmt.Println("Adding a bot") // Debugging: print action
                    targetCount++
                case rune(keyboard.KeyArrowLeft):
                    fmt.Println("Removing a bot") // Debugging: print action
                    if targetCount > 0 {
                        targetCount--
                    }
                }
                manageViewers(targetCount)
                updateStatus()
            }
        }
    }
}
