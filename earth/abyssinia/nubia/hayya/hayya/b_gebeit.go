package hayya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰拜特GebeitBarony struct {
	feud.BaseBarony
}

var BGebeit杰拜特 feud.Barony = &杰拜特GebeitBarony{}

func init() {
    f := BGebeit杰拜特.(*杰拜特GebeitBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gebeit",
		TitleName: "杰拜特",
		TitleCode: "b_gebeit",
	}
}
