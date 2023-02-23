package slesvig

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈泽斯莱乌HaderslevBarony struct {
	feud.BaseBarony
}

var BHaderslev哈泽斯莱乌 feud.Barony = &哈泽斯莱乌HaderslevBarony{}

func init() {
	f := BHaderslev哈泽斯莱乌.(*哈泽斯莱乌HaderslevBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "haderslev",
		TitleName: "哈泽斯莱乌",
		TitleCode: "b_haderslev",
	}
}
