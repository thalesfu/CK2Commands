package mainz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈瑙HanauBarony struct {
	feud.BaseBarony
}

var BHanau哈瑙 feud.Barony = &哈瑙HanauBarony{}

func init() {
	f := BHanau哈瑙.(*哈瑙HanauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hanau",
		TitleName: "哈瑙",
		TitleCode: "b_hanau",
	}
}
