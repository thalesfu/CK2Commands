package avhaz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊泽IdhajBarony struct {
	feud.BaseBarony
}

var BIdhaj伊泽 feud.Barony = &伊泽IdhajBarony{}

func init() {
	f := BIdhaj伊泽.(*伊泽IdhajBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "idhaj",
		TitleName: "伊泽",
		TitleCode: "b_idhaj",
	}
}
