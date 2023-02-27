package aror

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 查克ChakBarony struct {
	feud.BaseBarony
}

var BChak查克 feud.Barony = &查克ChakBarony{}

func init() {
    f := BChak查克.(*查克ChakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chak",
		TitleName: "查克",
		TitleCode: "b_chak",
	}
}
