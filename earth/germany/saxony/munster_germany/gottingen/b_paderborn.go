package gottingen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕德博恩PaderbornBarony struct {
	feud.BaseBarony
}

var BPaderborn帕德博恩 feud.Barony = &帕德博恩PaderbornBarony{}

func init() {
    f := BPaderborn帕德博恩.(*帕德博恩PaderbornBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "paderborn",
		TitleName: "帕德博恩",
		TitleCode: "b_paderborn",
	}
}
