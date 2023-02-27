package amalfi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡普里CapriBarony struct {
	feud.BaseBarony
}

var BCapri卡普里 feud.Barony = &卡普里CapriBarony{}

func init() {
    f := BCapri卡普里.(*卡普里CapriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "capri",
		TitleName: "卡普里",
		TitleCode: "b_capri",
	}
}
