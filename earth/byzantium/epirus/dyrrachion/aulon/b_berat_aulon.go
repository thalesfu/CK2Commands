package aulon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 培拉特Berat_aulonBarony struct {
	feud.BaseBarony
}

var BBerat_aulon培拉特 feud.Barony = &培拉特Berat_aulonBarony{}

func init() {
    f := BBerat_aulon培拉特.(*培拉特Berat_aulonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "berat_aulon",
		TitleName: "培拉特",
		TitleCode: "b_berat_aulon",
	}
}
