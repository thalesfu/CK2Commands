package dashhowuz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔赫塔TagtaBarony struct {
	feud.BaseBarony
}

var BTagta塔赫塔 feud.Barony = &塔赫塔TagtaBarony{}

func init() {
	f := BTagta塔赫塔.(*塔赫塔TagtaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tagta",
		TitleName: "塔赫塔",
		TitleCode: "b_tagta",
	}
}
