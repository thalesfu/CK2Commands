package rajrappa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦阇伽罗摩KajgaonBarony struct {
	feud.BaseBarony
}

var BKajgaon迦阇伽罗摩 feud.Barony = &迦阇伽罗摩KajgaonBarony{}

func init() {
	f := BKajgaon迦阇伽罗摩.(*迦阇伽罗摩KajgaonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kajgaon",
		TitleName: "迦阇伽罗摩",
		TitleCode: "b_kajgaon",
	}
}
