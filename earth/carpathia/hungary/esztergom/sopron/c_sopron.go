package sopron

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SopronCounty interface {
    feud.County
    BBorsmonostor博尔什莫诺什托尔() 	feud.Barony
    BCsepreg柴普赖格() 	feud.Barony
    BCsorna乔尔瑙() 	feud.Barony
    BGyor哲尔() 	feud.Barony
    BKapuvar考普堡() 	feud.Barony
    BKismarton基什毛尔通() 	feud.Barony
    BNagymarton瑙吉毛尔通() 	feud.Barony
    BSopron肖普朗() 	feud.Barony
}

type 肖普朗SopronCounty struct {
	feud.BaseCounty
	博尔什莫诺什托尔Borsmonostor 	feud.Barony
	柴普赖格Csepreg 	feud.Barony
	乔尔瑙Csorna 	feud.Barony
	哲尔Gyor 	feud.Barony
	考普堡Kapuvar 	feud.Barony
	基什毛尔通Kismarton 	feud.Barony
	瑙吉毛尔通Nagymarton 	feud.Barony
	肖普朗Sopron 	feud.Barony
}

func (c *肖普朗SopronCounty) BBorsmonostor博尔什莫诺什托尔() feud.Barony {
	return c.博尔什莫诺什托尔Borsmonostor
}
    
func (c *肖普朗SopronCounty) BCsepreg柴普赖格() feud.Barony {
	return c.柴普赖格Csepreg
}
    
func (c *肖普朗SopronCounty) BCsorna乔尔瑙() feud.Barony {
	return c.乔尔瑙Csorna
}
    
func (c *肖普朗SopronCounty) BGyor哲尔() feud.Barony {
	return c.哲尔Gyor
}
    
func (c *肖普朗SopronCounty) BKapuvar考普堡() feud.Barony {
	return c.考普堡Kapuvar
}
    
func (c *肖普朗SopronCounty) BKismarton基什毛尔通() feud.Barony {
	return c.基什毛尔通Kismarton
}
    
func (c *肖普朗SopronCounty) BNagymarton瑙吉毛尔通() feud.Barony {
	return c.瑙吉毛尔通Nagymarton
}
    
func (c *肖普朗SopronCounty) BSopron肖普朗() feud.Barony {
	return c.肖普朗Sopron
}
    
var CSopron肖普朗 SopronCounty = &肖普朗SopronCounty{}

func init() {
	f := CSopron肖普朗.(*肖普朗SopronCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "450",
		Title:     "sopron",
		TitleName: "肖普朗",
		TitleCode: "c_sopron",
		Baronies:  map[string]feud.Barony{},
	}

	f.博尔什莫诺什托尔Borsmonostor = BBorsmonostor博尔什莫诺什托尔
	f.博尔什莫诺什托尔Borsmonostor.SetParent(f)
	
	f.柴普赖格Csepreg = BCsepreg柴普赖格
	f.柴普赖格Csepreg.SetParent(f)
	
	f.乔尔瑙Csorna = BCsorna乔尔瑙
	f.乔尔瑙Csorna.SetParent(f)
	
	f.哲尔Gyor = BGyor哲尔
	f.哲尔Gyor.SetParent(f)
	
	f.考普堡Kapuvar = BKapuvar考普堡
	f.考普堡Kapuvar.SetParent(f)
	
	f.基什毛尔通Kismarton = BKismarton基什毛尔通
	f.基什毛尔通Kismarton.SetParent(f)
	
	f.瑙吉毛尔通Nagymarton = BNagymarton瑙吉毛尔通
	f.瑙吉毛尔通Nagymarton.SetParent(f)
	
	f.肖普朗Sopron = BSopron肖普朗
	f.肖普朗Sopron.SetParent(f)
	
}
