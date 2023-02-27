package faereyar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FaereyarCounty interface {
    feud.County
    BFunningur富宁古尔() 	feud.Barony
    BHov霍夫() 	feud.Barony
    BKirkjubour奇尔丘伯乌尔() 	feud.Barony
    BKlaksvik克拉克斯维克() 	feud.Barony
    BKvivik克维维克() 	feud.Barony
    BSandur桑杜尔() 	feud.Barony
    BSkansin斯卡尼锡() 	feud.Barony
    BTorshavn托尔斯豪恩() 	feud.Barony
}

type 法伊雷亚尔FaereyarCounty struct {
	feud.BaseCounty
	富宁古尔Funningur 	feud.Barony
	霍夫Hov 	feud.Barony
	奇尔丘伯乌尔Kirkjubour 	feud.Barony
	克拉克斯维克Klaksvik 	feud.Barony
	克维维克Kvivik 	feud.Barony
	桑杜尔Sandur 	feud.Barony
	斯卡尼锡Skansin 	feud.Barony
	托尔斯豪恩Torshavn 	feud.Barony
}

func (c *法伊雷亚尔FaereyarCounty) BFunningur富宁古尔() feud.Barony {
	return c.富宁古尔Funningur
}
    
func (c *法伊雷亚尔FaereyarCounty) BHov霍夫() feud.Barony {
	return c.霍夫Hov
}
    
func (c *法伊雷亚尔FaereyarCounty) BKirkjubour奇尔丘伯乌尔() feud.Barony {
	return c.奇尔丘伯乌尔Kirkjubour
}
    
func (c *法伊雷亚尔FaereyarCounty) BKlaksvik克拉克斯维克() feud.Barony {
	return c.克拉克斯维克Klaksvik
}
    
func (c *法伊雷亚尔FaereyarCounty) BKvivik克维维克() feud.Barony {
	return c.克维维克Kvivik
}
    
func (c *法伊雷亚尔FaereyarCounty) BSandur桑杜尔() feud.Barony {
	return c.桑杜尔Sandur
}
    
func (c *法伊雷亚尔FaereyarCounty) BSkansin斯卡尼锡() feud.Barony {
	return c.斯卡尼锡Skansin
}
    
func (c *法伊雷亚尔FaereyarCounty) BTorshavn托尔斯豪恩() feud.Barony {
	return c.托尔斯豪恩Torshavn
}
    
var CFaereyar法伊雷亚尔 FaereyarCounty = &法伊雷亚尔FaereyarCounty{}

func init() {
	f := CFaereyar法伊雷亚尔.(*法伊雷亚尔FaereyarCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "33",
		Title:     "faereyar",
		TitleName: "法伊雷亚尔",
		TitleCode: "c_faereyar",
		Baronies:  map[string]feud.Barony{},
	}

	f.富宁古尔Funningur = BFunningur富宁古尔
	f.富宁古尔Funningur.SetParent(f)
	
	f.霍夫Hov = BHov霍夫
	f.霍夫Hov.SetParent(f)
	
	f.奇尔丘伯乌尔Kirkjubour = BKirkjubour奇尔丘伯乌尔
	f.奇尔丘伯乌尔Kirkjubour.SetParent(f)
	
	f.克拉克斯维克Klaksvik = BKlaksvik克拉克斯维克
	f.克拉克斯维克Klaksvik.SetParent(f)
	
	f.克维维克Kvivik = BKvivik克维维克
	f.克维维克Kvivik.SetParent(f)
	
	f.桑杜尔Sandur = BSandur桑杜尔
	f.桑杜尔Sandur.SetParent(f)
	
	f.斯卡尼锡Skansin = BSkansin斯卡尼锡
	f.斯卡尼锡Skansin.SetParent(f)
	
	f.托尔斯豪恩Torshavn = BTorshavn托尔斯豪恩
	f.托尔斯豪恩Torshavn.SetParent(f)
	
}
