package kaneia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷西姆诺RethymnoBarony struct {
	feud.BaseBarony
}

var BRethymno雷西姆诺 feud.Barony = &雷西姆诺RethymnoBarony{}

func init() {
	f := BRethymno雷西姆诺.(*雷西姆诺RethymnoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rethymno",
		TitleName: "雷西姆诺",
		TitleCode: "b_rethymno",
	}
}
