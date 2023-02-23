package fejer

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙尔博加德SarbogardBarony struct {
	feud.BaseBarony
}

var BSarbogard沙尔博加德 feud.Barony = &沙尔博加德SarbogardBarony{}

func init() {
	f := BSarbogard沙尔博加德.(*沙尔博加德SarbogardBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarbogard",
		TitleName: "沙尔博加德",
		TitleCode: "b_sarbogard",
	}
}
