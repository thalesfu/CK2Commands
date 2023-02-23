package trier

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安德纳赫AndernachBarony struct {
	feud.BaseBarony
}

var BAndernach安德纳赫 feud.Barony = &安德纳赫AndernachBarony{}

func init() {
	f := BAndernach安德纳赫.(*安德纳赫AndernachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "andernach",
		TitleName: "安德纳赫",
		TitleCode: "b_andernach",
	}
}
