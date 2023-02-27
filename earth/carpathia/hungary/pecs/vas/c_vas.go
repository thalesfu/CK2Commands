package vas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VasCounty interface {
    feud.County
    BCelldomolk采尔德默尔克() 	feud.Barony
    BKormend克尔门德() 	feud.Barony
    BKoszeg克塞格() 	feud.Barony
    BNemetujvar内迈特新堡() 	feud.Barony
    BSarvar沙尔堡() 	feud.Barony
    BSzentgotthard圣戈特哈德() 	feud.Barony
    BSzombathely松博特海伊() 	feud.Barony
    BVasvar沃什堡() 	feud.Barony
}

type 沃什VasCounty struct {
	feud.BaseCounty
	采尔德默尔克Celldomolk 	feud.Barony
	克尔门德Kormend 	feud.Barony
	克塞格Koszeg 	feud.Barony
	内迈特新堡Nemetujvar 	feud.Barony
	沙尔堡Sarvar 	feud.Barony
	圣戈特哈德Szentgotthard 	feud.Barony
	松博特海伊Szombathely 	feud.Barony
	沃什堡Vasvar 	feud.Barony
}

func (c *沃什VasCounty) BCelldomolk采尔德默尔克() feud.Barony {
	return c.采尔德默尔克Celldomolk
}
    
func (c *沃什VasCounty) BKormend克尔门德() feud.Barony {
	return c.克尔门德Kormend
}
    
func (c *沃什VasCounty) BKoszeg克塞格() feud.Barony {
	return c.克塞格Koszeg
}
    
func (c *沃什VasCounty) BNemetujvar内迈特新堡() feud.Barony {
	return c.内迈特新堡Nemetujvar
}
    
func (c *沃什VasCounty) BSarvar沙尔堡() feud.Barony {
	return c.沙尔堡Sarvar
}
    
func (c *沃什VasCounty) BSzentgotthard圣戈特哈德() feud.Barony {
	return c.圣戈特哈德Szentgotthard
}
    
func (c *沃什VasCounty) BSzombathely松博特海伊() feud.Barony {
	return c.松博特海伊Szombathely
}
    
func (c *沃什VasCounty) BVasvar沃什堡() feud.Barony {
	return c.沃什堡Vasvar
}
    
var CVas沃什 VasCounty = &沃什VasCounty{}

func init() {
	f := CVas沃什.(*沃什VasCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "454",
		Title:     "vas",
		TitleName: "沃什",
		TitleCode: "c_vas",
		Baronies:  map[string]feud.Barony{},
	}

	f.采尔德默尔克Celldomolk = BCelldomolk采尔德默尔克
	f.采尔德默尔克Celldomolk.SetParent(f)
	
	f.克尔门德Kormend = BKormend克尔门德
	f.克尔门德Kormend.SetParent(f)
	
	f.克塞格Koszeg = BKoszeg克塞格
	f.克塞格Koszeg.SetParent(f)
	
	f.内迈特新堡Nemetujvar = BNemetujvar内迈特新堡
	f.内迈特新堡Nemetujvar.SetParent(f)
	
	f.沙尔堡Sarvar = BSarvar沙尔堡
	f.沙尔堡Sarvar.SetParent(f)
	
	f.圣戈特哈德Szentgotthard = BSzentgotthard圣戈特哈德
	f.圣戈特哈德Szentgotthard.SetParent(f)
	
	f.松博特海伊Szombathely = BSzombathely松博特海伊
	f.松博特海伊Szombathely.SetParent(f)
	
	f.沃什堡Vasvar = BVasvar沃什堡
	f.沃什堡Vasvar.SetParent(f)
	
}
