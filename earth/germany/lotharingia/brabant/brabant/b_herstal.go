package brabant

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫尔斯塔尔HerstalBarony struct {
	feud.BaseBarony
}

var BHerstal赫尔斯塔尔 feud.Barony = &赫尔斯塔尔HerstalBarony{}

func init() {
    f := BHerstal赫尔斯塔尔.(*赫尔斯塔尔HerstalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "herstal",
		TitleName: "赫尔斯塔尔",
		TitleCode: "b_herstal",
	}
}
