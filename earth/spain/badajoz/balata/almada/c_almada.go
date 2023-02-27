package almada

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AlmadaCounty interface {
    feud.County
    BAlcochete阿尔科谢泰() 	feud.Barony
    BAlmada阿尔马达() 	feud.Barony
    BMontijo蒙蒂霍() 	feud.Barony
    BPalmela帕尔梅拉() 	feud.Barony
    BPegoes佩戈埃斯() 	feud.Barony
    BSesimbra赛辛布拉() 	feud.Barony
    BSetubal塞图巴尔() 	feud.Barony
}

type 阿尔马达AlmadaCounty struct {
	feud.BaseCounty
	阿尔科谢泰Alcochete 	feud.Barony
	阿尔马达Almada 	feud.Barony
	蒙蒂霍Montijo 	feud.Barony
	帕尔梅拉Palmela 	feud.Barony
	佩戈埃斯Pegoes 	feud.Barony
	赛辛布拉Sesimbra 	feud.Barony
	塞图巴尔Setubal 	feud.Barony
}

func (c *阿尔马达AlmadaCounty) BAlcochete阿尔科谢泰() feud.Barony {
	return c.阿尔科谢泰Alcochete
}
    
func (c *阿尔马达AlmadaCounty) BAlmada阿尔马达() feud.Barony {
	return c.阿尔马达Almada
}
    
func (c *阿尔马达AlmadaCounty) BMontijo蒙蒂霍() feud.Barony {
	return c.蒙蒂霍Montijo
}
    
func (c *阿尔马达AlmadaCounty) BPalmela帕尔梅拉() feud.Barony {
	return c.帕尔梅拉Palmela
}
    
func (c *阿尔马达AlmadaCounty) BPegoes佩戈埃斯() feud.Barony {
	return c.佩戈埃斯Pegoes
}
    
func (c *阿尔马达AlmadaCounty) BSesimbra赛辛布拉() feud.Barony {
	return c.赛辛布拉Sesimbra
}
    
func (c *阿尔马达AlmadaCounty) BSetubal塞图巴尔() feud.Barony {
	return c.塞图巴尔Setubal
}
    
var CAlmada阿尔马达 AlmadaCounty = &阿尔马达AlmadaCounty{}

func init() {
	f := CAlmada阿尔马达.(*阿尔马达AlmadaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "2013",
		Title:     "almada",
		TitleName: "阿尔马达",
		TitleCode: "c_almada",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔科谢泰Alcochete = BAlcochete阿尔科谢泰
	f.阿尔科谢泰Alcochete.SetParent(f)
	
	f.阿尔马达Almada = BAlmada阿尔马达
	f.阿尔马达Almada.SetParent(f)
	
	f.蒙蒂霍Montijo = BMontijo蒙蒂霍
	f.蒙蒂霍Montijo.SetParent(f)
	
	f.帕尔梅拉Palmela = BPalmela帕尔梅拉
	f.帕尔梅拉Palmela.SetParent(f)
	
	f.佩戈埃斯Pegoes = BPegoes佩戈埃斯
	f.佩戈埃斯Pegoes.SetParent(f)
	
	f.赛辛布拉Sesimbra = BSesimbra赛辛布拉
	f.赛辛布拉Sesimbra.SetParent(f)
	
	f.塞图巴尔Setubal = BSetubal塞图巴尔
	f.塞图巴尔Setubal.SetParent(f)
	
}
