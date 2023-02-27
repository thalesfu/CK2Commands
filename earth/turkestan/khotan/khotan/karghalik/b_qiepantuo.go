package karghalik

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 朅盘陀QiepantuoBarony struct {
	feud.BaseBarony
}

var BQiepantuo朅盘陀 feud.Barony = &朅盘陀QiepantuoBarony{}

func init() {
    f := BQiepantuo朅盘陀.(*朅盘陀QiepantuoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qiepantuo",
		TitleName: "朅盘陀",
		TitleCode: "b_qiepantuo",
	}
}
