package toulouse

import (
	"github.com/thalesfu/CK2Commands/earth/france/aquitaine/toulouse/carcassonne"
	"github.com/thalesfu/CK2Commands/earth/france/aquitaine/toulouse/foix"
	"github.com/thalesfu/CK2Commands/earth/france/aquitaine/toulouse/montpellier"
	"github.com/thalesfu/CK2Commands/earth/france/aquitaine/toulouse/narbonne"
	"github.com/thalesfu/CK2Commands/earth/france/aquitaine/toulouse/rouergue"
	"github.com/thalesfu/CK2Commands/earth/france/aquitaine/toulouse/toulouse"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ToulouseDuke interface {
	feud.Duke
	CCarcassonne卡尔卡松() carcassonne.CarcassonneCounty
	CFoix富瓦() foix.FoixCounty
	CMontpellier莫吉奥() montpellier.MontpellierCounty
	CNarbonne纳博讷() narbonne.NarbonneCounty
	CRouergue鲁埃格() rouergue.RouergueCounty
	CToulouse图卢兹() toulouse.ToulouseCounty
}

type 图卢兹ToulouseDuke struct {
	feud.BaseDuke
	卡尔卡松Carcassonne carcassonne.CarcassonneCounty
	富瓦Foix          foix.FoixCounty
	莫吉奥Montpellier  montpellier.MontpellierCounty
	纳博讷Narbonne     narbonne.NarbonneCounty
	鲁埃格Rouergue     rouergue.RouergueCounty
	图卢兹Toulouse     toulouse.ToulouseCounty
}

func (d *图卢兹ToulouseDuke) CCarcassonne卡尔卡松() carcassonne.CarcassonneCounty {
	return d.卡尔卡松Carcassonne
}

func (d *图卢兹ToulouseDuke) CFoix富瓦() foix.FoixCounty {
	return d.富瓦Foix
}

func (d *图卢兹ToulouseDuke) CMontpellier莫吉奥() montpellier.MontpellierCounty {
	return d.莫吉奥Montpellier
}

func (d *图卢兹ToulouseDuke) CNarbonne纳博讷() narbonne.NarbonneCounty {
	return d.纳博讷Narbonne
}

func (d *图卢兹ToulouseDuke) CRouergue鲁埃格() rouergue.RouergueCounty {
	return d.鲁埃格Rouergue
}

func (d *图卢兹ToulouseDuke) CToulouse图卢兹() toulouse.ToulouseCounty {
	return d.图卢兹Toulouse
}

var DToulouse图卢兹 ToulouseDuke = &图卢兹ToulouseDuke{}

func init() {
	f := DToulouse图卢兹.(*图卢兹ToulouseDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "toulouse",
		TitleName: "图卢兹",
		TitleCode: "d_toulouse",
		Counties:  map[string]feud.County{},
	}

	f.卡尔卡松Carcassonne = carcassonne.CCarcassonne卡尔卡松
	f.卡尔卡松Carcassonne.SetParent(f)

	f.富瓦Foix = foix.CFoix富瓦
	f.富瓦Foix.SetParent(f)

	f.莫吉奥Montpellier = montpellier.CMontpellier莫吉奥
	f.莫吉奥Montpellier.SetParent(f)

	f.纳博讷Narbonne = narbonne.CNarbonne纳博讷
	f.纳博讷Narbonne.SetParent(f)

	f.鲁埃格Rouergue = rouergue.CRouergue鲁埃格
	f.鲁埃格Rouergue.SetParent(f)

	f.图卢兹Toulouse = toulouse.CToulouse图卢兹
	f.图卢兹Toulouse.SetParent(f)

}
