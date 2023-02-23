package pisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比萨PisaBarony struct {
	feud.BaseBarony
}

var BPisa比萨 feud.Barony = &比萨PisaBarony{}

func init() {
	f := BPisa比萨.(*比萨PisaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pisa",
		TitleName: "比萨",
		TitleCode: "b_pisa",
	}
}
