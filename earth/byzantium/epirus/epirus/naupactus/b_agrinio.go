package naupactus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿格里尼翁AgrinioBarony struct {
	feud.BaseBarony
}

var BAgrinio阿格里尼翁 feud.Barony = &阿格里尼翁AgrinioBarony{}

func init() {
    f := BAgrinio阿格里尼翁.(*阿格里尼翁AgrinioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "agrinio",
		TitleName: "阿格里尼翁",
		TitleCode: "b_agrinio",
	}
}
