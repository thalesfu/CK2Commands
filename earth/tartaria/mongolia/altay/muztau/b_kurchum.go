package muztau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库尔丘姆KurchumBarony struct {
	feud.BaseBarony
}

var BKurchum库尔丘姆 feud.Barony = &库尔丘姆KurchumBarony{}

func init() {
	f := BKurchum库尔丘姆.(*库尔丘姆KurchumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kurchum",
		TitleName: "库尔丘姆",
		TitleCode: "b_kurchum",
	}
}
