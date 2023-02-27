package szekezfehervar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伦杰尔托蒂LengyeltotiBarony struct {
	feud.BaseBarony
}

var BLengyeltoti伦杰尔托蒂 feud.Barony = &伦杰尔托蒂LengyeltotiBarony{}

func init() {
    f := BLengyeltoti伦杰尔托蒂.(*伦杰尔托蒂LengyeltotiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lengyeltoti",
		TitleName: "伦杰尔托蒂",
		TitleCode: "b_lengyeltoti",
	}
}
