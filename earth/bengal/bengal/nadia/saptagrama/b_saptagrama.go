package saptagrama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 七村SaptagramaBarony struct {
	feud.BaseBarony
}

var BSaptagrama七村 feud.Barony = &七村SaptagramaBarony{}

func init() {
	f := BSaptagrama七村.(*七村SaptagramaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saptagrama",
		TitleName: "七村",
		TitleCode: "b_saptagrama",
	}
}
