package begemder

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 玛丽亚穆MaryamuBarony struct {
	feud.BaseBarony
}

var BMaryamu玛丽亚穆 feud.Barony = &玛丽亚穆MaryamuBarony{}

func init() {
    f := BMaryamu玛丽亚穆.(*玛丽亚穆MaryamuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maryamu",
		TitleName: "玛丽亚穆",
		TitleCode: "b_maryamu",
	}
}
