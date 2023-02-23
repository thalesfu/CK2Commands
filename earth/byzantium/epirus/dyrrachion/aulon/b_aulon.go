package aulon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿夫洛纳斯AulonBarony struct {
	feud.BaseBarony
}

var BAulon阿夫洛纳斯 feud.Barony = &阿夫洛纳斯AulonBarony{}

func init() {
	f := BAulon阿夫洛纳斯.(*阿夫洛纳斯AulonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aulon",
		TitleName: "阿夫洛纳斯",
		TitleCode: "b_aulon",
	}
}
