package forde

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾于敦堡AudunborgBarony struct {
	feud.BaseBarony
}

var BAudunborg艾于敦堡 feud.Barony = &艾于敦堡AudunborgBarony{}

func init() {
	f := BAudunborg艾于敦堡.(*艾于敦堡AudunborgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "audunborg",
		TitleName: "艾于敦堡",
		TitleCode: "b_audunborg",
	}
}
