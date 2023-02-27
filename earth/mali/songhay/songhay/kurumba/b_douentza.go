package kurumba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜安扎DouentzaBarony struct {
	feud.BaseBarony
}

var BDouentza杜安扎 feud.Barony = &杜安扎DouentzaBarony{}

func init() {
    f := BDouentza杜安扎.(*杜安扎DouentzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "douentza",
		TitleName: "杜安扎",
		TitleCode: "b_douentza",
	}
}
