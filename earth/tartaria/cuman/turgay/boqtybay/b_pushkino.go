package boqtybay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普希金诺PushkinoBarony struct {
	feud.BaseBarony
}

var BPushkino普希金诺 feud.Barony = &普希金诺PushkinoBarony{}

func init() {
    f := BPushkino普希金诺.(*普希金诺PushkinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pushkino",
		TitleName: "普希金诺",
		TitleCode: "b_pushkino",
	}
}
