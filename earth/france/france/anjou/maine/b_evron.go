package maine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃夫龙EvronBarony struct {
	feud.BaseBarony
}

var BEvron埃夫龙 feud.Barony = &埃夫龙EvronBarony{}

func init() {
	f := BEvron埃夫龙.(*埃夫龙EvronBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "evron",
		TitleName: "埃夫龙",
		TitleCode: "b_evron",
	}
}
