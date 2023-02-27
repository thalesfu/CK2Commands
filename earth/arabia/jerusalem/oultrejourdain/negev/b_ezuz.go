package negev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃祖兹EzuzBarony struct {
	feud.BaseBarony
}

var BEzuz埃祖兹 feud.Barony = &埃祖兹EzuzBarony{}

func init() {
    f := BEzuz埃祖兹.(*埃祖兹EzuzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ezuz",
		TitleName: "埃祖兹",
		TitleCode: "b_ezuz",
	}
}
