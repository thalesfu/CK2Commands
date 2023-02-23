package irbid

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩拉PellaBarony struct {
	feud.BaseBarony
}

var BPella佩拉 feud.Barony = &佩拉PellaBarony{}

func init() {
	f := BPella佩拉.(*佩拉PellaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pella",
		TitleName: "佩拉",
		TitleCode: "b_pella",
	}
}
