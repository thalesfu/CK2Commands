package tura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 铁克TekeBarony struct {
	feud.BaseBarony
}

var BTeke铁克 feud.Barony = &铁克TekeBarony{}

func init() {
	f := BTeke铁克.(*铁克TekeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "teke",
		TitleName: "铁克",
		TitleCode: "b_teke",
	}
}
