package acre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿特利特AthlithBarony struct {
	feud.BaseBarony
}

var BAthlith阿特利特 feud.Barony = &阿特利特AthlithBarony{}

func init() {
    f := BAthlith阿特利特.(*阿特利特AthlithBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "athlith",
		TitleName: "阿特利特",
		TitleCode: "b_athlith",
	}
}
