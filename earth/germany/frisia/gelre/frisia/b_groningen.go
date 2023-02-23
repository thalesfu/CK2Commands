package frisia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格罗宁根GroningenBarony struct {
	feud.BaseBarony
}

var BGroningen格罗宁根 feud.Barony = &格罗宁根GroningenBarony{}

func init() {
	f := BGroningen格罗宁根.(*格罗宁根GroningenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "groningen",
		TitleName: "格罗宁根",
		TitleCode: "b_groningen",
	}
}
