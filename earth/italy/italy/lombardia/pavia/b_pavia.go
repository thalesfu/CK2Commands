package pavia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕维亚PaviaBarony struct {
	feud.BaseBarony
}

var BPavia帕维亚 feud.Barony = &帕维亚PaviaBarony{}

func init() {
	f := BPavia帕维亚.(*帕维亚PaviaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pavia",
		TitleName: "帕维亚",
		TitleCode: "b_pavia",
	}
}
