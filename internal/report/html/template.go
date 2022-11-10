package html

import (
	"fmt"
	"github.com/fdaines/arch-go/internal/common"
)

func htmlTemplate() string {
	template := `<!DOCTYPE html>
<html>
<head>
    <link rel="stylesheet" type="text/css" href="style.css">
<style>
%s
</style>
</head>
<body>

<h1>Arch-Go Verification Report</h1>

[RULE_LIST]
<br/>
[RULE_DETAILS]
<br/>
[UNCOVERED_PACKAGES]
<br/>
<hr/>
Report generated by <a href='http://arch-go.org'>Arch-Go</a> v%s
</body>
</html>`

	return fmt.Sprintf(template, cssStyle, common.Version)
}
