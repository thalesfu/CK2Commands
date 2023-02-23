package talakad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吠补罗VeppurBarony struct {
	feud.BaseBarony
}

var BVeppur吠补罗 feud.Barony = &吠补罗VeppurBarony{}

func init() {
	f := BVeppur吠补罗.(*吠补罗VeppurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "veppur",
		TitleName: "吠补罗",
		TitleCode: "b_veppur",
	}
}
