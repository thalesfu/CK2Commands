package suvraga_khairkhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫舍特Khushuut_suvragaBarony struct {
	feud.BaseBarony
}

var BKhushuut_suvraga赫舍特 feud.Barony = &赫舍特Khushuut_suvragaBarony{}

func init() {
    f := BKhushuut_suvraga赫舍特.(*赫舍特Khushuut_suvragaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khushuut_suvraga",
		TitleName: "赫舍特",
		TitleCode: "b_khushuut_suvraga",
	}
}
