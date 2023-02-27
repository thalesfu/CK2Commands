package asayita

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉瓦哈JewahaBarony struct {
	feud.BaseBarony
}

var BJewaha吉瓦哈 feud.Barony = &吉瓦哈JewahaBarony{}

func init() {
    f := BJewaha吉瓦哈.(*吉瓦哈JewahaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jewaha",
		TitleName: "吉瓦哈",
		TitleCode: "b_jewaha",
	}
}
