package rostov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 孔达科沃KondakovoBarony struct {
	feud.BaseBarony
}

var BKondakovo孔达科沃 feud.Barony = &孔达科沃KondakovoBarony{}

func init() {
    f := BKondakovo孔达科沃.(*孔达科沃KondakovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kondakovo",
		TitleName: "孔达科沃",
		TitleCode: "b_kondakovo",
	}
}
