package atlas_mnt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 撒玛娜SelmanaBarony struct {
	feud.BaseBarony
}

var BSelmana撒玛娜 feud.Barony = &撒玛娜SelmanaBarony{}

func init() {
    f := BSelmana撒玛娜.(*撒玛娜SelmanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "selmana",
		TitleName: "撒玛娜",
		TitleCode: "b_selmana",
	}
}
