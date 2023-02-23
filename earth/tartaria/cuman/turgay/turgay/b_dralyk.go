package turgay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德拉雷克DralykBarony struct {
	feud.BaseBarony
}

var BDralyk德拉雷克 feud.Barony = &德拉雷克DralykBarony{}

func init() {
	f := BDralyk德拉雷克.(*德拉雷克DralykBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dralyk",
		TitleName: "德拉雷克",
		TitleCode: "b_dralyk",
	}
}
