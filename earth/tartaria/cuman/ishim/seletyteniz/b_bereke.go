package seletyteniz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别列克BerekeBarony struct {
	feud.BaseBarony
}

var BBereke别列克 feud.Barony = &别列克BerekeBarony{}

func init() {
    f := BBereke别列克.(*别列克BerekeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bereke",
		TitleName: "别列克",
		TitleCode: "b_bereke",
	}
}
