package bam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法赫拉季FahrajBarony struct {
	feud.BaseBarony
}

var BFahraj法赫拉季 feud.Barony = &法赫拉季FahrajBarony{}

func init() {
	f := BFahraj法赫拉季.(*法赫拉季FahrajBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fahraj",
		TitleName: "法赫拉季",
		TitleCode: "b_fahraj",
	}
}
