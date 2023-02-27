package palmyra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盖斯尔AlqasrBarony struct {
	feud.BaseBarony
}

var BAlqasr盖斯尔 feud.Barony = &盖斯尔AlqasrBarony{}

func init() {
    f := BAlqasr盖斯尔.(*盖斯尔AlqasrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alqasr",
		TitleName: "盖斯尔",
		TitleCode: "b_alqasr",
	}
}
