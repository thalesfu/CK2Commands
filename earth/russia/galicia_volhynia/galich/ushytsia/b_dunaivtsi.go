package ushytsia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜纳伊夫齐DunaivtsiBarony struct {
	feud.BaseBarony
}

var BDunaivtsi杜纳伊夫齐 feud.Barony = &杜纳伊夫齐DunaivtsiBarony{}

func init() {
    f := BDunaivtsi杜纳伊夫齐.(*杜纳伊夫齐DunaivtsiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dunaivtsi",
		TitleName: "杜纳伊夫齐",
		TitleCode: "b_dunaivtsi",
	}
}
