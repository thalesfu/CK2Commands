package kongu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西陀曼迦楞SendamangalamBarony struct {
	feud.BaseBarony
}

var BSendamangalam西陀曼迦楞 feud.Barony = &西陀曼迦楞SendamangalamBarony{}

func init() {
    f := BSendamangalam西陀曼迦楞.(*西陀曼迦楞SendamangalamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sendamangalam",
		TitleName: "西陀曼迦楞",
		TitleCode: "b_sendamangalam",
	}
}
