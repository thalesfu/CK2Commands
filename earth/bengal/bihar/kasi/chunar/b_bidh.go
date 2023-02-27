package chunar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗陀BidhBarony struct {
	feud.BaseBarony
}

var BBidh毗陀 feud.Barony = &毗陀BidhBarony{}

func init() {
    f := BBidh毗陀.(*毗陀BidhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bidh",
		TitleName: "毗陀",
		TitleCode: "b_bidh",
	}
}
