package appseqPresent

import (
	"embed"
	"fmt"

	"github.com/ossf/scorecard/v5/checker"
	"github.com/ossf/scorecard/v5/finding"
	"github.com/ossf/scorecard/v5/internal/checknames"
	"github.com/ossf/scorecard/v5/internal/probes"
	"github.com/ossf/scorecard/v5/probes/internal/utils/uerror"
)

func init() {
	probes.MustRegister(Probe, Run, []checknames.CheckName{checknames.Appseq})
}

//go:embed *.yml
var fs embed.FS

const Probe = "appseqPresent"

func Run(raw *checker.RawResults) ([]finding.Finding, string, error) {
	if raw == nil {
		return nil, "", fmt.Errorf("%w: raw", uerror.ErrNil)
	}

	// Determine the outcome based on the presence of Appseq
	outcome := finding.OutcomeFalse
	message := "APPSEC.md file not detected"
	if raw.AppseqResults.FoundAppseq {
		outcome = finding.OutcomeTrue
		message = "APPSEC.md file detected"
	}

	// Create the finding based on the determined outcome
	f, err := finding.NewWith(fs, Probe, message, nil, outcome)
	if err != nil {
		return nil, Probe, fmt.Errorf("create finding: %w", err)
	}

	// Return the findings list
	return []finding.Finding{*f}, Probe, nil
}