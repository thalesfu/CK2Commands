package attaleia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基比拉CibyraBarony struct {
	feud.BaseBarony
}

var BCibyra基比拉 feud.Barony = &基比拉CibyraBarony{}

func init() {
    f := BCibyra基比拉.(*基比拉CibyraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cibyra",
		TitleName: "基比拉",
		TitleCode: "b_cibyra",
	}
}
