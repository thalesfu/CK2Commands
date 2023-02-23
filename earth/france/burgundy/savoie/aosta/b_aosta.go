package aosta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥斯塔AostaBarony struct {
	feud.BaseBarony
}

var BAosta奥斯塔 feud.Barony = &奥斯塔AostaBarony{}

func init() {
	f := BAosta奥斯塔.(*奥斯塔AostaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aosta",
		TitleName: "奥斯塔",
		TitleCode: "b_aosta",
	}
}
