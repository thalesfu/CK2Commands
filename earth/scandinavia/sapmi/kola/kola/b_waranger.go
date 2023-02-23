package kola

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦朗格WarangerBarony struct {
	feud.BaseBarony
}

var BWaranger瓦朗格 feud.Barony = &瓦朗格WarangerBarony{}

func init() {
	f := BWaranger瓦朗格.(*瓦朗格WarangerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "waranger",
		TitleName: "瓦朗格",
		TitleCode: "b_waranger",
	}
}
