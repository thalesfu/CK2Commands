package zoubtsov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维索基诺VysokinoBarony struct {
	feud.BaseBarony
}

var BVysokino维索基诺 feud.Barony = &维索基诺VysokinoBarony{}

func init() {
    f := BVysokino维索基诺.(*维索基诺VysokinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vysokino",
		TitleName: "维索基诺",
		TitleCode: "b_vysokino",
	}
}
