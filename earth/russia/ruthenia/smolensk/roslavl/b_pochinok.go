package roslavl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波奇诺克PochinokBarony struct {
	feud.BaseBarony
}

var BPochinok波奇诺克 feud.Barony = &波奇诺克PochinokBarony{}

func init() {
    f := BPochinok波奇诺克.(*波奇诺克PochinokBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pochinok",
		TitleName: "波奇诺克",
		TitleCode: "b_pochinok",
	}
}
