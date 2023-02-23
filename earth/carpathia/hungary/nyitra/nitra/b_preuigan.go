package nitra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普列维甘PreuiganBarony struct {
	feud.BaseBarony
}

var BPreuigan普列维甘 feud.Barony = &普列维甘PreuiganBarony{}

func init() {
	f := BPreuigan普列维甘.(*普列维甘PreuiganBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "preuigan",
		TitleName: "普列维甘",
		TitleCode: "b_preuigan",
	}
}
