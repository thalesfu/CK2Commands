package castellon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯特略CastellonBarony struct {
	feud.BaseBarony
}

var BCastellon卡斯特略 feud.Barony = &卡斯特略CastellonBarony{}

func init() {
    f := BCastellon卡斯特略.(*卡斯特略CastellonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "castellon",
		TitleName: "卡斯特略",
		TitleCode: "b_castellon",
	}
}
