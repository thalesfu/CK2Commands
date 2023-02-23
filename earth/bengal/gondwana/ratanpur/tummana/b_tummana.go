package tummana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 都摩那TummanaBarony struct {
	feud.BaseBarony
}

var BTummana都摩那 feud.Barony = &都摩那TummanaBarony{}

func init() {
	f := BTummana都摩那.(*都摩那TummanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tummana",
		TitleName: "都摩那",
		TitleCode: "b_tummana",
	}
}
