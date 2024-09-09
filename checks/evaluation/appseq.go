package evaluation

import (
	"fmt"

	"github.com/ossf/scorecard/v5/checker"
	sce "github.com/ossf/scorecard/v5/errors"
	"github.com/ossf/scorecard/v5/finding"
	"github.com/ossf/scorecard/v5/probes/appseqPresent"
)

// ReadmeCheck applies the score policy for the README.md check.
func Appseq(name string, findings []finding.Finding, dl checker.DetailLogger) checker.CheckResult {
	expectedProbes := []string{
		appseqPresent.Probe,
	}
	if !finding.UniqueProbesEqual(findings, expectedProbes) {
		e := sce.WithMessage(sce.ErrScorecardInternal, "invalid probe results")
		return checker.CreateRuntimeErrorResult(name, e)
	}
	//dump the findings t console
	fmt.Println(findings)

	score := 0
	m := make(map[string]bool)
	var logLevel checker.DetailType
	for i := range findings {
		f := &findings[i]
		fmt.Println("A " + f.Outcome)
		fmt.Println("B " + findings[i].Probe)
		logLevel = checker.DetailInfo
		switch f.Outcome {
			case finding.OutcomeTrue:
				m[f.Probe] = true
				score += scoreProbeOnce(f.Probe, m, 10)			
			case finding.OutcomeFalse:
				m[f.Probe] = false
				score += scoreProbeOnce(f.Probe, m, 0)			
			default:
				e := sce.WithMessage(sce.ErrScorecardInternal, "unknown probe results")
				return checker.CreateRuntimeErrorResult(name, e)
			}
		logLevel = checker.DetailDebug
		checker.LogFinding(dl, f, logLevel)
	}
	if !m[appseqPresent.Probe] {
		return checker.CreateMinScoreResult(name, "APPSEC.md file not detected")
	}		
	
	return checker.CreateResultWithScore(name, "APPSEC.md file detected", score)
}
