package bobo_dyulasso

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基凯莱KikeleBarony struct {
	feud.BaseBarony
}

var BKikele基凯莱 feud.Barony = &基凯莱KikeleBarony{}

func init() {
    f := BKikele基凯莱.(*基凯莱KikeleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kikele",
		TitleName: "基凯莱",
		TitleCode: "b_kikele",
	}
}
