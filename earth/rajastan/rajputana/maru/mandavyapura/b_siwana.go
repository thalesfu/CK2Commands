package mandavyapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯瓦那SiwanaBarony struct {
	feud.BaseBarony
}

var BSiwana斯瓦那 feud.Barony = &斯瓦那SiwanaBarony{}

func init() {
	f := BSiwana斯瓦那.(*斯瓦那SiwanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "siwana",
		TitleName: "斯瓦那",
		TitleCode: "b_siwana",
	}
}
