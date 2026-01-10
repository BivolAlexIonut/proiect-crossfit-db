-- =============================================================
-- SCRIPT FINAL DE ACTUALIZARE ȘI SECURIZARE A BAZEI DE DATE
-- Rulează acest script în DataGrip / SQL Developer
-- =============================================================

-- 1. ANTRENORI: Validare Specializare (Listă Fixă)
-- Mai întâi corectăm datele care nu respectă lista (le punem default 'Gimnastică')
UPDATE ANTRENORI 
SET Specializare = 'Gimnastică' 
WHERE Specializare NOT IN ('Gimnastică', 'Haltere', 'Crossfit L1', 'Crossfit L2', 'Mobilitate', 'Cardio');

-- Adăugăm constrângerea
ALTER TABLE ANTRENORI 
ADD CONSTRAINT CHK_SPECIALIZARE 
CHECK (SPECIALIZARE IN ('Gimnastică', 'Haltere', 'Crossfit L1', 'Crossfit L2', 'Mobilitate', 'Cardio'));

-- 2. ABONAMENTE: Preț Pozitiv
-- Corectăm eventualele prețuri negative
UPDATE ABONAMENTE SET Pret = 0 WHERE Pret < 0;

-- Adăugăm constrângerea
ALTER TABLE ABONAMENTE ADD CONSTRAINT CHK_ABONAMENT_PRET CHECK (PRET >= 0);

-- 3. PRODUSE: Preț și Stoc Pozitiv
UPDATE PRODUSE SET PretCurent = 0 WHERE PretCurent < 0;
UPDATE PRODUSE SET Stoc = 0 WHERE Stoc < 0;

ALTER TABLE PRODUSE ADD CONSTRAINT CHK_PRODUS_PRET CHECK (PRETCURENT >= 0);
ALTER TABLE PRODUSE ADD CONSTRAINT CHK_PRODUS_STOC CHECK (STOC >= 0);

-- 4. ECHIPAMENTE: Cantitate Pozitivă
UPDATE ECHIPAMENTE SET CantitateTotala = 0 WHERE CantitateTotala < 0;

ALTER TABLE ECHIPAMENTE ADD CONSTRAINT CHK_ECHIPAMENT_CANT CHECK (CANTITATETOTALA >= 0);

-- 5. MEMBRI: Validare Email (să conțină @)
-- Ștergem sau marcăm email-urile invalide (aici punem un placeholder dacă e invalid)
UPDATE MEMBRI SET Email = 'invalid_' || MEMBRUID || '@temp.com' WHERE Email NOT LIKE '%@%';

ALTER TABLE MEMBRI ADD CONSTRAINT CHK_MEMBRU_EMAIL CHECK (EMAIL LIKE '%@%');

-- 6. VIZUALIZARE COMPLEXĂ (pentru Raport)
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
    COUNT(m.MembruID) > 2;

-- 7. VIZUALIZARE PENTRU UPDATE (DML pe View)
CREATE OR REPLACE VIEW V_MEMBRI_ABONAMENTE as
SELECT
    m.Nume,
    m.Prenume,
    m.Email,
    a.TipAbonament,
    a.Pret
FROM
    MEMBRI m
        JOIN
    ABONAMENTE a ON m.AbonamentID = a.AbonamentID;

-- Confirmăm modificările
COMMIT;

-- Mesaj final (doar pentru informare, nu e SQL executabil)
-- Baza de date a fost actualizată cu succes!
