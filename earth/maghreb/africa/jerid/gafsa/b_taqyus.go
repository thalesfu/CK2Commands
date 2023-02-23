package gafsa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰格尤斯TaqyusBarony struct {
	feud.BaseBarony
}

var BTaqyus泰格尤斯 feud.Barony = &泰格尤斯TaqyusBarony{}

func init() {
	f := BTaqyus泰格尤斯.(*泰格尤斯TaqyusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taqyus",
		TitleName: "泰格尤斯",
		TitleCode: "b_taqyus",
	}
}
