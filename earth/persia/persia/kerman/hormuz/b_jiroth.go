package hormuz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉洛特JirothBarony struct {
	feud.BaseBarony
}

var BJiroth吉洛特 feud.Barony = &吉洛特JirothBarony{}

func init() {
    f := BJiroth吉洛特.(*吉洛特JirothBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jiroth",
		TitleName: "吉洛特",
		TitleCode: "b_jiroth",
	}
}
