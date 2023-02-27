package simaramapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 薄多伽罗摩BhaktagramaBarony struct {
	feud.BaseBarony
}

var BBhaktagrama薄多伽罗摩 feud.Barony = &薄多伽罗摩BhaktagramaBarony{}

func init() {
    f := BBhaktagrama薄多伽罗摩.(*薄多伽罗摩BhaktagramaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhaktagrama",
		TitleName: "薄多伽罗摩",
		TitleCode: "b_bhaktagrama",
	}
}
