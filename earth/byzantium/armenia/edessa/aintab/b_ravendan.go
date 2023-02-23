package aintab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷文丹RavendanBarony struct {
	feud.BaseBarony
}

var BRavendan雷文丹 feud.Barony = &雷文丹RavendanBarony{}

func init() {
	f := BRavendan雷文丹.(*雷文丹RavendanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ravendan",
		TitleName: "雷文丹",
		TitleCode: "b_ravendan",
	}
}
