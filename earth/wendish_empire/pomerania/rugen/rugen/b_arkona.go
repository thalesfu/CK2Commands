package rugen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿科纳ArkonaBarony struct {
	feud.BaseBarony
}

var BArkona阿科纳 feud.Barony = &阿科纳ArkonaBarony{}

func init() {
    f := BArkona阿科纳.(*阿科纳ArkonaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arkona",
		TitleName: "阿科纳",
		TitleCode: "b_arkona",
	}
}
