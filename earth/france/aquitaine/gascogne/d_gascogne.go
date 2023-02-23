package gascogne

import (
	"github.com/thalesfu/CK2Commands/earth/france/aquitaine/gascogne/albret"
	"github.com/thalesfu/CK2Commands/earth/france/aquitaine/gascogne/armagnac"
	"github.com/thalesfu/CK2Commands/earth/france/aquitaine/gascogne/bearn"
	"github.com/thalesfu/CK2Commands/earth/france/aquitaine/gascogne/labourd"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GascogneDuke interface {
	feud.Duke
	CAlbret马桑() albret.AlbretCounty
	CArmagnac阿马尼亚克() armagnac.ArmagnacCounty
	CBearn贝阿恩() bearn.BearnCounty
	CLabourd达克斯() labourd.LabourdCounty
}

type 加斯科涅GascogneDuke struct {
	feud.BaseDuke
	马桑Albret      albret.AlbretCounty
	阿马尼亚克Armagnac armagnac.ArmagnacCounty
	贝阿恩Bearn      bearn.BearnCounty
	达克斯Labourd    labourd.LabourdCounty
}

func (d *加斯科涅GascogneDuke) CAlbret马桑() albret.AlbretCounty {
	return d.马桑Albret
}

func (d *加斯科涅GascogneDuke) CArmagnac阿马尼亚克() armagnac.ArmagnacCounty {
	return d.阿马尼亚克Armagnac
}

func (d *加斯科涅GascogneDuke) CBearn贝阿恩() bearn.BearnCounty {
	return d.贝阿恩Bearn
}

func (d *加斯科涅GascogneDuke) CLabourd达克斯() labourd.LabourdCounty {
	return d.达克斯Labourd
}

var DGascogne加斯科涅 GascogneDuke = &加斯科涅GascogneDuke{}

func init() {
	f := DGascogne加斯科涅.(*加斯科涅GascogneDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "gascogne",
		TitleName: "加斯科涅",
		TitleCode: "d_gascogne",
		Counties:  map[string]feud.County{},
	}

	f.马桑Albret = albret.CAlbret马桑
	f.马桑Albret.SetParent(f)

	f.阿马尼亚克Armagnac = armagnac.CArmagnac阿马尼亚克
	f.阿马尼亚克Armagnac.SetParent(f)

	f.贝阿恩Bearn = bearn.CBearn贝阿恩
	f.贝阿恩Bearn.SetParent(f)

	f.达克斯Labourd = labourd.CLabourd达克斯
	f.达克斯Labourd.SetParent(f)

}
