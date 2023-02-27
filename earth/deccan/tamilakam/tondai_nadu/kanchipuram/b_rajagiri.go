package kanchipuram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗阇耆厘RajagiriBarony struct {
	feud.BaseBarony
}

var BRajagiri罗阇耆厘 feud.Barony = &罗阇耆厘RajagiriBarony{}

func init() {
    f := BRajagiri罗阇耆厘.(*罗阇耆厘RajagiriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rajagiri",
		TitleName: "罗阇耆厘",
		TitleCode: "b_rajagiri",
	}
}
