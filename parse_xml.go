package main

import (
    "encoding/xml"
    "fmt"
    "log"
)

type Asset struct {
    Ref  string `xml:"ref,attr"`
    Href string `xml:"href,attr"`
}

type Request struct {
    XMLName xml.Name `xml:"http://ns.vizrt.com/ardome/transfer request"`
    Asset   Asset    `xml:"http://www.vizrt.com/2010/mam asset"`
}

func main() {
    xmlStr := `<transfer:request xmlns:transfer="http://ns.vizrt.com/ardome/transfer">
                    <mam:asset ref="urn:vme:default:asset:2652302160000014521" href="https://onedev169.se.vizrt.internal/thirdparty/asset/item/2652302160000014521"
                        xmlns:mam="http://www.vizrt.com/2010/mam" />
                </transfer:request>`

    var req Request
    if err := xml.Unmarshal([]byte(xmlStr), &req); err != nil {
        log.Fatal(err)
    }

    fmt.Println(req.Asset.Href) // Output: https://onedev169.se.vizrt.internal/thirdparty/asset/item/2652302160000014521
}

