package kangly

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 田吉兹Tengiz_kanglyBarony struct {
	feud.BaseBarony
}

var BTengiz_kangly田吉兹 feud.Barony = &田吉兹Tengiz_kanglyBarony{}

func init() {
    f := BTengiz_kangly田吉兹.(*田吉兹Tengiz_kanglyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tengiz_kangly",
		TitleName: "田吉兹",
		TitleCode: "b_tengiz_kangly",
	}
}
