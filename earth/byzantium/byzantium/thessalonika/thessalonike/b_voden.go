package thessalonike

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃泽纳VodenBarony struct {
	feud.BaseBarony
}

var BVoden沃泽纳 feud.Barony = &沃泽纳VodenBarony{}

func init() {
	f := BVoden沃泽纳.(*沃泽纳VodenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "voden",
		TitleName: "沃泽纳",
		TitleCode: "b_voden",
	}
}
