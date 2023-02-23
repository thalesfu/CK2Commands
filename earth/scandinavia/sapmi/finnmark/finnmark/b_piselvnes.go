package finnmark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮瑟吕尼PiselvnesBarony struct {
	feud.BaseBarony
}

var BPiselvnes皮瑟吕尼 feud.Barony = &皮瑟吕尼PiselvnesBarony{}

func init() {
	f := BPiselvnes皮瑟吕尼.(*皮瑟吕尼PiselvnesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "piselvnes",
		TitleName: "皮瑟吕尼",
		TitleCode: "b_piselvnes",
	}
}
