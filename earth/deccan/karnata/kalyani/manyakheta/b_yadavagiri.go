package manyakheta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶陀婆耆厘YadavagiriBarony struct {
	feud.BaseBarony
}

var BYadavagiri耶陀婆耆厘 feud.Barony = &耶陀婆耆厘YadavagiriBarony{}

func init() {
    f := BYadavagiri耶陀婆耆厘.(*耶陀婆耆厘YadavagiriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yadavagiri",
		TitleName: "耶陀婆耆厘",
		TitleCode: "b_yadavagiri",
	}
}
