package dauphine_viennois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣昂图万StantoineBarony struct {
	feud.BaseBarony
}

var BStantoine圣昂图万 feud.Barony = &圣昂图万StantoineBarony{}

func init() {
    f := BStantoine圣昂图万.(*圣昂图万StantoineBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stantoine",
		TitleName: "圣昂图万",
		TitleCode: "b_stantoine",
	}
}
