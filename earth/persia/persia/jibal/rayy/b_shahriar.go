package rayy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙赫里亚尔ShahriarBarony struct {
	feud.BaseBarony
}

var BShahriar沙赫里亚尔 feud.Barony = &沙赫里亚尔ShahriarBarony{}

func init() {
    f := BShahriar沙赫里亚尔.(*沙赫里亚尔ShahriarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shahriar",
		TitleName: "沙赫里亚尔",
		TitleCode: "b_shahriar",
	}
}
