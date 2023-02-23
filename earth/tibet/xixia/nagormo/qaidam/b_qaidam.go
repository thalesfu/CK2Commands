package qaidam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 柴旦QaidamBarony struct {
	feud.BaseBarony
}

var BQaidam柴旦 feud.Barony = &柴旦QaidamBarony{}

func init() {
	f := BQaidam柴旦.(*柴旦QaidamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qaidam",
		TitleName: "柴旦",
		TitleCode: "b_qaidam",
	}
}
