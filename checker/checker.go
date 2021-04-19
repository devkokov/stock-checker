package checker

import (
	"fmt"
	"github.com/gen2brain/beeep"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Checker struct {
	Products          *[]Product
	HttpClient        *http.Client
	WaitBetweenChecks time.Duration
}

func (c *Checker) Run() {
	for {
		for _, prod := range *c.Products {
			c.checkProduct(prod)
		}

		log.Println(fmt.Sprintf("WaitBetweenChecks %s...", c.WaitBetweenChecks.String()))
		time.Sleep(c.WaitBetweenChecks)
	}
}

func (c *Checker) checkProduct(prod Product) {
	req, err := http.NewRequest("GET", prod.Url, nil)
	if err != nil {
		log.Println(err)
		return
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36")

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	stCh, err := prod.StockChecker()
	if err != nil {
		log.Println(err)
		return
	}

	inStock := stCh.IsInStock(string(body))
	outOfStock := stCh.IsOutOfStock(string(body))

	if outOfStock && !inStock {
		log.Println(prod.Name + " is OUT OF STOCK")
		return
	}

	if inStock && !outOfStock {
		msg := prod.Name + " is IN STOCK!!!"

		log.Println(msg)

		err = beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
		if err != nil {
			log.Println(err)
			return
		}

		err = beeep.Alert("Stock Checker", msg, "")
		if err != nil {
			log.Println(err)
			return
		}
	}
}
