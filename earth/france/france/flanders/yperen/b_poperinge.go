package yperen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波珀灵厄PoperingeBarony struct {
	feud.BaseBarony
}

var BPoperinge波珀灵厄 feud.Barony = &波珀灵厄PoperingeBarony{}

func init() {
	f := BPoperinge波珀灵厄.(*波珀灵厄PoperingeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "poperinge",
		TitleName: "波珀灵厄",
		TitleCode: "b_poperinge",
	}
}
