package murzuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舒拉夫Ash_shurrafBarony struct {
	feud.BaseBarony
}

var BAsh_shurraf舒拉夫 feud.Barony = &舒拉夫Ash_shurrafBarony{}

func init() {
    f := BAsh_shurraf舒拉夫.(*舒拉夫Ash_shurrafBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ash_shurraf",
		TitleName: "舒拉夫",
		TitleCode: "b_ash_shurraf",
	}
}
