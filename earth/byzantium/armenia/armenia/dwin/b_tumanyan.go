package dwin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图曼扬TumanyanBarony struct {
	feud.BaseBarony
}

var BTumanyan图曼扬 feud.Barony = &图曼扬TumanyanBarony{}

func init() {
	f := BTumanyan图曼扬.(*图曼扬TumanyanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tumanyan",
		TitleName: "图曼扬",
		TitleCode: "b_tumanyan",
	}
}
