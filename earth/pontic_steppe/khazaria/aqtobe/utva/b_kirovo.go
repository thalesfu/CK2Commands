package utva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基洛沃KirovoBarony struct {
	feud.BaseBarony
}

var BKirovo基洛沃 feud.Barony = &基洛沃KirovoBarony{}

func init() {
    f := BKirovo基洛沃.(*基洛沃KirovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kirovo",
		TitleName: "基洛沃",
		TitleCode: "b_kirovo",
	}
}
