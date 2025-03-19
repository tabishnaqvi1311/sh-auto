package main

import (
	"flag"
    "fmt"
)

type CmdFlags struct {
    Init string 
}

func NewCmdFlags() *CmdFlags {
    cf := &CmdFlags{}

    flag.StringVar(&cf.Init, "init", "", "type in your build command")

    flag.Parse()
    return cf
}

func (cf *CmdFlags) Exec() {
    if cf.Init != "" {
        WriteToSh(cf.Init)
    } else {
        fmt.Println("invalid command sent")
    }
}

