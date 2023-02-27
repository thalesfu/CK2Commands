package tagadur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯耶门伽愣SiyyamangalamBarony struct {
	feud.BaseBarony
}

var BSiyyamangalam斯耶门伽愣 feud.Barony = &斯耶门伽愣SiyyamangalamBarony{}

func init() {
    f := BSiyyamangalam斯耶门伽愣.(*斯耶门伽愣SiyyamangalamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "siyyamangalam",
		TitleName: "斯耶门伽愣",
		TitleCode: "b_siyyamangalam",
	}
}
