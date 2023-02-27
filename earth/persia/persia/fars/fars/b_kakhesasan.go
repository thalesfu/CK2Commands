package fars

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡赫斯撒恩KakhesasanBarony struct {
	feud.BaseBarony
}

var BKakhesasan卡赫斯撒恩 feud.Barony = &卡赫斯撒恩KakhesasanBarony{}

func init() {
    f := BKakhesasan卡赫斯撒恩.(*卡赫斯撒恩KakhesasanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kakhesasan",
		TitleName: "卡赫斯撒恩",
		TitleCode: "b_kakhesasan",
	}
}
