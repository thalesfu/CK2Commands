package nandapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩罗醯耶MarhiaBarony struct {
	feud.BaseBarony
}

var BMarhia摩罗醯耶 feud.Barony = &摩罗醯耶MarhiaBarony{}

func init() {
	f := BMarhia摩罗醯耶.(*摩罗醯耶MarhiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marhia",
		TitleName: "摩罗醯耶",
		TitleCode: "b_marhia",
	}
}
