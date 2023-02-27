package voin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊尔克利伊夫IrkliivBarony struct {
	feud.BaseBarony
}

var BIrkliiv伊尔克利伊夫 feud.Barony = &伊尔克利伊夫IrkliivBarony{}

func init() {
    f := BIrkliiv伊尔克利伊夫.(*伊尔克利伊夫IrkliivBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "irkliiv",
		TitleName: "伊尔克利伊夫",
		TitleCode: "b_irkliiv",
	}
}
