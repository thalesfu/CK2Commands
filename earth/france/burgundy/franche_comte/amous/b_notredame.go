package amous

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺特尔达姆NotredameBarony struct {
	feud.BaseBarony
}

var BNotredame诺特尔达姆 feud.Barony = &诺特尔达姆NotredameBarony{}

func init() {
    f := BNotredame诺特尔达姆.(*诺特尔达姆NotredameBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "notredame",
		TitleName: "诺特尔达姆",
		TitleCode: "b_notredame",
	}
}
