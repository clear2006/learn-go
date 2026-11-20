package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

//go:embed templates/*.html static/*.css
var siteFiles embed.FS

type Lap struct {
	Number int
	S1     string
	S2     string
	S3     string
	Total  string
}

type Driver struct {
	Position int
	Name     string
	Number   int
	Laps     int
	BestLap  string
	LastLap  string
	Gap      string
	LapTimes []Lap
}

type Race struct {
	ID       int
	Name     string
	Category string
	Status   string
	Date     string
	Duration string
	Drivers  []Driver
}

type PageData struct {
	Title       string
	Circuit     string
	Location    string
	Length      string
	TrackType   string
	Updated     string
	Races       []Race
	Selected    Race
	IsRacePage  bool
}

var races = []Race{
	{
		ID: 101, Name: "Session Sprint - Groupe A", Category: "SPRINT", Status: "Terminée",
		Date: "20/11/2026", Duration: "15 tours",
		Drivers: []Driver{
			{Position: 1, Name: "Alex Martin", Number: 12, Laps: 15, BestLap: "48.214", LastLap: "48.630", Gap: "—",
				LapTimes: []Lap{{1,"16.102","16.011","16.310","48.423"},{2,"16.050","15.980","16.184","48.214"},{3,"16.210","16.040","16.280","48.530"},{4,"16.180","16.120","16.330","48.630"},{5,"16.090","16.030","16.250","48.370"}}},
			{Position: 2, Name: "Camille Bernard", Number: 7, Laps: 15, BestLap: "48.390", LastLap: "48.910", Gap: "+ 2.184",
				LapTimes: []Lap{{1,"16.220","16.080","16.310","48.610"},{2,"16.130","16.020","16.240","48.390"},{3,"16.180","16.110","16.290","48.580"},{4,"16.250","16.140","16.520","48.910"},{5,"16.170","16.090","16.330","48.590"}}},
			{Position: 3, Name: "Noah Petit", Number: 21, Laps: 14, BestLap: "48.755", LastLap: "49.102", Gap: "+ 1 tour",
				LapTimes: []Lap{{1,"16.310","16.190","16.390","48.890"},{2,"16.200","16.110","16.445","48.755"},{3,"16.280","16.210","16.612","49.102"},{4,"16.190","16.180","16.410","48.780"},{5,"16.240","16.150","16.480","48.870"}}},
			{Position: 4, Name: "Lina Moreau", Number: 34, Laps: 14, BestLap: "49.021", LastLap: "49.330", Gap: "+ 1 tour",
				LapTimes: []Lap{{1,"16.410","16.220","16.510","49.140"},{2,"16.300","16.210","16.511","49.021"},{3,"16.390","16.280","16.660","49.330"},{4,"16.310","16.250","16.590","49.150"},{5,"16.280","16.240","16.540","49.060"}}},
		},
	},
	{
		ID: 102, Name: "Endurance découverte", Category: "ENDURANCE", Status: "Terminée",
		Date: "19/11/2026", Duration: "30 minutes",
		Drivers: []Driver{
			{Position: 1, Name: "Équipe Comète", Number: 5, Laps: 36, BestLap: "47.982", LastLap: "48.210", Gap: "—"},
			{Position: 2, Name: "Équipe Orion", Number: 9, Laps: 35, BestLap: "48.104", LastLap: "48.502", Gap: "+ 1 tour"},
			{Position: 3, Name: "Équipe Nova", Number: 18, Laps: 35, BestLap: "48.450", LastLap: "48.770", Gap: "+ 1 tour"},
		},
	},
}

func main() {
	tmpl := template.Must(template.ParseFS(siteFiles, "templates/*.html"))
	css, err := siteFiles.ReadFile("static/style.css")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/static/style.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
		_, _ = w.Write(css)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" && r.URL.Path != "/live-timing/cap-karting" {
			http.NotFound(w, r)
			return
		}
		data := PageData{
			Title: "Live Timing", Circuit: "CAP Karting (fictif)", Location: "Saint-Pierre, France",
			Length: "1 200 m", TrackType: "Indoor • 3 secteurs", Updated: time.Now().Format("15:04:05"),
			Races: races,
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tmpl.ExecuteTemplate(w, "live.html", data); err != nil {
			http.Error(w, "Erreur de rendu", http.StatusInternalServerError)
			log.Println(err)
		}
	})

	http.HandleFunc("/race", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		var selected Race
		found := false
		for _, race := range races {
			if race.ID == id {
				selected, found = race, true
				break
			}
		}
		if !found {
			http.NotFound(w, r)
			return
		}
		data := PageData{
			Title: selected.Name, Circuit: "CAP Karting (fictif)", Location: "Saint-Pierre, France",
			Length: "1 200 m", TrackType: "Indoor • 3 secteurs", Updated: time.Now().Format("15:04:05"),
			Selected: selected, IsRacePage: true,
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tmpl.ExecuteTemplate(w, "race.html", data); err != nil {
			http.Error(w, "Erreur de rendu", http.StatusInternalServerError)
			log.Println(err)
		}
	})

	log.Println("Mock timing site: http://localhost:8080/live-timing/cap-karting")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
