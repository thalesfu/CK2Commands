package ankyra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈伊马纳HaymanaBarony struct {
	feud.BaseBarony
}

var BHaymana哈伊马纳 feud.Barony = &哈伊马纳HaymanaBarony{}

func init() {
	f := BHaymana哈伊马纳.(*哈伊马纳HaymanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "haymana",
		TitleName: "哈伊马纳",
		TitleCode: "b_haymana",
	}
}
