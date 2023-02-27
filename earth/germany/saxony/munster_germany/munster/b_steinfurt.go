package munster

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 施泰因富特SteinfurtBarony struct {
	feud.BaseBarony
}

var BSteinfurt施泰因富特 feud.Barony = &施泰因富特SteinfurtBarony{}

func init() {
    f := BSteinfurt施泰因富特.(*施泰因富特SteinfurtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "steinfurt",
		TitleName: "施泰因富特",
		TitleCode: "b_steinfurt",
	}
}
