package sussex

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SussexCounty interface {
    feud.County
    BArundel阿伦德尔() 	feud.Barony
    BBodiam博迪亚姆() 	feud.Barony
    BBramber布兰伯() 	feud.Barony
    BChichester奇切斯特() 	feud.Barony
    BHastings黑斯廷斯() 	feud.Barony
    BLewes刘易斯() 	feud.Barony
    BPevensey佩文西() 	feud.Barony
    BRye拉伊() 	feud.Barony
}

type 萨塞克斯SussexCounty struct {
	feud.BaseCounty
	阿伦德尔Arundel 	feud.Barony
	博迪亚姆Bodiam 	feud.Barony
	布兰伯Bramber 	feud.Barony
	奇切斯特Chichester 	feud.Barony
	黑斯廷斯Hastings 	feud.Barony
	刘易斯Lewes 	feud.Barony
	佩文西Pevensey 	feud.Barony
	拉伊Rye 	feud.Barony
}

func (c *萨塞克斯SussexCounty) BArundel阿伦德尔() feud.Barony {
	return c.阿伦德尔Arundel
}
    
func (c *萨塞克斯SussexCounty) BBodiam博迪亚姆() feud.Barony {
	return c.博迪亚姆Bodiam
}
    
func (c *萨塞克斯SussexCounty) BBramber布兰伯() feud.Barony {
	return c.布兰伯Bramber
}
    
func (c *萨塞克斯SussexCounty) BChichester奇切斯特() feud.Barony {
	return c.奇切斯特Chichester
}
    
func (c *萨塞克斯SussexCounty) BHastings黑斯廷斯() feud.Barony {
	return c.黑斯廷斯Hastings
}
    
func (c *萨塞克斯SussexCounty) BLewes刘易斯() feud.Barony {
	return c.刘易斯Lewes
}
    
func (c *萨塞克斯SussexCounty) BPevensey佩文西() feud.Barony {
	return c.佩文西Pevensey
}
    
func (c *萨塞克斯SussexCounty) BRye拉伊() feud.Barony {
	return c.拉伊Rye
}
    
var CSussex萨塞克斯 SussexCounty = &萨塞克斯SussexCounty{}

func init() {
	f := CSussex萨塞克斯.(*萨塞克斯SussexCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "25",
		Title:     "sussex",
		TitleName: "萨塞克斯",
		TitleCode: "c_sussex",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿伦德尔Arundel = BArundel阿伦德尔
	f.阿伦德尔Arundel.SetParent(f)
	
	f.博迪亚姆Bodiam = BBodiam博迪亚姆
	f.博迪亚姆Bodiam.SetParent(f)
	
	f.布兰伯Bramber = BBramber布兰伯
	f.布兰伯Bramber.SetParent(f)
	
	f.奇切斯特Chichester = BChichester奇切斯特
	f.奇切斯特Chichester.SetParent(f)
	
	f.黑斯廷斯Hastings = BHastings黑斯廷斯
	f.黑斯廷斯Hastings.SetParent(f)
	
	f.刘易斯Lewes = BLewes刘易斯
	f.刘易斯Lewes.SetParent(f)
	
	f.佩文西Pevensey = BPevensey佩文西
	f.佩文西Pevensey.SetParent(f)
	
	f.拉伊Rye = BRye拉伊
	f.拉伊Rye.SetParent(f)
	
}
