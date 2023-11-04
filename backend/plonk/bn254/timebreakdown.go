package plonk

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type ProverTimeBreakdown struct {
	Total                    time.Duration
	ConstraintSolving        time.Duration
	PrepNumeratorData        time.Duration
	CompleteQK               time.Duration
	InitiateBlindingPoly     time.Duration
	DeriveBetaGamma          time.Duration
	ComputeAccumulatingRatio time.Duration
	ComputeH                 time.Duration
	OpenZ                    time.Duration
	FoldComH                 time.Duration
	ComputeLinearizedPoly    time.Duration
	BatchOpen                time.Duration
}

func ProcessJson(Input ProverTimeBreakdown, system string) {
	currentTime := time.Now()
	outputpath := fmt.Sprint(currentTime.Unix()) + system + "prover_timebreakdown.json"
	jsonfile, _ := json.MarshalIndent(Input, "", " ")
	_ = os.WriteFile(outputpath, jsonfile, 0644)
	fmt.Println("Prover breakdown saved in", outputpath)
}
