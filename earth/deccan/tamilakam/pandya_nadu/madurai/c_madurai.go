package madurai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MaduraiCounty interface {
    feud.County
    BAnaimalai阿奈马莱() 	feud.Barony
    BKilakarai吉罗伽罗() 	feud.Barony
    BMadurai摩菟来() 	feud.Barony
    BRamesvaram罗迷湿伐蓝() 	feud.Barony
    BRamnadhapuram愣那陀补愣() 	feud.Barony
    BSrivilliputtur斯里毗利布杜尔() 	feud.Barony
    BTiruparankunram提卢波兰贡愣() 	feud.Barony
}

type 摩菟来MaduraiCounty struct {
	feud.BaseCounty
	阿奈马莱Anaimalai 	feud.Barony
	吉罗伽罗Kilakarai 	feud.Barony
	摩菟来Madurai 	feud.Barony
	罗迷湿伐蓝Ramesvaram 	feud.Barony
	愣那陀补愣Ramnadhapuram 	feud.Barony
	斯里毗利布杜尔Srivilliputtur 	feud.Barony
	提卢波兰贡愣Tiruparankunram 	feud.Barony
}

func (c *摩菟来MaduraiCounty) BAnaimalai阿奈马莱() feud.Barony {
	return c.阿奈马莱Anaimalai
}
    
func (c *摩菟来MaduraiCounty) BKilakarai吉罗伽罗() feud.Barony {
	return c.吉罗伽罗Kilakarai
}
    
func (c *摩菟来MaduraiCounty) BMadurai摩菟来() feud.Barony {
	return c.摩菟来Madurai
}
    
func (c *摩菟来MaduraiCounty) BRamesvaram罗迷湿伐蓝() feud.Barony {
	return c.罗迷湿伐蓝Ramesvaram
}
    
func (c *摩菟来MaduraiCounty) BRamnadhapuram愣那陀补愣() feud.Barony {
	return c.愣那陀补愣Ramnadhapuram
}
    
func (c *摩菟来MaduraiCounty) BSrivilliputtur斯里毗利布杜尔() feud.Barony {
	return c.斯里毗利布杜尔Srivilliputtur
}
    
func (c *摩菟来MaduraiCounty) BTiruparankunram提卢波兰贡愣() feud.Barony {
	return c.提卢波兰贡愣Tiruparankunram
}
    
var CMadurai摩菟来 MaduraiCounty = &摩菟来MaduraiCounty{}

func init() {
	f := CMadurai摩菟来.(*摩菟来MaduraiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1112",
		Title:     "madurai",
		TitleName: "摩菟来",
		TitleCode: "c_madurai",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿奈马莱Anaimalai = BAnaimalai阿奈马莱
	f.阿奈马莱Anaimalai.SetParent(f)
	
	f.吉罗伽罗Kilakarai = BKilakarai吉罗伽罗
	f.吉罗伽罗Kilakarai.SetParent(f)
	
	f.摩菟来Madurai = BMadurai摩菟来
	f.摩菟来Madurai.SetParent(f)
	
	f.罗迷湿伐蓝Ramesvaram = BRamesvaram罗迷湿伐蓝
	f.罗迷湿伐蓝Ramesvaram.SetParent(f)
	
	f.愣那陀补愣Ramnadhapuram = BRamnadhapuram愣那陀补愣
	f.愣那陀补愣Ramnadhapuram.SetParent(f)
	
	f.斯里毗利布杜尔Srivilliputtur = BSrivilliputtur斯里毗利布杜尔
	f.斯里毗利布杜尔Srivilliputtur.SetParent(f)
	
	f.提卢波兰贡愣Tiruparankunram = BTiruparankunram提卢波兰贡愣
	f.提卢波兰贡愣Tiruparankunram.SetParent(f)
	
}
