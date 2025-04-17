package forms

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/ipoluianov/goforms/ui"
	"github.com/ipoluianov/suigo/client"
)

type MainForm struct {
	ui.Form

	// First line
	p1       *ui.Panel
	lineEdit *ui.LineEdit
	btn      *ui.Button

	// Coin Info
	p3 *ui.Panel

	lblCoinTypeValue        *ui.TextBlock
	lblCoinNameValue        *ui.TextBlock
	lblCoinSymbolValue      *ui.TextBlock
	lblCoinDecimalsValue    *ui.TextBlock
	lblCoinDescriptionValue *ui.TextBlock
	lblCoinIconUrlValue     *ui.TextBlock
	lblCoinObjectIdValue    *ui.TextBlock

	lvInfo *ui.ListView
}

func NewMainForm() *MainForm {
	var c MainForm
	return &c
}

func (c *MainForm) OnInit() {
	c.Resize(1000, 800)
	vpanel := c.Panel().AddVPanel()

	c.SetTitle("Sui Coin Info")

	// P1
	c.p1 = vpanel.AddHPanel()
	c.lineEdit = ui.NewLineEdit(c.p1)
	c.p1.AddWidget(c.lineEdit)

	btnPaste := ui.NewButton(c.p1, "Paste", func(event *ui.Event) {
		text, _ := clipboard.ReadAll()
		if text != "" {
			c.lineEdit.SetText(text)
		}
	})
	c.p1.AddWidget(btnPaste)

	btn := ui.NewButton(c.p1, "Get Info", func(event *ui.Event) {
		c.loadCoinInfo(c.lineEdit.Text())
	})
	c.p1.AddWidget(btn)

	c.p3 = vpanel.AddVPanel()

	c.lvInfo = c.p3.AddListView()
	c.lvInfo.AddColumn("Key", 200)
	c.lvInfo.AddColumn("Value", 900)

	c.lineEdit.Focus()
}

func (c *MainForm) loadCoinInfo(coinType string) {
	cl := client.NewClient(client.MAINNET_URL)
	coinInfo, err := cl.GetCoinMetadata(coinType)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	c.lvInfo.RemoveItems()
	c.lvInfo.AddItem2("Coin Type", coinType)
	c.lvInfo.AddItem2("Coin Name", coinInfo.Name)
	c.lvInfo.AddItem2("Coin Symbol", coinInfo.Symbol)
	c.lvInfo.AddItem2("Coin Decimals", fmt.Sprintf("%d", coinInfo.Decimals))
	c.lvInfo.AddItem2("Coin Description", coinInfo.Description)
	c.lvInfo.AddItem2("Coin Icon URL", coinInfo.IconUrl)
}
