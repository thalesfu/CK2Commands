package chernigov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切尔尼戈夫ChernigovBarony struct {
	feud.BaseBarony
}

var BChernigov切尔尼戈夫 feud.Barony = &切尔尼戈夫ChernigovBarony{}

func init() {
	f := BChernigov切尔尼戈夫.(*切尔尼戈夫ChernigovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chernigov",
		TitleName: "切尔尼戈夫",
		TitleCode: "b_chernigov",
	}
}
