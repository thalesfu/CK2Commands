package trencin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特伦琴TrencsenBarony struct {
	feud.BaseBarony
}

var BTrencsen特伦琴 feud.Barony = &特伦琴TrencsenBarony{}

func init() {
    f := BTrencsen特伦琴.(*特伦琴TrencsenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "trencsen",
		TitleName: "特伦琴",
		TitleCode: "b_trencsen",
	}
}
