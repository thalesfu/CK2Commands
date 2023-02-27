package koshma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利斯特马ListmaBarony struct {
	feud.BaseBarony
}

var BListma利斯特马 feud.Barony = &利斯特马ListmaBarony{}

func init() {
    f := BListma利斯特马.(*利斯特马ListmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "listma",
		TitleName: "利斯特马",
		TitleCode: "b_listma",
	}
}
