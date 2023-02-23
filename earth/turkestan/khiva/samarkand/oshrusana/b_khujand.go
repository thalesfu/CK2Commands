package oshrusana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苦盏KhujandBarony struct {
	feud.BaseBarony
}

var BKhujand苦盏 feud.Barony = &苦盏KhujandBarony{}

func init() {
	f := BKhujand苦盏.(*苦盏KhujandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khujand",
		TitleName: "苦盏",
		TitleCode: "b_khujand",
	}
}
