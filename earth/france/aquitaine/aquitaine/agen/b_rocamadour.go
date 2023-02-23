package agen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗卡马杜尔RocamadourBarony struct {
	feud.BaseBarony
}

var BRocamadour罗卡马杜尔 feud.Barony = &罗卡马杜尔RocamadourBarony{}

func init() {
	f := BRocamadour罗卡马杜尔.(*罗卡马杜尔RocamadourBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rocamadour",
		TitleName: "罗卡马杜尔",
		TitleCode: "b_rocamadour",
	}
}
