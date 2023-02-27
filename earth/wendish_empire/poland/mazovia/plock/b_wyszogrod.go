package plock

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维绍格鲁德WyszogrodBarony struct {
	feud.BaseBarony
}

var BWyszogrod维绍格鲁德 feud.Barony = &维绍格鲁德WyszogrodBarony{}

func init() {
    f := BWyszogrod维绍格鲁德.(*维绍格鲁德WyszogrodBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wyszogrod",
		TitleName: "维绍格鲁德",
		TitleCode: "b_wyszogrod",
	}
}
