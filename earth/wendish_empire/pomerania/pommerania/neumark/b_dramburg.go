package neumark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德兰堡DramburgBarony struct {
	feud.BaseBarony
}

var BDramburg德兰堡 feud.Barony = &德兰堡DramburgBarony{}

func init() {
    f := BDramburg德兰堡.(*德兰堡DramburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dramburg",
		TitleName: "德兰堡",
		TitleCode: "b_dramburg",
	}
}
