package tura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉木图AlmatyBarony struct {
	feud.BaseBarony
}

var BAlmaty阿拉木图 feud.Barony = &阿拉木图AlmatyBarony{}

func init() {
	f := BAlmaty阿拉木图.(*阿拉木图AlmatyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "almaty",
		TitleName: "阿拉木图",
		TitleCode: "b_almaty",
	}
}
