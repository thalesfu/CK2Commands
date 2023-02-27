package kuban

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科帕里亚CopariaBarony struct {
	feud.BaseBarony
}

var BCoparia科帕里亚 feud.Barony = &科帕里亚CopariaBarony{}

func init() {
    f := BCoparia科帕里亚.(*科帕里亚CopariaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "coparia",
		TitleName: "科帕里亚",
		TitleCode: "b_coparia",
	}
}
