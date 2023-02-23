package qayaliq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙甘ShaganBarony struct {
	feud.BaseBarony
}

var BShagan沙甘 feud.Barony = &沙甘ShaganBarony{}

func init() {
	f := BShagan沙甘.(*沙甘ShaganBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shagan",
		TitleName: "沙甘",
		TitleCode: "b_shagan",
	}
}
