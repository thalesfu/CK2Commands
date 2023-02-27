package oshrusana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊斯法纳IsfanaBarony struct {
	feud.BaseBarony
}

var BIsfana伊斯法纳 feud.Barony = &伊斯法纳IsfanaBarony{}

func init() {
    f := BIsfana伊斯法纳.(*伊斯法纳IsfanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "isfana",
		TitleName: "伊斯法纳",
		TitleCode: "b_isfana",
	}
}
