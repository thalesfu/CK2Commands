package suvraga_khairkhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哲格苏台ZegsteiBarony struct {
	feud.BaseBarony
}

var BZegstei哲格苏台 feud.Barony = &哲格苏台ZegsteiBarony{}

func init() {
    f := BZegstei哲格苏台.(*哲格苏台ZegsteiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zegstei",
		TitleName: "哲格苏台",
		TitleCode: "b_zegstei",
	}
}
