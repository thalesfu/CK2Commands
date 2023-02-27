package roma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图斯库卢姆TusculumBarony struct {
	feud.BaseBarony
}

var BTusculum图斯库卢姆 feud.Barony = &图斯库卢姆TusculumBarony{}

func init() {
    f := BTusculum图斯库卢姆.(*图斯库卢姆TusculumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tusculum",
		TitleName: "图斯库卢姆",
		TitleCode: "b_tusculum",
	}
}
