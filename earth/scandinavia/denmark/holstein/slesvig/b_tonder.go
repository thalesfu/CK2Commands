package slesvig

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 岑讷TonderBarony struct {
	feud.BaseBarony
}

var BTonder岑讷 feud.Barony = &岑讷TonderBarony{}

func init() {
	f := BTonder岑讷.(*岑讷TonderBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tonder",
		TitleName: "岑讷",
		TitleCode: "b_tonder",
	}
}
