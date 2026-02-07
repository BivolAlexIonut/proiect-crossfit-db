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

const dsn = "CROSSFIT_ADMIN/HTsjReTy12@localhost:1521/XE"

var db *sql.DB

type Membru struct {
	ID           int    `json:"id"`
	Nume         string `json:"nume"`
	Prenume      string `json:"prenume"`
	Email        string `json:"email"`
	AbonamentID  int    `json:"abonamentID"`
	TipAbonament string `json:"tipAbonament,omitempty"`
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
	ID             int     `json:"id"`
	MembruID       int     `json:"membruID"`
	ProdusID       int     `json:"produsID"`
	NumeMembru     string  `json:"numeMembru,omitempty"`
	NumeProdus     string  `json:"numeProdus,omitempty"`
	DataAchizitiei string  `json:"dataAchizitiei"`
	Cantitate      int     `json:"cantitate"`
	PretTotal      float64 `json:"pretTotal"`
}

type Mentorat struct {
	AntrenorID   int    `json:"antrenorID"`
	MembruID     int    `json:"membruID"`
	NumeAntrenor string `json:"numeAntrenor,omitempty"`
	NumeMembru   string `json:"numeMembru,omitempty"`
}

type Competitie struct {
	ID      int     `json:"id"`
	Nume    string  `json:"nume"`
	Data    string  `json:"data"`
	Locatie string  `json:"locatie"`
	Taxa    float64 `json:"taxa"`
}

type ParticipareCompetitie struct {
	CompetitieID   int    `json:"competitieID"`
	MembruID       int    `json:"membruID"`
	NumeCompetitie string `json:"numeCompetitie,omitempty"`
	NumeMembru     string `json:"numeMembru,omitempty"`
	LoculObtinut   int    `json:"loculObtinut"`
}

type OrarTemplate struct {
	ID               int    `json:"id"`
	ZiuaSaptamanii   int    `json:"ziuaSaptamanii"`
	Ora              string `json:"ora"`
	NumeWODTemplate  string `json:"numeWODTemplate"`
	AntrenorID       int    `json:"antrenorID"`
	TipAntrenamentID int    `json:"tipAntrenamentID"`
	NumeAntrenor     string `json:"numeAntrenor,omitempty"`
	NumeCategorie    string `json:"numeCategorie,omitempty"`
}

type OrarTemplateScan struct {
	TemplateID       int
	ZiuaSaptamanii   int
	Ora              string
	NumeWODTemplate  sql.NullString
	AntrenorID       sql.NullInt64
	TipAntrenamentID sql.NullInt64
}

// Structuri pentru Rapoarte
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

// =================================================================================
// MAIN FUNCTION - Configurare Server și Rute
// =================================================================================
func main() {
	fmt.Println("Inițializare conexiune Oracle Database...")
	var err error
	db, err = sql.Open("godror", dsn)
	if err != nil {
		log.Fatal("Eroare critică la sql.Open: ", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("Eroare critică la db.Ping (Database indisponibilă): ", err)
	}
	fmt.Println("Conexiune la baza de date reușită!")

	_, err = db.Exec(`
		CREATE OR REPLACE VIEW V_RAPORT_POPULARITATE AS
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
	`)
	if err != nil {
		log.Println("Avertisment: Nu s-a putut actualiza view-ul V_RAPORT_POPULARITATE:", err)
	} else {
		fmt.Println("Sistem: View-ul V_RAPORT_POPULARITATE a fost actualizat (HAVING > 2).")
	}

	// =============================================================================
	// DEFINIRE RUTE STATIC FILES (HTML/CSS/JS)
	// =============================================================================
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "html/index.html") })
	http.HandleFunc("/css/style.css", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "css/style.css") })
	http.HandleFunc("/js/script.js", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "js/script.js") })

	servePage := func(path string, file string) {
		http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "html/"+file) })
	}
	serveJS := func(path string, file string) {
		http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "js/"+file) })
	}

	servePage("/antrenori", "antrenori.html")
	serveJS("/js/antrenori.js", "antrenori.js")

	servePage("/abonamente", "abonamente.html")
	serveJS("/js/abonamente.js", "abonamente.js")

	servePage("/produse", "produse.html")
	serveJS("/js/produse.js", "produse.js")

	servePage("/echipamente", "echipamente.html")
	serveJS("/js/echipamente.js", "echipamente.js")

	servePage("/tipuri-antrenament", "tipuri-antrenament.html")
	serveJS("/js/tipuri-antrenament.js", "tipuri-antrenament.js")

	servePage("/clase", "clase.html")
	serveJS("/js/clase.js", "clase.js")

	servePage("/inscrieri", "inscrieri.html")
	serveJS("/js/inscrieri.js", "inscrieri.js")

	servePage("/orar", "orar.html")
	serveJS("/js/orar.js", "orar.js")

	servePage("/achizitii", "achizitii.html")
	serveJS("/js/achizitii.js", "achizitii.js")

	servePage("/mentorat", "mentorat.html")
	serveJS("/js/mentorat.js", "mentorat.js")

	servePage("/competitii", "competitii.html")
	serveJS("/js/competitii.js", "competitii.js")

	servePage("/rapoarte", "rapoarte.html")
	serveJS("/js/rapoarte.js", "rapoarte.js")

	// =============================================================================
	// DEFINIRE RUTE API (BACKEND)
	// =============================================================================

	// Membri
	http.HandleFunc("/api/membri", handlerGetMembri)
	http.HandleFunc("/api/membru", handlerGetUnMembru)
	http.HandleFunc("/api/membri/add", handlerAddMembru)
	http.HandleFunc("/api/membri/update", handlerUpdateMembru)
	http.HandleFunc("/api/membri/delete", handlerDeleteMembru)

	// Antrenori
	http.HandleFunc("/api/antrenori", handlerGetAntrenori)
	http.HandleFunc("/api/antrenor", handlerGetUnAntrenor)
	http.HandleFunc("/api/antrenori/add", handlerAddAntrenor)
	http.HandleFunc("/api/antrenori/update", handlerUpdateAntrenor)
	http.HandleFunc("/api/antrenori/delete", handlerDeleteAntrenor)

	// Abonamente
	http.HandleFunc("/api/abonamente", handlerGetAbonamente)
	http.HandleFunc("/api/abonament", handlerGetUnAbonament)
	http.HandleFunc("/api/abonamente/add", handlerAddAbonament)
	http.HandleFunc("/api/abonamente/update", handlerUpdateAbonament)
	http.HandleFunc("/api/abonamente/delete", handlerDeleteAbonament)

	// Produse
	http.HandleFunc("/api/produse", handlerGetProduse)
	http.HandleFunc("/api/produs", handlerGetUnProdus)
	http.HandleFunc("/api/produse/add", handlerAddProdus)
	http.HandleFunc("/api/produse/update", handlerUpdateProdus)
	http.HandleFunc("/api/produse/delete", handlerDeleteProdus)

	// Echipamente
	http.HandleFunc("/api/echipamente", handlerGetEchipamente)
	http.HandleFunc("/api/echipament", handlerGetUnEchipament)
	http.HandleFunc("/api/echipamente/add", handlerAddEchipament)
	http.HandleFunc("/api/echipamente/update", handlerUpdateEchipament)
	http.HandleFunc("/api/echipamente/delete", handlerDeleteEchipament)

	// Tipuri Antrenament
	http.HandleFunc("/api/tipuri-antrenament", handlerGetTipuri)
	http.HandleFunc("/api/tip-antrenament", handlerGetUnTip)
	http.HandleFunc("/api/tipuri-antrenament/add", handlerAddTip)
	http.HandleFunc("/api/tipuri-antrenament/update", handlerUpdateTip)
	http.HandleFunc("/api/tipuri-antrenament/delete", handlerDeleteTip)

	// Clase
	http.HandleFunc("/api/clase", handlerGetClase)
	http.HandleFunc("/api/clasa", handlerGetOClasa)
	http.HandleFunc("/api/clase/add", handlerAddClasa)
	http.HandleFunc("/api/clase/update", handlerUpdateClasa)
	http.HandleFunc("/api/clase/delete", handlerDeleteClasa)

	// Înscrieri
	http.HandleFunc("/api/inscrieri", handlerGetInscrieri)
	http.HandleFunc("/api/inscrieri/add", handlerAddInscriere)
	http.HandleFunc("/api/inscrieri/delete", handlerDeleteInscriere)

	// Orar
	http.HandleFunc("/api/orar", handlerGetOrar)
	http.HandleFunc("/api/orar/single", handlerGetUnOrar)
	http.HandleFunc("/api/orar/add", handlerAddOrar)
	http.HandleFunc("/api/orar/update", handlerUpdateOrar)
	http.HandleFunc("/api/orar/delete", handlerDeleteOrar)
	http.HandleFunc("/api/orar/generate", handlerGenerateOrar)

	// Achiziții
	http.HandleFunc("/api/achizitii", handlerGetAchizitii)
	http.HandleFunc("/api/achizitii/add", handlerAddAchizitie)
	http.HandleFunc("/api/achizitii/delete", handlerDeleteAchizitie)

	// Mentorat
	http.HandleFunc("/api/mentorat", handlerGetMentorat)
	http.HandleFunc("/api/mentorat/add", handlerAddMentorat)
	http.HandleFunc("/api/mentorat/delete", handlerDeleteMentorat)

	// Competiții
	http.HandleFunc("/api/competitii", handlerGetCompetitii)
	http.HandleFunc("/api/competitii/add", handlerAddCompetitie)
	http.HandleFunc("/api/competitii/delete", handlerDeleteCompetitie)
	http.HandleFunc("/api/competitii/participari", handlerGetParticipariCompetitie)
	http.HandleFunc("/api/competitii/participari/add", handlerAddParticipareCompetitie)

	// Rapoarte
	http.HandleFunc("/api/raport/abonamente", handlerRaportAbonamente)
	http.HandleFunc("/api/raport/vizualizare-membri", handlerRaportVizualizare)
	http.HandleFunc("/api/raport/complex-inscrieri", handlerRaportComplex)
	http.HandleFunc("/api/raport/update-view", handlerUpdateView)

	// --- Pornire Server ---
	port := ":8080"
	fmt.Println("Serverul web a pornit. Accesează http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}

// =================================================================================
// HANDLERE PENTRU MEMBRI
// =================================================================================

func handlerGetMembri(w http.ResponseWriter, _ *http.Request) {
	query := `
		SELECT m.MembruID, m.Nume, m.Prenume, m.Email, m.AbonamentID, a.TipAbonament 
		FROM MEMBRI m
		LEFT JOIN ABONAMENTE a ON m.AbonamentID = a.AbonamentID
		ORDER BY m.Nume
	`
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Eroare BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer rows.Close()

	membri := []Membru{} // Inițializare explicită pentru a returna [] în JSON dacă e gol
	for rows.Next() {
		var m Membru
		var tipAb sql.NullString
		if err := rows.Scan(&m.ID, &m.Nume, &m.Prenume, &m.Email, &m.AbonamentID, &tipAb); err != nil {
			http.Error(w, "Eroare Scan", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		if tipAb.Valid {
			m.TipAbonament = tipAb.String
		}
		membri = append(membri, m)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(membri)
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
	json.NewEncoder(w).Encode(m)
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
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Membru adăugat cu succes"})
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
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Membru actualizat cu succes"})
}

// Handler special pentru ștergerea în cascadă (Membru + Date Asociate)
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

	tx, err := db.Begin()
	if err != nil {
		http.Error(w, "Eroare internă (Transaction)", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Ștergem datele dependente
	steps := []string{
		"DELETE FROM INSCRIERI WHERE MembruID = :1",
		"DELETE FROM ACHIZITII WHERE MembruID = :1",
		"DELETE FROM MENTORAT WHERE MembruID = :1",
		"DELETE FROM PARTICIPARI_COMPETITIE WHERE MembruID = :1",
	}

	for _, query := range steps {
		if _, err := tx.Exec(query, payload.ID); err != nil {
			tx.Rollback()
			http.Error(w, "Eroare la ștergerea datelor asociate", http.StatusInternalServerError)
			log.Println("Eroare Query:", query, err)
			return
		}
	}

	// Ștergem membrul
	res, err := tx.Exec("DELETE FROM MEMBRI WHERE MembruID = :1", payload.ID)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Eroare la ștergerea membrului din BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		tx.Rollback()
		http.Error(w, "Membrul nu a fost găsit", http.StatusNotFound)
		return
	}

	if err := tx.Commit(); err != nil {
		http.Error(w, "Eroare la commit tranzacție", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Membru și datele asociate șterse cu succes"})
}

// =================================================================================
// HANDLERE PENTRU ANTRENORI
// =================================================================================

func handlerGetAntrenori(w http.ResponseWriter, _ *http.Request) {
	rows, err := db.Query("SELECT AntrenorID, Nume, Prenume, Specializare FROM ANTRENORI ORDER BY Nume")
	if err != nil {
		http.Error(w, "Eroare BD", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer rows.Close()
	var antrenori []Antrenor
	for rows.Next() {
		var a Antrenor
		if err := rows.Scan(&a.ID, &a.Nume, &a.Prenume, &a.Specializare); err != nil {
			http.Error(w, "Eroare Scan", http.StatusInternalServerError)
			return
		}
		antrenori = append(antrenori, a)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(antrenori)
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
		http.Error(w, "Antrenorul nu a fost găsit", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(a)
}

func handlerAddAntrenor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var a Antrenor
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		return
	}
	query := `INSERT INTO ANTRENORI (AntrenorID, Nume, Prenume, Specializare) VALUES (ANTRENORI_SEQ.NEXTVAL, :1, :2, :3)`
	_, err := db.Exec(query, a.Nume, a.Prenume, a.Specializare)
	if err != nil {
		http.Error(w, "Eroare la inserarea în BD", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Antrenor adăugat cu succes"})
}

func handlerUpdateAntrenor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var a Antrenor
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		return
	}
	query := `UPDATE ANTRENORI SET Nume = :1, Prenume = :2, Specializare = :3 WHERE AntrenorID = :4`
	_, err := db.Exec(query, a.Nume, a.Prenume, a.Specializare, a.ID)
	if err != nil {
		http.Error(w, "Eroare la actualizarea în BD", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Antrenor actualizat cu succes"})
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
		return
	}

	// Verificare integritate referențială
	var count int
	err := db.QueryRow("SELECT (SELECT COUNT(*) FROM CLASE WHERE AntrenorID = :1) + (SELECT COUNT(*) FROM MENTORAT WHERE AntrenorID = :1) FROM DUAL", payload.ID).Scan(&count)
	if err == nil && count > 0 {
		http.Error(w, "Antrenorul nu poate fi șters deoarece are clase sau mentorate asociate!", http.StatusConflict)
		return
	}

	query := `DELETE FROM ANTRENORI WHERE AntrenorID = :1`
	_, err = db.Exec(query, payload.ID)
	if err != nil {
		http.Error(w, "Eroare la ștergerea din BD", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Antrenor șters cu succes"})
}

// =================================================================================
// HANDLERE PENTRU ABONAMENTE
// =================================================================================

func handlerGetAbonamente(w http.ResponseWriter, _ *http.Request) {
	rows, err := db.Query("SELECT AbonamentID, TipAbonament, Pret FROM ABONAMENTE ORDER BY Pret")
	if err != nil {
		http.Error(w, "Eroare BD", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var abonamente []Abonament
	for rows.Next() {
		var a Abonament
		if err := rows.Scan(&a.ID, &a.Tip, &a.Pret); err != nil {
			http.Error(w, "Eroare Scan", http.StatusInternalServerError)
			return
		}
		abonamente = append(abonamente, a)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(abonamente)
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
		http.Error(w, "Abonamentul nu a fost găsit", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(a)
}

func handlerAddAbonament(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var a Abonament
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		return
	}
	query := `INSERT INTO ABONAMENTE (AbonamentID, TipAbonament, Pret) VALUES (ABONAMENTE_SEQ.NEXTVAL, :1, :2)`
	_, err := db.Exec(query, a.Tip, a.Pret)
	if err != nil {
		http.Error(w, "Eroare la inserarea în BD", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Abonament adăugat cu succes"})
}

func handlerUpdateAbonament(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var a Abonament
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		return
	}
	query := `UPDATE ABONAMENTE SET TipAbonament = :1, Pret = :2 WHERE AbonamentID = :3`
	_, err := db.Exec(query, a.Tip, a.Pret, a.ID)
	if err != nil {
		http.Error(w, "Eroare la actualizarea în BD", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Abonament actualizat cu succes"})
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
		return
	}
	query := `DELETE FROM ABONAMENTE WHERE AbonamentID = :1`
	_, err := db.Exec(query, payload.ID)
	if err != nil {
		http.Error(w, "Eroare la ștergere (Posibil utilizat de membri)", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Abonament șters cu succes"})
}

// =================================================================================
// HANDLERE PENTRU PRODUSE & ACHIZIȚII
// =================================================================================

func handlerGetProduse(w http.ResponseWriter, _ *http.Request) {
	rows, err := db.Query("SELECT ProdusID, NumeProdus, PretCurent, Stoc FROM PRODUSE ORDER BY NumeProdus")
	if err != nil {
		http.Error(w, "Eroare BD", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var produse []Produs
	for rows.Next() {
		var p Produs
		rows.Scan(&p.ID, &p.Nume, &p.Pret, &p.Stoc)
		produse = append(produse, p)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produse)
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
		http.Error(w, "Produsul nu a fost găsit", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func handlerAddProdus(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var p Produs
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		return
	}
	query := `INSERT INTO PRODUSE (ProdusID, NumeProdus, PretCurent, Stoc) VALUES (PRODUSE_SEQ.NEXTVAL, :1, :2, :3)`
	_, err := db.Exec(query, p.Nume, p.Pret, p.Stoc)
	if err != nil {
		http.Error(w, "Eroare la inserarea în BD", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Produs adăugat cu succes"})
}

func handlerUpdateProdus(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var p Produs
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		return
	}
	query := `UPDATE PRODUSE SET NumeProdus = :1, PretCurent = :2, Stoc = :3 WHERE ProdusID = :4`
	_, err := db.Exec(query, p.Nume, p.Pret, p.Stoc, p.ID)
	if err != nil {
		http.Error(w, "Eroare la actualizarea în BD", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Produs actualizat cu succes"})
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
		return
	}

	// Verificare dependențe (Achiziții)
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM ACHIZITII WHERE ProdusID = :1", payload.ID).Scan(&count)
	if err == nil && count > 0 {
		http.Error(w, "Produsul nu poate fi șters deoarece există achiziții în istoric!", http.StatusConflict)
		return
	}

	query := `DELETE FROM PRODUSE WHERE ProdusID = :1`
	_, err = db.Exec(query, payload.ID)
	if err != nil {
		http.Error(w, "Eroare la ștergerea din BD", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Produs șters cu succes"})
}

func handlerGetAchizitii(w http.ResponseWriter, _ *http.Request) {
	query := `
		SELECT 
			a.AchizitieID, a.MembruID, a.ProdusID,
			m.Nume || ' ' || m.Prenume AS NumeMembru,
			p.NumeProdus, 
			TO_CHAR(a.DataAchizitiei, 'YYYY-MM-DD') AS DataAchizitiei,
			a.Cantitate,
			(a.Cantitate * p.PretCurent) AS PretTotal
		FROM ACHIZITII a
		JOIN MEMBRI m ON a.MembruID = m.MembruID
		JOIN PRODUSE p ON a.ProdusID = p.ProdusID
		ORDER BY a.DataAchizitiei DESC
	`
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Eroare BD", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var achizitii []Achizitie
	for rows.Next() {
		var a Achizitie
		rows.Scan(&a.ID, &a.MembruID, &a.ProdusID, &a.NumeMembru, &a.NumeProdus, &a.DataAchizitiei, &a.Cantitate, &a.PretTotal)
		achizitii = append(achizitii, a)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(achizitii)
}

// Handler tranzacțional pentru achiziții (Gestionează stocul)
func handlerAddAchizitie(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var a Achizitie
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		return
	}
	if a.Cantitate <= 0 {
		a.Cantitate = 1
	}

	tx, err := db.Begin()
	if err != nil {
		http.Error(w, "Eroare internă", http.StatusInternalServerError)
		return
	}

	// Update stoc condiționat (parametrul cantitate trimis de 2 ori: pentru scădere și verificare)
	res, err := tx.Exec("UPDATE PRODUSE SET Stoc = Stoc - :1 WHERE ProdusID = :2 AND Stoc >= :3", a.Cantitate, a.ProdusID, a.Cantitate)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Eroare la actualizarea stocului", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		tx.Rollback()
		http.Error(w, "Stoc insuficient pentru acest produs!", http.StatusConflict)
		return
	}

	query := `INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, Cantitate) VALUES (ACHIZITII_SEQ.NEXTVAL, :1, :2, :3)`
	_, err = tx.Exec(query, a.MembruID, a.ProdusID, a.Cantitate)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Eroare la inserarea achiziției", http.StatusInternalServerError)
		return
	}

	if err := tx.Commit(); err != nil {
		http.Error(w, "Eroare la finalizarea tranzacției", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Achiziție adăugată cu succes. Stoc actualizat."})
}

// Handler tranzacțional pentru ștergere achiziție (Restituie stocul)
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
		return
	}

	tx, err := db.Begin()
	if err != nil {
		http.Error(w, "Eroare internă", http.StatusInternalServerError)
		return
	}

	var produsID, cantitate int
	queryGet := `SELECT ProdusID, Cantitate FROM ACHIZITII WHERE AchizitieID = :1`
	err = tx.QueryRow(queryGet, payload.ID).Scan(&produsID, &cantitate)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Achiziția nu a fost găsită", http.StatusNotFound)
		return
	}

	// Restituire stoc
	_, err = tx.Exec(`UPDATE PRODUSE SET Stoc = Stoc + :1 WHERE ProdusID = :2`, cantitate, produsID)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Eroare la restituirea stocului", http.StatusInternalServerError)
		return
	}

	_, err = tx.Exec(`DELETE FROM ACHIZITII WHERE AchizitieID = :1`, payload.ID)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Eroare la ștergerea din BD", http.StatusInternalServerError)
		return
	}

	if err := tx.Commit(); err != nil {
		http.Error(w, "Eroare la commit", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Achiziție ștearsă și stoc restituit cu succes"})
}

// =================================================================================
// HANDLERE PENTRU ECHIPAMENTE
// =================================================================================

func handlerGetEchipamente(w http.ResponseWriter, _ *http.Request) {
	rows, err := db.Query("SELECT EchipamentID, NumeEchipament, CantitateTotala FROM ECHIPAMENTE ORDER BY NumeEchipament")
	if err != nil {
		http.Error(w, "Eroare BD", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var echipamente []Echipament
	for rows.Next() {
		var e Echipament
		rows.Scan(&e.ID, &e.Nume, &e.Cantitate)
		echipamente = append(echipamente, e)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(echipamente)
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
		http.Error(w, "Echipamentul nu a fost găsit", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(e)
}

func handlerAddEchipament(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var e Echipament
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		return
	}
	query := `INSERT INTO ECHIPAMENTE (EchipamentID, NumeEchipament, CantitateTotala) VALUES (ECHIPAMENTE_SEQ.NEXTVAL, :1, :2)`
	_, err := db.Exec(query, e.Nume, e.Cantitate)
	if err != nil {
		http.Error(w, "Eroare la inserarea în BD", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Echipament adăugat cu succes"})
}

func handlerUpdateEchipament(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var e Echipament
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		return
	}
	query := `UPDATE ECHIPAMENTE SET NumeEchipament = :1, CantitateTotala = :2 WHERE EchipamentID = :3`
	_, err := db.Exec(query, e.Nume, e.Cantitate, e.ID)
	if err != nil {
		http.Error(w, "Eroare la actualizarea în BD", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Echipament actualizat cu succes"})
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
		return
	}
	query := `DELETE FROM ECHIPAMENTE WHERE EchipamentID = :1`
	_, err := db.Exec(query, payload.ID)
	if err != nil {
		http.Error(w, "Eroare la ștergerea din BD", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Echipament șters cu succes"})
}

// =================================================================================
// HANDLERE PENTRU TIPURI ANTRENAMENT
// =================================================================================

func handlerGetTipuri(w http.ResponseWriter, _ *http.Request) {
	rows, err := db.Query("SELECT TipAntrenamentID, NumeWOD, Descriere FROM TIPURI_ANTRENAMENT ORDER BY NumeWOD")
	if err != nil {
		http.Error(w, "Eroare BD", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var tipuri []TipAntrenament
	for rows.Next() {
		var t TipAntrenament
		var descriere sql.NullString
		rows.Scan(&t.ID, &t.Nume, &descriere)
		if descriere.Valid {
			t.Descriere = descriere.String
		}
		tipuri = append(tipuri, t)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tipuri)
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
		http.Error(w, "Nu a fost găsit", http.StatusNotFound)
		return
	}
	if descriere.Valid {
		t.Descriere = descriere.String
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)
}

func handlerAddTip(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var t TipAntrenament
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		return
	}
	query := `INSERT INTO TIPURI_ANTRENAMENT (TipAntrenamentID, NumeWOD, Descriere) VALUES (TIPURI_ANTRENAMENT_SEQ.NEXTVAL, :1, :2)`
	_, err := db.Exec(query, t.Nume, t.Descriere)
	if err != nil {
		http.Error(w, "Eroare la inserarea în BD", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Tip de antrenament adăugat cu succes"})
}

func handlerUpdateTip(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var t TipAntrenament
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		return
	}
	query := `UPDATE TIPURI_ANTRENAMENT SET NumeWOD = :1, Descriere = :2 WHERE TipAntrenamentID = :3`
	_, err := db.Exec(query, t.Nume, t.Descriere, t.ID)
	if err != nil {
		http.Error(w, "Eroare la actualizarea în BD", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Tip de antrenament actualizat cu succes"})
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
		return
	}
	query := `DELETE FROM TIPURI_ANTRENAMENT WHERE TipAntrenamentID = :1`
	_, err := db.Exec(query, payload.ID)
	if err != nil {
		http.Error(w, "Eroare la ștergerea din BD", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Tip de antrenament șters cu succes"})
}

// =================================================================================
// HANDLERE PENTRU CLASE & ORAR
// =================================================================================

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
		return
	}
	defer rows.Close()

	var clase []Clasa
	for rows.Next() {
		var c Clasa
		var numeAntrenor, numeCategorie, descriereWOD sql.NullString
		rows.Scan(&c.ID, &c.NumeWOD, &descriereWOD, &c.DataOra, &c.AntrenorID, &numeAntrenor, &c.TipAntrenamentID, &numeCategorie)

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
	json.NewEncoder(w).Encode(clase)
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
		http.Error(w, "Clasa nu a fost găsită", http.StatusNotFound)
		return
	}
	if descriereWOD.Valid {
		c.DescriereWOD = descriereWOD.String
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(c)
}

func handlerAddClasa(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var c Clasa
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		return
	}
	query := `INSERT INTO CLASE (ClasaID, NumeWOD, DescriereWOD, DataOra, AntrenorID, TipAntrenamentID)
			  VALUES (CLASE_SEQ.NEXTVAL, :1, :2, TO_TIMESTAMP(:3, 'YYYY-MM-DD"T"HH24:MI'), :4, :5)`
	_, err := db.Exec(query, c.NumeWOD, c.DescriereWOD, c.DataOra, c.AntrenorID, c.TipAntrenamentID)
	if err != nil {
		http.Error(w, "Eroare la inserarea în BD", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Clasă adăugată cu succes"})
}

func handlerUpdateClasa(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var c Clasa
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		return
	}
	query := `UPDATE CLASE 
			  SET NumeWOD = :1, DescriereWOD = :2, DataOra = TO_TIMESTAMP(:3, 'YYYY-MM-DD"T"HH24:MI'), 
				  AntrenorID = :4, TipAntrenamentID = :5
			  WHERE ClasaID = :6`
	_, err := db.Exec(query, c.NumeWOD, c.DescriereWOD, c.DataOra, c.AntrenorID, c.TipAntrenamentID, c.ID)
	if err != nil {
		http.Error(w, "Eroare la actualizarea în BD", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Clasă actualizată cu succes"})
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
		return
	}
	query := `DELETE FROM CLASE WHERE ClasaID = :1`
	_, err := db.Exec(query, payload.ID)
	if err != nil {
		http.Error(w, "Eroare la ștergerea din BD", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Clasă ștersă cu succes"})
}

// =================================================================================
// HANDLERE PENTRU ÎNSCRIERI
// =================================================================================

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
		return
	}
	defer rows.Close()

	var inscrieri []Inscriere
	for rows.Next() {
		var i Inscriere
		var numeAntrenor, dataOra sql.NullString
		rows.Scan(&i.MembruID, &i.ClasaID, &i.NumeMembru, &i.NumeWOD, &dataOra, &numeAntrenor)
		if numeAntrenor.Valid {
			i.NumeAntrenor = numeAntrenor.String
		}
		if dataOra.Valid {
			i.DataOra = dataOra.String
		}
		inscrieri = append(inscrieri, i)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inscrieri)
}

func handlerAddInscriere(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var i Inscriere
	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		return
	}

	// Verificare limită abonament
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
				WHEN a.TipAbonament LIKE '%Nelimitat%' OR a.TipAbonament LIKE '%Full Time%' THEN 1
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
		http.Error(w, "Membrul nu a fost găsit sau eroare abonament", http.StatusBadRequest)
		return
	}
	if check.PermiteInscriere == 0 {
		http.Error(w, fmt.Sprintf("Limita de ședințe atinsă! (%s: %d utilizate)", check.TipAbonament, check.InscrieriNumar), http.StatusForbidden)
		return
	}

	queryInsert := `INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (:1, :2)`
	_, err = db.Exec(queryInsert, i.MembruID, i.ClasaID)
	if err != nil {
		http.Error(w, "Eroare la înscriere (deja înscris?)", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Înscriere adăugată cu succes"})
}

func handlerDeleteInscriere(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var i Inscriere
	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		return
	}
	query := `DELETE FROM INSCRIERI WHERE MembruID = :1 AND ClasaID = :2`
	_, err := db.Exec(query, i.MembruID, i.ClasaID)
	if err != nil {
		http.Error(w, "Eroare la anularea înscrierii", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Înscriere anulată cu succes"})
}

// =================================================================================
// HANDLERE PENTRU ORAR (TEMPLATE & GENERARE)
// =================================================================================

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
		return
	}
	defer rows.Close()

	var orar []OrarTemplate
	for rows.Next() {
		var o OrarTemplate
		var numeAntrenor, numeCategorie, numeWOD sql.NullString
		rows.Scan(&o.ID, &o.ZiuaSaptamanii, &o.Ora, &numeWOD, &o.AntrenorID, &numeAntrenor, &o.TipAntrenamentID, &numeCategorie)
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
	json.NewEncoder(w).Encode(orar)
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
		http.Error(w, "Nu a fost găsit", http.StatusNotFound)
		return
	}
	if numeWOD.Valid {
		o.NumeWODTemplate = numeWOD.String
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(o)
}

func handlerAddOrar(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var o OrarTemplate
	if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		return
	}
	query := `INSERT INTO ORAR_TEMPLATE (TemplateID, ZiuaSaptamanii, Ora, NumeWOD_Template, AntrenorID, TipAntrenamentID)
			  VALUES (ORAR_TEMPLATE_SEQ.NEXTVAL, :1, :2, :3, :4, :5)`
	_, err := db.Exec(query, o.ZiuaSaptamanii, o.Ora, o.NumeWODTemplate, o.AntrenorID, o.TipAntrenamentID)
	if err != nil {
		http.Error(w, "Eroare la inserarea în BD", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Regulă de orar adăugată cu succes"})
}

func handlerUpdateOrar(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var o OrarTemplate
	if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		return
	}
	query := `UPDATE ORAR_TEMPLATE 
			  SET ZiuaSaptamanii = :1, Ora = :2, NumeWOD_Template = :3, AntrenorID = :4, TipAntrenamentID = :5
			  WHERE TemplateID = :6`
	_, err := db.Exec(query, o.ZiuaSaptamanii, o.Ora, o.NumeWODTemplate, o.AntrenorID, o.TipAntrenamentID, o.ID)
	if err != nil {
		http.Error(w, "Eroare la actualizarea în BD", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Regulă de orar actualizată cu succes"})
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
		return
	}
	query := `DELETE FROM ORAR_TEMPLATE WHERE TemplateID = :1`
	_, err := db.Exec(query, payload.ID)
	if err != nil {
		http.Error(w, "Eroare la ștergerea din BD", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Regulă de orar ștearsă cu succes"})
}

func handlerGenerateOrar(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	rows, err := db.Query(`SELECT TemplateID, ZiuaSaptamanii, Ora, NumeWOD_Template, AntrenorID, TipAntrenamentID FROM ORAR_TEMPLATE`)
	if err != nil {
		http.Error(w, "Eroare la citirea regulilor", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var reguli []OrarTemplateScan
	for rows.Next() {
		var r OrarTemplateScan
		rows.Scan(&r.TemplateID, &r.ZiuaSaptamanii, &r.Ora, &r.NumeWODTemplate, &r.AntrenorID, &r.TipAntrenamentID)
		reguli = append(reguli, r)
	}

	today := time.Now()
	// Calculăm data următoarei zile de Luni
	daysUntilNextMonday := (7 - int(today.Weekday()) + int(time.Monday)) % 7
	if daysUntilNextMonday == 0 {
		daysUntilNextMonday = 7
	}
	urmatorulLuni := today.AddDate(0, 0, daysUntilNextMonday)

	claseGenerate := 0
	tx, err := db.Begin()
	if err != nil {
		http.Error(w, "Eroare internă", http.StatusInternalServerError)
		return
	}

	queryInsert := `INSERT INTO CLASE (ClasaID, NumeWOD, DescriereWOD, DataOra, AntrenorID, TipAntrenamentID)
					VALUES (CLASE_SEQ.NEXTVAL, :1, :2, TO_TIMESTAMP(:3, 'YYYY-MM-DD"T"HH24:MI'), :4, :5)`

	for _, regula := range reguli {
		dataClasei := urmatorulLuni.AddDate(0, 0, regula.ZiuaSaptamanii-1)
		dataOraString := fmt.Sprintf("%sT%s", dataClasei.Format("2006-01-02"), regula.Ora)

		_, err := tx.Exec(queryInsert, regula.NumeWODTemplate, sql.NullString{}, dataOraString, regula.AntrenorID, regula.TipAntrenamentID)
		if err != nil {
			tx.Rollback()
			http.Error(w, "Eroare la generare (posibil duplicat)", http.StatusInternalServerError)
			return
		}
		claseGenerate++
	}

	if err := tx.Commit(); err != nil {
		http.Error(w, "Eroare la commit", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"mesaj": "Orarul a fost generat!", "claseGenerate": claseGenerate})
}

// =================================================================================
// HANDLERE PENTRU MENTORAT
// =================================================================================

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
		return
	}
	defer rows.Close()

	var mentorate []Mentorat
	for rows.Next() {
		var m Mentorat
		rows.Scan(&m.AntrenorID, &m.MembruID, &m.NumeAntrenor, &m.NumeMembru)
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
		return
	}
	query := `INSERT INTO MENTORAT (AntrenorID, MembruID) VALUES (:1, :2)`
	_, err := db.Exec(query, m.AntrenorID, m.MembruID)
	if err != nil {
		http.Error(w, "Eroare la inserarea în BD", http.StatusInternalServerError)
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
		return
	}
	query := `DELETE FROM MENTORAT WHERE AntrenorID = :1 AND MembruID = :2`
	_, err := db.Exec(query, m.AntrenorID, m.MembruID)
	if err != nil {
		http.Error(w, "Eroare la ștergerea din BD", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Mentorat șters cu succes"})
}

// =================================================================================
// HANDLERE PENTRU COMPETIȚII
// =================================================================================

func handlerGetCompetitii(w http.ResponseWriter, _ *http.Request) {
	rows, err := db.Query("SELECT CompetitieID, Nume, TO_CHAR(Data, 'YYYY-MM-DD'), Locatie, Taxa FROM COMPETITII ORDER BY Data DESC")
	if err != nil {
		http.Error(w, "Eroare BD", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var competitii []Competitie
	for rows.Next() {
		var c Competitie
		rows.Scan(&c.ID, &c.Nume, &c.Data, &c.Locatie, &c.Taxa)
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
		return
	}
	query := `INSERT INTO COMPETITII (CompetitieID, Nume, Data, Locatie, Taxa) VALUES (COMPETITII_SEQ.NEXTVAL, :1, TO_DATE(:2, 'YYYY-MM-DD'), :3, :4)`
	_, err := db.Exec(query, c.Nume, c.Data, c.Locatie, c.Taxa)
	if err != nil {
		http.Error(w, "Eroare la inserarea în BD", http.StatusInternalServerError)
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
		return
	}
	// Ștergere în cascadă manuală (participări)
	db.Exec("DELETE FROM PARTICIPARI_COMPETITIE WHERE CompetitieID = :1", payload.ID)

	query := `DELETE FROM COMPETITII WHERE CompetitieID = :1`
	_, err := db.Exec(query, payload.ID)
	if err != nil {
		http.Error(w, "Eroare la ștergerea din BD", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Competiție ștearsă cu succes"})
}

func handlerGetParticipariCompetitie(w http.ResponseWriter, r *http.Request) {
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
		return
	}
	defer rows.Close()
	var participari []ParticipareCompetitie
	for rows.Next() {
		var p ParticipareCompetitie
		var loc sql.NullInt64
		rows.Scan(&p.CompetitieID, &p.MembruID, &p.NumeCompetitie, &p.NumeMembru, &loc)
		if loc.Valid {
			p.LoculObtinut = int(loc.Int64)
		}
		participari = append(participari, p)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(participari)
}

func handlerAddParticipareCompetitie(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var p ParticipareCompetitie
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		return
	}
	query := `INSERT INTO PARTICIPARI_COMPETITIE (CompetitieID, MembruID) VALUES (:1, :2)`
	_, err := db.Exec(query, p.CompetitieID, p.MembruID)
	if err != nil {
		http.Error(w, "Eroare la înscriere (deja înscris?)", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"mesaj": "Membru înscris la competiție!"})
}

// =================================================================================
// HANDLERE PENTRU RAPOARTE
// =================================================================================

func handlerRaportAbonamente(w http.ResponseWriter, _ *http.Request) {
	query := `SELECT TipAbonament, Numar_Membri FROM V_RAPORT_POPULARITATE`
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Eroare BD", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var raport []RaportAbonamente
	for rows.Next() {
		var r RaportAbonamente
		rows.Scan(&r.TipAbonament, &r.NumarMembri)
		raport = append(raport, r)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(raport)
}

func handlerRaportVizualizare(w http.ResponseWriter, _ *http.Request) {
	query := `SELECT Nume, Prenume, Email, TipAbonament, Pret FROM V_MEMBRI_ABONAMENTE ORDER BY Nume`
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Eroare BD", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var raport []RaportVizualizare
	for rows.Next() {
		var r RaportVizualizare
		rows.Scan(&r.Nume, &r.Prenume, &r.Email, &r.TipAbonament, &r.Pret)
		raport = append(raport, r)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(raport)
}

func handlerRaportComplex(w http.ResponseWriter, _ *http.Request) {
	// Interogare Complexă: JOIN pe 5 tabele + Filtrare flexibilă
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
			(UPPER(ab.TipAbonament) LIKE '%NELIMITAT%' OR UPPER(ab.TipAbonament) LIKE '%FULL%')
			AND UPPER(a.Nume) LIKE '%POPESCU%'
	`
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Eroare BD", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var raport []RaportComplex
	for rows.Next() {
		var r RaportComplex
		rows.Scan(&r.NumeMembru, &r.PrenumeMembru, &r.NumeClasa, &r.NumeAntrenor)
		raport = append(raport, r)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(raport)
}

func handlerUpdateView(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metoda nu este permisă", http.StatusMethodNotAllowed)
		return
	}
	var payload struct {
		OldEmail string `json:"oldEmail"`
		NewEmail string `json:"newEmail"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "JSON invalid", http.StatusBadRequest)
		return
	}

	query := "UPDATE V_MEMBRI_ABONAMENTE SET Email = :1 WHERE Email = :2"
	result, err := db.Exec(query, payload.NewEmail, payload.OldEmail)
	if err != nil {
		http.Error(w, "Eroare la actualizarea prin VIEW", http.StatusInternalServerError)
		log.Println("Eroare Update View:", err)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"mesaj":        "Email actualizat cu succes prin VIEW!",
		"rowsAffected": rowsAffected,
	})
}
