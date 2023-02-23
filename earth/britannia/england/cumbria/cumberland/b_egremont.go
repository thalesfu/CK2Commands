package cumberland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃格勒蒙特EgremontBarony struct {
	feud.BaseBarony
}

var BEgremont埃格勒蒙特 feud.Barony = &埃格勒蒙特EgremontBarony{}

func init() {
	f := BEgremont埃格勒蒙特.(*埃格勒蒙特EgremontBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "egremont",
		TitleName: "埃格勒蒙特",
		TitleCode: "b_egremont",
	}
}
