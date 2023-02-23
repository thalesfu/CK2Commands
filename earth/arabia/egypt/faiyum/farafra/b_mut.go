package farafra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆特MutBarony struct {
	feud.BaseBarony
}

var BMut穆特 feud.Barony = &穆特MutBarony{}

func init() {
	f := BMut穆特.(*穆特MutBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mut",
		TitleName: "穆特",
		TitleCode: "b_mut",
	}
}
