package constantine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 君士坦丁ConstantineBarony struct {
	feud.BaseBarony
}

var BConstantine君士坦丁 feud.Barony = &君士坦丁ConstantineBarony{}

func init() {
    f := BConstantine君士坦丁.(*君士坦丁ConstantineBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "constantine",
		TitleName: "君士坦丁",
		TitleCode: "b_constantine",
	}
}
