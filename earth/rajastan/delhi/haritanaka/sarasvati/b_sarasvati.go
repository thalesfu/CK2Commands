package sarasvati

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨罗萨伐底SarasvatiBarony struct {
	feud.BaseBarony
}

var BSarasvati萨罗萨伐底 feud.Barony = &萨罗萨伐底SarasvatiBarony{}

func init() {
    f := BSarasvati萨罗萨伐底.(*萨罗萨伐底SarasvatiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarasvati",
		TitleName: "萨罗萨伐底",
		TitleCode: "b_sarasvati",
	}
}
