package kemi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌茨约基UtsjokiBarony struct {
	feud.BaseBarony
}

var BUtsjoki乌茨约基 feud.Barony = &乌茨约基UtsjokiBarony{}

func init() {
    f := BUtsjoki乌茨约基.(*乌茨约基UtsjokiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "utsjoki",
		TitleName: "乌茨约基",
		TitleCode: "b_utsjoki",
	}
}
