package kufa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈马尔HammarBarony struct {
	feud.BaseBarony
}

var BHammar哈马尔 feud.Barony = &哈马尔HammarBarony{}

func init() {
    f := BHammar哈马尔.(*哈马尔HammarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hammar",
		TitleName: "哈马尔",
		TitleCode: "b_hammar",
	}
}
