package byzantion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 君士坦丁堡ConstantinopleBarony struct {
	feud.BaseBarony
}

var BConstantinople君士坦丁堡 feud.Barony = &君士坦丁堡ConstantinopleBarony{}

func init() {
	f := BConstantinople君士坦丁堡.(*君士坦丁堡ConstantinopleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "constantinople",
		TitleName: "君士坦丁堡",
		TitleCode: "b_constantinople",
	}
}
