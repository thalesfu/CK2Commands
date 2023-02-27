package suvraga_khairkhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏布日嘎海尔汗Suvraga_khairkhanBarony struct {
	feud.BaseBarony
}

var BSuvraga_khairkhan苏布日嘎海尔汗 feud.Barony = &苏布日嘎海尔汗Suvraga_khairkhanBarony{}

func init() {
    f := BSuvraga_khairkhan苏布日嘎海尔汗.(*苏布日嘎海尔汗Suvraga_khairkhanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "suvraga_khairkhan",
		TitleName: "苏布日嘎海尔汗",
		TitleCode: "b_suvraga_khairkhan",
	}
}
