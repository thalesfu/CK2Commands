package trent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 施兰德斯SchlandersBarony struct {
	feud.BaseBarony
}

var BSchlanders施兰德斯 feud.Barony = &施兰德斯SchlandersBarony{}

func init() {
    f := BSchlanders施兰德斯.(*施兰德斯SchlandersBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "schlanders",
		TitleName: "施兰德斯",
		TitleCode: "b_schlanders",
	}
}
