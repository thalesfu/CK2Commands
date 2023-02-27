package akershus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊瑟格兰IsegranBarony struct {
	feud.BaseBarony
}

var BIsegran伊瑟格兰 feud.Barony = &伊瑟格兰IsegranBarony{}

func init() {
    f := BIsegran伊瑟格兰.(*伊瑟格兰IsegranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "isegran",
		TitleName: "伊瑟格兰",
		TitleCode: "b_isegran",
	}
}
