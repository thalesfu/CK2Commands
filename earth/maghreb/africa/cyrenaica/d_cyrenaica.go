package cyrenaica

import (
	"github.com/thalesfu/CK2Commands/earth/maghreb/africa/cyrenaica/benghazi"
	"github.com/thalesfu/CK2Commands/earth/maghreb/africa/cyrenaica/cyrenaica"
	"github.com/thalesfu/CK2Commands/earth/maghreb/africa/cyrenaica/senoussi"
	"github.com/thalesfu/CK2Commands/earth/maghreb/africa/cyrenaica/tobruk"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CyrenaicaDuke interface {
	feud.Duke
	CBenghazi班加西() benghazi.BenghaziCounty
	CCyrenaica昔兰尼加() cyrenaica.CyrenaicaCounty
	CSenoussi塞努西() senoussi.SenoussiCounty
	CTobruk图卜鲁格() tobruk.TobrukCounty
}

type 昔兰尼加CyrenaicaDuke struct {
	feud.BaseDuke
	班加西Benghazi   benghazi.BenghaziCounty
	昔兰尼加Cyrenaica cyrenaica.CyrenaicaCounty
	塞努西Senoussi   senoussi.SenoussiCounty
	图卜鲁格Tobruk    tobruk.TobrukCounty
}

func (d *昔兰尼加CyrenaicaDuke) CBenghazi班加西() benghazi.BenghaziCounty {
	return d.班加西Benghazi
}

func (d *昔兰尼加CyrenaicaDuke) CCyrenaica昔兰尼加() cyrenaica.CyrenaicaCounty {
	return d.昔兰尼加Cyrenaica
}

func (d *昔兰尼加CyrenaicaDuke) CSenoussi塞努西() senoussi.SenoussiCounty {
	return d.塞努西Senoussi
}

func (d *昔兰尼加CyrenaicaDuke) CTobruk图卜鲁格() tobruk.TobrukCounty {
	return d.图卜鲁格Tobruk
}

var DCyrenaica昔兰尼加 CyrenaicaDuke = &昔兰尼加CyrenaicaDuke{}

func init() {
	f := DCyrenaica昔兰尼加.(*昔兰尼加CyrenaicaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "cyrenaica",
		TitleName: "昔兰尼加",
		TitleCode: "d_cyrenaica",
		Counties:  map[string]feud.County{},
	}

	f.班加西Benghazi = benghazi.CBenghazi班加西
	f.班加西Benghazi.SetParent(f)

	f.昔兰尼加Cyrenaica = cyrenaica.CCyrenaica昔兰尼加
	f.昔兰尼加Cyrenaica.SetParent(f)

	f.塞努西Senoussi = senoussi.CSenoussi塞努西
	f.塞努西Senoussi.SetParent(f)

	f.图卜鲁格Tobruk = tobruk.CTobruk图卜鲁格
	f.图卜鲁格Tobruk.SetParent(f)

}
