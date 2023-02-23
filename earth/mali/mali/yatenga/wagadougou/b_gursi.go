package wagadougou

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古尔西GursiBarony struct {
	feud.BaseBarony
}

var BGursi古尔西 feud.Barony = &古尔西GursiBarony{}

func init() {
	f := BGursi古尔西.(*古尔西GursiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gursi",
		TitleName: "古尔西",
		TitleCode: "b_gursi",
	}
}
