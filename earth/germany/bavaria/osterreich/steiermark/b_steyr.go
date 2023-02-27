package steiermark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 施泰尔SteyrBarony struct {
	feud.BaseBarony
}

var BSteyr施泰尔 feud.Barony = &施泰尔SteyrBarony{}

func init() {
    f := BSteyr施泰尔.(*施泰尔SteyrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "steyr",
		TitleName: "施泰尔",
		TitleCode: "b_steyr",
	}
}
