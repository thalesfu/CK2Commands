package oppland

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway/oppland/gudbrandsdal"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway/oppland/hedmark"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway/oppland/oppland"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway/oppland/romerike"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OpplandDuke interface {
	feud.Duke
	CGudbrandsdal居德布兰河谷() gudbrandsdal.GudbrandsdalCounty
	CHedmark厄斯特谷地() hedmark.HedmarkCounty
	COppland海德马克() oppland.OpplandCounty
	CRomerike鲁默里克() romerike.RomerikeCounty
}

type 奥普兰OpplandDuke struct {
	feud.BaseDuke
	居德布兰河谷Gudbrandsdal gudbrandsdal.GudbrandsdalCounty
	厄斯特谷地Hedmark       hedmark.HedmarkCounty
	海德马克Oppland        oppland.OpplandCounty
	鲁默里克Romerike       romerike.RomerikeCounty
}

func (d *奥普兰OpplandDuke) CGudbrandsdal居德布兰河谷() gudbrandsdal.GudbrandsdalCounty {
	return d.居德布兰河谷Gudbrandsdal
}

func (d *奥普兰OpplandDuke) CHedmark厄斯特谷地() hedmark.HedmarkCounty {
	return d.厄斯特谷地Hedmark
}

func (d *奥普兰OpplandDuke) COppland海德马克() oppland.OpplandCounty {
	return d.海德马克Oppland
}

func (d *奥普兰OpplandDuke) CRomerike鲁默里克() romerike.RomerikeCounty {
	return d.鲁默里克Romerike
}

var DOppland奥普兰 OpplandDuke = &奥普兰OpplandDuke{}

func init() {
	f := DOppland奥普兰.(*奥普兰OpplandDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "oppland",
		TitleName: "奥普兰",
		TitleCode: "d_oppland",
		Counties:  map[string]feud.County{},
	}

	f.居德布兰河谷Gudbrandsdal = gudbrandsdal.CGudbrandsdal居德布兰河谷
	f.居德布兰河谷Gudbrandsdal.SetParent(f)

	f.厄斯特谷地Hedmark = hedmark.CHedmark厄斯特谷地
	f.厄斯特谷地Hedmark.SetParent(f)

	f.海德马克Oppland = oppland.COppland海德马克
	f.海德马克Oppland.SetParent(f)

	f.鲁默里克Romerike = romerike.CRomerike鲁默里克
	f.鲁默里克Romerike.SetParent(f)

}
