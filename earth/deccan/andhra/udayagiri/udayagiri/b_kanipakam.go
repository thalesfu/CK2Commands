package udayagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贾尼帕卡姆KanipakamBarony struct {
	feud.BaseBarony
}

var BKanipakam贾尼帕卡姆 feud.Barony = &贾尼帕卡姆KanipakamBarony{}

func init() {
    f := BKanipakam贾尼帕卡姆.(*贾尼帕卡姆KanipakamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kanipakam",
		TitleName: "贾尼帕卡姆",
		TitleCode: "b_kanipakam",
	}
}
