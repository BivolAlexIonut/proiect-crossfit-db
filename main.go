package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/godror/godror"
)

const dsn = "CROSSFIT_ADMIN/HTsjReTy12@localhost:1521/XEPDB1"

var db *sql.DB

type Membru struct {
	ID          int    `json:"id"`
	Nume        string `json:"nume"`
	Prenume     string `json:"prenume"`
	Email       string `json:"email"`
	AbonamentID int    `json:"abonamentID"`
}
type Abonament struct {
	ID   int     `json:"id"`
	Tip  string  `json:"tip"`
	Pret float64 `json:"pret"`
}
type Antrenor struct {
	ID           int    `json:"id"`
	Nume         string `json:"nume"`
	Prenume      string `json:"prenume"`
	Specializare string `json:"specializare"`
}
type Produs struct {
	ID   int     `json:"id"`
	Nume string  `json:"nume"`
	Pret float64 `json:"pret"`
	Stoc int     `json:"stoc"`
}
type Echipament struct {
	ID        int    `json:"id"`
	Nume      string `json:"nume"`
	Cantitate int    `json:"cantitate"`
}
type TipAntrenament struct {
	ID        int    `json:"id"`
	Nume      string `json:"nume"`
	Descriere string `json:"descriere"`
}
type Clasa struct {
	ID               int    `json:"id"`
	NumeWOD          string `json:"numeWOD"`
	DescriereWOD     string `json:"descriereWOD"`
	DataOra          string `json:"dataOra"`
	AntrenorID       int    `json:"antrenorID"`
	TipAntrenamentID int    `json:"tipAntrenamentID"`
	NumeAntrenor     string `json:"numeAntrenor,omitempty"`
	NumeCategorie    string `json:"numeCategorie,omitempty"`
}
type Inscriere struct {
	MembruID     int    `json:"membruID"`
	ClasaID      int    `json:"clasaID"`
	NumeMembru   string `json:"numeMembru"`
	NumeWOD      string `json:"numeWOD"`
	DataOra      string `json:"dataOra"`
	NumeAntrenor string `json:"numeAntrenor"`
}

type Achizitie struct {
	ID             int    `json:"id"`
	MembruID       int    `json:"membruID"`
	ProdusID       int    `json:"produsID"`
	NumeMembru     string `json:"numeMembru,omitempty"`
	NumeProdus     string `json:"numeProdus,omitempty"`
	DataAchizitiei string `json:"dataAchizitiei"`
	Cantitate      int    `json:"cantitate"`
}

type Mentorat struct {
	AntrenorID   int    `json:"antrenorID"`
	MembruID     int    `json:"membruID"`
	NumeAntrenor string `json:"numeAntrenor,omitempty"`
	NumeMembru   string `json:"numeMembru,omitempty"`
}

type Competitie struct {
	ID       int     `json:"id"`
	Nume     string  `json:"nume"`
	Data     string  `json:"data"`
	Locatie  string  `json:"locatie"`
	Taxa     float64 `json:"taxa"`
}

type ParticipareCompetitie struct {
	CompetitieID   int    `json:"competitieID"`
	MembruID       int    `json:"membruID"`
	NumeCompetitie string `json:"numeCompetitie,omitempty"`
	NumeMembru     string `json:"numeMembru,omitempty"`
	LoculObtinut   int    `json:"loculObtinut"` // 0 daca nu a participat inca
}

// --- STRUCTURĂ NOUĂ PENTRU ORAR_TEMPLATE ---
type OrarTemplate struct {
	ID               int    `json:"id"`
	ZiuaSaptamanii   int    `json:"ziuaSaptamanii"`
	Ora              string `json:"ora"`
	NumeWODTemplate  string `json:"numeWODTemplate"`
	AntrenorID       int    `json:"antrenorID"`
	TipAntrenamentID int    `json"tipAntrenamentID"`

	// Câmpuri suplimentare pentru afișare (din JOIN-uri)
	NumeAntrenor  string `json:"numeAntrenor,omitempty"`
	NumeCategorie string `json:"numeCategorie,omitempty"`
}

// === ADaugă ACEASTĂ STRUCTURĂ LIPSA ===
// Structură pentru a citi regulile din BD
type OrarTemplateScan struct {
	TemplateID       int
	ZiuaSaptamanii   int
	Ora              string
	NumeWODTemplate  sql.NullString // Poate fi NULL
	AntrenorID       sql.NullInt64  // Poate fi NULL
	TipAntrenamentID sql.NullInt64  // Poate fi NULL
}

// --- Structuri NOI pentru Rapoarte ---
type RaportAbonamente struct {
	TipAbonament string `json:"tipAbonament"`
	NumarMembri  int    `json:"numarMembri"`
}
type RaportVizualizare struct {
	Nume         string  `json:"nume"`
	Prenume      string  `json:"prenume"`
	Email        string  `json:"email"`
	TipAbonament string  `json:"tipAbonament"`
	Pret         float64 `json:"pret"`
}
type RaportComplex struct {
	NumeMembru    string `json:"numeMembru"`
	PrenumeMembru string `json:"prenumeMembru"`
	NumeClasa     string `json:"numeClasa"`
	NumeAntrenor  string `json:"numeAntrenor"`
}

// --- MAIN ---
func main() {
	fmt.Println("Conectare la baza de date Oracle...")
	var err error
	db, err = sql.Open("godror", dsn)
	if err != nil {
		log.Fatal("Eroare la sql.Open: ", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal("Eroare la db.Ping: ", err)
	}
	fmt.Println("Conexiune la baza de date reușită!")

	// --- Rute Fișiere (Frontend) ---
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "index.html") })
	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "style.css") })
	http.HandleFunc("/script.js", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "script.js") })

	http.HandleFunc("/antrenori", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "antrenori.html") })
	http.HandleFunc("/antrenori.js", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "antrenori.js") })

	http.HandleFunc("/abonamente", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "abonamente.html") })
	http.HandleFunc("/abonamente.js", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "abonamente.js") })

	http.HandleFunc("/produse", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "produse.html") })
	http.HandleFunc("/produse.js", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "produse.js") })

	http.HandleFunc("/echipamente", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "echipamente.html") })
	http.HandleFunc("/echipamente.js", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "echipamente.js") })

	http.HandleFunc("/tipuri-antrenament", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "tipuri-antrenament.html") })
	http.HandleFunc("/tipuri-antrenament.js", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "tipuri-antrenament.js") })

	http.HandleFunc("/clase", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "clase.html") })
	http.HandleFunc("/clase.js", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "clase.js") })

	http.HandleFunc("/inscrieri", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "inscrieri.html") })
	http.HandleFunc("/inscrieri.js", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "inscrieri.js") })

	http.HandleFunc("/orar", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "orar.html") })
	http.HandleFunc("/orar.js", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "orar.js") })

	http.HandleFunc("/achizitii", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "achizitii.html") })
	http.HandleFunc("/achizitii.js", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "achizitii.js") })

	http.HandleFunc("/mentorat", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "mentorat.html") })
	http.HandleFunc("/mentorat.js", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "mentorat.js") })

	http.HandleFunc("/competitii", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "competitii.html") })
	http.HandleFunc("/competitii.js", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "competitii.js") })

	http.HandleFunc("/rapoarte", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "rapoarte.html") })
	http.HandleFunc("/rapoarte.js", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "rapoarte.js") })

	// --- Rute API (Backend) ---
	http.HandleFunc("/api/membri", handlerGetMembri)
	http.HandleFunc("/api/membru", handlerGetUnMembru)
	http.HandleFunc("/api/membri/add", handlerAddMembru)
	http.HandleFunc("/api/membri/update", handlerUpdateMembru)
	http.HandleFunc("/api/membri/delete", handlerDeleteMembru)

	http.HandleFunc("/api/antrenori", handlerGetAntrenori)
	http.HandleFunc("/api/antrenor", handlerGetUnAntrenor)
	http.HandleFunc("/api/antrenori/add", handlerAddAntrenor)
	http.HandleFunc("/api/antrenori/update", handlerUpdateAntrenor)
	http.HandleFunc("/api/antrenori/delete", handlerDeleteAntrenor)

	http.HandleFunc("/api/abonamente", handlerGetAbonamente)
	http.HandleFunc("/api/abonament", handlerGetUnAbonament)
	http.HandleFunc("/api/abonamente/add", handlerAddAbonament)
	http.HandleFunc("/api/abonamente/update", handlerUpdateAbonament)
	http.HandleFunc("/api/abonamente/delete", handlerDeleteAbonament)

	http.HandleFunc("/api/produse", handlerGetProduse)
	http.HandleFunc("/api/produs", handlerGetUnProdus)
	http.HandleFunc("/api/produse/add", handlerAddProdus)
	http.HandleFunc("/api/produse/update", handlerUpdateProdus)
	http.HandleFunc("/api/produse/delete", handlerDeleteProdus)

	http.HandleFunc("/api/echipamente", handlerGetEchipamente)
	http.HandleFunc("/api/echipament", handlerGetUnEchipament)
	http.HandleFunc("/api/echipamente/add", handlerAddEchipament)
	http.HandleFunc("/api/echipamente/update", handlerUpdateEchipament)
	http.HandleFunc("/api/echipamente/delete", handlerDeleteEchipament)

	http.HandleFunc("/api/tipuri-antrenament", handlerGetTipuri)
	http.HandleFunc("/api/tip-antrenament", handlerGetUnTip)
	http.HandleFunc("/api/tipuri-antrenament/add", handlerAddTip)
	http.HandleFunc("/api/tipuri-antrenament/update", handlerUpdateTip)
	http.HandleFunc("/api/tipuri-antrenament/delete", handlerDeleteTip)

	http.HandleFunc("/api/clase", handlerGetClase)
	http.HandleFunc("/api/clasa", handlerGetOClasa)
	http.HandleFunc("/api/clase/add", handlerAddClasa)
	http.HandleFunc("/api/clase/update", handlerUpdateClasa)
	http.HandleFunc("/api/clase/delete", handlerDeleteClasa)

	http.HandleFunc("/api/inscrieri", handlerGetInscrieri)
	http.HandleFunc("/api/inscrieri/add", handlerAddInscriere)
	http.HandleFunc("/api/inscrieri/delete", handlerDeleteInscriere)

	http.HandleFunc("/api/orar", handlerGetOrar)
	http.HandleFunc("/api/orar/single", handlerGetUnOrar)
	http.HandleFunc("/api/orar/add", handlerAddOrar)
	http.HandleFunc("/api/orar/update", handlerUpdateOrar)
	http.HandleFunc("/api/orar/delete", handlerDeleteOrar)
	http.HandleFunc("/api/orar/generate", handlerGenerateOrar)

	http.HandleFunc("/api/achizitii", handlerGetAchizitii)
	http.HandleFunc("/api/achizitii/add", handlerAddAchizitie)
	http.HandleFunc("/api/achizitii/delete", handlerDeleteAchizitie)

	http.HandleFunc("/api/mentorat", handlerGetMentorat)
	http.HandleFunc("/api/mentorat/add", handlerAddMentorat)
	http.HandleFunc("/api/mentorat/delete", handlerDeleteMentorat)

	http.HandleFunc("/api/competitii", handlerGetCompetitii)
	http.HandleFunc("/api/competitii/add", handlerAddCompetitie)
	http.HandleFunc("/api/competitii/delete", handlerDeleteCompetitie)
	http.HandleFunc("/api/competitii/participari", handlerGetParticipariCompetitie)
	http.HandleFunc("/api/competitii/participari/add", handlerAddParticipareCompetitie)

	// --- API-uri NOI PENTRU RAPOARTE ---
	http.HandleFunc("/api/raport/abonamente", handlerRaportAbonamente)
	http.HandleFunc("/api/raport/vizualizare-membri", handlerRaportVizualizare)
	http.HandleFunc("/api/raport/complex-inscrieri", handlerRaportComplex)

	// --- Pornirea Serverului ---
	port := ":8080"
	fmt.Println("Serverul web a pornit. Accesează http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}

// ===================================================================
// HANDLERE PENTRU COMPETITII
// ===================================================================
func handlerGetCompetitii(w http.ResponseWriter, _ *http.Request) {
	rows, err := db.Query("SELECT CompetitieID, Nume, TO_CHAR(Data, 'YYYY-MM-DD'), Locatie, Taxa FROM COMPETITII ORDER BY Data DESC")
	if err != nil {
		http.Error(w, "Eroare BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer rows.Close()

	var competitii []Competitie
	for rows.Next() {
		var c Competitie
		if err := rows.Scan(&c.ID, &c.Nume, &c.Data, &c.Locatie, &c.Taxa); err != nil {
			http.Error(w, "Eroare Scan", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		competitii = append(competitii, c)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(competitii)
}

func handlerAddCompetitie(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var c Competitie
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		log.Println(err)
		return
	}
	query := `INSERT INTO COMPETITII (CompetitieID, Nume, Data, Locatie, Taxa) VALUES (COMPETITII_SEQ.NEXTVAL, :1, TO_DATE(:2, 'YYYY-MM-DD'), :3, :4)`
	_, err := db.Exec(query, c.Nume, c.Data, c.Locatie, c.Taxa)
	if err != nil {
		http.Error(w, "Eroare la inserarea în BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Competiție adăugată cu succes"})
}

func handlerDeleteCompetitie(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var payload struct {
		ID int `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		log.Println(err)
		return
	}
	query := `DELETE FROM COMPETITII WHERE CompetitieID = :1`
	_, err := db.Exec(query, payload.ID)
	if err != nil {
		http.Error(w, "Eroare la ștergerea din BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Competiție ștearsă cu succes"})
}

// HANDLERE PENTRU PARTICIPARI LA COMPETITII
func handlerGetParticipariCompetitie(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		// Daca nu e dat ID, returnam toate participarile (pentru un tabel general)
		query := `
			SELECT pc.CompetitieID, pc.MembruID, c.Nume, m.Nume || ' ' || m.Prenume, pc.LoculObtinut
			FROM PARTICIPARI_COMPETITIE pc
			JOIN COMPETITII c ON pc.CompetitieID = c.CompetitieID
			JOIN MEMBRI m ON pc.MembruID = m.MembruID
			ORDER BY c.Data DESC, pc.LoculObtinut ASC
		`
		rows, err := db.Query(query)
		if err != nil {
			http.Error(w, "Eroare BD", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		defer rows.Close()
		var participari []ParticipareCompetitie
		for rows.Next() {
			var p ParticipareCompetitie
			var loc sql.NullInt64
			if err := rows.Scan(&p.CompetitieID, &p.MembruID, &p.NumeCompetitie, &p.NumeMembru, &loc); err != nil {
				http.Error(w, "Eroare Scan", http.StatusInternalServerError)
				return
			}
			if loc.Valid {
				p.LoculObtinut = int(loc.Int64)
			}
			participari = append(participari, p)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(participari)
		return
	}

	// Daca e dat ID, putem returna participarile specifice unei competitii (nu e folosit inca in UI dar e bun de avut)
	// ... (implementare optionala)
}

func handlerAddParticipareCompetitie(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var p ParticipareCompetitie
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		log.Println(err)
		return
	}
	query := `INSERT INTO PARTICIPARI_COMPETITIE (CompetitieID, MembruID) VALUES (:1, :2)`
	_, err := db.Exec(query, p.CompetitieID, p.MembruID)
	if err != nil {
		http.Error(w, "Eroare la înscriere (posibil duplicat)", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Membru înscris la competiție!"})
}


// ===================================================================
// HANDLERE PENTRU ACHIZITII
// ===================================================================
func handlerGetAchizitii(w http.ResponseWriter, _ *http.Request) {
	query := `
		SELECT 
			a.AchizitieID, a.MembruID, a.ProdusID,
			m.Nume || ' ' || m.Prenume AS NumeMembru,
			p.NumeProdus, 
			TO_CHAR(a.DataAchizitiei, 'YYYY-MM-DD') AS DataAchizitiei,
			a.Cantitate
		FROM ACHIZITII a
		JOIN MEMBRI m ON a.MembruID = m.MembruID
		JOIN PRODUSE p ON a.ProdusID = p.ProdusID
		ORDER BY a.DataAchizitiei DESC
	`
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Eroare BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer rows.Close()

	var achizitii []Achizitie
	for rows.Next() {
		var a Achizitie
		if err := rows.Scan(&a.ID, &a.MembruID, &a.ProdusID, &a.NumeMembru, &a.NumeProdus, &a.DataAchizitiei, &a.Cantitate); err != nil {
			http.Error(w, "Eroare Scan", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		achizitii = append(achizitii, a)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(achizitii)
}

func handlerAddAchizitie(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var a Achizitie
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		log.Println(err)
		return
	}
	// Implicit cantitate 1 dacă nu e setată
	if a.Cantitate <= 0 {
		a.Cantitate = 1
	}

	query := `INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, Cantitate) VALUES (ACHIZITII_SEQ.NEXTVAL, :1, :2, :3)`
	_, err := db.Exec(query, a.MembruID, a.ProdusID, a.Cantitate)
	if err != nil {
		http.Error(w, "Eroare la inserarea în BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Achiziție adăugată cu succes"})
}

func handlerDeleteAchizitie(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var payload struct {
		ID int `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		log.Println(err)
		return
	}
	query := `DELETE FROM ACHIZITII WHERE AchizitieID = :1`
	_, err := db.Exec(query, payload.ID)
	if err != nil {
		http.Error(w, "Eroare la ștergerea din BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Achiziție ștearsă cu succes"})
}

// ===================================================================
// HANDLERE PENTRU MENTORAT
// ===================================================================
func handlerGetMentorat(w http.ResponseWriter, _ *http.Request) {
	query := `
		SELECT 
			men.AntrenorID, men.MembruID,
			a.Nume || ' ' || a.Prenume AS NumeAntrenor,
			m.Nume || ' ' || m.Prenume AS NumeMembru
		FROM MENTORAT men
		JOIN ANTRENORI a ON men.AntrenorID = a.AntrenorID
		JOIN MEMBRI m ON men.MembruID = m.MembruID
		ORDER BY a.Nume
	`
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Eroare BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer rows.Close()

	var mentorate []Mentorat
	for rows.Next() {
		var m Mentorat
		if err := rows.Scan(&m.AntrenorID, &m.MembruID, &m.NumeAntrenor, &m.NumeMembru); err != nil {
			http.Error(w, "Eroare Scan", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		mentorate = append(mentorate, m)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mentorate)
}

func handlerAddMentorat(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var m Mentorat
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		log.Println(err)
		return
	}
	query := `INSERT INTO MENTORAT (AntrenorID, MembruID) VALUES (:1, :2)`
	_, err := db.Exec(query, m.AntrenorID, m.MembruID)
	if err != nil {
		http.Error(w, "Eroare la inserarea în BD. Relația există deja?", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Mentorat adăugat cu succes"})
}

func handlerDeleteMentorat(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var m Mentorat
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		log.Println(err)
		return
	}
	query := `DELETE FROM MENTORAT WHERE AntrenorID = :1 AND MembruID = :2`
	_, err := db.Exec(query, m.AntrenorID, m.MembruID)
	if err != nil {
		http.Error(w, "Eroare la ștergerea din BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Mentorat șters cu succes"})
}


// ===================================================================
// HANDLERE PENTRU MEMBRI
// ===================================================================
func handlerGetMembri(w http.ResponseWriter, _ *http.Request) {
	rows, err := db.Query("SELECT MembruID, Nume, Prenume, Email FROM MEMBRI ORDER BY Nume")
	if err != nil {
		http.Error(w, "Eroare BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("Eroare la inchiderea rows:", err)
		}
	}()
	var membri []Membru
	for rows.Next() {
		var m Membru
		if err := rows.Scan(&m.ID, &m.Nume, &m.Prenume, &m.Email); err != nil {
			http.Error(w, "Eroare Scan", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		membri = append(membri, m)
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(membri); err != nil {
		log.Println("Eroare la encodare JSON membri:", err)
	}
}
func handlerGetUnMembru(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID invalid", http.StatusBadRequest)
		return
	}
	query := `SELECT MembruID, Nume, Prenume, Email, AbonamentID FROM MEMBRI WHERE MembruID = :1`
	var m Membru
	err = db.QueryRow(query, id).Scan(&m.ID, &m.Nume, &m.Prenume, &m.Email, &m.AbonamentID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Membrul nu a fost găsit", http.StatusNotFound)
			return
		}
		http.Error(w, "Eroare Scan", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(m); err != nil {
		log.Println("Eroare la encodare JSON membru:", err)
	}
}
func handlerAddMembru(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var m Membru
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		log.Println(err)
		return
	}
	query := `INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (MEMBRI_SEQ.NEXTVAL, :1, :2, :3, :4)`
	_, err := db.Exec(query, m.Nume, m.Prenume, m.Email, m.AbonamentID)
	if err != nil {
		http.Error(w, "Eroare la inserarea în BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(map[string]string{"mesaj": "Membru adăugat cu succes"}); err != nil {
		log.Println("Eroare la encodare JSON add membru:", err)
	}
}
func handlerUpdateMembru(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var m Membru
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		log.Println(err)
		return
	}
	query := `UPDATE MEMBRI SET Nume = :1, Prenume = :2, Email = :3, AbonamentID = :4 WHERE MembruID = :5`
	_, err := db.Exec(query, m.Nume, m.Prenume, m.Email, m.AbonamentID, m.ID)
	if err != nil {
		http.Error(w, "Eroare la actualizarea în BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"mesaj": "Membru actualizat cu succes"}); err != nil {
		log.Println("Eroare la encodare JSON update membru:", err)
	}
}
func handlerDeleteMembru(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var payload struct {
		ID int `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		log.Println(err)
		return
	}
	query := `DELETE FROM MEMBRI WHERE MembruID = :1`
	_, err := db.Exec(query, payload.ID)
	if err != nil {
		http.Error(w, "Eroare la ștergerea din BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"mesaj": "Membru șters cu succes"}); err != nil {
		log.Println("Eroare la encodare JSON delete membru:", err)
	}
}

// ===================================================================
// HANDLERE PENTRU ANTRENORI
// ===================================================================
func handlerGetAntrenori(w http.ResponseWriter, _ *http.Request) {
	rows, err := db.Query("SELECT AntrenorID, Nume, Prenume, Specializare FROM ANTRENORI ORDER BY Nume")
	if err != nil {
		http.Error(w, "Eroare BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("Eroare la inchiderea rows:", err)
		}
	}()
	var antrenori []Antrenor
	for rows.Next() {
		var a Antrenor
		if err := rows.Scan(&a.ID, &a.Nume, &a.Prenume, &a.Specializare); err != nil {
			http.Error(w, "Eroare Scan", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		antrenori = append(antrenori, a)
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(antrenori); err != nil {
		log.Println("Eroare la encodare JSON antrenori:", err)
	}
}
func handlerGetUnAntrenor(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID invalid", http.StatusBadRequest)
		return
	}
	query := `SELECT AntrenorID, Nume, Prenume, Specializare FROM ANTRENORI WHERE AntrenorID = :1`
	var a Antrenor
	err = db.QueryRow(query, id).Scan(&a.ID, &a.Nume, &a.Prenume, &a.Specializare)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Antrenorul nu a fost găsit", http.StatusNotFound)
			return
		}
		http.Error(w, "Eroare Scan", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(a); err != nil {
		log.Println("Eroare la encodare JSON antrenor:", err)
	}
}
func handlerAddAntrenor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var a Antrenor
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		log.Println(err)
		return
	}
	query := `INSERT INTO ANTRENORI (AntrenorID, Nume, Prenume, Specializare) VALUES (ANTRENORI_SEQ.NEXTVAL, :1, :2, :3)`
	_, err := db.Exec(query, a.Nume, a.Prenume, a.Specializare)
	if err != nil {
		http.Error(w, "Eroare la inserarea în BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(map[string]string{"mesaj": "Antrenor adăugat cu succes"}); err != nil {
		log.Println("Eroare la encodare JSON add antrenor:", err)
	}
}
func handlerUpdateAntrenor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var a Antrenor
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		log.Println(err)
		return
	}
	query := `UPDATE ANTRENORI SET Nume = :1, Prenume = :2, Specializare = :3 WHERE AntrenorID = :4`
	_, err := db.Exec(query, a.Nume, a.Prenume, a.Specializare, a.ID)
	if err != nil {
		http.Error(w, "Eroare la actualizarea în BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"mesaj": "Antrenor actualizat cu succes"}); err != nil {
		log.Println("Eroare la encodare JSON update antrenor:", err)
	}
}
func handlerDeleteAntrenor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var payload struct {
		ID int `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		log.Println(err)
		return
	}
	query := `DELETE FROM ANTRENORI WHERE AntrenorID = :1`
	_, err := db.Exec(query, payload.ID)
	if err != nil {
		http.Error(w, "Eroare la ștergerea din BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"mesaj": "Antrenor șters cu succes"}); err != nil {
		log.Println("Eroare la encodare JSON delete antrenor:", err)
	}
}

// ===================================================================
// HANDLERE PENTRU ABONAMENTE
// ===================================================================
func handlerGetAbonamente(w http.ResponseWriter, _ *http.Request) {
	rows, err := db.Query("SELECT AbonamentID, TipAbonament, Pret FROM ABONAMENTE ORDER BY Pret")
	if err != nil {
		http.Error(w, "Eroare BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("Eroare la inchiderea rows:", err)
		}
	}()
	var abonamente []Abonament
	for rows.Next() {
		var a Abonament
		if err := rows.Scan(&a.ID, &a.Tip, &a.Pret); err != nil {
			http.Error(w, "Eroare Scan", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		abonamente = append(abonamente, a)
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(abonamente); err != nil {
		log.Println("Eroare la encodare JSON abonamente:", err)
	}
}
func handlerGetUnAbonament(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID invalid", http.StatusBadRequest)
		return
	}
	query := `SELECT AbonamentID, TipAbonament, Pret FROM ABONAMENTE WHERE AbonamentID = :1`
	var a Abonament
	err = db.QueryRow(query, id).Scan(&a.ID, &a.Tip, &a.Pret)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Abonamentul nu a fost găsit", http.StatusNotFound)
			return
		}
		http.Error(w, "Eroare Scan", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(a); err != nil {
		log.Println("Eroare la encodare JSON abonament:", err)
	}
}
func handlerAddAbonament(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var a Abonament
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		log.Println(err)
		return
	}
	query := `INSERT INTO ABONAMENTE (AbonamentID, TipAbonament, Pret) VALUES (ABONAMENTE_SEQ.NEXTVAL, :1, :2)`
	_, err := db.Exec(query, a.Tip, a.Pret)
	if err != nil {
		http.Error(w, "Eroare la inserarea în BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(map[string]string{"mesaj": "Abonament adăugat cu succes"}); err != nil {
		log.Println("Eroare la encodare JSON add abonament:", err)
	}
}
func handlerUpdateAbonament(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var a Abonament
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		log.Println(err)
		return
	}
	query := `UPDATE ABONAMENTE SET TipAbonament = :1, Pret = :2 WHERE AbonamentID = :3`
	_, err := db.Exec(query, a.Tip, a.Pret, a.ID)
	if err != nil {
		http.Error(w, "Eroare la actualizarea în BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"mesaj": "Abonament actualizat cu succes"}); err != nil {
		log.Println("Eroare la encodare JSON update abonament:", err)
	}
}
func handlerDeleteAbonament(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var payload struct {
		ID int `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		log.Println(err)
		return
	}
	query := `DELETE FROM ABONAMENTE WHERE AbonamentID = :1`
	_, err := db.Exec(query, payload.ID)
	if err != nil {
		http.Error(w, "Eroare la ștergerea din BD. Asigură-te că niciun membru nu folosește acest abonament.", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"mesaj": "Abonament șters cu succes"}); err != nil {
		log.Println("Eroare la encodare JSON delete abonament:", err)
	}
}

// ===================================================================
// HANDLERE PENTRU PRODUSE
// ===================================================================
func handlerGetProduse(w http.ResponseWriter, _ *http.Request) {
	rows, err := db.Query("SELECT ProdusID, NumeProdus, PretCurent, Stoc FROM PRODUSE ORDER BY NumeProdus")
	if err != nil {
		http.Error(w, "Eroare BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("Eroare la inchiderea rows:", err)
		}
	}()
	var produse []Produs
	for rows.Next() {
		var p Produs
		if err := rows.Scan(&p.ID, &p.Nume, &p.Pret, &p.Stoc); err != nil {
			http.Error(w, "Eroare Scan", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		produse = append(produse, p)
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(produse); err != nil {
		log.Println("Eroare la encodare JSON produse:", err)
	}
}
func handlerGetUnProdus(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID invalid", http.StatusBadRequest)
		return
	}
	query := `SELECT ProdusID, NumeProdus, PretCurent, Stoc FROM PRODUSE WHERE ProdusID = :1`
	var p Produs
	err = db.QueryRow(query, id).Scan(&p.ID, &p.Nume, &p.Pret, &p.Stoc)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Produsul nu a fost găsit", http.StatusNotFound)
			return
		}
		http.Error(w, "Eroare Scan", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(p); err != nil {
		log.Println("Eroare la encodare JSON produs:", err)
	}
}
func handlerAddProdus(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var p Produs
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		log.Println(err)
		return
	}
	query := `INSERT INTO PRODUSE (ProdusID, NumeProdus, PretCurent, Stoc) VALUES (PRODUSE_SEQ.NEXTVAL, :1, :2, :3)`
	_, err := db.Exec(query, p.Nume, p.Pret, p.Stoc)
	if err != nil {
		http.Error(w, "Eroare la inserarea în BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(map[string]string{"mesaj": "Produs adăugat cu succes"}); err != nil {
		log.Println("Eroare la encodare JSON add produs:", err)
	}
}
func handlerUpdateProdus(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var p Produs
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		log.Println(err)
		return
	}
	query := `UPDATE PRODUSE SET NumeProdus = :1, PretCurent = :2, Stoc = :3 WHERE ProdusID = :4`
	_, err := db.Exec(query, p.Nume, p.Pret, p.Stoc, p.ID)
	if err != nil {
		http.Error(w, "Eroare la actualizarea în BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"mesaj": "Produs actualizat cu succes"}); err != nil {
		log.Println("Eroare la encodare JSON update produs:", err)
	}
}
func handlerDeleteProdus(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var payload struct {
		ID int `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		log.Println(err)
		return
	}
	query := `DELETE FROM PRODUSE WHERE ProdusID = :1`
	_, err := db.Exec(query, payload.ID)
	if err != nil {
		http.Error(w, "Eroare la ștergerea din BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"mesaj": "Produs șters cu succes"}); err != nil {
		log.Println("Eroare la encodare JSON delete produs:", err)
	}
}

// ===================================================================
// HANDLERE PENTRU ECHIPAMENTE
// ===================================================================
func handlerGetEchipamente(w http.ResponseWriter, _ *http.Request) {
	rows, err := db.Query("SELECT EchipamentID, NumeEchipament, CantitateTotala FROM ECHIPAMENTE ORDER BY NumeEchipament")
	if err != nil {
		http.Error(w, "Eroare BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("Eroare la inchiderea rows:", err)
		}
	}()
	var echipamente []Echipament
	for rows.Next() {
		var e Echipament
		if err := rows.Scan(&e.ID, &e.Nume, &e.Cantitate); err != nil {
			http.Error(w, "Eroare Scan", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		echipamente = append(echipamente, e)
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(echipamente); err != nil {
		log.Println("Eroare la encodare JSON echipamente:", err)
	}
}
func handlerGetUnEchipament(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID invalid", http.StatusBadRequest)
		return
	}
	query := `SELECT EchipamentID, NumeEchipament, CantitateTotala FROM ECHIPAMENTE WHERE EchipamentID = :1`
	var e Echipament
	err = db.QueryRow(query, id).Scan(&e.ID, &e.Nume, &e.Cantitate)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Echipamentul nu a fost găsit", http.StatusNotFound)
			return
		}
		http.Error(w, "Eroare Scan", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(e); err != nil {
		log.Println("Eroare la encodare JSON echipament:", err)
	}
}
func handlerAddEchipament(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var e Echipament
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		log.Println(err)
		return
	}
	query := `INSERT INTO ECHIPAMENTE (EchipamentID, NumeEchipament, CantitateTotala) VALUES (ECHIPAMENTE_SEQ.NEXTVAL, :1, :2)`
	_, err := db.Exec(query, e.Nume, e.Cantitate)
	if err != nil {
		http.Error(w, "Eroare la inserarea în BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(map[string]string{"mesaj": "Echipament adăugat cu succes"}); err != nil {
		log.Println("Eroare la encodare JSON add echipament:", err)
	}
}
func handlerUpdateEchipament(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var e Echipament
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		log.Println(err)
		return
	}
	query := `UPDATE ECHIPAMENTE SET NumeEchipament = :1, CantitateTotala = :2 WHERE EchipamentID = :3`
	_, err := db.Exec(query, e.Nume, e.Cantitate, e.ID)
	if err != nil {
		http.Error(w, "Eroare la actualizarea în BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"mesaj": "Echipament actualizat cu succes"}); err != nil {
		log.Println("Eroare la encodare JSON update echipament:", err)
	}
}
func handlerDeleteEchipament(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var payload struct {
		ID int `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		log.Println(err)
		return
	}
	query := `DELETE FROM ECHIPAMENTE WHERE EchipamentID = :1`
	_, err := db.Exec(query, payload.ID)
	if err != nil {
		http.Error(w, "Eroare la ștergerea din BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"mesaj": "Echipament șters cu succes"}); err != nil {
		log.Println("Eroare la encodare JSON delete echipament:", err)
	}
}

// ===================================================================
// HANDLERE PENTRU TIPURI_ANTRENAMENT
// ===================================================================
func handlerGetTipuri(w http.ResponseWriter, _ *http.Request) {
	rows, err := db.Query("SELECT TipAntrenamentID, NumeWOD, Descriere FROM TIPURI_ANTRENAMENT ORDER BY NumeWOD")
	if err != nil {
		http.Error(w, "Eroare BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("Eroare la inchiderea rows:", err)
		}
	}()
	var tipuri []TipAntrenament
	for rows.Next() {
		var t TipAntrenament
		var descriere sql.NullString
		if err := rows.Scan(&t.ID, &t.Nume, &descriere); err != nil {
			http.Error(w, "Eroare Scan", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		if descriere.Valid {
			t.Descriere = descriere.String
		}
		tipuri = append(tipuri, t)
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tipuri); err != nil {
		log.Println("Eroare la encodare JSON tipuri:", err)
	}
}
func handlerGetUnTip(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID invalid", http.StatusBadRequest)
		return
	}
	query := `SELECT TipAntrenamentID, NumeWOD, Descriere FROM TIPURI_ANTRENAMENT WHERE TipAntrenamentID = :1`
	var t TipAntrenament
	var descriere sql.NullString
	err = db.QueryRow(query, id).Scan(&t.ID, &t.Nume, &descriere)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Tipul de antrenament nu a fost găsit", http.StatusNotFound)
			return
		}
		http.Error(w, "Eroare Scan", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if descriere.Valid {
		t.Descriere = descriere.String
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(t); err != nil {
		log.Println("Eroare la encodare JSON tip:", err)
	}
}
func handlerAddTip(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var t TipAntrenament
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		log.Println(err)
		return
	}
	query := `INSERT INTO TIPURI_ANTRENAMENT (TipAntrenamentID, NumeWOD, Descriere) VALUES (TIPURI_ANTRENAMENT_SEQ.NEXTVAL, :1, :2)`
	_, err := db.Exec(query, t.Nume, t.Descriere)
	if err != nil {
		http.Error(w, "Eroare la inserarea în BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(map[string]string{"mesaj": "Tip de antrenament adăugat cu succes"}); err != nil {
		log.Println("Eroare la encodare JSON add tip:", err)
	}
}
func handlerUpdateTip(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var t TipAntrenament
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		log.Println(err)
		return
	}
	query := `UPDATE TIPURI_ANTRENAMENT SET NumeWOD = :1, Descriere = :2 WHERE TipAntrenamentID = :3`
	_, err := db.Exec(query, t.Nume, t.Descriere, t.ID)
	if err != nil {
		http.Error(w, "Eroare la actualizarea în BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"mesaj": "Tip de antrenament actualizat cu succes"}); err != nil {
		log.Println("Eroare la encodare JSON update tip:", err)
	}
}
func handlerDeleteTip(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var payload struct {
		ID int `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		log.Println(err)
		return
	}
	query := `DELETE FROM TIPURI_ANTRENAMENT WHERE TipAntrenamentID = :1`
	_, err := db.Exec(query, payload.ID)
	if err != nil {
		http.Error(w, "Eroare la ștergerea din BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"mesaj": "Tip de antrenament șters cu succes"}); err != nil {
		log.Println("Eroare la encodare JSON delete tip:", err)
	}
}

// ===================================================================
// HANDLERE PENTRU CLASE
// ===================================================================
func handlerGetClase(w http.ResponseWriter, _ *http.Request) {
	query := `
		SELECT 
			c.ClasaID, c.NumeWOD, c.DescriereWOD, TO_CHAR(c.DataOra, 'YYYY-MM-DD"T"HH24:MI') AS DataOraFormatata,
			a.AntrenorID, a.Nume || ' ' || a.Prenume AS NumeAntrenor,
			t.TipAntrenamentID, t.NumeWOD AS NumeCategorie
		FROM CLASE c
		LEFT JOIN ANTRENORI a ON c.AntrenorID = a.AntrenorID
		LEFT JOIN TIPURI_ANTRENAMENT t ON c.TipAntrenamentID = t.TipAntrenamentID
		ORDER BY c.DataOra DESC
	`
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Eroare BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("Eroare la inchiderea rows:", err)
		}
	}()

	var clase []Clasa
	for rows.Next() {
		var c Clasa
		var numeAntrenor sql.NullString
		var numeCategorie sql.NullString
		var descriereWOD sql.NullString

		if err := rows.Scan(
			&c.ID, &c.NumeWOD, &descriereWOD, &c.DataOra,
			&c.AntrenorID, &numeAntrenor,
			&c.TipAntrenamentID, &numeCategorie,
		); err != nil {
			http.Error(w, "Eroare Scan", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		if numeAntrenor.Valid {
			c.NumeAntrenor = numeAntrenor.String
		}
		if numeCategorie.Valid {
			c.NumeCategorie = numeCategorie.String
		}
		if descriereWOD.Valid {
			c.DescriereWOD = descriereWOD.String
		}

		clase = append(clase, c)
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(clase); err != nil {
		log.Println("Eroare la encodare JSON clase:", err)
	}
}
func handlerGetOClasa(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID invalid", http.StatusBadRequest)
		return
	}
	query := `
		SELECT ClasaID, NumeWOD, DescriereWOD, 
		       TO_CHAR(DataOra, 'YYYY-MM-DD"T"HH24:MI') AS DataOraFormatata, 
		       AntrenorID, TipAntrenamentID 
		FROM CLASE 
		WHERE ClasaID = :1
	`
	var c Clasa
	var descriereWOD sql.NullString

	err = db.QueryRow(query, id).Scan(&c.ID, &c.NumeWOD, &descriereWOD, &c.DataOra, &c.AntrenorID, &c.TipAntrenamentID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Clasa nu a fost găsită", http.StatusNotFound)
			return
		}
		http.Error(w, "Eroare Scan", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	if descriereWOD.Valid {
		c.DescriereWOD = descriereWOD.String
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(c); err != nil {
		log.Println("Eroare la encodare JSON clasa:", err)
	}
}
func handlerAddClasa(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var c Clasa
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		log.Println(err)
		return
	}
	query := `
		INSERT INTO CLASE (ClasaID, NumeWOD, DescriereWOD, DataOra, AntrenorID, TipAntrenamentID)
		VALUES (CLASE_SEQ.NEXTVAL, :1, :2, TO_TIMESTAMP(:3, 'YYYY-MM-DD"T"HH24:MI'), :4, :5)
	`
	_, err := db.Exec(query, c.NumeWOD, c.DescriereWOD, c.DataOra, c.AntrenorID, c.TipAntrenamentID)
	if err != nil {
		http.Error(w, "Eroare la inserarea în BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(map[string]string{"mesaj": "Clasă adăugată cu succes"}); err != nil {
		log.Println("Eroare la encodare JSON add clasa:", err)
	}
}
func handlerUpdateClasa(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var c Clasa
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		log.Println(err)
		return
	}
	query := `
		UPDATE CLASE 
		SET NumeWOD = :1, DescriereWOD = :2, DataOra = TO_TIMESTAMP(:3, 'YYYY-MM-DD"T"HH24:MI'), 
		    AntrenorID = :4, TipAntrenamentID = :5
		WHERE ClasaID = :6
	`
	_, err := db.Exec(query, c.NumeWOD, c.DescriereWOD, c.DataOra, c.AntrenorID, c.TipAntrenamentID, c.ID)
	if err != nil {
		http.Error(w, "Eroare la actualizarea în BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"mesaj": "Clasă actualizată cu succes"}); err != nil {
		log.Println("Eroare la encodare JSON update clasa:", err)
	}
}
func handlerDeleteClasa(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var payload struct {
		ID int `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		log.Println(err)
		return
	}
	query := `DELETE FROM CLASE WHERE ClasaID = :1`
	_, err := db.Exec(query, payload.ID)
	if err != nil {
		http.Error(w, "Eroare la ștergerea din BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"mesaj": "Clasă ștersă cu succes"}); err != nil {
		log.Println("Eroare la encodare JSON delete clasa:", err)
	}
}

// ===================================================================
// HANDLERE PENTRU ÎNSCRIERI
// ===================================================================
func handlerGetInscrieri(w http.ResponseWriter, _ *http.Request) {
	query := `
		SELECT 
			i.MembruID, i.ClasaID,
			m.Nume || ' ' || m.Prenume AS NumeMembru,
			c.NumeWOD, TO_CHAR(c.DataOra, 'YYYY-MM-DD HH24:MI') AS DataOraFormatata,
			a.Nume || ' ' || a.Prenume AS NumeAntrenor
		FROM INSCRIERI i
		JOIN MEMBRI m ON i.MembruID = m.MembruID
		JOIN CLASE c ON i.ClasaID = c.ClasaID
		LEFT JOIN ANTRENORI a ON c.AntrenorID = a.AntrenorID
		ORDER BY c.DataOra DESC
	`
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Eroare BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("Eroare la inchiderea rows:", err)
		}
	}()

	var inscrieri []Inscriere
	for rows.Next() {
		var i Inscriere
		var numeAntrenor sql.NullString
		var dataOra sql.NullString

		if err := rows.Scan(
			&i.MembruID, &i.ClasaID,
			&i.NumeMembru,
			&i.NumeWOD, &dataOra,
			&numeAntrenor,
		); err != nil {
			http.Error(w, "Eroare Scan", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		if numeAntrenor.Valid {
			i.NumeAntrenor = numeAntrenor.String
		}
		if dataOra.Valid {
			i.DataOra = dataOra.String
		}
		inscrieri = append(inscrieri, i)
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(inscrieri); err != nil {
		log.Println("Eroare la encodare JSON inscrieri:", err)
	}
}
func handlerAddInscriere(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var i Inscriere
	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		log.Println(err)
		return
	}
	var check struct {
		TipAbonament     string
		InscrieriNumar   int
		PermiteInscriere int
	}
	queryCheck := `
		SELECT
			a.TipAbonament,
			COUNT(insc.MembruID) AS NumarInscrieri,
			CASE
				WHEN a.TipAbonament LIKE '%Nelimitat%' THEN 1
				WHEN a.TipAbonament = '12 Sedinte' AND COUNT(insc.MembruID) < 12 THEN 1
				WHEN a.TipAbonament = '10 Sedinte' AND COUNT(insc.MembruID) < 10 THEN 1
				WHEN a.TipAbonament = '8 Sedinte' AND COUNT(insc.MembruID) < 8 THEN 1
				WHEN a.TipAbonament = '4 Sedinte' AND COUNT(insc.MembruID) < 4 THEN 1
				ELSE 0
			END AS PermiteInscriere
		FROM MEMBRI m
		JOIN ABONAMENTE a ON m.AbonamentID = a.AbonamentID
		LEFT JOIN INSCRIERI insc ON m.MembruID = insc.MembruID
		WHERE m.MembruID = :1
		GROUP BY a.TipAbonament
	`
	err := db.QueryRow(queryCheck, i.MembruID).Scan(&check.TipAbonament, &check.InscrieriNumar, &check.PermiteInscriere)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Eroare: Membrul sau Abonamentul nu a fost găsit.", http.StatusNotFound)
			return
		}
		http.Error(w, "Eroare la verificarea abonamentului", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if check.PermiteInscriere == 0 {
		mesajEroare := fmt.Sprintf("Eroare: Limita de ședințe a fost atinsă! (Abonament: %s, Înscrieri: %d)", check.TipAbonament, check.InscrieriNumar)
		http.Error(w, mesajEroare, http.StatusForbidden)
		log.Println(mesajEroare)
		return
	}
	queryInsert := `INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (:1, :2)`
	_, err = db.Exec(queryInsert, i.MembruID, i.ClasaID)
	if err != nil {
		http.Error(w, "Eroare la inserarea în BD. Membrul este deja înscris la această clasă?", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(map[string]string{"mesaj": "Înscriere adăugată cu succes"}); err != nil {
		log.Println("Eroare la encodare JSON add inscriere:", err)
	}
}
func handlerDeleteInscriere(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var i Inscriere
	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		log.Println(err)
		return
	}
	query := `DELETE FROM INSCRIERI WHERE MembruID = :1 AND ClasaID = :2`
	_, err := db.Exec(query, i.MembruID, i.ClasaID)
	if err != nil {
		http.Error(w, "Eroare la ștergerea din BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"mesaj": "Înscriere anulată cu succes"}); err != nil {
		log.Println("Eroare la encodare JSON delete inscriere:", err)
	}
}

// ===================================================================
// HANDLERE PENTRU ORAR_TEMPLATE
// ===================================================================
func handlerGetOrar(w http.ResponseWriter, _ *http.Request) {
	query := `
		SELECT 
			ot.TemplateID, ot.ZiuaSaptamanii, ot.Ora, ot.NumeWOD_Template,
			a.AntrenorID, a.Nume || ' ' || a.Prenume AS NumeAntrenor,
			t.TipAntrenamentID, t.NumeWOD AS NumeCategorie
		FROM ORAR_TEMPLATE ot
		LEFT JOIN ANTRENORI a ON ot.AntrenorID = a.AntrenorID
		LEFT JOIN TIPURI_ANTRENAMENT t ON ot.TipAntrenamentID = t.TipAntrenamentID
		ORDER BY ot.ZiuaSaptamanii, ot.Ora
	`
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Eroare BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("Eroare la inchiderea rows:", err)
		}
	}()

	var orar []OrarTemplate
	for rows.Next() {
		var o OrarTemplate
		var numeAntrenor sql.NullString
		var numeCategorie sql.NullString
		var numeWOD sql.NullString

		if err := rows.Scan(
			&o.ID, &o.ZiuaSaptamanii, &o.Ora, &numeWOD,
			&o.AntrenorID, &numeAntrenor,
			&o.TipAntrenamentID, &numeCategorie,
		); err != nil {
			http.Error(w, "Eroare Scan", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		if numeAntrenor.Valid {
			o.NumeAntrenor = numeAntrenor.String
		}
		if numeCategorie.Valid {
			o.NumeCategorie = numeCategorie.String
		}
		if numeWOD.Valid {
			o.NumeWODTemplate = numeWOD.String
		}

		orar = append(orar, o)
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(orar); err != nil {
		log.Println("Eroare la encodare JSON orar:", err)
	}
}
func handlerGetUnOrar(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID invalid", http.StatusBadRequest)
		return
	}
	query := `SELECT TemplateID, ZiuaSaptamanii, Ora, NumeWOD_Template, AntrenorID, TipAntrenamentID FROM ORAR_TEMPLATE WHERE TemplateID = :1`
	var o OrarTemplate
	var numeWOD sql.NullString
	err = db.QueryRow(query, id).Scan(&o.ID, &o.ZiuaSaptamanii, &o.Ora, &numeWOD, &o.AntrenorID, &o.TipAntrenamentID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Regula de orar nu a fost găsită", http.StatusNotFound)
			return
		}
		http.Error(w, "Eroare Scan", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if numeWOD.Valid {
		o.NumeWODTemplate = numeWOD.String
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(o); err != nil {
		log.Println("Eroare la encodare JSON orar:", err)
	}
}
func handlerAddOrar(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var o OrarTemplate
	if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		log.Println(err)
		return
	}
	query := `
		INSERT INTO ORAR_TEMPLATE (TemplateID, ZiuaSaptamanii, Ora, NumeWOD_Template, AntrenorID, TipAntrenamentID)
		VALUES (ORAR_TEMPLATE_SEQ.NEXTVAL, :1, :2, :3, :4, :5)
	`
	_, err := db.Exec(query, o.ZiuaSaptamanii, o.Ora, o.NumeWODTemplate, o.AntrenorID, o.TipAntrenamentID)
	if err != nil {
		http.Error(w, "Eroare la inserarea în BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(map[string]string{"mesaj": "Regulă de orar adăugată cu succes"}); err != nil {
		log.Println("Eroare la encodare JSON add orar:", err)
	}
}
func handlerUpdateOrar(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var o OrarTemplate
	if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		log.Println(err)
		return
	}
	query := `
		UPDATE ORAR_TEMPLATE 
		SET ZiuaSaptamanii = :1, Ora = :2, NumeWOD_Template = :3, AntrenorID = :4, TipAntrenamentID = :5
		WHERE TemplateID = :6
	`
	_, err := db.Exec(query, o.ZiuaSaptamanii, o.Ora, o.NumeWODTemplate, o.AntrenorID, o.TipAntrenamentID, o.ID)
	if err != nil {
		http.Error(w, "Eroare la actualizarea în BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"mesaj": "Regulă de orar actualizată cu succes"}); err != nil {
		log.Println("Eroare la encodare JSON update orar:", err)
	}
}
func handlerDeleteOrar(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var payload struct {
		ID int `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		log.Println(err)
		return
	}
	query := `DELETE FROM ORAR_TEMPLATE WHERE TemplateID = :1`
	_, err := db.Exec(query, payload.ID)
	if err != nil {
		http.Error(w, "Eroare la ștergerea din BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"mesaj": "Regulă de orar ștersă cu succes"}); err != nil {
		log.Println("Eroare la encodare JSON delete orar:", err)
	}
}
func handlerGenerateOrar(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	querySelect := `SELECT TemplateID, ZiuaSaptamanii, Ora, NumeWOD_Template, AntrenorID, TipAntrenamentID FROM ORAR_TEMPLATE`
	rows, err := db.Query(querySelect)
	if err != nil {
		http.Error(w, "Eroare la citirea regulilor de orar", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer rows.Close()
	var reguli []OrarTemplateScan
	for rows.Next() {
		var regula OrarTemplateScan
		if err := rows.Scan(&regula.TemplateID, &regula.ZiuaSaptamanii, &regula.Ora, &regula.NumeWODTemplate, &regula.AntrenorID, &regula.TipAntrenamentID); err != nil {
			http.Error(w, "Eroare la scanarea regulilor", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		reguli = append(reguli, regula)
	}
	today := time.Now()
	daysUntilNextMonday := (7 - int(today.Weekday()) + int(time.Monday)) % 7
	if daysUntilNextMonday == 0 {
		daysUntilNextMonday = 7
	}
	urmatorulLuni := today.AddDate(0, 0, daysUntilNextMonday)
	claseGenerate := 0
	tx, err := db.Begin()
	if err != nil {
		http.Error(w, "Eroare la pornirea tranzacției", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	queryInsert := `
		INSERT INTO CLASE (ClasaID, NumeWOD, DescriereWOD, DataOra, AntrenorID, TipAntrenamentID)
		VALUES (CLASE_SEQ.NEXTVAL, :1, :2, TO_TIMESTAMP(:3, 'YYYY-MM-DD"T"HH24:MI'), :4, :5)
	`
	for _, regula := range reguli {
		dataClasei := urmatorulLuni.AddDate(0, 0, regula.ZiuaSaptamanii-1)
		dataOraString := fmt.Sprintf("%sT%s", dataClasei.Format("2006-01-02"), regula.Ora)
		_, err := tx.Exec(queryInsert,
			regula.NumeWODTemplate,
			sql.NullString{},
			dataOraString,
			regula.AntrenorID,
			regula.TipAntrenamentID,
		)
		if err != nil {
			tx.Rollback()
			http.Error(w, "Eroare la inserarea unei clase (posibil duplicat?). Tranzacție anulată.", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		claseGenerate++
	}
	if err := tx.Commit(); err != nil {
		http.Error(w, "Eroare la salvarea tranzacției (commit)", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"mesaj":         "Orarul a fost generat cu succes!",
		"claseGenerate": claseGenerate,
	})
}

// ===================================================================
// HANDLERE PENTRU RAPOARTE
// ===================================================================
func handlerRaportAbonamente(w http.ResponseWriter, _ *http.Request) {
	query := `
		SELECT
			a.TipAbonament,
			COUNT(m.MembruID) AS Numar_Membri
		FROM
			ABONAMENTE a
		JOIN
			MEMBRI m ON a.AbonamentID = m.AbonamentID
		GROUP BY
			a.TipAbonament
		HAVING
			COUNT(m.MembruID) > 2
	`
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Eroare BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("Eroare la inchiderea rows:", err)
		}
	}()

	var raport []RaportAbonamente
	for rows.Next() {
		var r RaportAbonamente
		if err := rows.Scan(&r.TipAbonament, &r.NumarMembri); err != nil {
			http.Error(w, "Eroare Scan", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		raport = append(raport, r)
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(raport); err != nil {
		log.Println("Eroare la encodare JSON raport abonamente:", err)
	}
}
func handlerRaportVizualizare(w http.ResponseWriter, _ *http.Request) {
	// Preluăm din vizualizarea V_MEMBRI_ABONAMENTE pe care am creat-o în DataGrip
	query := `SELECT Nume, Prenume, Email, TipAbonament, Pret FROM V_MEMBRI_ABONAMENTE ORDER BY Nume`
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Eroare BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("Eroare la inchiderea rows:", err)
		}
	}()

	var raport []RaportVizualizare
	for rows.Next() {
		var r RaportVizualizare
		if err := rows.Scan(&r.Nume, &r.Prenume, &r.Email, &r.TipAbonament, &r.Pret); err != nil {
			http.Error(w, "Eroare Scan", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		raport = append(raport, r)
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(raport); err != nil {
		log.Println("Eroare la encodare JSON raport vizualizare:", err)
	}
}
func handlerRaportComplex(w http.ResponseWriter, _ *http.Request) {
	// Interogarea III.c
	query := `
		SELECT
			m.Nume AS Nume_Membru,
			m.Prenume AS Prenume_Membru,
			c.NumeWOD AS NumeClasa,
			a.Nume AS Nume_Antrenor
		FROM
			MEMBRI m
		JOIN
			ABONAMENTE ab ON m.AbonamentID = ab.AbonamentID
		JOIN
			INSCRIERI i ON m.MembruID = i.MembruID
		JOIN
			CLASE c ON i.ClasaID = c.ClasaID
		JOIN
			ANTRENORI a ON c.AntrenorID = a.AntrenorID
		WHERE
			ab.TipAbonament = 'Nelimitat 1 Luna'
			AND a.Nume = 'Popescu'
	`
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Eroare BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("Eroare la inchiderea rows:", err)
		}
	}()

	var raport []RaportComplex
	for rows.Next() {
		var r RaportComplex
		if err := rows.Scan(&r.NumeMembru, &r.PrenumeMembru, &r.NumeClasa, &r.NumeAntrenor); err != nil {
			http.Error(w, "Eroare Scan", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		raport = append(raport, r)
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(raport); err != nil {
		log.Println("Eroare la encodare JSON raport complex:", err)
	}
}
