package main

import (
    "flag"
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/itsryuku/gorumbot/botgen"
    "github.com/eiannone/keyboard"
)

func banner() {
    fmt.Println(`
            ┳┓         ┓    
            ┣┫┓┏┏┳┓┣┓┏┓╋
            ┛┗┗┻┛┗┗┗┛┗┛┗ V3
                    Forked from Ryuku ^_^
                `)
}

func manageViewers(viewerIDs *map[string]string, videoID string, targetCount int, verbose bool) {
    currentCount := len(*viewerIDs)
    diff := targetCount - currentCount

    if diff > 0 {
        newUserAgents := botgen.GenerateUserAgents(diff)
        for _, ua := range newUserAgents {
            (*viewerIDs)[ua] = ua // Using user agent as viewer ID for simplicity
            botgen.Viewbot(map[string]string{ua: ua}, videoID, verbose)
        }
    } else if diff < 0 {
        for id := range *viewerIDs {
            if targetCount <= 0 {
                break
            }
            delete(*viewerIDs, id)
            targetCount--
        }
    }
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
    videoID, err := botgen.ExtractVideoID(*urlFlag)
    if err != nil {
        fmt.Println(err)
        return
    }

    viewerIDs, extractedVideoID, channelName := botgen.GetViewerIds(videoID, *botsFlag)
    botgen.Viewbot(viewerIDs, extractedVideoID, *verboseFlag)
    fmt.Println("(+) Viewbotting Channel:", channelName)
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

    targetCount := *botsFlag
    realViewers := 0 // Simulating real viewers count

    // Simulate real viewers count change (for demonstration purposes)
    go func() {
        for {
            time.Sleep(10 * time.Second)
            realViewers = 3 // Change this value to simulate actual real viewers
        }
    }()

    for {
        select {
        case <-sigChan:
            fmt.Println("Shutting down gracefully...")
            return
        default:
            // Handle key press events
            if key, _, err := keyboard.GetKey(); err == nil {
                switch key {
                case keyboard.KeyArrowUp:
                    targetCount++
                case keyboard.KeyArrowDown:
                    if targetCount > 0 {
                        targetCount--
                    }
                }
                manageViewers(&viewerIDs, extractedVideoID, targetCount, *verboseFlag)
                fmt.Printf("Total viewers: %d (Real: %d, Bot: %d)\n", targetCount+realViewers, realViewers, len(viewerIDs))
            }
        }
    }
}
