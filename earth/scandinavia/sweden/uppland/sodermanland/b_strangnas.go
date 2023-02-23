package sodermanland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯特兰奈斯StrangnasBarony struct {
	feud.BaseBarony
}

var BStrangnas斯特兰奈斯 feud.Barony = &斯特兰奈斯StrangnasBarony{}

func init() {
	f := BStrangnas斯特兰奈斯.(*斯特兰奈斯StrangnasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "strangnas",
		TitleName: "斯特兰奈斯",
		TitleCode: "b_strangnas",
	}
}
