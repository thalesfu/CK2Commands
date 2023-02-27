package mallabhum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MallabhumCounty interface {
    feud.County
    BArch阿遮() 	feud.Barony
    BBaghmundi婆迦蒙提() 	feud.Barony
    BBishnupur毗湿奴城() 	feud.Barony
    BChandrakona旃陀罗拘拿() 	feud.Barony
    BDihar提诃罗() 	feud.Barony
    BGarhbeta迦罗吠多() 	feud.Barony
    BMandaran曼陀兰() 	feud.Barony
}

type 末罗浮摩MallabhumCounty struct {
	feud.BaseCounty
	阿遮Arch 	feud.Barony
	婆迦蒙提Baghmundi 	feud.Barony
	毗湿奴城Bishnupur 	feud.Barony
	旃陀罗拘拿Chandrakona 	feud.Barony
	提诃罗Dihar 	feud.Barony
	迦罗吠多Garhbeta 	feud.Barony
	曼陀兰Mandaran 	feud.Barony
}

func (c *末罗浮摩MallabhumCounty) BArch阿遮() feud.Barony {
	return c.阿遮Arch
}
    
func (c *末罗浮摩MallabhumCounty) BBaghmundi婆迦蒙提() feud.Barony {
	return c.婆迦蒙提Baghmundi
}
    
func (c *末罗浮摩MallabhumCounty) BBishnupur毗湿奴城() feud.Barony {
	return c.毗湿奴城Bishnupur
}
    
func (c *末罗浮摩MallabhumCounty) BChandrakona旃陀罗拘拿() feud.Barony {
	return c.旃陀罗拘拿Chandrakona
}
    
func (c *末罗浮摩MallabhumCounty) BDihar提诃罗() feud.Barony {
	return c.提诃罗Dihar
}
    
func (c *末罗浮摩MallabhumCounty) BGarhbeta迦罗吠多() feud.Barony {
	return c.迦罗吠多Garhbeta
}
    
func (c *末罗浮摩MallabhumCounty) BMandaran曼陀兰() feud.Barony {
	return c.曼陀兰Mandaran
}
    
var CMallabhum末罗浮摩 MallabhumCounty = &末罗浮摩MallabhumCounty{}

func init() {
	f := CMallabhum末罗浮摩.(*末罗浮摩MallabhumCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1238",
		Title:     "mallabhum",
		TitleName: "末罗浮摩",
		TitleCode: "c_mallabhum",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿遮Arch = BArch阿遮
	f.阿遮Arch.SetParent(f)
	
	f.婆迦蒙提Baghmundi = BBaghmundi婆迦蒙提
	f.婆迦蒙提Baghmundi.SetParent(f)
	
	f.毗湿奴城Bishnupur = BBishnupur毗湿奴城
	f.毗湿奴城Bishnupur.SetParent(f)
	
	f.旃陀罗拘拿Chandrakona = BChandrakona旃陀罗拘拿
	f.旃陀罗拘拿Chandrakona.SetParent(f)
	
	f.提诃罗Dihar = BDihar提诃罗
	f.提诃罗Dihar.SetParent(f)
	
	f.迦罗吠多Garhbeta = BGarhbeta迦罗吠多
	f.迦罗吠多Garhbeta.SetParent(f)
	
	f.曼陀兰Mandaran = BMandaran曼陀兰
	f.曼陀兰Mandaran.SetParent(f)
	
}
