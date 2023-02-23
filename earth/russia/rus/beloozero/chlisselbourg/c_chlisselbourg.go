package chlisselbourg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ChlisselbourgCounty interface {
	feud.County
	BChlisselbourg什利谢利堡() feud.Barony
	BMokrishvitsa莫克里什维察() feud.Barony
	BNevdubstroy涅夫杜斯特罗伊() feud.Barony
	BNossok诺索克() feud.Barony
	BStorozhno斯托罗日诺() feud.Barony
	BVolkhov沃尔霍夫() feud.Barony
}

type 什利谢利堡ChlisselbourgCounty struct {
	feud.BaseCounty
	什利谢利堡Chlisselbourg feud.Barony
	莫克里什维察Mokrishvitsa feud.Barony
	涅夫杜斯特罗伊Nevdubstroy feud.Barony
	诺索克Nossok          feud.Barony
	斯托罗日诺Storozhno     feud.Barony
	沃尔霍夫Volkhov        feud.Barony
}

func (c *什利谢利堡ChlisselbourgCounty) BChlisselbourg什利谢利堡() feud.Barony {
	return c.什利谢利堡Chlisselbourg
}

func (c *什利谢利堡ChlisselbourgCounty) BMokrishvitsa莫克里什维察() feud.Barony {
	return c.莫克里什维察Mokrishvitsa
}

func (c *什利谢利堡ChlisselbourgCounty) BNevdubstroy涅夫杜斯特罗伊() feud.Barony {
	return c.涅夫杜斯特罗伊Nevdubstroy
}

func (c *什利谢利堡ChlisselbourgCounty) BNossok诺索克() feud.Barony {
	return c.诺索克Nossok
}

func (c *什利谢利堡ChlisselbourgCounty) BStorozhno斯托罗日诺() feud.Barony {
	return c.斯托罗日诺Storozhno
}

func (c *什利谢利堡ChlisselbourgCounty) BVolkhov沃尔霍夫() feud.Barony {
	return c.沃尔霍夫Volkhov
}

var CChlisselbourg什利谢利堡 ChlisselbourgCounty = &什利谢利堡ChlisselbourgCounty{}

func init() {
	f := CChlisselbourg什利谢利堡.(*什利谢利堡ChlisselbourgCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1660",
		Title:     "chlisselbourg",
		TitleName: "什利谢利堡",
		TitleCode: "c_chlisselbourg",
		Baronies:  map[string]feud.Barony{},
	}

	f.什利谢利堡Chlisselbourg = BChlisselbourg什利谢利堡
	f.什利谢利堡Chlisselbourg.SetParent(f)

	f.莫克里什维察Mokrishvitsa = BMokrishvitsa莫克里什维察
	f.莫克里什维察Mokrishvitsa.SetParent(f)

	f.涅夫杜斯特罗伊Nevdubstroy = BNevdubstroy涅夫杜斯特罗伊
	f.涅夫杜斯特罗伊Nevdubstroy.SetParent(f)

	f.诺索克Nossok = BNossok诺索克
	f.诺索克Nossok.SetParent(f)

	f.斯托罗日诺Storozhno = BStorozhno斯托罗日诺
	f.斯托罗日诺Storozhno.SetParent(f)

	f.沃尔霍夫Volkhov = BVolkhov沃尔霍夫
	f.沃尔霍夫Volkhov.SetParent(f)

}
