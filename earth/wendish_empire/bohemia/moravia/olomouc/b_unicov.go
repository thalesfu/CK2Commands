package olomouc

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌尼乔夫UnicovBarony struct {
	feud.BaseBarony
}

var BUnicov乌尼乔夫 feud.Barony = &乌尼乔夫UnicovBarony{}

func init() {
    f := BUnicov乌尼乔夫.(*乌尼乔夫UnicovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "unicov",
		TitleName: "乌尼乔夫",
		TitleCode: "b_unicov",
	}
}
