package belukha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆利塔MultaBarony struct {
	feud.BaseBarony
}

var BMulta穆利塔 feud.Barony = &穆利塔MultaBarony{}

func init() {
	f := BMulta穆利塔.(*穆利塔MultaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "multa",
		TitleName: "穆利塔",
		TitleCode: "b_multa",
	}
}
