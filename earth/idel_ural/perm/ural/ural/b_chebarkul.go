package ural

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切巴尔库尔ChebarkulBarony struct {
	feud.BaseBarony
}

var BChebarkul切巴尔库尔 feud.Barony = &切巴尔库尔ChebarkulBarony{}

func init() {
    f := BChebarkul切巴尔库尔.(*切巴尔库尔ChebarkulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chebarkul",
		TitleName: "切巴尔库尔",
		TitleCode: "b_chebarkul",
	}
}
