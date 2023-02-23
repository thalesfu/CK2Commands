package mesembria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MesembriaCounty interface {
	feud.County
	BAetos艾托斯() feud.Barony
	BAnchilios安基里奥斯() feud.Barony
	BBourgas布尔加斯() feud.Barony
	BMesembria墨森布里亚() feud.Barony
	BOdessos奥德索斯() feud.Barony
	BSozopolis索佐波利斯() feud.Barony
	BValchidol沃尔奇多尔() feud.Barony
	BVarna瓦尔纳() feud.Barony
}

type 墨森布里亚MesembriaCounty struct {
	feud.BaseCounty
	艾托斯Aetos       feud.Barony
	安基里奥斯Anchilios feud.Barony
	布尔加斯Bourgas    feud.Barony
	墨森布里亚Mesembria feud.Barony
	奥德索斯Odessos    feud.Barony
	索佐波利斯Sozopolis feud.Barony
	沃尔奇多尔Valchidol feud.Barony
	瓦尔纳Varna       feud.Barony
}

func (c *墨森布里亚MesembriaCounty) BAetos艾托斯() feud.Barony {
	return c.艾托斯Aetos
}

func (c *墨森布里亚MesembriaCounty) BAnchilios安基里奥斯() feud.Barony {
	return c.安基里奥斯Anchilios
}

func (c *墨森布里亚MesembriaCounty) BBourgas布尔加斯() feud.Barony {
	return c.布尔加斯Bourgas
}

func (c *墨森布里亚MesembriaCounty) BMesembria墨森布里亚() feud.Barony {
	return c.墨森布里亚Mesembria
}

func (c *墨森布里亚MesembriaCounty) BOdessos奥德索斯() feud.Barony {
	return c.奥德索斯Odessos
}

func (c *墨森布里亚MesembriaCounty) BSozopolis索佐波利斯() feud.Barony {
	return c.索佐波利斯Sozopolis
}

func (c *墨森布里亚MesembriaCounty) BValchidol沃尔奇多尔() feud.Barony {
	return c.沃尔奇多尔Valchidol
}

func (c *墨森布里亚MesembriaCounty) BVarna瓦尔纳() feud.Barony {
	return c.瓦尔纳Varna
}

var CMesembria墨森布里亚 MesembriaCounty = &墨森布里亚MesembriaCounty{}

func init() {
	f := CMesembria墨森布里亚.(*墨森布里亚MesembriaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "498",
		Title:     "mesembria",
		TitleName: "墨森布里亚",
		TitleCode: "c_mesembria",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾托斯Aetos = BAetos艾托斯
	f.艾托斯Aetos.SetParent(f)

	f.安基里奥斯Anchilios = BAnchilios安基里奥斯
	f.安基里奥斯Anchilios.SetParent(f)

	f.布尔加斯Bourgas = BBourgas布尔加斯
	f.布尔加斯Bourgas.SetParent(f)

	f.墨森布里亚Mesembria = BMesembria墨森布里亚
	f.墨森布里亚Mesembria.SetParent(f)

	f.奥德索斯Odessos = BOdessos奥德索斯
	f.奥德索斯Odessos.SetParent(f)

	f.索佐波利斯Sozopolis = BSozopolis索佐波利斯
	f.索佐波利斯Sozopolis.SetParent(f)

	f.沃尔奇多尔Valchidol = BValchidol沃尔奇多尔
	f.沃尔奇多尔Valchidol.SetParent(f)

	f.瓦尔纳Varna = BVarna瓦尔纳
	f.瓦尔纳Varna.SetParent(f)

}
