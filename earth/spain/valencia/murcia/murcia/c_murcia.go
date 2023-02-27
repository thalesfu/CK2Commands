package murcia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MurciaCounty interface {
    feud.County
    BAlcantarilla阿尔坎塔里利亚() 	feud.Barony
    BCartagena卡塔赫纳() 	feud.Barony
    BLorca洛尔卡() 	feud.Barony
    BMedinasiyasa锡亚萨城() 	feud.Barony
    BMolinalaseca拉塞卡() 	feud.Barony
    BMurcia穆尔西亚() 	feud.Barony
    BNogalte诺加尔特() 	feud.Barony
    BYecla耶克拉() 	feud.Barony
}

type 穆尔西亚MurciaCounty struct {
	feud.BaseCounty
	阿尔坎塔里利亚Alcantarilla 	feud.Barony
	卡塔赫纳Cartagena 	feud.Barony
	洛尔卡Lorca 	feud.Barony
	锡亚萨城Medinasiyasa 	feud.Barony
	拉塞卡Molinalaseca 	feud.Barony
	穆尔西亚Murcia 	feud.Barony
	诺加尔特Nogalte 	feud.Barony
	耶克拉Yecla 	feud.Barony
}

func (c *穆尔西亚MurciaCounty) BAlcantarilla阿尔坎塔里利亚() feud.Barony {
	return c.阿尔坎塔里利亚Alcantarilla
}
    
func (c *穆尔西亚MurciaCounty) BCartagena卡塔赫纳() feud.Barony {
	return c.卡塔赫纳Cartagena
}
    
func (c *穆尔西亚MurciaCounty) BLorca洛尔卡() feud.Barony {
	return c.洛尔卡Lorca
}
    
func (c *穆尔西亚MurciaCounty) BMedinasiyasa锡亚萨城() feud.Barony {
	return c.锡亚萨城Medinasiyasa
}
    
func (c *穆尔西亚MurciaCounty) BMolinalaseca拉塞卡() feud.Barony {
	return c.拉塞卡Molinalaseca
}
    
func (c *穆尔西亚MurciaCounty) BMurcia穆尔西亚() feud.Barony {
	return c.穆尔西亚Murcia
}
    
func (c *穆尔西亚MurciaCounty) BNogalte诺加尔特() feud.Barony {
	return c.诺加尔特Nogalte
}
    
func (c *穆尔西亚MurciaCounty) BYecla耶克拉() feud.Barony {
	return c.耶克拉Yecla
}
    
var CMurcia穆尔西亚 MurciaCounty = &穆尔西亚MurciaCounty{}

func init() {
	f := CMurcia穆尔西亚.(*穆尔西亚MurciaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "169",
		Title:     "murcia",
		TitleName: "穆尔西亚",
		TitleCode: "c_murcia",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔坎塔里利亚Alcantarilla = BAlcantarilla阿尔坎塔里利亚
	f.阿尔坎塔里利亚Alcantarilla.SetParent(f)
	
	f.卡塔赫纳Cartagena = BCartagena卡塔赫纳
	f.卡塔赫纳Cartagena.SetParent(f)
	
	f.洛尔卡Lorca = BLorca洛尔卡
	f.洛尔卡Lorca.SetParent(f)
	
	f.锡亚萨城Medinasiyasa = BMedinasiyasa锡亚萨城
	f.锡亚萨城Medinasiyasa.SetParent(f)
	
	f.拉塞卡Molinalaseca = BMolinalaseca拉塞卡
	f.拉塞卡Molinalaseca.SetParent(f)
	
	f.穆尔西亚Murcia = BMurcia穆尔西亚
	f.穆尔西亚Murcia.SetParent(f)
	
	f.诺加尔特Nogalte = BNogalte诺加尔特
	f.诺加尔特Nogalte.SetParent(f)
	
	f.耶克拉Yecla = BYecla耶克拉
	f.耶克拉Yecla.SetParent(f)
	
}
