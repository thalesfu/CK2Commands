package maragha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌詹UjanBarony struct {
	feud.BaseBarony
}

var BUjan乌詹 feud.Barony = &乌詹UjanBarony{}

func init() {
    f := BUjan乌詹.(*乌詹UjanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ujan",
		TitleName: "乌詹",
		TitleCode: "b_ujan",
	}
}
