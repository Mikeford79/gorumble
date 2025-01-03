package main

import (
    "flag"
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/Mikeford79/gorumble/botgen"
    "github.com/eiannone/keyboard"
    "github.com/andlabs/ui"
    _ "github.com/andlabs/ui/winmanifest"
)

var (
    targetCount int
    realViewers int
    viewerIDs   map[string]string
    videoID     string
    verbose     bool
)

func banner() {
    fmt.Println(`
            ┳┓         ┓    
            ┣┫┓┏┏┳┓┣┓┏┓╋
            ┛┗┗┻┛┗┗┗┛┗┛┗ V3
                    Forked from Ryuku ^_^
                `)
}

func manageViewers(targetCount int) {
    currentCount := len(viewerIDs)
    diff := targetCount - currentCount

    if diff > 0 {
        newUserAgents := botgen.GenerateUserAgents(diff)
        for _, ua := range newUserAgents {
            viewerIDs[ua] = ua // Using user agent as viewer ID for simplicity
            botgen.Viewbot(map[string]string{ua: ua}, videoID, verbose)
        }
    } else if diff < 0 {
        for id := range viewerIDs {
            if targetCount <= 0 {
                break
            }
            delete(viewerIDs, id)
            targetCount--
        }
    }
}

func updateLabel(label *ui.Label) {
    label.SetText(fmt.Sprintf("Total viewers: %d (Real: %d, Bot: %d)", targetCount+realViewers, realViewers, len(viewerIDs)))
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

    // Initialize GUI
    err = ui.Main(func() {
        window := ui.NewWindow("Viewbot Controller", 300, 100, true)
        label := ui.NewLabel("Total viewers: 0 (Real: 0, Bot: 0)")

        upButton := ui.NewButton("Add Bot")
        downButton := ui.NewButton("Remove Bot")

        upButton.OnClicked(func(*ui.Button) {
            targetCount++
            manageViewers(targetCount)
            updateLabel(label)
        })

        downButton.OnClicked(func(*ui.Button) {
            if targetCount > 0 {
                targetCount--
                manageViewers(targetCount)
                updateLabel(label)
            }
        })

        box := ui.NewVerticalBox()
        box.Append(label, false)
        box.Append(upButton, false)
        box.Append(downButton, false)
        window.SetChild(box)

        // Update label initially
        updateLabel(label)

        window.OnClosing(func(*ui.Window) bool {
            ui.Quit()
            return true
        })

        window.Show()
    })
    if err != nil {
        panic(err)
    }

    // Simulate real viewers count change (for demonstration purposes)
    go func() {
        for {
            time.Sleep(10 * time.Second)
            realViewers = 3 // Change this value to simulate actual real viewers
        }
    }()

    // Handle keyboard inputs for adding/removing bots
    go func() {
        if err := keyboard.Open(); err != nil {
            fmt.Println("Failed to open keyboard: ", err)
            return
        }
        defer keyboard.Close()

        for {
            select {
            case <-sigChan:
                fmt.Println("Shutting down gracefully...")
                return
            default:
                // Handle key press events
                if key, _, err := keyboard.GetKey(); err == nil {
                    switch key {
                    case '↑': // KeyArrowUp
                        targetCount++
                    case '↓': // KeyArrowDown
                        if targetCount > 0 {
                            targetCount--
                        }
                    }
                    manageViewers(targetCount)
                    updateLabel(nil) // Update GUI label if needed
                }
            }
        }
    }()

    select {} // Keep the program running
}
