package dorylaion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加卢拉CaruraBarony struct {
	feud.BaseBarony
}

var BCarura加卢拉 feud.Barony = &加卢拉CaruraBarony{}

func init() {
    f := BCarura加卢拉.(*加卢拉CaruraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "carura",
		TitleName: "加卢拉",
		TitleCode: "b_carura",
	}
}
