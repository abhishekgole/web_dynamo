package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"regexp"
	"strings"
)

func homePage(w http.ResponseWriter, r *http.Request) {

	req, err := http.NewRequest("GET", "https://www.booking.com/hotel/gb/thistlemarblearch.html?label=&sid=&checkin=2021-11-04&checkout=2021-11-08&no_rooms=1&group_adults=2&selected_currency=hotel_currency", nil)
	if err != nil {
		log.Fatalln(err)
	}

	//req.Header.Set("Referer", "https://www.booking.com/searchresults.html?label=gen173nr-1FCAEoggI46AdIM1gEaGyIAQGYATG4AQjIAQzYAQHoAQH4AQKIAgGoAgO4Ao-V8IoGwAIB0gIkMDVmYzU5Y2UtMDc5Yi00YzdmLWFmYzUtZTExMDFiOWQ1YWU42AIF4AIB&sid=7cc9ce8dee09c6c286bb113f0c2b3a5d&sb=1&sb_lp=1&src=index&src_elem=sb&error_url=https%3A%2F%2Fwww.booking.com%2Findex.html%3Flabel%3Dgen173nr-1FCAEoggI46AdIM1gEaGyIAQGYATG4AQjIAQzYAQHoAQH4AQKIAgGoAgO4Ao-V8IoGwAIB0gIkMDVmYzU5Y2UtMDc5Yi00YzdmLWFmYzUtZTExMDFiOWQ1YWU42AIF4AIB%3Bsid%3D7cc9ce8dee09c6c286bb113f0c2b3a5d%3Bsb_price_type%3Dtotal%26%3B&ss=Cartagena+de+Indias%2C+Bolivar%2C+Colombia&is_ski_area=&checkin_year=2021&checkin_month=11&checkin_monthday=8&checkout_year=2021&checkout_month=11&checkout_monthday=11&group_adults=2&group_children=0&no_rooms=1&b_h4u_keep_filters=&from_sf=1&ss_raw=ind&ac_position=4&ac_langcode=en&ac_click_type=b&dest_id=-579943&dest_type=city&iata=CTG&place_id_lat=10.425008&place_id_lon=-75.546905&search_pageview_id=b2753a878bef0048&search_selected=true&search_pageview_id=b2753a878bef0048&ac_suggestion_list_length=5&ac_suggestion_theme_list_length=0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.71 Safari/537.36")
	req.Header.Set("Host", "www.booking.com")
	req.Header.Set("Connection", "keep-alive")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	b, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatalln(err)
	}

	result := strings.Split(string(b), "e2e-hprt-table-row")
	for i, _ := range result {

		if i == 0 {
			i = i + 1
			continue
		}
		var foo []string

		space := regexp.MustCompile(`\s+`)
		str := space.ReplaceAllString(result[i], " ")

		rev_amt := regexp.MustCompile("bui-u-sr-only\">\\sPrice\\s₹(.*?)\\s</span>")
		rev_amt2 := rev_amt.FindStringSubmatch(str)
		rev_amt3 := rev_amt2[1]
		foo = append(foo, rev_amt3)
		///// Extraction For ROOM TYPE //////
		room_type := regexp.MustCompile("hprt-roomtype-icon-link\\s\">\\s(.*?)\\s</span>")
		room_type2 := room_type.FindStringSubmatch(str)
		room_type3 := room_type2[0:]

		if len(room_type3) != 0 {

			foo = append(foo, room_type3[1])
		}
		///// Extraction For Currency Code //////
		currency := regexp.MustCompile("bui-u-sr-only\">\\sPrice\\s(.*?)\\d+,\\d+\\s</span>")
		currency2 := currency.FindStringSubmatch(str)
		currency3 := currency2[0:]

		if len(currency3) != 0 {

			foo = append(foo, currency3[1])
		}
		///// Extraction For Cancellation Policy //////
		cancellation := regexp.MustCompile("hprt-item--emphasised\">(.*?)</span></span>\\s(.*?)\\s</li")
		cancellation2 := cancellation.FindStringSubmatch(str)
		cancellation3 := cancellation2[0:]

		if len(cancellation3) != 0 {

			foo = append(foo, cancellation3[1])
		}
		///// Extraction For Tax Info //////
		tax_info := regexp.MustCompile("Included:</span>\\s+(.*?)\\s+</div>")
		tax_info2 := tax_info.FindStringSubmatch(str)
		tax_info3 := tax_info2[0:]

		if len(tax_info3) != 0 {

			foo = append(foo, tax_info3[1])
		}
		///// Extraction For Bed Type //////
		bed_type := regexp.MustCompile("hprt-roomtype-bed.*?bed-types-wrapper.*?rt-bed-types.*?class=\"rt-bed-type.*?n>\\s+(.*?)\\s+<i.*?singles.*?<span>\\s+(.*?)\\s+<i.*?bicon-couch")
		bed_type2 := bed_type.FindStringSubmatch(str)
		bed_type3 := bed_type2[0:]

		if len(bed_type3) != 0 {

			foo = append(foo, bed_type3[1])
		}
		///// Extraction For Max Guest //////
		max_guests := regexp.MustCompile("bui-u-sr-only\">\\s+Max\\s+people.*?(\\d+).*?</s")
		max_guests2 := max_guests.FindStringSubmatch(str)
		max_guests3 := max_guests2[0:]

		if len(max_guests3) != 0 {

			foo = append(foo, max_guests3[1])
		}

		for _, value := range foo {
			if len(value) != 0 {
				fmt.Fprintln(w, value)
			}

		}
		fmt.Fprintln(w, "New Room")

	}
	if err != nil {
		log.Fatal(err)
	}
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
