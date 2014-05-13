package main

import (
	"net/http"
	"html/template"
	"fmt"
)

type Pair struct {
    Key string
    Value string
}

func main() {
	flashvars := []Pair{
		Pair{"fb_source", "bookmark"},
		Pair{"ref", "bookmarks"},
		Pair{"count", "0"},
		Pair{"fb_bmpos", "2_0"},
		Pair{"code", "AQA9KrPoSxjyTjNoG9B1mUHoZ8ooUxusmPWV6Aa17SfgHIkSVubBwLxCKC5EO7fkfIiC9LvnrDOY35pzlPwyasKVe6q1dcOZzvyQeTmrlf-rhjjMH0FZGh0PWwMp0k3IqZTg1tKunkFFGMALgo4Vf8vSzGFG2r8DiJq5-N7K-5MIg7j3VnR0A0-EzaM5kATvrM5FmI1XWmxHbHFmHpS52rKuTvSuH27Ipwt4p2V2DGayvPDjnvvfs6d5-hdaCtoxoOvBJDfecDakToecSzr3kAU6zF4QiMCIC1MtihaH_3C7a9BeLdVqMhr4w4q33WqEso0"},
		Pair{"sys_id", "4"},
		Pair{"network", "facebook"},
		Pair{"uid", "675063875"},
		Pair{"app_id", "1424411677784893"},
		Pair{"interface_lang", "uk"},
		Pair{"access_token", "CAAUPfrARez0BAKDb5H1uds5kLg3794HyPAbTYRZAA1H2i43NPl8sSjpxl77gIqDapYZB4QxWrZAK1H6VQUVAFbWuTWr4VYbXagirvciMba7FhyYKSUboICrvSJKYgBndShSZA0n4ZA5JRZBqigVbMRdCsHrjl8AQEmcWfbJqkHflqmv8XEBarKEVJRfLp56ksLZCO7TBzkfVQZDZD"},
		Pair{"auth_key", "6514a97ae525f196b8060337380e0cbb"},
		Pair{"requestLoadingInfoTimeout", "3000"},
		Pair{"ref", "bookmark"},
		Pair{"rpc_url", "https%3A%2F%2Fepicwar-facebook.progrestar.net%2Frpc%2F"},
		Pair{"preloader_asset", "preloader%2Fpreloader_dwarf.swf"},
		Pair{"browser", "chrome"},
		Pair{"country_code", "UA"},
		Pair{"geoip_city", "Kiev"},
		Pair{"index_version", "1398240421"},
		Pair{"static_url", "https%3A%2F%2Fepicwar-a.akamaihd.net%2F"},
		Pair{"preloader", "https%3A%2F%2Fepicwar-a.akamaihd.net%2Ffacebook%2Fv042%2FFbLoader.swf%3Fv%3D2"},
		Pair{"rpc_url", "https%3A%2F%2Fepicwar-facebook.progrestar.net%2Frpc%2F"},
		Pair{"stat_url", "https://stat.progrestar.net/collector/client/"},
		Pair{"error_url", "https://error.progrestar.net/client/"},
	}

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("index.html")
    	err := t.Execute(w, struct{Flashvars []Pair}{Flashvars: flashvars})
    	if err != nil {
	        fmt.Println("There was an error:", err)
	    }
	})
    http.ListenAndServe(":8080", nil )
}