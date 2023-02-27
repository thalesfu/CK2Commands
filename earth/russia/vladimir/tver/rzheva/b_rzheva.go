package rzheva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勒热瓦RzhevaBarony struct {
	feud.BaseBarony
}

var BRzheva勒热瓦 feud.Barony = &勒热瓦RzhevaBarony{}

func init() {
    f := BRzheva勒热瓦.(*勒热瓦RzhevaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rzheva",
		TitleName: "勒热瓦",
		TitleCode: "b_rzheva",
	}
}
