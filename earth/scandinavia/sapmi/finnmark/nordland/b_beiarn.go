package nordland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝亚恩BeiarnBarony struct {
	feud.BaseBarony
}

var BBeiarn贝亚恩 feud.Barony = &贝亚恩BeiarnBarony{}

func init() {
    f := BBeiarn贝亚恩.(*贝亚恩BeiarnBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beiarn",
		TitleName: "贝亚恩",
		TitleCode: "b_beiarn",
	}
}
