package bidar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗陀罗BidarBarony struct {
	feud.BaseBarony
}

var BBidar毗陀罗 feud.Barony = &毗陀罗BidarBarony{}

func init() {
    f := BBidar毗陀罗.(*毗陀罗BidarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bidar",
		TitleName: "毗陀罗",
		TitleCode: "b_bidar",
	}
}
