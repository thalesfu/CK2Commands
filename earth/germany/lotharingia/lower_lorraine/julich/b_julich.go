package julich

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 于利希JulichBarony struct {
	feud.BaseBarony
}

var BJulich于利希 feud.Barony = &于利希JulichBarony{}

func init() {
    f := BJulich于利希.(*于利希JulichBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "julich",
		TitleName: "于利希",
		TitleCode: "b_julich",
	}
}
