package quena

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基纳QuenaBarony struct {
	feud.BaseBarony
}

var BQuena基纳 feud.Barony = &基纳QuenaBarony{}

func init() {
    f := BQuena基纳.(*基纳QuenaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "quena",
		TitleName: "基纳",
		TitleCode: "b_quena",
	}
}
