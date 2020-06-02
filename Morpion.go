package main

import (
    "strconv"
    "github.com/tadvi/winc"
)

func main() {
    mainWindow := winc.NewForm(nil)
    mainWindow.SetSize(400, 300) 
    mainWindow.SetText("TpMorpion")

    var player = "X"
    var Case = 0
    var CaseCocher = 0
    var PosX = 50
    var PosY = 50
    var m map[string]string
    m = make(map[string]string)

    for x := 0; x < 3; x++ {
        for y := 0; y < 3; y++ {
            Case = Case + 1
            Button := winc.NewPushButton(mainWindow)
            Button.SetSize(50, 50)
            Button.SetPos(PosX+x*50, PosY+y*50)
            name := strconv.Itoa(Case)
            Button.SetText(name)

            Button.OnClick().Bind(func(e *winc.Event) {
                if m[name] == "" {
                    CaseCocher++
                    Button.SetText(player)
                    m[name] = player

                    if player == "X" {
                        player = "O"
                    } else {
                        player = "X"
                    }

                    
                }
            })
        }
    }

    mainWindow.Center()
    mainWindow.Show()
    mainWindow.OnClose().Bind(wndOnClose)

    winc.RunMainLoop()
}

func wndOnClose(arg *winc.Event) {
    winc.Exit()
}
