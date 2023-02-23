package anglesey

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兰法斯LlanfaesBarony struct {
	feud.BaseBarony
}

var BLlanfaes兰法斯 feud.Barony = &兰法斯LlanfaesBarony{}

func init() {
	f := BLlanfaes兰法斯.(*兰法斯LlanfaesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "llanfaes",
		TitleName: "兰法斯",
		TitleCode: "b_llanfaes",
	}
}
