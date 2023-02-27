package gorlitz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格尔利茨GorlitzBarony struct {
	feud.BaseBarony
}

var BGorlitz格尔利茨 feud.Barony = &格尔利茨GorlitzBarony{}

func init() {
    f := BGorlitz格尔利茨.(*格尔利茨GorlitzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gorlitz",
		TitleName: "格尔利茨",
		TitleCode: "b_gorlitz",
	}
}
