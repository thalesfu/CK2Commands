package debul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舍般陀ShahbandarBarony struct {
	feud.BaseBarony
}

var BShahbandar舍般陀 feud.Barony = &舍般陀ShahbandarBarony{}

func init() {
	f := BShahbandar舍般陀.(*舍般陀ShahbandarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shahbandar",
		TitleName: "舍般陀",
		TitleCode: "b_shahbandar",
	}
}
