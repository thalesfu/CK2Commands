package kerman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎兰德ZarandBarony struct {
	feud.BaseBarony
}

var BZarand扎兰德 feud.Barony = &扎兰德ZarandBarony{}

func init() {
    f := BZarand扎兰德.(*扎兰德ZarandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zarand",
		TitleName: "扎兰德",
		TitleCode: "b_zarand",
	}
}
