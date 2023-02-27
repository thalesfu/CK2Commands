package brugge

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅尔MaleBarony struct {
	feud.BaseBarony
}

var BMale梅尔 feud.Barony = &梅尔MaleBarony{}

func init() {
    f := BMale梅尔.(*梅尔MaleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "male",
		TitleName: "梅尔",
		TitleCode: "b_male",
	}
}
