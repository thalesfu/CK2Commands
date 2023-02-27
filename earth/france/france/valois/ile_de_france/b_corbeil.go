package ile_de_france

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尔贝CorbeilBarony struct {
	feud.BaseBarony
}

var BCorbeil科尔贝 feud.Barony = &科尔贝CorbeilBarony{}

func init() {
    f := BCorbeil科尔贝.(*科尔贝CorbeilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "corbeil",
		TitleName: "科尔贝",
		TitleCode: "b_corbeil",
	}
}
