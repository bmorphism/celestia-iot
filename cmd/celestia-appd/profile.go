package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"

	"github.com/spf13/cobra"
	tmrand "github.com/tendermint/tendermint/libs/rand"
)

const (
	profileFlag         = "profile"
	profileDurationFlag = "profile-duration"
)

func addProfileFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().Bool(profileFlag, false, "determine wether to profile or not")
	cmd.PersistentFlags().Int(profileDurationFlag, 300, "duration to profile the cpu before saving the profile file and restarting")
}

func profileConstantly(cmd *cobra.Command) {
	ctx := cmd.Context()

	p, err := cmd.Flags().GetBool(profileFlag)
	if err != nil {
		log.Fatal("failure to get profile flag", err)
	}

	if !p {
		return
	}

	dur, err := cmd.Flags().GetInt(profileDurationFlag)
	if err != nil {
		log.Fatal(err)
	}

	t := time.Second * time.Duration(dur)

	go func(ctx context.Context, t time.Duration) {
		cursor := 0
		session := tmrand.Str(4)
		cont := true
		for cont {
			filename := fmt.Sprintf("%s-%d-%d.prof", session, cursor, time.Now().Unix())
			fmt.Println(filename, "-------- profile started -----------------------------------------------------------")
			cont = profile(ctx, filename, t)
			fmt.Println(filename, "-------- profile finished ----------------------------------------------------------")
			cursor++
		}
	}(ctx, t)

}

func profile(ctx context.Context, filename string, dur time.Duration) bool {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal("failed to open profile file", err)
	}
	defer closeProfile(file)
	pprof.StartCPUProfile(file)
	select {
	case <-ctx.Done():
		return false
	case <-time.After(dur):
		return true
	}
}

func closeProfile(file *os.File) {
	pprof.StopCPUProfile()
	file.Close()
}
