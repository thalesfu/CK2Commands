package alabuga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨拉普尔SarapulBarony struct {
	feud.BaseBarony
}

var BSarapul萨拉普尔 feud.Barony = &萨拉普尔SarapulBarony{}

func init() {
    f := BSarapul萨拉普尔.(*萨拉普尔SarapulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarapul",
		TitleName: "萨拉普尔",
		TitleCode: "b_sarapul",
	}
}
