package serdica

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦尔伯日德VelbazhdBarony struct {
	feud.BaseBarony
}

var BVelbazhd韦尔伯日德 feud.Barony = &韦尔伯日德VelbazhdBarony{}

func init() {
    f := BVelbazhd韦尔伯日德.(*韦尔伯日德VelbazhdBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "velbazhd",
		TitleName: "韦尔伯日德",
		TitleCode: "b_velbazhd",
	}
}
