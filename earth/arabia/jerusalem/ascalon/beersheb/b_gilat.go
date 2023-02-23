package beersheb

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉来特GilatBarony struct {
	feud.BaseBarony
}

var BGilat吉来特 feud.Barony = &吉来特GilatBarony{}

func init() {
	f := BGilat吉来特.(*吉来特GilatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gilat",
		TitleName: "吉来特",
		TitleCode: "b_gilat",
	}
}
