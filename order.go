package sfcc_export_parser

import (
	"encoding/xml"
	"io"
	"time"
)

type Orders struct {
	XMLName xml.Name `xml:"orders"`
	Orders  []Order  `xml:"order" json:"orders"`
}

type Order struct {
	XMLName            xml.Name           `xml:"order" json:"order"`
	Date               time.Time          `xml:"order-date" json:"orderDate"`
	OrderNumber        string             `xml:"original-order-no" json:"originalOrderNo"`
	InvoiceNumber      string             `xml:"invoice-no" json:"invoiceNo"`
	Customer           Customer           `xml:"customer" json:"customer"`
	Status             Status             `xml:"status" json:"status"`
	BusinessType       string             `xml:"business-type" json:"businessType"`
	ChannelType        string             `xml:"channel-type" json:"channelType"`
	CurrentOrderNumber string             `xml:"current-order-no" json:"currentOrderNo"`
	ProductLineItems   []ProductLineItem  `xml:"product-lineitems>product-lineitem" json:"productLineItems"`
	ShippingLineItems  []ShippingLineItem `xml:"shipping-lineitems>shipping-lineitem" json:"shippingLineItems"`
	Shipments          []Shipment         `xml:"shipments>shipment" json:"shipments"`
	Totals             Totals             `xml:"totals" json:"totals"`
	Payments           []Payment          `xml:"payments>payment" json:"payments"`
	RemoteHost         string             `xml:"remoteHost" json:"remoteHost"`
	Notes              []Note             `xml:"notes>note" json:"notes"`
	CustomAttributes   CustomAttributes   `xml:"custom-attributes" json:"customAttributes"`
}

type Customer struct {
	XMLName        xml.Name       `xml:"customer"`
	Guest          bool           `xml:"guest" json:"guest"`
	Name           string         `xml:"customer-name" json:"name"`
	Email          string         `xml:"customer-email" json:"email"`
	BillingAddress BillingAddress `xml:"billing-address" json:"billingAddress"`
}

type Address struct {
	FirstName        string           `xml:"first-name" json:"firstName"`
	LastName         string           `xml:"last-name" json:"lastName"`
	Address          string           `xml:"address1" json:"address"`
	City             string           `xml:"city" json:"city"`
	PostalCode       string           `xml:"postal-code" json:"postalCode"`
	CountryCode      string           `xml:"country-code" json:"countryCode"`
	Phone            string           `xml:"phone" json:"phone"`
	CustomAttributes CustomAttributes `xml:"custom-attributes" json:"customAttributes"`
}

type BillingAddress Address
type ShippingAddress Address

type CustomAttributes struct {
	Elements map[string]string
}

type Status struct {
	OrderStatus        string `xml:"order-status" json:"orderStatus"`
	ShippingStatus     string `xml:"shipping-status" json:"shippingStatus"`
	ConfirmationStatus string `xml:"confirmation-status" json:"confirmationStatus"`
	PaymentStatus      string `xml:"payment-status" json:"paymentStatus"`
}

type ProductLineItem struct {
	NetPrice         float64           `xml:"net-price" json:"netPrice"`
	Tax              float64           `xml:"tax" json:"tax"`
	GrossPrice       float64           `xml:"Gross-price" json:"grossPrice"`
	BasePrice        float64           `xml:"base-price" json:"basePrice"`
	Text             string            `xml:"lineitem-text" json:"text"`
	TaxBasis         float64           `xml:"tax-basis" json:"taxBasis"`
	Position         int               `xml:"position" json:"position"`
	ProductID        string            `xml:"product-id" json:"productID"`
	ProductName      string            `xml:"product-name" json:"productName"`
	Quantity         float64           `xml:"quantity" json:"quantity"`
	TaxRate          float64           `xml:"tax-rate" json:"taxRate"`
	ShipmentID       string            `xml:"shipment-id" json:"shipmentID"`
	Gift             bool              `xml:"gift" json:"gift"`
	CustomAttributes CustomAttributes  `xml:"custom-attributes" json:"customAttributes"`
	PriceAdjustments []PriceAdjustment `xml:"price-adjustments>price-adjustment" json:"priceAdjustments"`
}

type PriceAdjustment struct {
	NetPrice    float64 `xml:"net-price" json:"netPrice"`
	Tax         float64 `xml:"tax" json:"tax"`
	GrossPrice  float64 `xml:"Gross-price" json:"grossPrice"`
	BasePrice   float64 `xml:"base-price" json:"basePrice"`
	Text        string  `xml:"lineitem-text" json:"text"`
	TaxBasis    float64 `xml:"tax-basis" json:"taxBasis"`
	PromotionID string  `xml:"promotion-id" json:"promotionID"`
	Discount    float64 `xml:"discount>amount" json:"amount"`
}

type ShippingLineItem struct {
	NetPrice         float64           `xml:"net-price" json:"netPrice"`
	Tax              float64           `xml:"tax" json:"tax"`
	GrossPrice       float64           `xml:"Gross-price" json:"grossPrice"`
	BasePrice        float64           `xml:"base-price" json:"basePrice"`
	TaxBasis         float64           `xml:"tax-basis" json:"taxBasis"`
	PriceAdjustments []PriceAdjustment `xml:"price-adjustments>price-adjustment"`
	ItemID           string            `xml:"item-id" json:"itemID"`
	ShipmentID       string            `xml:"shipment-id" json:"shipmentID"`
	TaxRate          float64           `xml:"tax-rate" json:"taxRate"`
}

type Shipment struct {
	ID               string           `xml:"shipment-id,attr" json:"ID"`
	Status           string           `xml:"status>shipping-status" json:"status"`
	Method           string           `xml:"shipping-method" json:"method"`
	TrackingNumber   string           `xml:"tracking-number" json:"trackingNumber"`
	ShippingAddress  ShippingAddress  `xml:"shipping-address" json:"shippingAddress"`
	Gift             bool             `xml:"gift" json:"gift"`
	Totals           Totals           `xml:"totals" json:"totals"`
	CustomAttributes CustomAttributes `xml:"custom-attributes" json:"customAttributes"`
}

type Total struct {
	NetPrice         float64           `xml:"net-price" json:"netPrice"`
	Tax              float64           `xml:"tax" json:"tax"`
	GrossPrice       float64           `xml:"gross-price" json:"grossPrice"`
	PriceAdjustments []PriceAdjustment `xml:"price-adjustments>price-adjustment,omitempty"`
}

type Totals struct {
	MerchandiseTotal         MerchandiseTotal         `xml:"merchandize-total" json:"merchandiseTotal"`
	AdjustedMerchandiseTotal AdjustedMerchandiseTotal `xml:"adjusted-merchandize-total" json:"adjustedMerchandiseTotal"`
	ShippingTotal            ShippingTotal            `xml:"shipping-total" json:"shippingTotal"`
	AdjustedShippingTotal    AdjustedShippingTotal    `xml:"adjusted-shipping-total" json:"adjustedShippingTotal"`
	ShipmentTotal            ShipmentTotal            `xml:"shipment-total" json:"shipmentTotal"`
	OrderTotal               OrderTotal               `xml:"order-total,omitempty" json:"orderTotal,omitempty"`
}

type MerchandiseTotal Total
type AdjustedMerchandiseTotal Total
type ShippingTotal Total
type AdjustedShippingTotal Total
type ShipmentTotal Total
type OrderTotal Total

type Payment struct {
	CustomMethod     string           `xml:"custom-method>method-name" json:"customMethod"`
	Amount           float64          `xml:"amount" json:"amont"`
	ProcessorID      string           `xml:"processor-id" json:"processorID"`
	TransactionID    string           `xml:"transaction-id" json:"transactionID"`
	CustomAttributes CustomAttributes `xml:"custom-attributes" json:"customAttributes"`
}

type Note struct {
	CreatedBy    string    `xml:"created-by" json:"createdBy"`
	CreationDate time.Time `xml:"creation-date" json:"creationDate"`
	Subject      string    `xml:"subject" json:"subject"`
	Text         string    `xml:"text" json:"text"`
}

func (c *CustomAttributes) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	type customAttribute struct {
		Key   string `xml:"attribute-id,attr"`
		Value string `xml:",chardata"`
	}

	e := customAttribute{}
	c.Elements = map[string]string{}

	for err = d.Decode(&e); err == nil; err = d.Decode(&e) {
		c.Elements[e.Key] = e.Value
	}

	if err != nil && err != io.EOF {
		return err
	}

	return nil
}
