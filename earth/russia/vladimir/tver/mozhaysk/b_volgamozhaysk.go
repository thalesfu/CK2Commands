package mozhaysk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫扎伊斯克VolgamozhayskBarony struct {
	feud.BaseBarony
}

var BVolgamozhaysk莫扎伊斯克 feud.Barony = &莫扎伊斯克VolgamozhayskBarony{}

func init() {
    f := BVolgamozhaysk莫扎伊斯克.(*莫扎伊斯克VolgamozhayskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "volgamozhaysk",
		TitleName: "莫扎伊斯克",
		TitleCode: "b_volgamozhaysk",
	}
}
