package oualata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比如BiruBarony struct {
	feud.BaseBarony
}

var BBiru比如 feud.Barony = &比如BiruBarony{}

func init() {
    f := BBiru比如.(*比如BiruBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "biru",
		TitleName: "比如",
		TitleCode: "b_biru",
	}
}
