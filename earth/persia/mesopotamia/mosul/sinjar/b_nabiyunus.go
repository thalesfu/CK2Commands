package sinjar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳比尤努斯NabiyunusBarony struct {
	feud.BaseBarony
}

var BNabiyunus纳比尤努斯 feud.Barony = &纳比尤努斯NabiyunusBarony{}

func init() {
	f := BNabiyunus纳比尤努斯.(*纳比尤努斯NabiyunusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nabiyunus",
		TitleName: "纳比尤努斯",
		TitleCode: "b_nabiyunus",
	}
}
