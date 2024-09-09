https://github.com/ossf/scorecard/blob/main/checks/write.md


1. Define the check
/docs/checks/internal/checks.yml

2. 
checks/mycheck.go
  The main check, refers to evaluation and raw
  checker.checkrequest -> checker.checkresult

3. Check runs the raw request, gather data and extract into raw data
checks/raw/mycheck.go
  implementation of the request and populates the raw data for the request

checker/raw_result.go
  Definition of raw results (Data) for the check
  ```go
	  type AppseqData struct {
			Desc     string
			IsGood   bool
	  }
  ```
  probes/entries.go
    Declare probes for check
  probes/
    Operate on raw data from request and prepare findings as results

4. Evaluate results from all probes
checks/evaluation/mycheck.go



bare minimum
```go

```
