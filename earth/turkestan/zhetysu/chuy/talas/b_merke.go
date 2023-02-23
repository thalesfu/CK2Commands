package talas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 灭尔基MerkeBarony struct {
	feud.BaseBarony
}

var BMerke灭尔基 feud.Barony = &灭尔基MerkeBarony{}

func init() {
	f := BMerke灭尔基.(*灭尔基MerkeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "merke",
		TitleName: "灭尔基",
		TitleCode: "b_merke",
	}
}
