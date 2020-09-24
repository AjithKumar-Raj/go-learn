package main

import (
	"fmt"
	"os"
	"time"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

const logopath = "logo.png"

func main() {
	begin := time.Now()

	// darkGrayColor := getDarkGrayColor()
	// grayColor := getGrayColor()
	// whiteColor := color.NewWhite()
	header := getHeader()
	contents := getContents()

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 20)

	m.RegisterHeader(func() {
		m.Row(20, func() {
			m.Col(4, func() {
				_ = m.FileImage(logopath, props.Rect{
					Center:  true,
					Percent: 80,
				})
			})
			m.ColSpace(4)
			m.Col(4, func() {
				m.Text(customerName, props.Text{
					Size:        10,
					Align:       consts.Left,
					Extrapolate: false,
				})
				m.Text("Phone: "+customerNumber, props.Text{
					Top:   4,
					Size:  8,
					Style: consts.Bold,
					Align: consts.Left,
				})
				m.Text("Payment: "+paymentMethod, props.Text{
					Top:   7,
					Size:  8,
					Align: consts.Left,
				})
				m.Text("Method: "+orderMethod, props.Text{
					Top:   10,
					Size:  8,
					Align: consts.Left,
				})
				if orderMethod == "Delivery" {
					m.Text("Address: "+customerAddress, props.Text{
						Top:   13,
						Size:  8,
						Align: consts.Left,
					})
				}
			})
		})
	})
	m.Row(8, func() {
		m.Col(12, func() {
			m.Text(businessName, props.Text{
				Top:   4,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
	})
	m.Row(8, func() {
		m.Col(12, func() {
			m.Text(businessAddress, props.Text{
				Size:  8,
				Align: consts.Center,
			})
		})
	})
	m.Row(8, func() {
		m.Col(12, func() {
			m.Text("Invoice ID: "+invoiceID, props.Text{
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
	})
	m.SetBorder(true)
	m.TableList(header, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      10,
			GridSizes: []uint{4, 3, 2, 3},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{4, 3, 2, 3},
		},
		Align:              consts.Center,
		HeaderContentSpace: 1,
	})

	m.Row(5, func() {
		m.Col(9, func() {
			m.Text("SubTotal     ", props.Text{
				Style: consts.Bold,
				Size:  8,
				Align: consts.Right,
			})
		})
		m.Col(3, func() {
			m.Text("$ 2,567.00", props.Text{
				Style: consts.Bold,
				Size:  8,
				Align: consts.Center,
			})
		})
	})

	m.Row(5, func() {
		m.Col(9, func() {
			m.Text("Tax     ", props.Text{
				Style: consts.Bold,
				Size:  8,
				Align: consts.Right,
			})
		})
		m.Col(3, func() {
			m.Text("$ 3.50", props.Text{
				Style: consts.Bold,
				Size:  8,
				Align: consts.Center,
			})
		})
	})

	m.Row(5, func() {
		m.Col(9, func() {
			m.Text("Total     ", props.Text{
				Style: consts.Bold,
				Size:  8,
				Align: consts.Right,
			})
		})
		m.Col(3, func() {
			m.Text("$ 2,570.50", props.Text{
				Style: consts.Bold,
				Size:  8,
				Align: consts.Center,
			})
		})
	})

	err := m.OutputFileAndClose("billing.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))
}

func getHeader() []string {
	return []string{"Product", "Unit Price", "Quantity", "Price"}
}

func getContents() [][]string {
	return [][]string{
		{"Swamp", "$1.0", "12", "$ 4.00"},
		{"Sorin, A Planeswalker", "$1.0", "4", "$ 90.00"},
		{"Tassa", "$1.0", "4", "$ 30.00"},
		{"Skinrender", "$1.0", "4", "$ 9.00"},
		{"Island", "$1.0", "12", "$ 4.00"},
		{"Mountain", "$1.0", "12", "$ 4.00"},
		{"Plain", "$1.0", "12", "$ 4.00"},
		{"Black Lotus", "$1.0", "1", "$ 1.000.00"},
		{"Time Walk", "$1.0", "1", "$ 1.000.00"},
		{"Emberclave", "$1.0", "4", "$ 44.00"},
		{"Anax", "$1.0", "4", "$ 32.00"},
		{"Murderous Rider", "$1.0", "4", "$ 22.00"},
		{"Gray Merchant of Asphodel \n - Spicy & medium", "$1.0", "4", "$ 2.00"},
		{"Ajani's Pridemate", "$1.0", "4", "$ 2.00"},
		{"Renan, Chatuba", "$1.0", "4", "$ 19.00"},
		{"Tymarett", "$1.0", "4", "$ 13.00"},
		{"Doom Blade", "$1.0", "4", "$ 5.00"},
		{"Dark Lord", "$1.0", "3", "$ 7.00"},
		{"Memory of Thanatos", "$1.0", "3", "$ 32.00"},
		{"Poring", "$1.0", "4", "$ 1.00"},
		{"Deviling", "$1.0", "4", "$ 99.00"},
		{"Seiya", "$1.0", "4", "$ 45.00"},
		{"Harry Potter", "$1.0", "4", "$ 62.00"},
		{"Goku", "$1.0", "4", "$ 77.00"},
		{"Phreoni", "$1.0", "4", "$ 22.00"},
		{"Katheryn High Wizard", "$1.0", "4", "$ 25.00"},
		{"Lord Seyren", "$1.0", "4", "$ 55.00"},
	}
}

func getDarkGrayColor() color.Color {
	return color.Color{
		Red:   144,
		Green: 144,
		Blue:  144,
	}
}

func getGrayColor() color.Color {
	return color.Color{
		Red:   200,
		Green: 200,
		Blue:  200,
	}
}
