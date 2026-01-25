-- 1. POPULARE ABONAMENTE
INSERT INTO ABONAMENTE (AbonamentID, TipAbonament, Pret) VALUES (1, 'Standard Crossfit', 250);
INSERT INTO ABONAMENTE (AbonamentID, TipAbonament, Pret) VALUES (2, 'Student Crossfit', 150);
INSERT INTO ABONAMENTE (AbonamentID, TipAbonament, Pret) VALUES (3, 'Open Gym', 200);
INSERT INTO ABONAMENTE (AbonamentID, TipAbonament, Pret) VALUES (4, 'Premium All Access', 400);
INSERT INTO ABONAMENTE (AbonamentID, TipAbonament, Pret) VALUES (5, 'Full Time', 350);
INSERT INTO ABONAMENTE (AbonamentID, TipAbonament, Pret) VALUES (6, 'Weekend Only', 120);

-- 2. POPULARE ANTRENORI
INSERT INTO ANTRENORI (AntrenorID, Nume, Prenume, Specializare) VALUES (1, 'Popescu', 'Ion', 'Crossfit L2');
INSERT INTO ANTRENORI (AntrenorID, Nume, Prenume, Specializare) VALUES (2, 'Ionescu', 'Maria', 'Gimnastica');
INSERT INTO ANTRENORI (AntrenorID, Nume, Prenume, Specializare) VALUES (3, 'Vasilescu', 'Andrei', 'Haltere');
INSERT INTO ANTRENORI (AntrenorID, Nume, Prenume, Specializare) VALUES (4, 'Radu', 'Elena', 'Cardio');
INSERT INTO ANTRENORI (AntrenorID, Nume, Prenume, Specializare) VALUES (5, 'Dumitrescu', 'Cristian', 'Crossfit L1');
INSERT INTO ANTRENORI (AntrenorID, Nume, Prenume, Specializare) VALUES (6, 'Stoica', 'Laura', 'Mobilitate');
INSERT INTO ANTRENORI (AntrenorID, Nume, Prenume, Specializare) VALUES (7, 'Manea', 'Victor', 'Crossfit L1');
INSERT INTO ANTRENORI (AntrenorID, Nume, Prenume, Specializare) VALUES (8, 'Diaconu', 'Alina', 'Gimnastica');
INSERT INTO ANTRENORI (AntrenorID, Nume, Prenume, Specializare) VALUES (9, 'Popa', 'Sergiu', 'Haltere');
INSERT INTO ANTRENORI (AntrenorID, Nume, Prenume, Specializare) VALUES (10, 'Voinea', 'Carmen', 'Cardio');
INSERT INTO ANTRENORI (AntrenorID, Nume, Prenume, Specializare) VALUES (11, 'Florea', 'Mihai', 'Crossfit L2');
INSERT INTO ANTRENORI (AntrenorID, Nume, Prenume, Specializare) VALUES (12, 'Dinu', 'Robert', 'Crossfit L1');
INSERT INTO ANTRENORI (AntrenorID, Nume, Prenume, Specializare) VALUES (13, 'Toma', 'Diana', 'Mobilitate');
INSERT INTO ANTRENORI (AntrenorID, Nume, Prenume, Specializare) VALUES (14, 'Stanciu', 'Gabriel', 'Haltere');
INSERT INTO ANTRENORI (AntrenorID, Nume, Prenume, Specializare) VALUES (15, 'Neagu', 'Oana', 'Gimnastica');
INSERT INTO ANTRENORI (AntrenorID, Nume, Prenume, Specializare) VALUES (16, 'Preda', 'Alexandru', 'Crossfit L1');
INSERT INTO ANTRENORI (AntrenorID, Nume, Prenume, Specializare) VALUES (17, 'Barbu', 'Silviu', 'Cardio');
INSERT INTO ANTRENORI (AntrenorID, Nume, Prenume, Specializare) VALUES (18, 'Nistor', 'Irina', 'Mobilitate');
INSERT INTO ANTRENORI (AntrenorID, Nume, Prenume, Specializare) VALUES (19, 'Mocanu', 'Florin', 'Crossfit L2');
INSERT INTO ANTRENORI (AntrenorID, Nume, Prenume, Specializare) VALUES (20, 'Oprea', 'Simona', 'Gimnastica');

-- 3. POPULARE PRODUSE
INSERT INTO PRODUSE (ProdusID, NumeProdus, PretCurent, Stoc) VALUES (1, 'Proteine Whey 1kg', 120, 50);
INSERT INTO PRODUSE (ProdusID, NumeProdus, PretCurent, Stoc) VALUES (2, 'Tricou Crossfit', 85, 100);
INSERT INTO PRODUSE (ProdusID, NumeProdus, PretCurent, Stoc) VALUES (3, 'Baton Proteic', 12, 200);
INSERT INTO PRODUSE (ProdusID, NumeProdus, PretCurent, Stoc) VALUES (4, 'Coarda Sarit', 45, 30);
INSERT INTO PRODUSE (ProdusID, NumeProdus, PretCurent, Stoc) VALUES (5, 'Magneziu Lichid', 35, 40);
INSERT INTO PRODUSE (ProdusID, NumeProdus, PretCurent, Stoc) VALUES (6, 'Shaker 700ml', 25, 60);
INSERT INTO PRODUSE (ProdusID, NumeProdus, PretCurent, Stoc) VALUES (7, 'Centura Haltere', 150, 15);
INSERT INTO PRODUSE (ProdusID, NumeProdus, PretCurent, Stoc) VALUES (8, 'Ghetere Crossfit', 450, 20);
INSERT INTO PRODUSE (ProdusID, NumeProdus, PretCurent, Stoc) VALUES (9, 'Mansete Maini', 60, 50);
INSERT INTO PRODUSE (ProdusID, NumeProdus, PretCurent, Stoc) VALUES (10, 'Genunchiere 5mm', 130, 25);
INSERT INTO PRODUSE (ProdusID, NumeProdus, PretCurent, Stoc) VALUES (11, 'Genunchiere 7mm', 140, 20);
INSERT INTO PRODUSE (ProdusID, NumeProdus, PretCurent, Stoc) VALUES (12, 'Banda Elastica Rosie', 40, 30);
INSERT INTO PRODUSE (ProdusID, NumeProdus, PretCurent, Stoc) VALUES (13, 'Banda Elastica Verde', 55, 25);
INSERT INTO PRODUSE (ProdusID, NumeProdus, PretCurent, Stoc) VALUES (14, 'Creatina Monohidrata', 80, 45);
INSERT INTO PRODUSE (ProdusID, NumeProdus, PretCurent, Stoc) VALUES (15, 'Bautura Izotonica', 10, 150);
INSERT INTO PRODUSE (ProdusID, NumeProdus, PretCurent, Stoc) VALUES (16, 'Prosop Sala', 30, 80);
INSERT INTO PRODUSE (ProdusID, NumeProdus, PretCurent, Stoc) VALUES (17, 'Geanta Sport', 180, 15);
INSERT INTO PRODUSE (ProdusID, NumeProdus, PretCurent, Stoc) VALUES (18, 'Curea Piele', 160, 10);
INSERT INTO PRODUSE (ProdusID, NumeProdus, PretCurent, Stoc) VALUES (19, 'Pre-Workout', 110, 40);
INSERT INTO PRODUSE (ProdusID, NumeProdus, PretCurent, Stoc) VALUES (20, 'Multivitamine', 65, 55);

-- 4. POPULARE ECHIPAMENTE
INSERT INTO ECHIPAMENTE (EchipamentID, NumeEchipament, CantitateTotala) VALUES (1, 'Bara Olimpica M', 15);
INSERT INTO ECHIPAMENTE (EchipamentID, NumeEchipament, CantitateTotala) VALUES (2, 'Discuri 10kg', 40);
INSERT INTO ECHIPAMENTE (EchipamentID, NumeEchipament, CantitateTotala) VALUES (3, 'Concept2 Rower', 8);
INSERT INTO ECHIPAMENTE (EchipamentID, NumeEchipament, CantitateTotala) VALUES (4, 'Kettlebell 16kg', 25);
INSERT INTO ECHIPAMENTE (EchipamentID, NumeEchipament, CantitateTotala) VALUES (5, 'Inele Gimnastica', 10);
INSERT INTO ECHIPAMENTE (EchipamentID, NumeEchipament, CantitateTotala) VALUES (6, 'Assault AirBike', 5);
INSERT INTO ECHIPAMENTE (EchipamentID, NumeEchipament, CantitateTotala) VALUES (7, 'Box Jump (Lemn)', 12);
INSERT INTO ECHIPAMENTE (EchipamentID, NumeEchipament, CantitateTotala) VALUES (8, 'Bara Olimpica F', 10);
INSERT INTO ECHIPAMENTE (EchipamentID, NumeEchipament, CantitateTotala) VALUES (9, 'Discuri 5kg', 40);
INSERT INTO ECHIPAMENTE (EchipamentID, NumeEchipament, CantitateTotala) VALUES (10, 'Discuri 15kg', 30);
INSERT INTO ECHIPAMENTE (EchipamentID, NumeEchipament, CantitateTotala) VALUES (11, 'Discuri 20kg', 30);
INSERT INTO ECHIPAMENTE (EchipamentID, NumeEchipament, CantitateTotala) VALUES (12, 'SkiErg', 4);
INSERT INTO ECHIPAMENTE (EchipamentID, NumeEchipament, CantitateTotala) VALUES (13, 'Kettlebell 24kg', 20);
INSERT INTO ECHIPAMENTE (EchipamentID, NumeEchipament, CantitateTotala) VALUES (14, 'Kettlebell 8kg', 15);
INSERT INTO ECHIPAMENTE (EchipamentID, NumeEchipament, CantitateTotala) VALUES (15, 'Dumbbell 15kg', 12);
INSERT INTO ECHIPAMENTE (EchipamentID, NumeEchipament, CantitateTotala) VALUES (16, 'Dumbbell 22.5kg', 12);
INSERT INTO ECHIPAMENTE (EchipamentID, NumeEchipament, CantitateTotala) VALUES (17, 'MedBall 6kg', 10);
INSERT INTO ECHIPAMENTE (EchipamentID, NumeEchipament, CantitateTotala) VALUES (18, 'MedBall 9kg', 10);
INSERT INTO ECHIPAMENTE (EchipamentID, NumeEchipament, CantitateTotala) VALUES (19, 'AbMat', 20);
INSERT INTO ECHIPAMENTE (EchipamentID, NumeEchipament, CantitateTotala) VALUES (20, 'Coarda Catarare', 4);

-- 5. POPULARE TIPURI_ANTRENAMENT
INSERT INTO TIPURI_ANTRENAMENT (TipAntrenamentID, NumeWOD, Descriere) VALUES (1, 'Murph', '1 mile run, 100 pullups, 200 pushups, 300 squats');
INSERT INTO TIPURI_ANTRENAMENT (TipAntrenamentID, NumeWOD, Descriere) VALUES (2, 'Fran', '21-15-9 Thrusters and Pullups');
INSERT INTO TIPURI_ANTRENAMENT (TipAntrenamentID, NumeWOD, Descriere) VALUES (3, 'EMOM Haltere', 'Every Minute On the Minute: 3 Clean and Jerks');
INSERT INTO TIPURI_ANTRENAMENT (TipAntrenamentID, NumeWOD, Descriere) VALUES (4, 'Yoga Mobility', 'Sesiune de mobilitate si stretching');
INSERT INTO TIPURI_ANTRENAMENT (TipAntrenamentID, NumeWOD, Descriere) VALUES (5, 'Cindy', '20 min AMRAP: 5 pullups, 10 pushups, 15 squats');
INSERT INTO TIPURI_ANTRENAMENT (TipAntrenamentID, NumeWOD, Descriere) VALUES (6, 'Linda', '10-9-8...1 reps: Deadlift, Bench Press, Clean');
INSERT INTO TIPURI_ANTRENAMENT (TipAntrenamentID, NumeWOD, Descriere) VALUES (7, 'Grace', '30 Clean and Jerks for time (60kg/43kg)');
INSERT INTO TIPURI_ANTRENAMENT (TipAntrenamentID, NumeWOD, Descriere) VALUES (8, 'Isabel', '30 Snatches for time (60kg/43kg)');
INSERT INTO TIPURI_ANTRENAMENT (TipAntrenamentID, NumeWOD, Descriere) VALUES (9, 'Karen', '150 Wall Balls for time');
INSERT INTO TIPURI_ANTRENAMENT (TipAntrenamentID, NumeWOD, Descriere) VALUES (10, 'Annie', '50-40-30-20-10 Double Unders and Sit-ups');
INSERT INTO TIPURI_ANTRENAMENT (TipAntrenamentID, NumeWOD, Descriere) VALUES (11, 'Fight Gone Bad', '3 rounds, 1 min per station, max reps');
INSERT INTO TIPURI_ANTRENAMENT (TipAntrenamentID, NumeWOD, Descriere) VALUES (12, 'Helen', '3 rounds: 400m run, 21 KB swings, 12 Pull-ups');
INSERT INTO TIPURI_ANTRENAMENT (TipAntrenamentID, NumeWOD, Descriere) VALUES (13, 'Diane', '21-15-9 Deadlift and Handstand Push-ups');
INSERT INTO TIPURI_ANTRENAMENT (TipAntrenamentID, NumeWOD, Descriere) VALUES (14, 'Elizabeth', '21-15-9 Cleans and Ring Dips');
INSERT INTO TIPURI_ANTRENAMENT (TipAntrenamentID, NumeWOD, Descriere) VALUES (15, 'Tabata This', 'Tabata intervals: Row, Squat, Pull-up, Push-up');
INSERT INTO TIPURI_ANTRENAMENT (TipAntrenamentID, NumeWOD, Descriere) VALUES (16, 'DT', '5 rounds: 12 Deadlift, 9 Hang Power Clean');

-- 6. POPULARE COMPETITII
INSERT INTO COMPETITII (CompetitieID, Nume, Data, Locatie, Taxa) VALUES (1, 'National Crossfit Games', TO_DATE('2024-06-15', 'YYYY-MM-DD'), 'Bucuresti Arena', 100);
INSERT INTO COMPETITII (CompetitieID, Nume, Data, Locatie, Taxa) VALUES (2, 'Summer Throwdown', TO_DATE('2024-08-20', 'YYYY-MM-DD'), 'Constanta Beach', 150);
INSERT INTO COMPETITII (CompetitieID, Nume, Data, Locatie, Taxa) VALUES (3, 'Winter Warrior', TO_DATE('2024-12-05', 'YYYY-MM-DD'), 'Brasov Gym', 80);
INSERT INTO COMPETITII (CompetitieID, Nume, Data, Locatie, Taxa) VALUES (4, 'Cluj Napoca Challenge', TO_DATE('2024-09-10', 'YYYY-MM-DD'), 'Cluj Polyvalent', 120);
INSERT INTO COMPETITII (CompetitieID, Nume, Data, Locatie, Taxa) VALUES (5, 'Iasi Fitness Cup', TO_DATE('2024-10-25', 'YYYY-MM-DD'), 'Iasi Sport Hall', 90);
INSERT INTO COMPETITII (CompetitieID, Nume, Data, Locatie, Taxa) VALUES (6, 'Timisoara Strongest', TO_DATE('2024-07-01', 'YYYY-MM-DD'), 'Timisoara Expo', 110);
INSERT INTO COMPETITII (CompetitieID, Nume, Data, Locatie, Taxa) VALUES (7, 'Bucharest Easter Cup', TO_DATE('2024-04-20', 'YYYY-MM-DD'), 'Bucuresti', 70);

-- 7. POPULARE MEMBRI
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (1, 'Dumitru', 'Alex', 'alex.dumitru@gmail.com', 5);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (2, 'Stan', 'Mihai', 'mihai.stan@gmail.com', 2);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (3, 'Popa', 'Ana', 'ana.popa@gmail.com', 4);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (4, 'Gheorghe', 'Cristina', 'cris.gheorghe@gmail.com', 5);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (5, 'Marin', 'Vlad', 'vlad.marin@gmail.com', 3);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (6, 'Dobre', 'Ioana', 'ioana.dobre@gmail.com', 2);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (7, 'Enache', 'Stefan', 'stefan.enache@gmail.com', 1);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (8, 'Lazar', 'George', 'george.lazar@gmail.com', 5);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (9, 'Munteanu', 'Andreea', 'andreea.munteanu@gmail.com', 4);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (10, 'Iancu', 'Razvan', 'razvan.iancu@gmail.com', 3);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (11, 'Badea', 'Ionut', 'ionut.badea@gmail.com', 1);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (12, 'Costea', 'Elena', 'elena.costea@gmail.com', 6);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (13, 'Serban', 'Adrian', 'adrian.serban@gmail.com', 4);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (14, 'Ungureanu', 'Marius', 'marius.ungureanu@gmail.com', 5);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (15, 'Dragan', 'Roxana', 'roxana.dragan@gmail.com', 2);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (16, 'Avram', 'Daniel', 'daniel.avram@gmail.com', 3);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (17, 'Ciobanu', 'Gabriela', 'gabriela.ciobanu@gmail.com', 1);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (18, 'Pavel', 'Lucian', 'lucian.pavel@gmail.com', 6);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (19, 'Tudor', 'Simona', 'simona.tudor@gmail.com', 5);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (20, 'Vasile', 'Bogdan', 'bogdan.vasile@gmail.com', 4);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (21, 'Cristea', 'Octavian', 'octavian.cristea@gmail.com', 1);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (22, 'Nica', 'Valentina', 'valentina.nica@gmail.com', 2);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (23, 'Fodor', 'Eduard', 'eduard.fodor@gmail.com', 3);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (24, 'Sandu', 'Laura', 'laura.sandu@gmail.com', 4);
INSERT INTO MEMBRI (MembruID, Nume, Prenume, Email, AbonamentID) VALUES (25, 'Olteanu', 'Victor', 'victor.olteanu@gmail.com', 5);

-- 8. POPULARE CLASE
INSERT INTO CLASE (ClasaID, NumeWOD, DescriereWOD, DataOra, AntrenorID, TipAntrenamentID) 
VALUES (1, 'Morning Murph', 'Rezistenta', TO_TIMESTAMP('2024-05-10 08:00:00', 'YYYY-MM-DD HH24:MI:SS'), 1, 1);
INSERT INTO CLASE (ClasaID, NumeWOD, DescriereWOD, DataOra, AntrenorID, TipAntrenamentID) 
VALUES (2, 'Evening Fran', 'Intensitate', TO_TIMESTAMP('2024-05-10 18:30:00', 'YYYY-MM-DD HH24:MI:SS'), 3, 2);
INSERT INTO CLASE (ClasaID, NumeWOD, DescriereWOD, DataOra, AntrenorID, TipAntrenamentID) 
VALUES (3, 'Weightlifting Tech', 'Smuls/Aruncat', TO_TIMESTAMP('2024-05-11 10:00:00', 'YYYY-MM-DD HH24:MI:SS'), 3, 3);
INSERT INTO CLASE (ClasaID, NumeWOD, DescriereWOD, DataOra, AntrenorID, TipAntrenamentID) 
VALUES (4, 'Sunday Yoga', 'Mobilitate', TO_TIMESTAMP('2024-05-12 09:00:00', 'YYYY-MM-DD HH24:MI:SS'), 6, 4);
INSERT INTO CLASE (ClasaID, NumeWOD, DescriereWOD, DataOra, AntrenorID, TipAntrenamentID) 
VALUES (5, 'Cindy AMRAP', 'Bodyweight', TO_TIMESTAMP('2024-05-13 17:00:00', 'YYYY-MM-DD HH24:MI:SS'), 5, 5);
INSERT INTO CLASE (ClasaID, NumeWOD, DescriereWOD, DataOra, AntrenorID, TipAntrenamentID) 
VALUES (6, 'Heavy Day', 'Deadlifts', TO_TIMESTAMP('2024-05-14 18:00:00', 'YYYY-MM-DD HH24:MI:SS'), 11, 6);
INSERT INTO CLASE (ClasaID, NumeWOD, DescriereWOD, DataOra, AntrenorID, TipAntrenamentID) 
VALUES (7, 'Fast Grace', 'Cycling Barbell', TO_TIMESTAMP('2024-05-14 19:00:00', 'YYYY-MM-DD HH24:MI:SS'), 7, 7);
INSERT INTO CLASE (ClasaID, NumeWOD, DescriereWOD, DataOra, AntrenorID, TipAntrenamentID) 
VALUES (8, 'Snatch Clinic', 'Tehnica', TO_TIMESTAMP('2024-05-15 08:00:00', 'YYYY-MM-DD HH24:MI:SS'), 9, 8);
INSERT INTO CLASE (ClasaID, NumeWOD, DescriereWOD, DataOra, AntrenorID, TipAntrenamentID) 
VALUES (9, 'Leg Burner', 'Wall Balls', TO_TIMESTAMP('2024-05-15 17:30:00', 'YYYY-MM-DD HH24:MI:SS'), 12, 9);
INSERT INTO CLASE (ClasaID, NumeWOD, DescriereWOD, DataOra, AntrenorID, TipAntrenamentID) 
VALUES (10, 'Cardio Madness', 'Double Unders', TO_TIMESTAMP('2024-05-16 07:00:00', 'YYYY-MM-DD HH24:MI:SS'), 4, 10);
INSERT INTO CLASE (ClasaID, NumeWOD, DescriereWOD, DataOra, AntrenorID, TipAntrenamentID) 
VALUES (11, 'Fight Gone Bad', 'Circuit', TO_TIMESTAMP('2024-05-16 18:30:00', 'YYYY-MM-DD HH24:MI:SS'), 19, 11);
INSERT INTO CLASE (ClasaID, NumeWOD, DescriereWOD, DataOra, AntrenorID, TipAntrenamentID) 
VALUES (12, 'Helen Test', 'Running & KB', TO_TIMESTAMP('2024-05-17 12:00:00', 'YYYY-MM-DD HH24:MI:SS'), 16, 12);
INSERT INTO CLASE (ClasaID, NumeWOD, DescriereWOD, DataOra, AntrenorID, TipAntrenamentID) 
VALUES (13, 'Diane HSPU', 'Gymnastics strength', TO_TIMESTAMP('2024-05-17 19:00:00', 'YYYY-MM-DD HH24:MI:SS'), 15, 13);
INSERT INTO CLASE (ClasaID, NumeWOD, DescriereWOD, DataOra, AntrenorID, TipAntrenamentID) 
VALUES (14, 'Weekend Warrior', 'Team WOD', TO_TIMESTAMP('2024-05-18 10:00:00', 'YYYY-MM-DD HH24:MI:SS'), 1, 1);
INSERT INTO CLASE (ClasaID, NumeWOD, DescriereWOD, DataOra, AntrenorID, TipAntrenamentID) 
VALUES (15, 'Mobility Flow', 'Recovery', TO_TIMESTAMP('2024-05-18 11:30:00', 'YYYY-MM-DD HH24:MI:SS'), 18, 4);
INSERT INTO CLASE (ClasaID, NumeWOD, DescriereWOD, DataOra, AntrenorID, TipAntrenamentID) 
VALUES (16, 'Elizabeth Fast', 'Sprint', TO_TIMESTAMP('2024-05-19 18:00:00', 'YYYY-MM-DD HH24:MI:SS'), 5, 14);
INSERT INTO CLASE (ClasaID, NumeWOD, DescriereWOD, DataOra, AntrenorID, TipAntrenamentID) 
VALUES (17, 'Tabata Sweat', 'Intervals', TO_TIMESTAMP('2024-05-20 07:00:00', 'YYYY-MM-DD HH24:MI:SS'), 17, 15);
INSERT INTO CLASE (ClasaID, NumeWOD, DescriereWOD, DataOra, AntrenorID, TipAntrenamentID) 
VALUES (18, 'DT Hero', 'Barbell Cycling', TO_TIMESTAMP('2024-05-20 18:30:00', 'YYYY-MM-DD HH24:MI:SS'), 14, 16);
INSERT INTO CLASE (ClasaID, NumeWOD, DescriereWOD, DataOra, AntrenorID, TipAntrenamentID) 
VALUES (19, 'Rowing Tech', 'Engine building', TO_TIMESTAMP('2024-05-21 08:00:00', 'YYYY-MM-DD HH24:MI:SS'), 10, 4);
INSERT INTO CLASE (ClasaID, NumeWOD, DescriereWOD, DataOra, AntrenorID, TipAntrenamentID) 
VALUES (20, 'Open Prep', 'Strategy', TO_TIMESTAMP('2024-05-21 19:00:00', 'YYYY-MM-DD HH24:MI:SS'), 2, 2);

-- 9. POPULARE ACHIZITII
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, DataAchizitiei, Cantitate) VALUES (1, 1, 1, SYSDATE-5, 1);
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, DataAchizitiei, Cantitate) VALUES (2, 2, 2, SYSDATE-2, 2);
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, DataAchizitiei, Cantitate) VALUES (3, 3, 3, SYSDATE, 5);
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, DataAchizitiei, Cantitate) VALUES (4, 1, 4, SYSDATE-10, 1);
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, DataAchizitiei, Cantitate) VALUES (5, 7, 7, SYSDATE-1, 1);
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, DataAchizitiei, Cantitate) VALUES (6, 9, 5, SYSDATE-3, 2);
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, DataAchizitiei, Cantitate) VALUES (7, 11, 1, SYSDATE-15, 1);
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, DataAchizitiei, Cantitate) VALUES (8, 12, 10, SYSDATE-12, 2);
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, DataAchizitiei, Cantitate) VALUES (9, 15, 15, SYSDATE-1, 10);
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, DataAchizitiei, Cantitate) VALUES (10, 18, 19, SYSDATE-4, 1);
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, DataAchizitiei, Cantitate) VALUES (11, 20, 14, SYSDATE-20, 1);
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, DataAchizitiei, Cantitate) VALUES (12, 4, 8, SYSDATE-25, 1);
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, DataAchizitiei, Cantitate) VALUES (13, 22, 16, SYSDATE-2, 1);
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, DataAchizitiei, Cantitate) VALUES (14, 25, 2, SYSDATE-6, 1);
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, DataAchizitiei, Cantitate) VALUES (15, 5, 20, SYSDATE-30, 1);
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, DataAchizitiei, Cantitate) VALUES (16, 8, 9, SYSDATE-3, 1);
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, DataAchizitiei, Cantitate) VALUES (17, 13, 13, SYSDATE-8, 1);
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, DataAchizitiei, Cantitate) VALUES (18, 16, 6, SYSDATE-1, 1);
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, DataAchizitiei, Cantitate) VALUES (19, 10, 3, SYSDATE, 3);
INSERT INTO ACHIZITII (AchizitieID, MembruID, ProdusID, DataAchizitiei, Cantitate) VALUES (20, 24, 1, SYSDATE-2, 1);

-- 10. POPULARE INSCRIERI
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (1, 1);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (2, 1);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (3, 2);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (1, 3);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (4, 4);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (8, 5);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (9, 5);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (10, 2);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (11, 6);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (12, 6);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (13, 7);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (14, 8);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (15, 9);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (16, 10);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (17, 11);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (18, 12);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (19, 13);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (20, 14);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (21, 15);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (22, 16);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (23, 17);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (24, 18);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (25, 19);
INSERT INTO INSCRIERI (MembruID, ClasaID) VALUES (5, 20);

-- 11. POPULARE MENTORAT
INSERT INTO MENTORAT (AntrenorID, MembruID) VALUES (1, 1);
INSERT INTO MENTORAT (AntrenorID, MembruID) VALUES (3, 2);
INSERT INTO MENTORAT (AntrenorID, MembruID) VALUES (2, 4);
INSERT INTO MENTORAT (AntrenorID, MembruID) VALUES (5, 8);
INSERT INTO MENTORAT (AntrenorID, MembruID) VALUES (7, 11);
INSERT INTO MENTORAT (AntrenorID, MembruID) VALUES (8, 12);
INSERT INTO MENTORAT (AntrenorID, MembruID) VALUES (9, 15);
INSERT INTO MENTORAT (AntrenorID, MembruID) VALUES (10, 18);
INSERT INTO MENTORAT (AntrenorID, MembruID) VALUES (11, 20);
INSERT INTO MENTORAT (AntrenorID, MembruID) VALUES (14, 25);

-- 12. POPULARE NECESAR_ECHIPAMENT
INSERT INTO NECESAR_ECHIPAMENT (TipAntrenamentID, EchipamentID) VALUES (2, 1);
INSERT INTO NECESAR_ECHIPAMENT (TipAntrenamentID, EchipamentID) VALUES (3, 1);
INSERT INTO NECESAR_ECHIPAMENT (TipAntrenamentID, EchipamentID) VALUES (3, 2);
INSERT INTO NECESAR_ECHIPAMENT (TipAntrenamentID, EchipamentID) VALUES (6, 1);
INSERT INTO NECESAR_ECHIPAMENT (TipAntrenamentID, EchipamentID) VALUES (6, 2);
INSERT INTO NECESAR_ECHIPAMENT (TipAntrenamentID, EchipamentID) VALUES (7, 1);
INSERT INTO NECESAR_ECHIPAMENT (TipAntrenamentID, EchipamentID) VALUES (8, 1);
INSERT INTO NECESAR_ECHIPAMENT (TipAntrenamentID, EchipamentID) VALUES (9, 17);
INSERT INTO NECESAR_ECHIPAMENT (TipAntrenamentID, EchipamentID) VALUES (10, 4);
INSERT INTO NECESAR_ECHIPAMENT (TipAntrenamentID, EchipamentID) VALUES (10, 19);
INSERT INTO NECESAR_ECHIPAMENT (TipAntrenamentID, EchipamentID) VALUES (12, 4);
INSERT INTO NECESAR_ECHIPAMENT (TipAntrenamentID, EchipamentID) VALUES (16, 1);

-- 13. POPULARE ORAR_TEMPLATE
INSERT INTO ORAR_TEMPLATE (TemplateID, ZiuaSaptamanii, Ora, NumeWOD_Template, AntrenorID, TipAntrenamentID) 
VALUES (1, 1, '08:00', 'Morning Cardio', 4, 1);
INSERT INTO ORAR_TEMPLATE (TemplateID, ZiuaSaptamanii, Ora, NumeWOD_Template, AntrenorID, TipAntrenamentID) 
VALUES (2, 3, '18:00', 'Miercuri Haltere', 3, 3);
INSERT INTO ORAR_TEMPLATE (TemplateID, ZiuaSaptamanii, Ora, NumeWOD_Template, AntrenorID, TipAntrenamentID) 
VALUES (3, 5, '17:00', 'Friday Fun AMRAP', 5, 5);
INSERT INTO ORAR_TEMPLATE (TemplateID, ZiuaSaptamanii, Ora, NumeWOD_Template, AntrenorID, TipAntrenamentID) 
VALUES (4, 2, '19:00', 'Gymnastics Class', 2, 2);
INSERT INTO ORAR_TEMPLATE (TemplateID, ZiuaSaptamanii, Ora, NumeWOD_Template, AntrenorID, TipAntrenamentID) 
VALUES (5, 4, '07:00', 'Early Bird', 1, 6);
INSERT INTO ORAR_TEMPLATE (TemplateID, ZiuaSaptamanii, Ora, NumeWOD_Template, AntrenorID, TipAntrenamentID) 
VALUES (6, 6, '10:00', 'Team Saturday', 11, 11);
INSERT INTO ORAR_TEMPLATE (TemplateID, ZiuaSaptamanii, Ora, NumeWOD_Template, AntrenorID, TipAntrenamentID) 
VALUES (7, 1, '18:30', 'Monday Grind', 7, 7);
INSERT INTO ORAR_TEMPLATE (TemplateID, ZiuaSaptamanii, Ora, NumeWOD_Template, AntrenorID, TipAntrenamentID) 
VALUES (8, 3, '09:00', 'Midweek Mobility', 18, 4);

-- 14. POPULARE PARTICIPARI_COMPETITIE
INSERT INTO PARTICIPARI_COMPETITIE (CompetitieID, MembruID, LoculObtinut) VALUES (1, 1, 5);
INSERT INTO PARTICIPARI_COMPETITIE (CompetitieID, MembruID, LoculObtinut) VALUES (1, 3, 12);
INSERT INTO PARTICIPARI_COMPETITIE (CompetitieID, MembruID, LoculObtinut) VALUES (2, 2, 2);
INSERT INTO PARTICIPARI_COMPETITIE (CompetitieID, MembruID, LoculObtinut) VALUES (3, 8, 1);
INSERT INTO PARTICIPARI_COMPETITIE (CompetitieID, MembruID, LoculObtinut) VALUES (4, 15, 8);
INSERT INTO PARTICIPARI_COMPETITIE (CompetitieID, MembruID, LoculObtinut) VALUES (4, 20, 10);
INSERT INTO PARTICIPARI_COMPETITIE (CompetitieID, MembruID, LoculObtinut) VALUES (5, 11, 3);
INSERT INTO PARTICIPARI_COMPETITIE (CompetitieID, MembruID, LoculObtinut) VALUES (6, 25, 1);
INSERT INTO PARTICIPARI_COMPETITIE (CompetitieID, MembruID, LoculObtinut) VALUES (7, 4, 15);
INSERT INTO PARTICIPARI_COMPETITIE (CompetitieID, MembruID, LoculObtinut) VALUES (2, 18, 20);

-- SINCRONIZARE SECVENTE
ALTER SEQUENCE ABONAMENTE_SEQ RESTART START WITH 20;
ALTER SEQUENCE ANTRENORI_SEQ RESTART START WITH 30;
ALTER SEQUENCE PRODUSE_SEQ RESTART START WITH 30;
ALTER SEQUENCE ECHIPAMENTE_SEQ RESTART START WITH 30;
ALTER SEQUENCE TIPURI_ANTRENAMENT_SEQ RESTART START WITH 30;
ALTER SEQUENCE COMPETITII_SEQ RESTART START WITH 20;
ALTER SEQUENCE MEMBRI_SEQ RESTART START WITH 50;
ALTER SEQUENCE CLASE_SEQ RESTART START WITH 50;
ALTER SEQUENCE ACHIZITII_SEQ RESTART START WITH 50;
ALTER SEQUENCE ORAR_TEMPLATE_SEQ RESTART START WITH 30;

COMMIT;
