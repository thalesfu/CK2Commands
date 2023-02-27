package sonda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 孙陀SondaBarony struct {
	feud.BaseBarony
}

var BSonda孙陀 feud.Barony = &孙陀SondaBarony{}

func init() {
    f := BSonda孙陀.(*孙陀SondaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sonda",
		TitleName: "孙陀",
		TitleCode: "b_sonda",
	}
}
