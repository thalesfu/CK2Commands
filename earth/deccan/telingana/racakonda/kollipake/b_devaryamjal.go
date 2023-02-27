package kollipake

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德瓦尔亚加利DevaryamjalBarony struct {
	feud.BaseBarony
}

var BDevaryamjal德瓦尔亚加利 feud.Barony = &德瓦尔亚加利DevaryamjalBarony{}

func init() {
    f := BDevaryamjal德瓦尔亚加利.(*德瓦尔亚加利DevaryamjalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "devaryamjal",
		TitleName: "德瓦尔亚加利",
		TitleCode: "b_devaryamjal",
	}
}
