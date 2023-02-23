package kexholm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托克索沃ToksovoBarony struct {
	feud.BaseBarony
}

var BToksovo托克索沃 feud.Barony = &托克索沃ToksovoBarony{}

func init() {
	f := BToksovo托克索沃.(*托克索沃ToksovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "toksovo",
		TitleName: "托克索沃",
		TitleCode: "b_toksovo",
	}
}
