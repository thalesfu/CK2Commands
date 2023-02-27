package burgos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯特罗巴尔特CastrobarteBarony struct {
	feud.BaseBarony
}

var BCastrobarte卡斯特罗巴尔特 feud.Barony = &卡斯特罗巴尔特CastrobarteBarony{}

func init() {
    f := BCastrobarte卡斯特罗巴尔特.(*卡斯特罗巴尔特CastrobarteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "castrobarte",
		TitleName: "卡斯特罗巴尔特",
		TitleCode: "b_castrobarte",
	}
}
