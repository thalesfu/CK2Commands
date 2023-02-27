package lindisfarne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特威德茅斯TweedmouthBarony struct {
	feud.BaseBarony
}

var BTweedmouth特威德茅斯 feud.Barony = &特威德茅斯TweedmouthBarony{}

func init() {
    f := BTweedmouth特威德茅斯.(*特威德茅斯TweedmouthBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tweedmouth",
		TitleName: "特威德茅斯",
		TitleCode: "b_tweedmouth",
	}
}
