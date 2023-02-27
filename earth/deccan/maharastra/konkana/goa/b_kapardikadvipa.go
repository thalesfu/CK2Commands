package goa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦钵地迦提鞞波KapardikadvipaBarony struct {
	feud.BaseBarony
}

var BKapardikadvipa迦钵地迦提鞞波 feud.Barony = &迦钵地迦提鞞波KapardikadvipaBarony{}

func init() {
    f := BKapardikadvipa迦钵地迦提鞞波.(*迦钵地迦提鞞波KapardikadvipaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kapardikadvipa",
		TitleName: "迦钵地迦提鞞波",
		TitleCode: "b_kapardikadvipa",
	}
}
