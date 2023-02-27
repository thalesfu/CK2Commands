package khasagt_khairkhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈萨格特海尔汗Khasagt_khairkhanBarony struct {
	feud.BaseBarony
}

var BKhasagt_khairkhan哈萨格特海尔汗 feud.Barony = &哈萨格特海尔汗Khasagt_khairkhanBarony{}

func init() {
    f := BKhasagt_khairkhan哈萨格特海尔汗.(*哈萨格特海尔汗Khasagt_khairkhanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khasagt_khairkhan",
		TitleName: "哈萨格特海尔汗",
		TitleCode: "b_khasagt_khairkhan",
	}
}
