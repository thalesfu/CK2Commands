package leptis_magna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奈伊迈NaimahBarony struct {
	feud.BaseBarony
}

var BNaimah奈伊迈 feud.Barony = &奈伊迈NaimahBarony{}

func init() {
    f := BNaimah奈伊迈.(*奈伊迈NaimahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "naimah",
		TitleName: "奈伊迈",
		TitleCode: "b_naimah",
	}
}
