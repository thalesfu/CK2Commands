package silves

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿沃尔AlvorBarony struct {
	feud.BaseBarony
}

var BAlvor阿沃尔 feud.Barony = &阿沃尔AlvorBarony{}

func init() {
    f := BAlvor阿沃尔.(*阿沃尔AlvorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alvor",
		TitleName: "阿沃尔",
		TitleCode: "b_alvor",
	}
}
