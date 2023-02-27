package sakala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 富兰那揭罗Puran_nagarBarony struct {
	feud.BaseBarony
}

var BPuran_nagar富兰那揭罗 feud.Barony = &富兰那揭罗Puran_nagarBarony{}

func init() {
    f := BPuran_nagar富兰那揭罗.(*富兰那揭罗Puran_nagarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "puran_nagar",
		TitleName: "富兰那揭罗",
		TitleCode: "b_puran_nagar",
	}
}
