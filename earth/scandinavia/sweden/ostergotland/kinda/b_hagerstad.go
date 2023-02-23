package kinda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海格斯塔德HagerstadBarony struct {
	feud.BaseBarony
}

var BHagerstad海格斯塔德 feud.Barony = &海格斯塔德HagerstadBarony{}

func init() {
	f := BHagerstad海格斯塔德.(*海格斯塔德HagerstadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hagerstad",
		TitleName: "海格斯塔德",
		TitleCode: "b_hagerstad",
	}
}
