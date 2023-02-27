package sthanisvara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 娑摩那SamanaBarony struct {
	feud.BaseBarony
}

var BSamana娑摩那 feud.Barony = &娑摩那SamanaBarony{}

func init() {
    f := BSamana娑摩那.(*娑摩那SamanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "samana",
		TitleName: "娑摩那",
		TitleCode: "b_samana",
	}
}
