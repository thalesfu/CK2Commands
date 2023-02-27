package lopnor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 壁边BibianBarony struct {
	feud.BaseBarony
}

var BBibian壁边 feud.Barony = &壁边BibianBarony{}

func init() {
    f := BBibian壁边.(*壁边BibianBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bibian",
		TitleName: "壁边",
		TitleCode: "b_bibian",
	}
}
