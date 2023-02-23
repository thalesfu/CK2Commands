package empuries

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫罗纳GironaBarony struct {
	feud.BaseBarony
}

var BGirona赫罗纳 feud.Barony = &赫罗纳GironaBarony{}

func init() {
	f := BGirona赫罗纳.(*赫罗纳GironaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "girona",
		TitleName: "赫罗纳",
		TitleCode: "b_girona",
	}
}
