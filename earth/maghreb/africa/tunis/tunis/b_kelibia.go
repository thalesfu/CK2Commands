package tunis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古莱比耶KelibiaBarony struct {
	feud.BaseBarony
}

var BKelibia古莱比耶 feud.Barony = &古莱比耶KelibiaBarony{}

func init() {
	f := BKelibia古莱比耶.(*古莱比耶KelibiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kelibia",
		TitleName: "古莱比耶",
		TitleCode: "b_kelibia",
	}
}
