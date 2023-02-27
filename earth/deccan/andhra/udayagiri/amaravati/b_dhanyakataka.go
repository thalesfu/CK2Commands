package amaravati

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 驮那羯磔迦DhanyakatakaBarony struct {
	feud.BaseBarony
}

var BDhanyakataka驮那羯磔迦 feud.Barony = &驮那羯磔迦DhanyakatakaBarony{}

func init() {
    f := BDhanyakataka驮那羯磔迦.(*驮那羯磔迦DhanyakatakaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dhanyakataka",
		TitleName: "驮那羯磔迦",
		TitleCode: "b_dhanyakataka",
	}
}
