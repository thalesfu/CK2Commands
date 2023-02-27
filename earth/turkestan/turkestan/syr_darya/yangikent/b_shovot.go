package yangikent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙瓦特ShovotBarony struct {
	feud.BaseBarony
}

var BShovot沙瓦特 feud.Barony = &沙瓦特ShovotBarony{}

func init() {
    f := BShovot沙瓦特.(*沙瓦特ShovotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shovot",
		TitleName: "沙瓦特",
		TitleCode: "b_shovot",
	}
}
