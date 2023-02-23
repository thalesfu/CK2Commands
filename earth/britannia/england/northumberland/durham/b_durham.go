package durham

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达勒姆DurhamBarony struct {
	feud.BaseBarony
}

var BDurham达勒姆 feud.Barony = &达勒姆DurhamBarony{}

func init() {
	f := BDurham达勒姆.(*达勒姆DurhamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "durham",
		TitleName: "达勒姆",
		TitleCode: "b_durham",
	}
}
