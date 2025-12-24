package main

import (
	"fmt"
	"strings"
	"time"
)

// Receipt は fmt.Stringer を実装し、整形済みの表示を返す
type Receipt struct {
	ID        string
	Items     []LineItem
	CreatedAt time.Time
}

type LineItem struct {
	Name  string
	Price int
}

func (r Receipt) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("Receipt #%s @ %s\n", r.ID, r.CreatedAt.Format(time.RFC3339)))
	total := 0
	for _, it := range r.Items {
		b.WriteString(fmt.Sprintf("- %-12s %5d\n", it.Name, it.Price))
		total += it.Price
	}
	b.WriteString(fmt.Sprintf("TOTAL: %d", total))
	return b.String()
}

// 用途: Stringer を活かしてそのまま fmt.Println する
func main() {
	rcpt := Receipt{
		ID: "A-42",
		Items: []LineItem{
			{Name: "Coffee", Price: 480},
			{Name: "Donut", Price: 260},
			{Name: "Tax", Price: 50},
		},
		CreatedAt: time.Date(2024, 12, 24, 12, 34, 56, 0, time.UTC),
	}

	fmt.Println(rcpt)
}
