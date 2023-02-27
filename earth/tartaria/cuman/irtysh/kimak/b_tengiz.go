package kimak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 田吉兹TengizBarony struct {
	feud.BaseBarony
}

var BTengiz田吉兹 feud.Barony = &田吉兹TengizBarony{}

func init() {
    f := BTengiz田吉兹.(*田吉兹TengizBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tengiz",
		TitleName: "田吉兹",
		TitleCode: "b_tengiz",
	}
}
