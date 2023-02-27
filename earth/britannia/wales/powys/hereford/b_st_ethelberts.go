package hereford

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣艾塞尔伯特St_ethelbertsBarony struct {
	feud.BaseBarony
}

var BSt_ethelberts圣艾塞尔伯特 feud.Barony = &圣艾塞尔伯特St_ethelbertsBarony{}

func init() {
    f := BSt_ethelberts圣艾塞尔伯特.(*圣艾塞尔伯特St_ethelbertsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_ethelberts",
		TitleName: "圣艾塞尔伯特",
		TitleCode: "b_st_ethelberts",
	}
}
