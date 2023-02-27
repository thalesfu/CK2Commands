package la_marche

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 茹雅JouillatBarony struct {
	feud.BaseBarony
}

var BJouillat茹雅 feud.Barony = &茹雅JouillatBarony{}

func init() {
    f := BJouillat茹雅.(*茹雅JouillatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jouillat",
		TitleName: "茹雅",
		TitleCode: "b_jouillat",
	}
}
