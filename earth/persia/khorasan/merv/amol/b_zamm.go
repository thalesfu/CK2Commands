package amol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎木木ZammBarony struct {
	feud.BaseBarony
}

var BZamm扎木木 feud.Barony = &扎木木ZammBarony{}

func init() {
	f := BZamm扎木木.(*扎木木ZammBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zamm",
		TitleName: "扎木木",
		TitleCode: "b_zamm",
	}
}
