package belgorod

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡普里亚纳CaprianaBarony struct {
	feud.BaseBarony
}

var BCapriana卡普里亚纳 feud.Barony = &卡普里亚纳CaprianaBarony{}

func init() {
	f := BCapriana卡普里亚纳.(*卡普里亚纳CaprianaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "capriana",
		TitleName: "卡普里亚纳",
		TitleCode: "b_capriana",
	}
}
