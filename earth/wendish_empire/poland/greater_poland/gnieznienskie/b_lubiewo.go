package gnieznienskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢别沃LubiewoBarony struct {
	feud.BaseBarony
}

var BLubiewo卢别沃 feud.Barony = &卢别沃LubiewoBarony{}

func init() {
    f := BLubiewo卢别沃.(*卢别沃LubiewoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lubiewo",
		TitleName: "卢别沃",
		TitleCode: "b_lubiewo",
	}
}
