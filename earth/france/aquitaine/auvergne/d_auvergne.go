package auvergne

import (
	"github.com/thalesfu/CK2Commands/earth/france/aquitaine/auvergne/aurillac"
	"github.com/thalesfu/CK2Commands/earth/france/aquitaine/auvergne/auvergne"
	"github.com/thalesfu/CK2Commands/earth/france/aquitaine/auvergne/gevaudan"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AuvergneDuke interface {
	feud.Duke
	CAurillac欧里亚克() aurillac.AurillacCounty
	CAuvergne奥弗涅() auvergne.AuvergneCounty
	CGevaudan热沃当() gevaudan.GevaudanCounty
}

type 奥弗涅AuvergneDuke struct {
	feud.BaseDuke
	欧里亚克Aurillac aurillac.AurillacCounty
	奥弗涅Auvergne  auvergne.AuvergneCounty
	热沃当Gevaudan  gevaudan.GevaudanCounty
}

func (d *奥弗涅AuvergneDuke) CAurillac欧里亚克() aurillac.AurillacCounty {
	return d.欧里亚克Aurillac
}

func (d *奥弗涅AuvergneDuke) CAuvergne奥弗涅() auvergne.AuvergneCounty {
	return d.奥弗涅Auvergne
}

func (d *奥弗涅AuvergneDuke) CGevaudan热沃当() gevaudan.GevaudanCounty {
	return d.热沃当Gevaudan
}

var DAuvergne奥弗涅 AuvergneDuke = &奥弗涅AuvergneDuke{}

func init() {
	f := DAuvergne奥弗涅.(*奥弗涅AuvergneDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "auvergne",
		TitleName: "奥弗涅",
		TitleCode: "d_auvergne",
		Counties:  map[string]feud.County{},
	}

	f.欧里亚克Aurillac = aurillac.CAurillac欧里亚克
	f.欧里亚克Aurillac.SetParent(f)

	f.奥弗涅Auvergne = auvergne.CAuvergne奥弗涅
	f.奥弗涅Auvergne.SetParent(f)

	f.热沃当Gevaudan = gevaudan.CGevaudan热沃当
	f.热沃当Gevaudan.SetParent(f)

}
