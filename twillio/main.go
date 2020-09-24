package main

import (
	"context"
	"fmt"
	"net/url"

	"github.com/kevinburke/twilio-go"
)

func main() {

	pdfurl, _ := url.Parse(pdffile)

	client := twilio.NewClient(sid, tok, nil)

	// Send a SMS
	msg, err := client.Messages.SendMessage(fromnumber, msgto, "Hi, Your order got confirmed. Food will reach you within 30 min. Stay connect with eater24.com", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(msg.Sid, msg.FriendlyPrice(), msg.Price)

	// Send a Fax via SendFax function
	fax, err := client.Faxes.SendFax(fromnumber, faxto, pdfurl)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(fax.Sid, fax.FriendlyPrice(), msg.Price)

	// Send a Fax via Create function
	data := url.Values{
		"To":             []string{faxto},
		"From":           []string{fromnumber},
		"MediaUrl":       []string{pdffile},
		"Quality":        []string{"fine"},
		"StatusCallback": []string{""},
	}
	fax, err = client.Fax.Faxes.Create(context.Background(), data)
	if err != nil {
		fmt.Println(err.Error())
	}
	if fax.Quality != "fine" {
		fmt.Println("Fax quality is incorrect", fax.Quality)
	}
	fmt.Println(fax.Sid, fax.FriendlyPrice(), fax.Price)

}
