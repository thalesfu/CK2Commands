package zanjan_abhar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏赫拉瓦德SuhrawardBarony struct {
	feud.BaseBarony
}

var BSuhraward苏赫拉瓦德 feud.Barony = &苏赫拉瓦德SuhrawardBarony{}

func init() {
    f := BSuhraward苏赫拉瓦德.(*苏赫拉瓦德SuhrawardBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "suhraward",
		TitleName: "苏赫拉瓦德",
		TitleCode: "b_suhraward",
	}
}
