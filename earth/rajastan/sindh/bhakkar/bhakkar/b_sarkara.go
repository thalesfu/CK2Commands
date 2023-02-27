package bhakkar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 煞割令SarkaraBarony struct {
	feud.BaseBarony
}

var BSarkara煞割令 feud.Barony = &煞割令SarkaraBarony{}

func init() {
    f := BSarkara煞割令.(*煞割令SarkaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarkara",
		TitleName: "煞割令",
		TitleCode: "b_sarkara",
	}
}
