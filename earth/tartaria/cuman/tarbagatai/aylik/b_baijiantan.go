package aylik

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 白碱滩BaijiantanBarony struct {
	feud.BaseBarony
}

var BBaijiantan白碱滩 feud.Barony = &白碱滩BaijiantanBarony{}

func init() {
	f := BBaijiantan白碱滩.(*白碱滩BaijiantanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baijiantan",
		TitleName: "白碱滩",
		TitleCode: "b_baijiantan",
	}
}
