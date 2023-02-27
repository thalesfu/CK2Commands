package birlad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特尔古普特纳Targu_putneiBarony struct {
	feud.BaseBarony
}

var BTargu_putnei特尔古普特纳 feud.Barony = &特尔古普特纳Targu_putneiBarony{}

func init() {
    f := BTargu_putnei特尔古普特纳.(*特尔古普特纳Targu_putneiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "targu_putnei",
		TitleName: "特尔古普特纳",
		TitleCode: "b_targu_putnei",
	}
}
