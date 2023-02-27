package caceres

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CaceresCounty interface {
    feud.County
    BAlburquerque阿尔布开克() 	feud.Barony
    BAlcuescar阿尔库埃斯卡尔() 	feud.Barony
    BAlia阿利亚() 	feud.Barony
    BArroyodelalluz阿罗约德拉卢斯() 	feud.Barony
    BCaceres卡塞雷斯() 	feud.Barony
    BGuadalupe瓜达卢佩() 	feud.Barony
    BLogrosan洛格罗桑() 	feud.Barony
    BTrujillo特鲁希略() 	feud.Barony
}

type 卡塞雷斯CaceresCounty struct {
	feud.BaseCounty
	阿尔布开克Alburquerque 	feud.Barony
	阿尔库埃斯卡尔Alcuescar 	feud.Barony
	阿利亚Alia 	feud.Barony
	阿罗约德拉卢斯Arroyodelalluz 	feud.Barony
	卡塞雷斯Caceres 	feud.Barony
	瓜达卢佩Guadalupe 	feud.Barony
	洛格罗桑Logrosan 	feud.Barony
	特鲁希略Trujillo 	feud.Barony
}

func (c *卡塞雷斯CaceresCounty) BAlburquerque阿尔布开克() feud.Barony {
	return c.阿尔布开克Alburquerque
}
    
func (c *卡塞雷斯CaceresCounty) BAlcuescar阿尔库埃斯卡尔() feud.Barony {
	return c.阿尔库埃斯卡尔Alcuescar
}
    
func (c *卡塞雷斯CaceresCounty) BAlia阿利亚() feud.Barony {
	return c.阿利亚Alia
}
    
func (c *卡塞雷斯CaceresCounty) BArroyodelalluz阿罗约德拉卢斯() feud.Barony {
	return c.阿罗约德拉卢斯Arroyodelalluz
}
    
func (c *卡塞雷斯CaceresCounty) BCaceres卡塞雷斯() feud.Barony {
	return c.卡塞雷斯Caceres
}
    
func (c *卡塞雷斯CaceresCounty) BGuadalupe瓜达卢佩() feud.Barony {
	return c.瓜达卢佩Guadalupe
}
    
func (c *卡塞雷斯CaceresCounty) BLogrosan洛格罗桑() feud.Barony {
	return c.洛格罗桑Logrosan
}
    
func (c *卡塞雷斯CaceresCounty) BTrujillo特鲁希略() feud.Barony {
	return c.特鲁希略Trujillo
}
    
var CCaceres卡塞雷斯 CaceresCounty = &卡塞雷斯CaceresCounty{}

func init() {
	f := CCaceres卡塞雷斯.(*卡塞雷斯CaceresCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "195",
		Title:     "caceres",
		TitleName: "卡塞雷斯",
		TitleCode: "c_caceres",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔布开克Alburquerque = BAlburquerque阿尔布开克
	f.阿尔布开克Alburquerque.SetParent(f)
	
	f.阿尔库埃斯卡尔Alcuescar = BAlcuescar阿尔库埃斯卡尔
	f.阿尔库埃斯卡尔Alcuescar.SetParent(f)
	
	f.阿利亚Alia = BAlia阿利亚
	f.阿利亚Alia.SetParent(f)
	
	f.阿罗约德拉卢斯Arroyodelalluz = BArroyodelalluz阿罗约德拉卢斯
	f.阿罗约德拉卢斯Arroyodelalluz.SetParent(f)
	
	f.卡塞雷斯Caceres = BCaceres卡塞雷斯
	f.卡塞雷斯Caceres.SetParent(f)
	
	f.瓜达卢佩Guadalupe = BGuadalupe瓜达卢佩
	f.瓜达卢佩Guadalupe.SetParent(f)
	
	f.洛格罗桑Logrosan = BLogrosan洛格罗桑
	f.洛格罗桑Logrosan.SetParent(f)
	
	f.特鲁希略Trujillo = BTrujillo特鲁希略
	f.特鲁希略Trujillo.SetParent(f)
	
}
