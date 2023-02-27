package srirangapatna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 难加罗亚帕提塔纳NanjarayapattanaBarony struct {
	feud.BaseBarony
}

var BNanjarayapattana难加罗亚帕提塔纳 feud.Barony = &难加罗亚帕提塔纳NanjarayapattanaBarony{}

func init() {
    f := BNanjarayapattana难加罗亚帕提塔纳.(*难加罗亚帕提塔纳NanjarayapattanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nanjarayapattana",
		TitleName: "难加罗亚帕提塔纳",
		TitleCode: "b_nanjarayapattana",
	}
}
