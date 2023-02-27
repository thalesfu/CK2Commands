package lower_volga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科托沃KotovoBarony struct {
	feud.BaseBarony
}

var BKotovo科托沃 feud.Barony = &科托沃KotovoBarony{}

func init() {
    f := BKotovo科托沃.(*科托沃KotovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kotovo",
		TitleName: "科托沃",
		TitleCode: "b_kotovo",
	}
}
