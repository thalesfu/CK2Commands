package kildare

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿斯隆AthloneBarony struct {
	feud.BaseBarony
}

var BAthlone阿斯隆 feud.Barony = &阿斯隆AthloneBarony{}

func init() {
    f := BAthlone阿斯隆.(*阿斯隆AthloneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "athlone",
		TitleName: "阿斯隆",
		TitleCode: "b_athlone",
	}
}
