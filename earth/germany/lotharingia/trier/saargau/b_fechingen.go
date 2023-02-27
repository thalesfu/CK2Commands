package saargau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费兴根FechingenBarony struct {
	feud.BaseBarony
}

var BFechingen费兴根 feud.Barony = &费兴根FechingenBarony{}

func init() {
    f := BFechingen费兴根.(*费兴根FechingenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fechingen",
		TitleName: "费兴根",
		TitleCode: "b_fechingen",
	}
}
