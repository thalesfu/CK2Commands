package korosten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切尔诺贝利ChernobylBarony struct {
	feud.BaseBarony
}

var BChernobyl切尔诺贝利 feud.Barony = &切尔诺贝利ChernobylBarony{}

func init() {
	f := BChernobyl切尔诺贝利.(*切尔诺贝利ChernobylBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chernobyl",
		TitleName: "切尔诺贝利",
		TitleCode: "b_chernobyl",
	}
}
