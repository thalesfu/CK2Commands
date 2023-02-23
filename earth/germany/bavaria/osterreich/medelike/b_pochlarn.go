package medelike

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 珀希拉恩PochlarnBarony struct {
	feud.BaseBarony
}

var BPochlarn珀希拉恩 feud.Barony = &珀希拉恩PochlarnBarony{}

func init() {
	f := BPochlarn珀希拉恩.(*珀希拉恩PochlarnBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pochlarn",
		TitleName: "珀希拉恩",
		TitleCode: "b_pochlarn",
	}
}
