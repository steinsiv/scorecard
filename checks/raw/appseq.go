package raw

import (
	"github.com/ossf/scorecard/v5/checker"
	"github.com/ossf/scorecard/v5/checks/fileparser"
)

// Check if the filename is APPSEC.md.
var isAppseqFile fileparser.DoWhileTrueOnFilename = func(name string, args ...interface{}) (bool, error) {
	foundReadme, ok := args[0].(*bool)
	if !ok {
		return true, nil
	}
	if name == "APPSEC.md" {
		*foundReadme = true
		return false, nil // Stop looking once found
	}
	return true, nil
}

// ReadmeCheck checks for the presence of APPSEC.md.
func Appseq(c *checker.CheckRequest, expectedText string) (checker.AppseqData, error) {
	var foundAppseq bool

	// Look for APPSEC.md
	err := fileparser.OnAllFilesDo(c.RepoClient, isAppseqFile, &foundAppseq)
	if err != nil {
		return checker.AppseqData{}, err
	}

	return checker.AppseqData{FoundAppseq: foundAppseq}, nil
}
