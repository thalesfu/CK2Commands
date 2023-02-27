package kotthasara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伐多耆厘VatagiriBarony struct {
	feud.BaseBarony
}

var BVatagiri伐多耆厘 feud.Barony = &伐多耆厘VatagiriBarony{}

func init() {
    f := BVatagiri伐多耆厘.(*伐多耆厘VatagiriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vatagiri",
		TitleName: "伐多耆厘",
		TitleCode: "b_vatagiri",
	}
}
