package candhoba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯毗SipriBarony struct {
	feud.BaseBarony
}

var BSipri斯毗 feud.Barony = &斯毗SipriBarony{}

func init() {
    f := BSipri斯毗.(*斯毗SipriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sipri",
		TitleName: "斯毗",
		TitleCode: "b_sipri",
	}
}
