package porto

import (
	"github.com/thalesfu/CK2Commands/earth/spain/portugal/porto/braganza"
	"github.com/thalesfu/CK2Commands/earth/spain/portugal/porto/coimbra"
	"github.com/thalesfu/CK2Commands/earth/spain/portugal/porto/porto"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PortoDuke interface {
	feud.Duke
	CBraganza布拉干萨() braganza.BraganzaCounty
	CCoimbra科英布拉() coimbra.CoimbraCounty
	CPorto波尔图() porto.PortoCounty
}

type 波图斯卡莱PortoDuke struct {
	feud.BaseDuke
	布拉干萨Braganza braganza.BraganzaCounty
	科英布拉Coimbra  coimbra.CoimbraCounty
	波尔图Porto     porto.PortoCounty
}

func (d *波图斯卡莱PortoDuke) CBraganza布拉干萨() braganza.BraganzaCounty {
	return d.布拉干萨Braganza
}

func (d *波图斯卡莱PortoDuke) CCoimbra科英布拉() coimbra.CoimbraCounty {
	return d.科英布拉Coimbra
}

func (d *波图斯卡莱PortoDuke) CPorto波尔图() porto.PortoCounty {
	return d.波尔图Porto
}

var DPorto波图斯卡莱 PortoDuke = &波图斯卡莱PortoDuke{}

func init() {
	f := DPorto波图斯卡莱.(*波图斯卡莱PortoDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "porto",
		TitleName: "波图斯卡莱",
		TitleCode: "d_porto",
		Counties:  map[string]feud.County{},
	}

	f.布拉干萨Braganza = braganza.CBraganza布拉干萨
	f.布拉干萨Braganza.SetParent(f)

	f.科英布拉Coimbra = coimbra.CCoimbra科英布拉
	f.科英布拉Coimbra.SetParent(f)

	f.波尔图Porto = porto.CPorto波尔图
	f.波尔图Porto.SetParent(f)

}
