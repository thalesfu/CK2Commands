package suakin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰威TaiwiBarony struct {
	feud.BaseBarony
}

var BTaiwi泰威 feud.Barony = &泰威TaiwiBarony{}

func init() {
	f := BTaiwi泰威.(*泰威TaiwiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taiwi",
		TitleName: "泰威",
		TitleCode: "b_taiwi",
	}
}
