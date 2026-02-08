
DELETE FROM ACHIZITII;
DELETE FROM PARTICIPARI_COMPETITIE;
DELETE FROM INSCRIERI;
DELETE FROM MENTORAT;
DELETE FROM NECESAR_ECHIPAMENT;
DELETE FROM CLASE;
DELETE FROM COMPETITII;
DELETE FROM MEMBRI;
DELETE FROM PRODUSE;
DELETE FROM ECHIPAMENTE;
DELETE FROM TIPURI_ANTRENAMENT;
DELETE FROM ANTRENORI;
DELETE FROM ABONAMENTE;

--1. POPULARE ABONAMENTE
INSERT INTO ABONAMENTE (AbonamentID, TipAbonament, Pret) VALUES (1, 'Standard Crossfit', 250);
INSERT INTO ABONAMENTE (AbonamentID, TipAbonament, Pret) VALUES (2, 'Student Crossfit', 150);
INSERT INTO ABONAMENTE (AbonamentID, TipAbonament, Pret) VALUES (3, 'Open Gym', 200);
INSERT INTO ABONAMENTE (AbonamentID, TipAbonament, Pret) VALUES (4, 'Premium All Access', 400);
INSERT INTO ABONAMENTE (AbonamentID, TipAbonament, Pret) VALUES (5, 'Full Time', 350);
INSERT INTO ABONAMENTE (AbonamentID, TipAbonament, Pret) VALUES (6, 'Weekend Only', 120);

--2. POPULARE ANTRENORI
INSERT INTO ANTRENORI (AntrenorID, Nume, Prenume, Specializare) VALUES (1, 'Popescu', 'Ion', 'Crossfit L2');
INSERT INTO ANTRENORI (AntrenorID, Nume, Prenume, Specializare) VALUES (2, 'Ionescu', 'Maria', 'Gimnastica');
INSERT INTO ANTRENORI (AntrenorID, Nume, Prenume, Specializare) VALUES (3, 'Vasilescu', 'Andrei', 'Haltere');
INSERT INTO ANTRENORI (AntrenorID, Nume, Prenume, Specializare) VALUES (4, 'Radu', 'Elena', 'Cardio');
INSERT INTO ANTRENORI (AntrenorID, Nume, Prenume, Specializare) VALUES (5, 'Mihai', 'George', 'Mobilitate');

-- 3. POPULARE MEMBRI
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (1, 'Dumitrescu', 'Alex', 'alex.d@email.com', 1);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (2, 'Stan', 'Ana', 'ana.stan@email.com', 2);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (3, 'Popa', 'Vlad', 'vlad.popa@email.com', 4);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (4, 'Gheorghe', 'Ioana', 'ioana.gh@email.com', 1);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (5, 'Matei', 'Cristian', 'cristi.m@email.com', 3);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (6, 'Dobre', 'Raluca', 'raluca.d@email.com', 5);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (7, 'Enache', 'Mihai', 'mihai.e@email.com', 2);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (8, 'Diaconu', 'Elena', 'elena.dia@email.com', 4);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (9, 'Stoica', 'Bogdan', 'bogdan.s@email.com', 1);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (10, 'Florea', 'Andreea', 'andreea.f@email.com', 6);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (11, 'Manole', 'Dan', 'dan.manole@email.com', 1);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (12, 'Voinea', 'Carmen', 'carmen.v@email.com', 2);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (13, 'Nistor', 'Adrian', 'adrian.n@email.com', 3);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (14, 'Marin', 'Gabriela', 'gabi.m@email.com', 4);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (15, 'Toma', 'Lucian', 'lucian.t@email.com', 5);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (16, 'Serban', 'Roxana', 'roxana.s@email.com', 1);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (17, 'Lupu', 'Victor', 'victor.l@email.com', 2);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (18, 'Mocanu', 'Silvia', 'silvia.m@email.com', 6);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (19, 'Oprea', 'Florin', 'florin.o@email.com', 3);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (20, 'Sandu', 'Irina', 'irina.s@email.com', 4);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (21, 'Barbu', 'Costin', 'costin.b@email.com', 1);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (22, 'Constantin', 'Diana', 'diana.c@email.com', 5);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (23, 'Neagu', 'Paul', 'paul.n@email.com', 2);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (24, 'Grigore', 'Simona', 'simona.g@email.com', 3);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (25, 'Dumitru', 'Razvan', 'razvan.d@email.com', 4);

-- 4. POPULARE MENTORAT
INSERT INTO MENTORAT (AntrenorID, MembruID) VALUES (1, 1);
INSERT INTO MENTORAT (AntrenorID, MembruID) VALUES (2, 2);
INSERT INTO MENTORAT (AntrenorID, MembruID) VALUES (3, 3);
INSERT INTO MENTORAT (AntrenorID, MembruID) VALUES (1, 4);
INSERT INTO MENTORAT (AntrenorID, MembruID) VALUES (4, 5);
INSERT INTO MENTORAT (AntrenorID, MembruID) VALUES (5, 6);
INSERT INTO MENTORAT (AntrenorID, MembruID) VALUES (2, 7);
INSERT INTO MENTORAT (AntrenorID, MembruID) VALUES (3, 8);
INSERT INTO MENTORAT (AntrenorID, MembruID) VALUES (1, 9);
INSERT INTO MENTORAT (AntrenorID, MembruID) VALUES (4, 10);

-- 5. POPULARE TIPURI_ANTRENAMENT
INSERT INTO TIPURI_ANTRENAMENT (TipID, NumeWOD, Descriere) VALUES (1, 'Murph', 'Alergare 1.6km, 100 Tractiuni, 200 Flotari, 300 Genuflexiuni, Alergare 1.6km. Cu vesta 10kg.');
INSERT INTO TIPURI_ANTRENAMENT (TipID, NumeWOD, Descriere) VALUES (2, 'Fran', '21-15-9 repetari de Thrusters (43kg) si Tractiuni.');
INSERT INTO TIPURI_ANTRENAMENT (TipID, NumeWOD, Descriere) VALUES (3, 'Grace', '30 repetari Clean and Jerk (61kg) contra cronometru.');
INSERT INTO TIPURI_ANTRENAMENT (TipID, NumeWOD, Descriere) VALUES (4, 'Cindy', 'AMRAP 20 min: 5 Tractiuni, 10 Flotari, 15 Genuflexiuni.');
INSERT INTO TIPURI_ANTRENAMENT (TipID, NumeWOD, Descriere) VALUES (5, 'Helen', '3 runde: Alergare 400m, 21 Kettlebell Swings, 12 Tractiuni.');
INSERT INTO TIPURI_ANTRENAMENT (TipID, NumeWOD, Descriere) VALUES (6, 'Linda', '10-9-8-7-6-5-4-3-2-1 repetari de Deadlift, Bench Press, Clean.');

-- 6. POPULARE CLASE
INSERT INTO CLASE (ClasaID, AntrenorID, TipID, DataOra) VALUES (1, 1, 1, TO_DATE('2024-02-10 08:00', 'YYYY-MM-DD HH24:MI'));
INSERT INTO CLASE (ClasaID, AntrenorID, TipID, DataOra) VALUES (2, 2, 2, TO_DATE('2024-02-10 09:30', 'YYYY-MM-DD HH24:MI'));
INSERT INTO CLASE (ClasaID, AntrenorID, TipID, DataOra) VALUES (3, 3, 3, TO_DATE('2024-02-10 11:00', 'YYYY-MM-DD HH24:MI'));
INSERT INTO CLASE (ClasaID, AntrenorID, TipID, DataOra) VALUES (4, 4, 4, TO_DATE('2024-02-10 17:00', 'YYYY-MM-DD HH24:MI'));
INSERT INTO CLASE (ClasaID, AntrenorID, TipID, DataOra) VALUES (5, 5, 5, TO_DATE('2024-02-10 18:30', 'YYYY-MM-DD HH24:MI'));
INSERT INTO CLASE (ClasaID, AntrenorID, TipID, DataOra) VALUES (6, 1, 6, TO_DATE('2024-02-11 08:00', 'YYYY-MM-DD HH24:MI'));
INSERT INTO CLASE (ClasaID, AntrenorID, TipID, DataOra) VALUES (7, 2, 1, TO_DATE('2024-02-11 09:30', 'YYYY-MM-DD HH24:MI'));
INSERT INTO CLASE (ClasaID, AntrenorID, TipID, DataOra) VALUES (8, 3, 2, TO_DATE('2024-02-11 11:00', 'YYYY-MM-DD HH24:MI'));
INSERT INTO CLASE (ClasaID, AntrenorID, TipID, DataOra) VALUES (9, 4, 3, TO_DATE('2024-02-11 17:00', 'YYYY-MM-DD HH24:MI'));
INSERT INTO CLASE (ClasaID, AntrenorID, TipID, DataOra) VALUES (10, 5, 4, TO_DATE('2024-02-11 18:30', 'YYYY-MM-DD HH24:MI'));
INSERT INTO CLASE (ClasaID, AntrenorID, TipID, DataOra) VALUES (11, 1, 5, TO_DATE('2024-02-12 08:00', 'YYYY-MM-DD HH24:MI'));
INSERT INTO CLASE (ClasaID, AntrenorID, TipID, DataOra) VALUES (12, 2, 6, TO_DATE('2024-02-12 09:30', 'YYYY-MM-DD HH24:MI'));
INSERT INTO CLASE (ClasaID, AntrenorID, TipID, DataOra) VALUES (13, 3, 1, TO_DATE('2024-02-12 11:00', 'YYYY-MM-DD HH24:MI'));
INSERT INTO CLASE (ClasaID, AntrenorID, TipID, DataOra) VALUES (14, 4, 2, TO_DATE('2024-02-12 17:00', 'YYYY-MM-DD HH24:MI'));
INSERT INTO CLASE (ClasaID, AntrenorID, TipID, DataOra) VALUES (15, 5, 3, TO_DATE('2024-02-12 18:30', 'YYYY-MM-DD HH24:MI'));
INSERT INTO CLASE (ClasaID, AntrenorID, TipID, DataOra) VALUES (16, 1, 4, TO_DATE('2024-02-13 08:00', 'YYYY-MM-DD HH24:MI'));
INSERT INTO CLASE (ClasaID, AntrenorID, TipID, DataOra) VALUES (17, 2, 5, TO_DATE('2024-02-13 09:30', 'YYYY-MM-DD HH24:MI'));
INSERT INTO CLASE (ClasaID, AntrenorID, TipID, DataOra) VALUES (18, 3, 6, TO_DATE('2024-02-13 11:00', 'YYYY-MM-DD HH24:MI'));
INSERT INTO CLASE (ClasaID, AntrenorID, TipID, DataOra) VALUES (19, 4, 1, TO_DATE('2024-02-13 17:00', 'YYYY-MM-DD HH24:MI'));
INSERT INTO CLASE (ClasaID, AntrenorID, TipID, DataOra) VALUES (20, 5, 2, TO_DATE('2024-02-13 18:30', 'YYYY-MM-DD HH24:MI'));


-- 7. POPULARE INSCRIERI
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (1, 1);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (2, 1);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (3, 2);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (4, 2);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (5, 3);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (6, 4);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (7, 5);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (8, 6);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (9, 7);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (10, 8);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (1, 11);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (2, 12);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (15, 1);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (16, 2);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (20, 5);

-- 8. POPULARE COMPETITII
INSERT INTO COMPETITII (CompetitieID, Nume, Data, Taxa) VALUES (1, 'Winter Games', TO_DATE('2024-01-20', 'YYYY-MM-DD'), 100);
INSERT INTO COMPETITII (CompetitieID, Nume, Data, Taxa) VALUES (2, 'National Championship', TO_DATE('2024-05-15', 'YYYY-MM-DD'), 250);
INSERT INTO COMPETITII (CompetitieID, Nume, Data, Taxa) VALUES (3, 'Summer Throwdown', TO_DATE('2024-07-20', 'YYYY-MM-DD'), 150);
INSERT INTO COMPETITII (CompetitieID, Nume, Data, Taxa) VALUES (4, 'Rookie Challenge', TO_DATE('2024-09-10', 'YYYY-MM-DD'), 80);
INSERT INTO COMPETITII (CompetitieID, Nume, Data, Taxa) VALUES (5, 'Team Series', TO_DATE('2024-11-05', 'YYYY-MM-DD'), 120);
INSERT INTO COMPETITII (CompetitieID, Nume, Data, Taxa) VALUES (6, 'Masters Cup', TO_DATE('2024-03-30', 'YYYY-MM-DD'), 150);
INSERT INTO COMPETITII (CompetitieID, Nume, Data, Taxa) VALUES (7, 'Junior League', TO_DATE('2024-06-01', 'YYYY-MM-DD'), 50);

-- 9. POPULARE PARTICIPARI_COMPETITIE
INSERT INTO PARTICIPARI_COMPETITIE (CompetitieID, MembruID, LoculObtinut) VALUES (1, 1, 5);
INSERT INTO PARTICIPARI_COMPETITIE (CompetitieID, MembruID, LoculObtinut) VALUES (1, 3, 12);
INSERT INTO PARTICIPARI_COMPETITIE (CompetitieID, MembruID, LoculObtinut) VALUES (2, 2, 2);
INSERT INTO PARTICIPARI_COMPETITIE (CompetitieID, MembruID, LoculObtinut) VALUES (3, 8, 1);
INSERT INTO PARTICIPARI_COMPETITIE (CompetitieID, MembruID, LoculObtinut) VALUES (4, 15, 8);
INSERT INTO PARTICIPARI_COMPETITIE (CompetitieID, MembruID, LoculObtinut) VALUES (4, 20, 10);
INSERT INTO PARTICIPARI_COMPETITIE (CompetitieID, MembruID, LoculObtinut) VALUES (5, 11, 3);
INSERT INTO PARTICIPARI_COMPETITIE (CompetitieID, MembruID, LoculObtinut) VALUES (6, 25, 1);
INSERT INTO PARTICIPARI_COMPETITIE (CompetitieID, MembruID, LoculObtinut) VALUES (7, 4, 15);
INSERT INTO PARTICIPARI_COMPETITIE (CompetitieID, MembruID, LoculObtinut) VALUES (2, 1, 10);

-- 10. POPULARE ECHIPAMENTE
INSERT INTO ECHIPAMENTE (EchipamentID, NumeEchipament, CantitateTotala) VALUES (1, 'Bara Olimpica', 20);
INSERT INTO ECHIPAMENTE (EchipamentID, NumeEchipament, CantitateTotala) VALUES (2, 'Discuri 10kg', 50);
INSERT INTO ECHIPAMENTE (EchipamentID, NumeEchipament, CantitateTotala) VALUES (3, 'Bara Tractiuni', 30);
INSERT INTO ECHIPAMENTE (EchipamentID, NumeEchipament, CantitateTotala) VALUES (4, 'Kettlebell 16kg', 25);
INSERT INTO ECHIPAMENTE (EchipamentID, NumeEchipament, CantitateTotala) VALUES (5, 'Cutie Plyo', 15);
INSERT INTO ECHIPAMENTE (EchipamentID, NumeEchipament, CantitateTotala) VALUES (6, 'Coarda Sarit', 40);
INSERT INTO ECHIPAMENTE (EchipamentID, NumeEchipament, CantitateTotala) VALUES (7, 'Inele Gimnastica', 10);
INSERT INTO ECHIPAMENTE (EchipamentID, NumeEchipament, CantitateTotala) VALUES (8, 'Banca Impins', 12);

-- 11. POPULARE NECESAR_ECHIPAMENT
INSERT INTO NECESAR_ECHIPAMENT (TipID, EchipamentID) VALUES (1, 3); -- Murph: Tractiuni
INSERT INTO NECESAR_ECHIPAMENT (TipID, EchipamentID) VALUES (2, 3); -- Fran: Tractiuni
INSERT INTO NECESAR_ECHIPAMENT (TipID, EchipamentID) VALUES (2, 1); -- Fran: Bara
INSERT INTO NECESAR_ECHIPAMENT (TipID, EchipamentID) VALUES (3, 1); -- Grace: Bara
INSERT INTO NECESAR_ECHIPAMENT (TipID, EchipamentID) VALUES (4, 3); -- Cindy: Tractiuni
INSERT INTO NECESAR_ECHIPAMENT (TipID, EchipamentID) VALUES (5, 3); -- Helen: Tractiuni
INSERT INTO NECESAR_ECHIPAMENT (TipID, EchipamentID) VALUES (5, 4); -- Helen: KB
INSERT INTO NECESAR_ECHIPAMENT (TipID, EchipamentID) VALUES (6, 1); -- Linda: Bara
INSERT INTO NECESAR_ECHIPAMENT (TipID, EchipamentID) VALUES (6, 8); -- Linda: Banca

-- 12. POPULARE PRODUSE
INSERT INTO PRODUSE (ProdusID, NumeProdus, PretCurent, Stoc) VALUES (1, 'Proteine Whey', 180, 50);
INSERT INTO PRODUSE (ProdusID, NumeProdus, PretCurent, Stoc) VALUES (2, 'Creatina', 90, 40);
INSERT INTO PRODUSE (ProdusID, NumeProdus, PretCurent, Stoc) VALUES (3, 'Baton Proteic', 12, 200);
INSERT INTO PRODUSE (ProdusID, NumeProdus, PretCurent, Stoc) VALUES (4, 'Tricou Sala', 85, 30);
INSERT INTO PRODUSE (ProdusID, NumeProdus, PretCurent, Stoc) VALUES (5, 'Shaker', 25, 60);
INSERT INTO PRODUSE (ProdusID, NumeProdus, PretCurent, Stoc) VALUES (6, 'Magneziu Lichid', 35, 25);
INSERT INTO PRODUSE (ProdusID, NumeProdus, PretCurent, Stoc) VALUES (7, 'Genunchiere', 110, 15);

-- 13. POPULARE ACHIZITII
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, Cantitate) VALUES (1, 1, 1, 1);
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, Cantitate) VALUES (2, 2, 3, 5);
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, Cantitate) VALUES (3, 5, 2, 1);
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, Cantitate) VALUES (4, 10, 4, 2);
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, Cantitate) VALUES (5, 3, 3, 10);
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, Cantitate) VALUES (6, 8, 1, 2);
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, Cantitate) VALUES (7, 12, 6, 1);
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, Cantitate) VALUES (8, 1, 5, 1);
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, Cantitate) VALUES (9, 20, 7, 1);
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, Cantitate) VALUES (10, 4, 3, 3);
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, Cantitate) VALUES (11, 15, 1, 1);
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, Cantitate) VALUES (12, 6, 2, 1);
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, Cantitate) VALUES (13, 22, 4, 1);
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, Cantitate) VALUES (14, 9, 3, 4);
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, Cantitate) VALUES (15, 11, 1, 1);
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, Cantitate) VALUES (16, 1, 2, 1);

COMMIT;
