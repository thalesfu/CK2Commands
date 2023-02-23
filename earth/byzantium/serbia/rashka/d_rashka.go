package rashka

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/serbia/rashka/belgrade"
	"github.com/thalesfu/CK2Commands/earth/byzantium/serbia/rashka/hum"
	"github.com/thalesfu/CK2Commands/earth/byzantium/serbia/rashka/onogost"
	"github.com/thalesfu/CK2Commands/earth/byzantium/serbia/rashka/rashka"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RashkaDuke interface {
	feud.Duke
	CBelgrade贝尔格莱德() belgrade.BelgradeCounty
	CHum拉斯() hum.HumCounty
	COnogost奥诺戈什特() onogost.OnogostCounty
	CRashka拉什卡() rashka.RashkaCounty
}

type 拉什卡RashkaDuke struct {
	feud.BaseDuke
	贝尔格莱德Belgrade belgrade.BelgradeCounty
	拉斯Hum         hum.HumCounty
	奥诺戈什特Onogost  onogost.OnogostCounty
	拉什卡Rashka     rashka.RashkaCounty
}

func (d *拉什卡RashkaDuke) CBelgrade贝尔格莱德() belgrade.BelgradeCounty {
	return d.贝尔格莱德Belgrade
}

func (d *拉什卡RashkaDuke) CHum拉斯() hum.HumCounty {
	return d.拉斯Hum
}

func (d *拉什卡RashkaDuke) COnogost奥诺戈什特() onogost.OnogostCounty {
	return d.奥诺戈什特Onogost
}

func (d *拉什卡RashkaDuke) CRashka拉什卡() rashka.RashkaCounty {
	return d.拉什卡Rashka
}

var DRashka拉什卡 RashkaDuke = &拉什卡RashkaDuke{}

func init() {
	f := DRashka拉什卡.(*拉什卡RashkaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "rashka",
		TitleName: "拉什卡",
		TitleCode: "d_rashka",
		Counties:  map[string]feud.County{},
	}

	f.贝尔格莱德Belgrade = belgrade.CBelgrade贝尔格莱德
	f.贝尔格莱德Belgrade.SetParent(f)

	f.拉斯Hum = hum.CHum拉斯
	f.拉斯Hum.SetParent(f)

	f.奥诺戈什特Onogost = onogost.COnogost奥诺戈什特
	f.奥诺戈什特Onogost.SetParent(f)

	f.拉什卡Rashka = rashka.CRashka拉什卡
	f.拉什卡Rashka.SetParent(f)

}
