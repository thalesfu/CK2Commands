package midnapore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦尸城KashipurBarony struct {
	feud.BaseBarony
}

var BKashipur迦尸城 feud.Barony = &迦尸城KashipurBarony{}

func init() {
	f := BKashipur迦尸城.(*迦尸城KashipurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kashipur",
		TitleName: "迦尸城",
		TitleCode: "b_kashipur",
	}
}
