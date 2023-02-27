package swetaka_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 湿吠多迦城SwetakapuraBarony struct {
	feud.BaseBarony
}

var BSwetakapura湿吠多迦城 feud.Barony = &湿吠多迦城SwetakapuraBarony{}

func init() {
    f := BSwetakapura湿吠多迦城.(*湿吠多迦城SwetakapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "swetakapura",
		TitleName: "湿吠多迦城",
		TitleCode: "b_swetakapura",
	}
}
