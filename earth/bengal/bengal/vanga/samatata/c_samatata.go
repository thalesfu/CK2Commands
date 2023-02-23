package samatata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SamatataCounty interface {
	feud.County
	BCandranatha旃陀罗那他() feud.Barony
	BChatigama车底伽摩() feud.Barony
	BDevaparvata提婆钵伐多() feud.Barony
	BGarmali伽罗摩梨() feud.Barony
	BGond贡德() feud.Barony
	BKanaoudi伽那奥提() feud.Barony
	BKandalike乾陀梨艺() feud.Barony
}

type 三摩呾吒SamatataCounty struct {
	feud.BaseCounty
	旃陀罗那他Candranatha feud.Barony
	车底伽摩Chatigama    feud.Barony
	提婆钵伐多Devaparvata feud.Barony
	伽罗摩梨Garmali      feud.Barony
	贡德Gond           feud.Barony
	伽那奥提Kanaoudi     feud.Barony
	乾陀梨艺Kandalike    feud.Barony
}

func (c *三摩呾吒SamatataCounty) BCandranatha旃陀罗那他() feud.Barony {
	return c.旃陀罗那他Candranatha
}

func (c *三摩呾吒SamatataCounty) BChatigama车底伽摩() feud.Barony {
	return c.车底伽摩Chatigama
}

func (c *三摩呾吒SamatataCounty) BDevaparvata提婆钵伐多() feud.Barony {
	return c.提婆钵伐多Devaparvata
}

func (c *三摩呾吒SamatataCounty) BGarmali伽罗摩梨() feud.Barony {
	return c.伽罗摩梨Garmali
}

func (c *三摩呾吒SamatataCounty) BGond贡德() feud.Barony {
	return c.贡德Gond
}

func (c *三摩呾吒SamatataCounty) BKanaoudi伽那奥提() feud.Barony {
	return c.伽那奥提Kanaoudi
}

func (c *三摩呾吒SamatataCounty) BKandalike乾陀梨艺() feud.Barony {
	return c.乾陀梨艺Kandalike
}

var CSamatata三摩呾吒 SamatataCounty = &三摩呾吒SamatataCounty{}

func init() {
	f := CSamatata三摩呾吒.(*三摩呾吒SamatataCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1318",
		Title:     "samatata",
		TitleName: "三摩呾吒",
		TitleCode: "c_samatata",
		Baronies:  map[string]feud.Barony{},
	}

	f.旃陀罗那他Candranatha = BCandranatha旃陀罗那他
	f.旃陀罗那他Candranatha.SetParent(f)

	f.车底伽摩Chatigama = BChatigama车底伽摩
	f.车底伽摩Chatigama.SetParent(f)

	f.提婆钵伐多Devaparvata = BDevaparvata提婆钵伐多
	f.提婆钵伐多Devaparvata.SetParent(f)

	f.伽罗摩梨Garmali = BGarmali伽罗摩梨
	f.伽罗摩梨Garmali.SetParent(f)

	f.贡德Gond = BGond贡德
	f.贡德Gond.SetParent(f)

	f.伽那奥提Kanaoudi = BKanaoudi伽那奥提
	f.伽那奥提Kanaoudi.SetParent(f)

	f.乾陀梨艺Kandalike = BKandalike乾陀梨艺
	f.乾陀梨艺Kandalike.SetParent(f)

}
