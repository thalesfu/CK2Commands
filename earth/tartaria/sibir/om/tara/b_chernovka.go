package tara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切尔诺夫卡ChernovkaBarony struct {
	feud.BaseBarony
}

var BChernovka切尔诺夫卡 feud.Barony = &切尔诺夫卡ChernovkaBarony{}

func init() {
	f := BChernovka切尔诺夫卡.(*切尔诺夫卡ChernovkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chernovka",
		TitleName: "切尔诺夫卡",
		TitleCode: "b_chernovka",
	}
}
