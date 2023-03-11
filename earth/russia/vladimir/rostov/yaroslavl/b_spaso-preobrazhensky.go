package yaroslavl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯帕索普列奥布拉任斯基Spaso_preobrazhenskyBarony struct {
	feud.BaseBarony
}

var BSpaso_preobrazhensky斯帕索普列奥布拉任斯基 feud.Barony = &斯帕索普列奥布拉任斯基Spaso_preobrazhenskyBarony{}

func init() {
    f := BSpaso_preobrazhensky斯帕索普列奥布拉任斯基.(*斯帕索普列奥布拉任斯基Spaso_preobrazhenskyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "spaso_preobrazhensky",
		TitleName: "斯帕索普列奥布拉任斯基",
		TitleCode: "b_spaso-preobrazhensky",
	}
}
