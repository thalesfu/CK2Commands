package holstein

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HolsteinCounty interface {
	feud.County
	BGluckstadt格吕克施塔特() feud.Barony
	BGottorp戈托普() feud.Barony
	BKiel基尔() feud.Barony
	BLauenburg劳恩堡() feud.Barony
	BReinholdsburg莱因霍尔德堡() feud.Barony
	BSegeberg塞格贝格() feud.Barony
}

type 荷尔斯泰因HolsteinCounty struct {
	feud.BaseCounty
	格吕克施塔特Gluckstadt    feud.Barony
	戈托普Gottorp          feud.Barony
	基尔Kiel              feud.Barony
	劳恩堡Lauenburg        feud.Barony
	莱因霍尔德堡Reinholdsburg feud.Barony
	塞格贝格Segeberg        feud.Barony
}

func (c *荷尔斯泰因HolsteinCounty) BGluckstadt格吕克施塔特() feud.Barony {
	return c.格吕克施塔特Gluckstadt
}

func (c *荷尔斯泰因HolsteinCounty) BGottorp戈托普() feud.Barony {
	return c.戈托普Gottorp
}

func (c *荷尔斯泰因HolsteinCounty) BKiel基尔() feud.Barony {
	return c.基尔Kiel
}

func (c *荷尔斯泰因HolsteinCounty) BLauenburg劳恩堡() feud.Barony {
	return c.劳恩堡Lauenburg
}

func (c *荷尔斯泰因HolsteinCounty) BReinholdsburg莱因霍尔德堡() feud.Barony {
	return c.莱因霍尔德堡Reinholdsburg
}

func (c *荷尔斯泰因HolsteinCounty) BSegeberg塞格贝格() feud.Barony {
	return c.塞格贝格Segeberg
}

var CHolstein荷尔斯泰因 HolsteinCounty = &荷尔斯泰因HolsteinCounty{}

func init() {
	f := CHolstein荷尔斯泰因.(*荷尔斯泰因HolsteinCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "263",
		Title:     "holstein",
		TitleName: "荷尔斯泰因",
		TitleCode: "c_holstein",
		Baronies:  map[string]feud.Barony{},
	}

	f.格吕克施塔特Gluckstadt = BGluckstadt格吕克施塔特
	f.格吕克施塔特Gluckstadt.SetParent(f)

	f.戈托普Gottorp = BGottorp戈托普
	f.戈托普Gottorp.SetParent(f)

	f.基尔Kiel = BKiel基尔
	f.基尔Kiel.SetParent(f)

	f.劳恩堡Lauenburg = BLauenburg劳恩堡
	f.劳恩堡Lauenburg.SetParent(f)

	f.莱因霍尔德堡Reinholdsburg = BReinholdsburg莱因霍尔德堡
	f.莱因霍尔德堡Reinholdsburg.SetParent(f)

	f.塞格贝格Segeberg = BSegeberg塞格贝格
	f.塞格贝格Segeberg.SetParent(f)

}
