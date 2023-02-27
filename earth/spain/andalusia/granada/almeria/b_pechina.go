package almeria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩奇纳PechinaBarony struct {
	feud.BaseBarony
}

var BPechina佩奇纳 feud.Barony = &佩奇纳PechinaBarony{}

func init() {
    f := BPechina佩奇纳.(*佩奇纳PechinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pechina",
		TitleName: "佩奇纳",
		TitleCode: "b_pechina",
	}
}
