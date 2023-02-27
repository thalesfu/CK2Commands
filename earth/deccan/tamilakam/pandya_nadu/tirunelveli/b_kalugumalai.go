package tirunelveli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽卢古摩罗KalugumalaiBarony struct {
	feud.BaseBarony
}

var BKalugumalai伽卢古摩罗 feud.Barony = &伽卢古摩罗KalugumalaiBarony{}

func init() {
    f := BKalugumalai伽卢古摩罗.(*伽卢古摩罗KalugumalaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalugumalai",
		TitleName: "伽卢古摩罗",
		TitleCode: "b_kalugumalai",
	}
}
