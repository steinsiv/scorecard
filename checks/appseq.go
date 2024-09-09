package checks

import (
	"github.com/ossf/scorecard/v5/checker"
	"github.com/ossf/scorecard/v5/checks/evaluation"
	"github.com/ossf/scorecard/v5/checks/raw"
	sce "github.com/ossf/scorecard/v5/errors"
	"github.com/ossf/scorecard/v5/probes"
	"github.com/ossf/scorecard/v5/probes/zrunner"
)

const CheckAppseq string = "Appseq"

//nolint:gochecknoinits
func init() {
	supportedRequestTypes := []checker.RequestType{
		checker.FileBased,
		checker.CommitBased,
	}
	if err := registerCheck(CheckAppseq, Appseq, supportedRequestTypes); err != nil {
		// this should never happen
		panic(err)
	}
}

// Appseq  will check the repository contains binary artifacts.
func Appseq(c *checker.CheckRequest) checker.CheckResult {
	rawData, err := raw.Appseq(c, "All Good!")
	if err != nil {
		e := sce.WithMessage(sce.ErrScorecardInternal, err.Error())
		return checker.CreateRuntimeErrorResult(CheckAppseq, e)
	}

	// Set the raw results.
	pRawResults := getRawResults(c)
	pRawResults.AppseqResults = rawData

	// Evaluate the probes.
	findings, err := zrunner.Run(pRawResults, probes.Appseq)
	if err != nil {
		e := sce.WithMessage(sce.ErrScorecardInternal, err.Error())
		return checker.CreateRuntimeErrorResult(CheckAppseq, e)
	}

	ret := evaluation.Appseq(CheckAppseq, findings, c.Dlogger)
	ret.Findings = findings
	return ret
}
