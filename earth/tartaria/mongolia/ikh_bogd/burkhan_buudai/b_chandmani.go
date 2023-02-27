package burkhan_buudai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 钱德曼ChandmaniBarony struct {
	feud.BaseBarony
}

var BChandmani钱德曼 feud.Barony = &钱德曼ChandmaniBarony{}

func init() {
    f := BChandmani钱德曼.(*钱德曼ChandmaniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chandmani",
		TitleName: "钱德曼",
		TitleCode: "b_chandmani",
	}
}
