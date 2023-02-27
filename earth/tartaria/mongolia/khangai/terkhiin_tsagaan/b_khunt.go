package terkhiin_tsagaan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洪特KhuntBarony struct {
	feud.BaseBarony
}

var BKhunt洪特 feud.Barony = &洪特KhuntBarony{}

func init() {
    f := BKhunt洪特.(*洪特KhuntBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khunt",
		TitleName: "洪特",
		TitleCode: "b_khunt",
	}
}
