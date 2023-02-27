package manyakheta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 崛跋伽GulbargaBarony struct {
	feud.BaseBarony
}

var BGulbarga崛跋伽 feud.Barony = &崛跋伽GulbargaBarony{}

func init() {
    f := BGulbarga崛跋伽.(*崛跋伽GulbargaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gulbarga",
		TitleName: "崛跋伽",
		TitleCode: "b_gulbarga",
	}
}
