package loon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈瑟尔特HasseltBarony struct {
	feud.BaseBarony
}

var BHasselt哈瑟尔特 feud.Barony = &哈瑟尔特HasseltBarony{}

func init() {
    f := BHasselt哈瑟尔特.(*哈瑟尔特HasseltBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hasselt",
		TitleName: "哈瑟尔特",
		TitleCode: "b_hasselt",
	}
}
