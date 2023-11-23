package groth16

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type ProverTimeBReakdown struct {
	Total           time.Duration
	WitnesReduction time.Duration
	ComputeDetla    time.Duration
	ComputeAR1      time.Duration
	ComputeBS1      time.Duration
	ComputeKRS      time.Duration
	ComputeBS2      time.Duration
}

func ProcessJson(Input ProverTimeBReakdown, system string) {

	currentTime := time.Now()
	outputpath := fmt.Sprint(currentTime.Unix()) + system + "prover_timebreakdown.json"
	jsonfile, _ := json.MarshalIndent(Input, "", " ")
	_ = os.WriteFile(outputpath, jsonfile, 0644)
	fmt.Println("Prover breakdown saved in", outputpath)
}
