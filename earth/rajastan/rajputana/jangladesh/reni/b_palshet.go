package reni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波湿怛PalshetBarony struct {
	feud.BaseBarony
}

var BPalshet波湿怛 feud.Barony = &波湿怛PalshetBarony{}

func init() {
    f := BPalshet波湿怛.(*波湿怛PalshetBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "palshet",
		TitleName: "波湿怛",
		TitleCode: "b_palshet",
	}
}
