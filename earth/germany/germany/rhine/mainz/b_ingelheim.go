package mainz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 英格尔海姆IngelheimBarony struct {
	feud.BaseBarony
}

var BIngelheim英格尔海姆 feud.Barony = &英格尔海姆IngelheimBarony{}

func init() {
    f := BIngelheim英格尔海姆.(*英格尔海姆IngelheimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ingelheim",
		TitleName: "英格尔海姆",
		TitleCode: "b_ingelheim",
	}
}
