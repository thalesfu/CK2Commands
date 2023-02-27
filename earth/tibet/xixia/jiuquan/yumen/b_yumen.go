package yumen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 玉门YumenBarony struct {
	feud.BaseBarony
}

var BYumen玉门 feud.Barony = &玉门YumenBarony{}

func init() {
    f := BYumen玉门.(*玉门YumenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yumen",
		TitleName: "玉门",
		TitleCode: "b_yumen",
	}
}
