package malaga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MalagaCounty interface {
    feud.County
    BAntequera安特克拉() 	feud.Barony
    BBenalmadena贝纳尔马德纳() 	feud.Barony
    BCartajima卡尔塔希马() 	feud.Barony
    BCoin科因() 	feud.Barony
    BMalaga马拉加() 	feud.Barony
    BSuel苏埃尔() 	feud.Barony
    BTamisa塔米萨() 	feud.Barony
    BVelezmalaga贝莱斯_马拉加() 	feud.Barony
}

type 马拉加MalagaCounty struct {
	feud.BaseCounty
	安特克拉Antequera 	feud.Barony
	贝纳尔马德纳Benalmadena 	feud.Barony
	卡尔塔希马Cartajima 	feud.Barony
	科因Coin 	feud.Barony
	马拉加Malaga 	feud.Barony
	苏埃尔Suel 	feud.Barony
	塔米萨Tamisa 	feud.Barony
	贝莱斯_马拉加Velezmalaga 	feud.Barony
}

func (c *马拉加MalagaCounty) BAntequera安特克拉() feud.Barony {
	return c.安特克拉Antequera
}
    
func (c *马拉加MalagaCounty) BBenalmadena贝纳尔马德纳() feud.Barony {
	return c.贝纳尔马德纳Benalmadena
}
    
func (c *马拉加MalagaCounty) BCartajima卡尔塔希马() feud.Barony {
	return c.卡尔塔希马Cartajima
}
    
func (c *马拉加MalagaCounty) BCoin科因() feud.Barony {
	return c.科因Coin
}
    
func (c *马拉加MalagaCounty) BMalaga马拉加() feud.Barony {
	return c.马拉加Malaga
}
    
func (c *马拉加MalagaCounty) BSuel苏埃尔() feud.Barony {
	return c.苏埃尔Suel
}
    
func (c *马拉加MalagaCounty) BTamisa塔米萨() feud.Barony {
	return c.塔米萨Tamisa
}
    
func (c *马拉加MalagaCounty) BVelezmalaga贝莱斯_马拉加() feud.Barony {
	return c.贝莱斯_马拉加Velezmalaga
}
    
var CMalaga马拉加 MalagaCounty = &马拉加MalagaCounty{}

func init() {
	f := CMalaga马拉加.(*马拉加MalagaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "167",
		Title:     "malaga",
		TitleName: "马拉加",
		TitleCode: "c_malaga",
		Baronies:  map[string]feud.Barony{},
	}

	f.安特克拉Antequera = BAntequera安特克拉
	f.安特克拉Antequera.SetParent(f)
	
	f.贝纳尔马德纳Benalmadena = BBenalmadena贝纳尔马德纳
	f.贝纳尔马德纳Benalmadena.SetParent(f)
	
	f.卡尔塔希马Cartajima = BCartajima卡尔塔希马
	f.卡尔塔希马Cartajima.SetParent(f)
	
	f.科因Coin = BCoin科因
	f.科因Coin.SetParent(f)
	
	f.马拉加Malaga = BMalaga马拉加
	f.马拉加Malaga.SetParent(f)
	
	f.苏埃尔Suel = BSuel苏埃尔
	f.苏埃尔Suel.SetParent(f)
	
	f.塔米萨Tamisa = BTamisa塔米萨
	f.塔米萨Tamisa.SetParent(f)
	
	f.贝莱斯_马拉加Velezmalaga = BVelezmalaga贝莱斯_马拉加
	f.贝莱斯_马拉加Velezmalaga.SetParent(f)
	
}
