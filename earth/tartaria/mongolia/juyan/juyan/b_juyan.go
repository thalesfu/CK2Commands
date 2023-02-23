package juyan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 居延JuyanBarony struct {
	feud.BaseBarony
}

var BJuyan居延 feud.Barony = &居延JuyanBarony{}

func init() {
	f := BJuyan居延.(*居延JuyanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "juyan",
		TitleName: "居延",
		TitleCode: "b_juyan",
	}
}
