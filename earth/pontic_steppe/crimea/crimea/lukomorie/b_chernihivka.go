package lukomorie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切尔尼戈夫卡ChernihivkaBarony struct {
	feud.BaseBarony
}

var BChernihivka切尔尼戈夫卡 feud.Barony = &切尔尼戈夫卡ChernihivkaBarony{}

func init() {
    f := BChernihivka切尔尼戈夫卡.(*切尔尼戈夫卡ChernihivkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chernihivka",
		TitleName: "切尔尼戈夫卡",
		TitleCode: "b_chernihivka",
	}
}
