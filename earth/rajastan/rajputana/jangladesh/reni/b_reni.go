package reni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梨尼ReniBarony struct {
	feud.BaseBarony
}

var BReni梨尼 feud.Barony = &梨尼ReniBarony{}

func init() {
    f := BReni梨尼.(*梨尼ReniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "reni",
		TitleName: "梨尼",
		TitleCode: "b_reni",
	}
}
