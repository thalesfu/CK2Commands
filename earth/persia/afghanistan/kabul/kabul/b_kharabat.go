package kabul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈拉巴特KharabatBarony struct {
	feud.BaseBarony
}

var BKharabat哈拉巴特 feud.Barony = &哈拉巴特KharabatBarony{}

func init() {
    f := BKharabat哈拉巴特.(*哈拉巴特KharabatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kharabat",
		TitleName: "哈拉巴特",
		TitleCode: "b_kharabat",
	}
}
