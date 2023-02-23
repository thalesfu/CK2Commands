package mustang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曲桑ChhusangBarony struct {
	feud.BaseBarony
}

var BChhusang曲桑 feud.Barony = &曲桑ChhusangBarony{}

func init() {
	f := BChhusang曲桑.(*曲桑ChhusangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chhusang",
		TitleName: "曲桑",
		TitleCode: "b_chhusang",
	}
}
