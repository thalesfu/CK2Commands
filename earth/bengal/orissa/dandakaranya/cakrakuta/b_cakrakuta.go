package cakrakuta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 遮迦罗矩吒CakrakutaBarony struct {
	feud.BaseBarony
}

var BCakrakuta遮迦罗矩吒 feud.Barony = &遮迦罗矩吒CakrakutaBarony{}

func init() {
	f := BCakrakuta遮迦罗矩吒.(*遮迦罗矩吒CakrakutaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cakrakuta",
		TitleName: "遮迦罗矩吒",
		TitleCode: "b_cakrakuta",
	}
}
