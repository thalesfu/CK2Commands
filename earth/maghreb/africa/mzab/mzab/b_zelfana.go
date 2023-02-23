package mzab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 宰勒法纳ZelfanaBarony struct {
	feud.BaseBarony
}

var BZelfana宰勒法纳 feud.Barony = &宰勒法纳ZelfanaBarony{}

func init() {
	f := BZelfana宰勒法纳.(*宰勒法纳ZelfanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zelfana",
		TitleName: "宰勒法纳",
		TitleCode: "b_zelfana",
	}
}
